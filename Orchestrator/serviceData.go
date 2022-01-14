package Orchestrator

import (
	forms "arrowhead/Orchestrator/forms"
	//"encoding/json"
	//"fmt"
)

type ServiceData struct {
	ServiceRequestForm    forms.ServiceRequestForm
	OrchestrationResponse forms.OrchestrationResponse
	Discover              forms.Discover
	IntraCloudResult      forms.IntraCloudResult
	IntraCloudRule        forms.IntraCloudRule
	TokenRule             forms.TokenRule
	TokenResult           forms.TokenData
}

func NewServiceData(serviceRequest interface{}) *ServiceData {

	s := &ServiceData{ServiceRequestForm: serviceRequest.(forms.ServiceRequestForm)}
	forms.ConstructServiceQueryForm(&s.ServiceRequestForm, &s.Discover.ServiceQueryForm)
	return s
}
