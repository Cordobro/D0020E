package Orchestrator

import (
	forms "arrowhead/Orchestrator/forms"
	"encoding/json"
)

type request struct {
	data           ServiceData
	serviceRegAdrs string
	authAdrs       string
}

func newRequest(data ServiceData, serviceRegAdrs string, authAdrs string) *request {
	r := request{data: data,
		serviceRegAdrs: serviceRegAdrs,
		authAdrs:       authAdrs}
	return &r
}

func sendServiceRequest(r *request) {

	serviceRegClient := newClient(r.serviceRegAdrs)
	serviceQueryListJson := exchangeJson(serviceRegClient, r.data.Discover.ServiceQueryForm)
	var serviceQueryList forms.ServiceQueryList
	err := json.Unmarshal(serviceQueryListJson, &serviceQueryList)
	errorHandler(err)
	r.data.Discover.ServiceQueryList = serviceQueryList
}

/*
func sendAuthQuery(r *request) {
	authClient := NewClient(r.authAdrs)
	intraCloudResultJson := ExchangeJson(authClient, r.data.IntraCloudRule)
	var intraCloudResult forms.IntraCloudResult
	err := json.Unmarshal(intraCloudResultJson, &intraCloudResult)
	errorHandler(err)
	r.data.IntraCloudResult = intraCloudResult
}


func sendTokenQuery(r *request) {
	authClient := NewClient(r.authAdrs)
	tokenResultJson := ExchangeJson(authClient, r.data.TokenRule)
	var tokenResult forms.TokenResult
	err := json.Unmarshal(tokenResultJson, &tokenResult)
	errorHandler(err)
	r.data.TokenResult = tokenResult
}
*/
func matchmaker(r *request) {
	forms.ConstructOrchestrationResponse(&r.data.Discover.ServiceQueryList, &r.data.OrchestrationResponse)
}
