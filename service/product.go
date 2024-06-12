package service

import (
	"app/config"
	"app/grpc/proto"
	"app/model"
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/rabbitmq/amqp091-go"
)

type productService struct {
	rabbitConnection           *amqp091.Connection
	elasticClient              *elasticsearch.TypedClient
	grpcWarehouseService       proto.WarehouseServiceClient
	grpcTypeInWarehouseService proto.TypeInWarehouseServiceClient
}

type ProductService interface {
	PushProductToElasticsearch() error
	UpdateProductInElasticsearch() error
	DeleteProductInElasticsearch() error

	UpCountWarehouse() error
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

			dataConvert := map[string]types.Property{}
			for key, value := range dataMap {
				if key == "_id" {
					continue
				}
				dataConvert[key] = value
			}

			_, err := s.elasticClient.
				Create(string(model.PRODUCT_INDEX), dataMap["_id"].(string)).
				Request(dataConvert).
				Do(context.Background())

			if err != nil {
				log.Println("Error: ", err)
			}
			wg.Done()
		}(msg.Body)
	}
	wg.Wait()

	return nil
}

func (s *productService) UpdateProductInElasticsearch() error {
	ch, errCh := s.rabbitConnection.Channel()

	if errCh != nil {
		return errCh
	}

	q, errQ := ch.QueueDeclare(
		string(model.UPDATE_PRODUCT_TO_ELASTIC), // name
		true,                                    // durable
		false,                                   // delete when unused
		false,                                   // exclusive
		false,                                   // no-wait
		nil,                                     // arguments
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
			var product map[string]interface{}
			if err := json.Unmarshal(data, &product); err != nil {
				wg.Done()
			}

			dataConvert := map[string]interface{}{}
			for key, value := range product {
				if key == "_id" {
					continue
				}
				dataConvert[key] = value
			}

			dataConvertBytes, errConvertBytes := json.Marshal(dataConvert)
			if errConvertBytes != nil {
				log.Println(errConvertBytes)
				wg.Done()
			}

			request := update.Request{
				Doc: dataConvertBytes,
			}
			_, err := s.elasticClient.
				Update(string(model.PRODUCT_INDEX), product["_id"].(string)).
				Request(&request).
				Do(context.Background())

			if err != nil {
				log.Println("error update: ", err)
				_, errCreate := s.elasticClient.
					Create(string(model.PRODUCT_INDEX), product["_id"].(string)).
					Request(dataConvert).
					Do(context.Background())

				if errCreate != nil {
					log.Println("error create: ", errCreate)
				}
			}

			wg.Done()
		}(msg.Body)
	}
	return nil
}

func (s *productService) DeleteProductInElasticsearch() error {
	ch, errCh := s.rabbitConnection.Channel()

	if errCh != nil {
		return errCh
	}

	q, errQ := ch.QueueDeclare(
		string(model.DELETE_PRODUCT_TO_ELASTIC), // name
		true,                                    // durable
		false,                                   // delete when unused
		false,                                   // exclusive
		false,                                   // no-wait
		nil,                                     // arguments
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
			var product map[string]interface{}
			if err := json.Unmarshal(data, &product); err != nil {
				wg.Done()
			}

			_, err := s.elasticClient.
				Delete(string(model.PRODUCT_INDEX), product["productId"].(string)).
				Do(context.Background())

			if err != nil {
				log.Println(err)
			}

			wg.Done()
		}(msg.Body)
	}
	return nil
}

func (s *productService) UpCountWarehouse() error {
	ch, errCh := s.rabbitConnection.Channel()

	if errCh != nil {
		return errCh
	}

	q, errQ := ch.QueueDeclare(
		string(model.UP_COUNT_WAREHOUSE), // name
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
			var checkCountPayload []model.CheckcountPayload
			if err := json.Unmarshal(data, &checkCountPayload); err != nil {
				wg.Done()
			}

			for _, item := range checkCountPayload {
				if item.TypeInWarehouseId != nil {
					s.grpcTypeInWarehouseService.UpCount(context.Background(), &proto.UpCountTypeInWarehouseReq{
						Id:     uint64(*item.TypeInWarehouseId),
						Amount: 1,
					})
				} else {
					s.grpcWarehouseService.UpCount(context.Background(), &proto.UpCountWarehouseReq{
						Id:     uint64(item.WarehouseId),
						Amount: 1,
					})
				}
			}
		}(msg.Body)
	}
	return nil
}

func NewProductService() ProductService {
	return &productService{
		rabbitConnection:           config.GetRabbitConnection(),
		elasticClient:              config.GetElasticClient(),
		grpcWarehouseService:       proto.NewWarehouseServiceClient(config.GetConnWarehouseGRPC()),
		grpcTypeInWarehouseService: proto.NewTypeInWarehouseServiceClient(config.GetConnWarehouseGRPC()),
	}
}
