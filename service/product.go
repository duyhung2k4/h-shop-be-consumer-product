package service

import (
	"app/config"
	"app/model"
	"bytes"
	"context"
	"encoding/json"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/rabbitmq/amqp091-go"
)

type productService struct {
	rabbitConnection *amqp091.Connection
	es               *elasticsearch.Client
}

type ProductService interface {
	PushProductToElasticsearch() error
}

func (s *productService) PushProductToElasticsearch() error {
	ch, errCh := s.rabbitConnection.Channel()

	if errCh != nil {
		return errCh
	}

	q, errQ := ch.QueueDeclare(
		string(model.PRODUCT_TO_ELASTIC), // name
		true,                             // durable
		false,                            // delete when unused
		false,                            // exclusive
		false,                            // no-wait
		nil,                              // arguments
	)

	if errQ != nil {
		return errQ
	}

	msgs, errMsgs := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if errMsgs != nil {
		return errMsgs
	}

	var wg sync.WaitGroup
	for msg := range msgs {
		wg.Add(1)
		go func(data []byte) {
			dataMap := map[string]interface{}{}

			if err := json.Unmarshal(data, &dataMap); err != nil {
				wg.Done()
			}

			dataConvert := map[string]interface{}{}
			for key, value := range dataMap {
				if key != "_id" {
					dataConvert[key] = value
				}
			}

			dataJSON, errDataJSON := json.Marshal(dataConvert)
			if errDataJSON != nil {
				wg.Done()
			}

			req := esapi.IndexRequest{
				Index:      string(model.PRODUCT_INDEX), // Replace with your index name
				DocumentID: dataMap["_id"].(string),
				Body:       bytes.NewReader(dataJSON),
				Refresh:    "true",
			}

			_, err := req.Do(context.Background(), s.es)
			if err != nil {
				wg.Done()
			}

			wg.Done()
		}(msg.Body)
	}
	wg.Wait()

	return nil
}

func NewProductService() ProductService {
	return &productService{
		rabbitConnection: config.GetRabbitConnection(),
		es:               config.GetEs(),
	}
}
