package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

func Export() {
	fmt.Println("Export being called")

	var hits []algoliasearch.Map

	it, err := AlgoliaSortedIndex.BrowseAll(algoliasearch.Map{"query": ""})
	if err != nil {
		panic(err)
	}

	res, err := it.Next()
	for err != algoliasearch.NoMoreHitsErr {
		if err != nil {
			panic(err)
		}

		hits = append(hits, res)
		res, err = it.Next()
	}

	json, err := json.Marshal(hits)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("your_filename.json", json, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
