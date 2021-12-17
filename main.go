package main

import (
	o "arrowhead/Orchestrator"
	forms "arrowhead/Orchestrator/forms"
	"fmt"
	//forms "arrowhead/Orchestrator/forms"
	//"fmt"
)

func main() {
	testServiceData()
}

func testServiceData() {

	//A Request comes from consumer
	r1 := new(forms.ServiceRequestForm)
	r1.RequestedService.ServiceDefinitionRequirement = "GE MIG TEMP"

	r2 := new(forms.ServiceRequestForm)
	r2.RequestedService.ServiceDefinitionRequirement = "GE MIG TID"

	//New serviceData is created for each request
	s1 := o.NewServiceData(r1)
	s2 := o.NewServiceData(r2)

	//Form the request a queryForm is constructed
	forms.ConstructServiceQueryForm(&s1.ServiceRequestForm)
	fmt.Println("S1: " + s1.ServiceRequestForm.RequestedService.ServiceDefinitionRequirement)

	forms.ConstructServiceQueryForm(&s2.ServiceRequestForm)
	fmt.Println("S2: " + s2.ServiceRequestForm.RequestedService.ServiceDefinitionRequirement)

}
