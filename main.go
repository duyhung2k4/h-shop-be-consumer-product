package main

import (
	_ "app/config"
	"app/service"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		service.NewProductService().PushProductToElasticsearch()
		wg.Done()
	}()

	go func() {
		service.NewProductService().UpdateProductInElasticsearch()
		wg.Done()
	}()

	go func() {
		service.NewProductService().DeleteProductInElasticsearch()
		wg.Done()
	}()

	go func() {
		service.NewProductService().UpCountWarehouse()
		wg.Done()
	}()
	wg.Wait()
}
