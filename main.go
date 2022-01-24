package main

import (
	o "arrowhead/Orchestrator"
	"arrowhead/Orchestrator/forms"
	"encoding/json"
	"fmt"
	//"time"
	//"arrowhead/Orchestrator/forms"
	//"fmt"
)

func main() {

	CONN_PORT := o.ReadFile("CONN_PORT")
	LISTEN_HANDLE := o.ReadFile("LISTEN_HANDLE")
	o.InitOrchestrator()

	/* http.HandleFunc(LISTEN_HANDLE, o.Listen)
	http.ListenAndServe(CONN_PORT, nil) */

	o.SetupServer(CONN_PORT, LISTEN_HANDLE)

	//testServiceData()

	/*
			http.HandleFunc("/ServiceRegistry", serviceRegistry)
		    go http.ListenAndServe(":8001", nil)

			var serviceRequestForm forms.ServiceRequestForm
			serviceRequestForm.RequestedService.MetadataRequirements.AdditionalProp1 = "Icecream"

			requestor("https://localhost:8000/", serviceRequestForm) */

}

/* func serviceRegistry(rw http.ResponseWriter, req *http.Request) {
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        panic(err)
    }



    var t interface{}
    err = json.Unmarshal(body, &t)
	if err != nil {
        panic(err)
    }

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(t)


}

//Client requestor
func requestor(url string, struc interface{}) {
	jsonStr, err := json.Marshal(struc)
	errorHandler(err)
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	errorHandler(err);
	defer resp.Body.Close()


	jsonBody, err := ioutil.ReadAll(resp.Body)
	errorHandler(err)

	fmt.Println(jsonBody)
}
*/

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
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

func NewServiceQueryList(i int) {
	panic("unimplemented")
}
