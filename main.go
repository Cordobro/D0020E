package main

import (
<<<<<<< HEAD
	o "arrowhead/Orchestrator"
	forms "arrowhead/Orchestrator/forms"
	"encoding/json"
	"fmt"
=======
	 o "arrowhead/Orchestrator"
<<<<<<< HEAD
	//forms "arrowhead/Orchestrator/forms"
	//"encoding/json"
	//"fmt" 
=======
	forms "arrowhead/Orchestrator/forms"
	"encoding/json"
	"fmt" 
>>>>>>> a91626ade7e5d22b94e57f176a096036ac951b30
>>>>>>> fb29c1863f8ec0251b6bfaf75eff812182345dea
	//forms "arrowhead/Orchestrator/forms"
	//"fmt"
)

func main() {
	//testServiceData()
	s := o.NewServer("123")
	o.Listen(s)
}
/*
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

	var sql = forms.TestNewServiceQueryList(5)

	fmt.Println("sql EMPTY")
	fmt.Println(sql)
	fmt.Println("")

	sql.ServiceQueryData[0].Provider.SystemName = "Jacks Dator"
	sql.ServiceQueryData[0].Provider.Address = "192.168.1.111"

	fmt.Println("sql:")
	fmt.Println(sql)
	fmt.Println("")

	s1.Discover.ServiceQueryList = sql

	fmt.Println("s1.Discover.ServiceQueryList")
	fmt.Println(s1.Discover.ServiceQueryList)
	fmt.Println("")

	//construct intraCloudRule

	s1.IntraCloudRule = forms.ConstructIntraCloudRule(&s1.ServiceRequestForm, &s1.Discover)

	//get result

	icres := new(forms.IntraCloudResult)
	s1.IntraCloudResult = icres

	//construct tokenRule

	//get result

	//matchmake

	//construct orchestrationResponse

<<<<<<< HEAD
}

func NewServiceQueryList(i int) {
	panic("unimplemented")
=======
>>>>>>> a91626ade7e5d22b94e57f176a096036ac951b30
}
*/