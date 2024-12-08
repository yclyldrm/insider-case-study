package pkg

import (
	"encoding/json"
	"fmt"
	"insider-case-study/internal/api"
	"insider-case-study/internal/application"
	"insider-case-study/internal/domain/message"
	"insider-case-study/internal/infrastructure"
	"log"
	"sync"
	"time"
)

type JobService struct {
	ms          application.MessageService
	rClient     *infrastructure.RedisClient
	isJobEnable bool
	stopChannel chan bool
	wg          sync.WaitGroup
	mut         sync.Mutex
}

func NewJobService(ms application.MessageService, redisClient *infrastructure.RedisClient) *JobService {
	return &JobService{
		ms:          ms,
		rClient:     redisClient,
		isJobEnable: true,
		stopChannel: make(chan bool, 1),
	}
}

func (js *JobService) SendMessagesJob() error {
	js.mut.Lock()
	if js.isJobEnable {
		js.mut.Unlock()
		messages, err := js.ms.GetUnsentMessages()
		if err != nil {
			log.Println("error while getting messages", err)
			return err
		}
		ticker := time.NewTicker(2 * time.Minute)
		client := api.NewClient()
		go func(ticker *time.Ticker, client *api.Client, messages []*message.Message) {
			js.runJob(*ticker, *client, messages)
		}(ticker, client, messages)
	} else {
		js.mut.Unlock()
	}
	return nil
}

func (js *JobService) SetJobStatus(status bool) {
	js.mut.Lock()
	js.isJobEnable = status
	if !status {
		select {
		case js.stopChannel <- false:
		default:
		}
	}
	js.mut.Unlock()
}

func (js *JobService) runJob(ticker time.Ticker, client api.Client, messages []*message.Message) {
	if len(messages) == 0 {
		fmt.Println("there is no messages to send")
		return
	}

	defer ticker.Stop()
	js.wg.Add(1)
	defer js.wg.Done()

	for {
		select {
		case <-ticker.C:
			for i := 0; i < len(messages); i++ {
				timeNow := time.Now()
				messageID, err := client.SendMessage(messages[i])
				if err != nil {
					log.Println(err.Error())
					continue
				}
				log.Println("MessageID:", messageID)
				messages[i].MessageID = messageID
				messages[i].SentAt = timeNow
				if err := js.ms.UpdateMessage(messages[i]); err != nil {
					log.Println("message couldn't update", err)
					continue
				}

				data := map[string]interface{}{
					"messageID":   messageID,
					"sendingTime": messages[i].SentAt.Format("02-01-2006 15:04:05"),
				}

				cacheKey := fmt.Sprintf("message:%s", messageID)
				jsonData, _ := json.Marshal(data)

				if err := js.rClient.Set(cacheKey, jsonData); err != nil {
					log.Printf("message couldn't save to cache. messageID: %s\n", messageID)
				}
			}
		case <-js.stopChannel:
			fmt.Println("job status false")
			return
		}
	}
}
