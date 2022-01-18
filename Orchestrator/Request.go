package Orchestrator

import forms "arrowhead/Orchestrator/forms"

type request struct {
	data           ServiceData
	serviceRegAdrs string
	authAdrs       string
}

func NewRequest(data ServiceData, serviceRegAdrs string, authAdrs string) *request {
	r := request{data: data,
		serviceRegAdrs: serviceRegAdrs,
		authAdrs:       authAdrs}
	return &r
}

func sendServiceRequest(r *request) {

	serviceRegClient := NewClient(r.serviceRegAdrs)
	serviceQueryList := ExchangeJson(serviceRegClient, r.data.Discover.ServiceQueryForm)

	r.data.Discover.ServiceQueryList = serviceQueryList.(forms.ServiceQueryList)
}

func sendAuthQuery(r *request) {
	authClient := NewClient(r.authAdrs)
	intraCloudResult := ExchangeJson(authClient, r.data.IntraCloudRule)
	r.data.IntraCloudResult = intraCloudResult.(forms.IntraCloudResult)
}

/*
func sendTokenQuery(r *request) {
	authClient := NewClient(r.authAdrs)
	tokenResult := ExchangeJson(*authClient, r.data.TokenRule)
	r.data.IntraCloudResult = tokenResult.(forms.IntraCloudResult)
}
*/

//func Matchmaker(r *request) {}
