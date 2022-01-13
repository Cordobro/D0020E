package main

import (
	 o "arrowhead/Orchestrator"
	forms "arrowhead/Orchestrator/forms"
	"encoding/json"
	"fmt" 
	//forms "arrowhead/Orchestrator/forms"
	//"fmt"
)

func main() {
	testServiceData()
}

func testServiceData() {

	//A Request comes from consumer
	request1 := new(forms.ServiceRequestForm)
	request1.RequestedService.ServiceDefinitionRequirement = "GE MIG TEMP"
	msg1, _ := json.Marshal(request1)

	//New serviceData is created for each request
	s1 := o.NewServiceData(msg1)

	//Form the request a queryForm is constructed
	s1.Discover.ServiceQueryForm = forms.ConstructServiceQueryForm(&s1.ServiceRequestForm)
	fmt.Println("S1: ServiceQueryForm")
	fmt.Println(s1.Discover.ServiceQueryForm)
	fmt.Println("")

	//marshal the form
	o.ComposeServiceQueryForm(s1.Discover.ServiceQueryForm)

	//pretend to send query

	//get result back

	sql := new(forms.ServiceQueryList)
	sql.UnfilteredHits = 10

	fmt.Println("sql")
	fmt.Println(sql)
	fmt.Println("")

	s1.Discover.ServiceQueryList = sql

	fmt.Println("S1: ServiceQueryList")
	fmt.Println(s1.Discover.ServiceQueryList)
	fmt.Println("")

	//construct intraCloudRule

	s1.IntraCloudRule = forms.ConstructIntraCloudRule(&s1.ServiceRequestForm, *s1.Discover.ServiceQueryList)

	//get result

	icres := new(forms.IntraCloudResult)
	s1.IntraCloudResult = icres

	//construct tokenRule

	//get result

	//matchmake

	//construct orchestrationResponse

}
