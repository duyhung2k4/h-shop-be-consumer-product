package main

import (
	_ "app/config"
	"app/service"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		service.NewProductService().PushProductToElasticsearch()
		wg.Done()
	}()
	wg.Wait()
}
