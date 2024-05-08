package config

import (
	"app/model"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

func connectElastic() error {
	cert, errCert := os.ReadFile("cert/http_ca.crt")

	if errCert != nil {
		return errCert
	}

	var errEs error
	elasticClient, errEs = elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{urlElastic},
		Username:  "elastic",
		Password:  "oouLuchH8ymSYzie_+Fs",
		CACert:    cert,
	})

	if errEs != nil {
		return errEs
	}

	initIndex(elasticClient)

	return nil
}

func initIndex(elasticClient *elasticsearch.TypedClient) {
	elasticClient.Indices.Create(string(model.PRODUCT_INDEX))
}
