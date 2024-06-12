package config

import (
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
		Username:  userElasticSearch,
		Password:  passwordElasticSearch,
		CACert:    cert,
	})

	if errEs != nil {
		return errEs
	}

	return nil
}
