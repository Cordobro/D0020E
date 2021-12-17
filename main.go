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

	//New serviceData is created for each request
	s1 := o.NewServiceData(r1)

	//Form the request a queryForm is constructed
	s1.Discover.ServiceQueryForm = forms.ConstructServiceQueryForm(&s1.ServiceRequestForm)
	fmt.Println("S1")
	fmt.Println(s1.Discover.ServiceQueryForm)
	fmt.Println("")

	//marshal the form
	o.ComposeServiceQueryForm(s1.Discover.ServiceQueryForm)

	//pretend to send query

	//get result back

	//construct intraCloudRule

	//get result

	//construct tokenRule

	//get result

	//matchmake

	//construct orchestrationResponse

}
