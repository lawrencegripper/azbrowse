package expanders

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/lawrencegripper/azbrowse/pkg/armclient"
	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
)

func Test_Expanders(t *testing.T) {

	InitializeExpanders(nil, nil, nil, nil)
	expanders := getRegisteredExpanders()

	// Add special case of default expander
	expanders = append(expanders, &DefaultExpander{})

	for _, expander := range expanders {

		hasTests, testCases := expander.testCases()
		if !hasTests {
			t.Log("!!!! No tests defined for expander, there should be. Provider Name:" + expander.Name())
			continue
		}

		t.Log("Testing expander. Name:" + expander.Name())

		for _, tt := range *testCases {
			t.Run(expander.Name()+":"+tt.name, func(t *testing.T) {

				const testServer = "https://management.azure.com"
				testPath := tt.urlPath

				expectedJSONResponse := "Error No response content supplied to test framework"
				if tt.responseFile != "" {
					dat, err := ioutil.ReadFile(tt.responseFile)
					if err != nil {
						t.Error(err)
						t.FailNow()
					}
					expectedJSONResponse = string(dat)
				}

				defer gock.Off()
				if tt.configureGockFunc != nil {
					setupGock := *tt.configureGockFunc
					setupGock(t)
				} else {
					gock.New(testServer).
						Get(testPath).
						Reply(tt.statusCode).
						JSON(expectedJSONResponse)
				}

				httpClient := &http.Client{Transport: &http.Transport{}}
				gock.InterceptClient(httpClient)

				// Set the ARM client to use out test server
				client := armclient.NewClientFromConfig(httpClient, DummyTokenFunc(), 5000)
				// set dummy client
				expander.setClient(client)

				ctx := context.Background()

				result := expander.Expand(ctx, tt.nodeToExpand)

				tt.treeNodeCheckerFunc(t, result)

				// Verify that we don't have pending mocks
				st.Expect(t, gock.IsDone(), true)
			})
		}
	}
}
