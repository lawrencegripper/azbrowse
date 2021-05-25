package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
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
	queryFileLocation := "/root/.azbrowse-query.config"
	user, err := user.Current()
	if err == nil {
		queryFileLocation = user.HomeDir + "/.azbrowse-query.config"
	}

	_, err = os.Stat(queryFileLocation)
	if err != nil {
		// don't error on no config file
		return []GraphQuery{}, nil
	}

	configFile, err := os.Open(queryFileLocation)
	if err != nil {
		return []GraphQuery{}, err
	}
	defer configFile.Close() //nolint: errcheck
	bytes, _ := ioutil.ReadAll(configFile)

	rawQueries := strings.Split(string(bytes), "---\n")
	graphQueries := make([]GraphQuery, 0, len(rawQueries))
	for _, rawQuery := range rawQueries {
		if len(rawQuery) == 0 {
			// Skip any empty sections
			continue
		}
		lines := strings.Split(rawQuery, "\n")
		if len(lines) < 2 {
			return nil, fmt.Errorf("query file in incorrect format, section with less than 2 lines, see docs. Found %q", rawQuery)
		}
		// First line must be query name prefixed with a #
		if !strings.HasPrefix(lines[0], "#") {
			return nil, fmt.Errorf("query file in incorrect format, must contain #namehere in each query section. Found %q", lines[0])
		}
		graphQueries = append(graphQueries, GraphQuery{
			Name:  lines[0],
			Query: strings.Join(lines[1:], ""), // Flatten the query to a single line for ease as kusto doesn't care
		})
	}

	return graphQueries, nil
}
