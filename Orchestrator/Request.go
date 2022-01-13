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

func sendServiceRequest(r *request) *forms.IntraCloudRule {
	/*
		serviceRegClient := NewClient(r.serviceRegAdrs)
		serviceQueryList := ExchangeJson(*serviceRegClient, r.data.Discover.ServiceQueryForm)
	*/

	//Ger error need fix
	//r.data.Discover.ServiceQueryList = &serviceQueryList.(*forms.ServiceQueryList)

	return forms.ConstructIntraCloudRule(&r.data.ServiceRequestForm, &r.data.Discover)
}

func sendAuthQuery(r *request) *forms.InterCloudResult {
	authClient := NewClient(r.authAdrs)
	intraCloudResult := ExchangeJson(*authClient, r.data.ServiceRequestForm)
	return intraCloudResult.(*forms.InterCloudResult)
}

//func Matchmaker(r *request) {}
