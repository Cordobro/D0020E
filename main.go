package main

import (
	o "arrowhead/Orchestrator"
	forms "arrowhead/Orchestrator/forms"
	"encoding/json"
	"fmt"
)

func main() {
	testServiceData()
	//s := o.NewServer("123")
	//o.Listen(s)
}

func testServiceData() {

	//A Request comes from consumer
	var request1 forms.ServiceRequestForm

	request1.RequestedService.ServiceDefinitionRequirement = "GE MIG TEMP"
	request1.RequesterSystem.Address = "199.999.999.99"

	msg1, _ := json.Marshal(request1)

	//New serviceData is created for each request
	s1 := o.NewServiceData(msg1)

	//Form the request a queryForm is constructed
	forms.ConstructServiceQueryForm(&s1.ServiceRequestForm, &s1.Discover.ServiceQueryForm)
	fmt.Println("S1: ServiceQueryForm")
	fmt.Println(s1.Discover.ServiceQueryForm)
	fmt.Println("")

	//marshal the form

	//pretend to send query

	//get result back

	var sql = forms.TestNewServiceQueryList(1)

	fmt.Println("sql EMPTY")
	fmt.Println(sql)
	fmt.Println("")

	sql.ServiceQueryData[0].Provider.SystemName = "Jacks Dator"
	sql.ServiceQueryData[0].Provider.Address = "192.168.1.111"

	fmt.Println("sql:")
	fmt.Println(sql)
	fmt.Println("")

	s1.Discover.ServiceQueryList = *sql

	fmt.Println("s1.Discover.ServiceQueryList")
	fmt.Println(s1.Discover.ServiceQueryList)
	fmt.Println("")

	//construct intraCloudRule

	forms.ConstructIntraCloudRule(&s1.ServiceRequestForm, &s1.Discover, &s1.IntraCloudRule)

	//get result

	//construct tokenRule

	//get result

	//matchmake

	//construct orchestrationResponse
	forms.ConstructOrchestrationResponse(&s1.Discover.ServiceQueryList, &s1.OrchestrationResponse)
	fmt.Println("s1.OrchestrationResponse")
	fmt.Println(s1.OrchestrationResponse)
	fmt.Println("")

}
