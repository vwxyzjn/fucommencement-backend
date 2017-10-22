package backend

import (
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Client, Index
var (
	AlgoliaClient = algoliasearch.NewClient("TH20RENZY1", "f6fc0cc56e0b7af1fc5e5d71ff207bf6")
	AlgoliaIndex  = AlgoliaClient.InitIndex("student")
)
