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
	IntraCloudResult      *forms.IntraCloudResult
	IntraCloudRule        *forms.IntraCloudRule
	TokenRule             *forms.TokenRule
	TokenResult           forms.TokenData
}

func NewServiceData(byteValue []byte) *ServiceData {

	var request forms.ServiceRequestForm
	var _ = json.Unmarshal(byteValue, &request)

	fmt.Println("---Inside NewServiceData---")
	fmt.Println("")
	fmt.Println(request)
	fmt.Println("")

	s := new(ServiceData)
	s.ServiceRequestForm = request

	return s
}

func ComposeServiceQueryForm(sqf *forms.ServiceQueryForm) []byte {

	msg, _ := json.Marshal(sqf)

	fmt.Println("---Inside ComposeServiceQueryForm---")
	fmt.Println("")
	fmt.Println(string(msg))
	fmt.Println("")

	return msg
}

func TestParseDiscover(byteValue []byte) {

	var discover forms.Discover

	var _ = json.Unmarshal(byteValue, &discover.ServiceQueryForm)

	fmt.Println("___INSIDE ParseMessage___")
	fmt.Println(discover.ServiceQueryForm.ServiceDefinitionRequirement)

}
