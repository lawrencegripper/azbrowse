package expanders

import (
	"context"
	"encoding/json"
	"net/url"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/interfaces"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// ActivityLogExpander expands activity logs under an RG
type ActivityLogExpander struct {
	ExpanderBase
	client *armclient.Client
}

func (e *ActivityLogExpander) setClient(c *armclient.Client) {
	e.client = c
}

// Name returns the name of the expander
func (e *ActivityLogExpander) Name() string {
	return "ActivityLogExpander"
}

// DoesExpand checks if this is an RG
func (e *ActivityLogExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	if currentItem.ItemType == activityLogType {
		return true, nil
	}
	return false, nil
}

// Expand returns Resources in the RG
func (e *ActivityLogExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	method := "GET"
	data, err := e.client.DoRequest(ctx, method, currentItem.ExpandURL)
	if err != nil {
		return ExpanderResult{
			Err:               err,
			Response:          ExpanderResponse{Response: string(data), ResponseType: interfaces.ResponseJSON},
			SourceDescription: "Activity Log",
			IsPrimaryResponse: true,
		}
	}
	newItems := []*TreeNode{}

	var activityLogs ActivityLogResource
	err = json.Unmarshal([]byte(data), &activityLogs)
	if err != nil {
		panic(err)
	}

	value, err := fastJSONParser.Parse(data)
	if err != nil {
		panic(err)
	}

	for i, log := range activityLogs.Value {
		if len(value.GetArray("value"))-1 < i {
			break
		}
		// Update the existing state as we have more up-to-date info
		objectJSON := string(value.GetArray("value")[i].MarshalTo([]byte("")))

		newItems = append(newItems, &TreeNode{
			Name:            log.OperationName.Value,
			Display:         log.OperationName.LocalizedValue + "\n   " + style.Subtle("At:  "+log.EventTimestamp.String()) + "\n   " + style.Subtle("ResourceType: "+log.ResourceType.Value) + "\n   " + style.Subtle("Status: "+log.Status.Value+""),
			ID:              log.ID,
			Parentid:        currentItem.ID,
			ExpandURL:       ExpandURLNotSupported,
			ItemType:        subActivityLogType,
			SubscriptionID:  currentItem.SubscriptionID,
			StatusIndicator: DrawStatus(log.Status.Value),
			Metadata: map[string]string{
				"jsonItem": objectJSON,
			},
		})
	}

	return ExpanderResult{
		Err:               err,
		Response:          ExpanderResponse{Response: string(data), ResponseType: interfaces.ResponseJSON},
		SourceDescription: "Deployments request",
		Nodes:             newItems,
		IsPrimaryResponse: true,
	}
}

// GetActivityLogExpandURL gets the urls which should be used to get activity logs
func GetActivityLogExpandURL(subscriptionID, resourceName string) string {
	queryString := `eventTimestamp ge '` + time.Now().AddDate(0, 0, -30).Format("2006-01-02T15:04:05Z07:00") + `' and eventTimestamp le '` +
		time.Now().Format("2006-01-02T15:04:05Z07:00") + `' and eventChannels eq 'Admin, Operation' and resourceGroupName eq '` +
		resourceName + `' and levels eq 'Critical,Error,Warning,Informational' | orderby eventTimestamp desc`
	return `/subscriptions/` + subscriptionID + `/providers/microsoft.insights/eventtypes/management/values?api-version=2017-03-01-preview&$filter=` +
		url.QueryEscape(queryString)
}

// ActivityLogResource is returned when requesting activity logs for an RG
type ActivityLogResource struct {
	Value []struct {
		Authorization struct {
			Action string `json:"action"`
			Scope  string `json:"scope"`
		} `json:"authorization"`
		Caller   string `json:"caller"`
		Channels string `json:"channels"`
		Claims   struct {
			Aud                                                       string `json:"aud"`
			Iss                                                       string `json:"iss"`
			Iat                                                       string `json:"iat"`
			Nbf                                                       string `json:"nbf"`
			Exp                                                       string `json:"exp"`
			Aio                                                       string `json:"aio"`
			Appid                                                     string `json:"appid"`
			Appidacr                                                  string `json:"appidacr"`
			HTTPSchemasMicrosoftComIdentityClaimsIdentityprovider     string `json:"http://schemas.microsoft.com/identity/claims/identityprovider"`
			HTTPSchemasMicrosoftComIdentityClaimsObjectidentifier     string `json:"http://schemas.microsoft.com/identity/claims/objectidentifier"`
			HTTPSchemasXmlsoapOrgWs200505IdentityClaimsNameidentifier string `json:"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/nameidentifier"`
			HTTPSchemasMicrosoftComIdentityClaimsTenantid             string `json:"http://schemas.microsoft.com/identity/claims/tenantid"`
			Uti                                                       string `json:"uti"`
			Ver                                                       string `json:"ver"`
		} `json:"claims"`
		CorrelationID string `json:"correlationId"`
		Description   string `json:"description"`
		EventDataID   string `json:"eventDataId"`
		EventName     struct {
			Value          string `json:"value"`
			LocalizedValue string `json:"localizedValue"`
		} `json:"eventName"`
		Category struct {
			Value          string `json:"value"`
			LocalizedValue string `json:"localizedValue"`
		} `json:"category"`
		ID                   string `json:"id"`
		Level                string `json:"level"`
		ResourceGroupName    string `json:"resourceGroupName"`
		ResourceProviderName struct {
			Value          string `json:"value"`
			LocalizedValue string `json:"localizedValue"`
		} `json:"resourceProviderName"`
		ResourceID   string `json:"resourceId"`
		ResourceType struct {
			Value          string `json:"value"`
			LocalizedValue string `json:"localizedValue"`
		} `json:"resourceType"`
		OperationID   string `json:"operationId"`
		OperationName struct {
			Value          string `json:"value"`
			LocalizedValue string `json:"localizedValue"`
		} `json:"operationName"`
		Properties struct {
			IsComplianceCheck string `json:"isComplianceCheck"`
			ResourceLocation  string `json:"resourceLocation"`
			Ancestors         string `json:"ancestors"`
			Policies          string `json:"policies"`
		} `json:"properties,omitempty"`
		Status struct {
			Value          string `json:"value"`
			LocalizedValue string `json:"localizedValue"`
		} `json:"status"`
		SubStatus struct {
			Value          string `json:"value"`
			LocalizedValue string `json:"localizedValue"`
		} `json:"subStatus"`
		EventTimestamp      time.Time `json:"eventTimestamp"`
		SubmissionTimestamp time.Time `json:"submissionTimestamp"`
		SubscriptionID      string    `json:"subscriptionId"`
		TenantID            string    `json:"tenantId"`
		HTTPRequest         struct {
			ClientRequestID string `json:"clientRequestId"`
			ClientIPAddress string `json:"clientIpAddress"`
			Method          string `json:"method"`
		} `json:"httpRequest,omitempty"`
	} `json:"value"`
}
