package backend

import (
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

type AlgoliaInstance struct {
	AlgoliaClient algoliasearch.Client
	AlgoliaIndex  algoliasearch.Index
}

func (a *AlgoliaInstance) Initialize(AlgoliaAppID string, AlgoliaKey string, AlgoliaIndexName string) {
	a.AlgoliaClient = algoliasearch.NewClient(AlgoliaAppID, AlgoliaKey)
	a.AlgoliaIndex = a.AlgoliaClient.InitIndex(AlgoliaIndexName)
}
