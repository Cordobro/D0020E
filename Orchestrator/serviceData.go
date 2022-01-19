package Orchestrator

import (
	forms "arrowhead/Orchestrator/forms"
	"encoding/json"
	"fmt"
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

func NewServiceData(bytevalue []byte) *ServiceData {

	var request forms.ServiceRequestForm

	if err := json.Unmarshal(bytevalue, &request); err != nil {
		panic(err)
	}

	fmt.Println("NEW Request:")
	fmt.Print(request)
	fmt.Println("____________________")
	fmt.Println("")

	s := new(ServiceData)
	s.ServiceRequestForm = request

	forms.ConstructServiceQueryForm(&s.ServiceRequestForm, &s.Discover.ServiceQueryForm)
	return s
}
