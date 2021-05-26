package config

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/lawrencegripper/azbrowse/internal/pkg/storage"
)

// GraphQuery represents a custom query for Azure Resource Graph which should be presented
// in a tenant
type GraphQuery struct {
	Name  string
	Query string
}

// GetCustomResourceGraphQueries retreives custom resource queries from the ~/.azbrowse-queries.config file
// The file should be
func GetCustomResourceGraphQueries() ([]GraphQuery, error) {
	queryDirLocation := storage.GetStorageDir()
	// Create the dir if it doesn't exist
	err := os.MkdirAll(queryDirLocation, 0644)
	if err != nil {
		return []GraphQuery{}, nil
	}

	_, err = os.Stat(queryDirLocation)
	if err != nil {
		// don't error on no config file
		return []GraphQuery{}, nil
	}

	files, err := ioutil.ReadDir(queryDirLocation)
	if err != nil {
		return []GraphQuery{}, err
	}
	queries := make([]GraphQuery, 0, len(files))
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".kql") {
			continue
		}
		content, err := os.ReadFile(path.Join(queryDirLocation, file.Name()))
		if err != nil {
			return []GraphQuery{}, err
		}
		queries = append(queries, GraphQuery{Name: file.Name(), Query: string(content)})
	}

	return queries, nil
}
