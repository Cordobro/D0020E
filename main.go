package main

import (
	o "arrowhead/Orchestrator"
<<<<<<< HEAD
	forms "arrowhead/Orchestrator/forms"
	"encoding/json"
	"fmt"
=======
	/*forms "arrowhead/Orchestrator/forms"
	"encoding/json"
	"fmt"
	*/
>>>>>>> 27f1ebff15046ed2dff07de43fa20722ca49ba37
)

func main() {
	testServiceData()
	//s := o.NewServer("123")
	//o.Listen(s)
}
/*
func testServiceData() {

	//A Request comes from consumer
	request1 := new(forms.ServiceRequestForm)
	request1.RequestedService.ServiceDefinitionRequirement = "GE MIG TEMP"
	msg1, _ := json.Marshal(request1)

	var res1 interface{}
	if err := json.Unmarshal(msg1, &res1); err != nil {
		panic(err)
	}

	//New serviceData is created for each request
	s1 := o.NewServiceData(res1)

	//Form the request a queryForm is constructed
	forms.ConstructServiceQueryForm(&s1.ServiceRequestForm, &s1.Discover.ServiceQueryForm)
	fmt.Println("S1: ServiceQueryForm")
	fmt.Println(s1.Discover.ServiceQueryForm)
	fmt.Println("")

	//marshal the form

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

}
*/