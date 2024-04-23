package config

import (
	"app/model"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

func connectElastic() error {
	cert, errCert := os.ReadFile("cert/http_ca.crt")

	if errCert != nil {
		return errCert
	}

	var errEs error
	es, errEs = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{urlElastic},
		Username:  "elastic",
		Password:  "oouLuchH8ymSYzie_+Fs",
		CACert:    cert,
	})

	if errEs != nil {
		return errEs
	}

	initIndex(es)

	log.Println("Elastic done")
	return nil
}

func initIndex(es *elasticsearch.Client) {
	es.Indices.Create(string(model.PRODUCT_INDEX))
}
