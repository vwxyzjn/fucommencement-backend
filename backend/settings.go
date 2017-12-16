package backend

import "github.com/algolia/algoliasearch-client-go/algoliasearch"

func GetSettings() {
	addReplica()
	settings, err := AlgoliaIndex.GetSettings()
	if err != nil {
		panic(err)
	}
	PrettyPrint(settings.ToMap())
}

func addReplica() {
	settings := algoliasearch.Map{
		"replicas": []string{"student_by_custom_sorting"},
	}
	if _, err := AlgoliaIndex.SetSettings(settings); err != nil {
		panic(err)
	}
}
