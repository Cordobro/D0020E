package Orchestrator

import (
	forms "arrowhead/Orchestrator/forms"
	"encoding/json"
	"fmt"
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

func NewServiceData(request *forms.ServiceRequestForm) *ServiceData {

	s := new(ServiceData)
	s.ServiceRequestForm = *request

	return s
}

func ComposeMessage() {

}

func ParseMessage(byteValue []byte) {

	//var dat map[string]interface{}
	/*
		var requestForm forms.ServiceRequestForm

		var _ = json.Unmarshal(byteValue, &requestForm)

		fmt.Println("___INSIDE ParseMessage___")
		fmt.Println(requestForm)

	*/
	var discover forms.Discover

	var _ = json.Unmarshal(byteValue, &discover.ServiceQueryForm)

	fmt.Println("___INSIDE ParseMessage___")
	fmt.Println(discover.ServiceQueryForm.ServiceDefinitionRequirement)

}
