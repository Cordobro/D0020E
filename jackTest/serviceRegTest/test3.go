package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/******************************************* STRUCTS ******************************************/

type ServiceQueryList struct {
	ServiceQueryData []ServiceQueryData `json:"serviceQueryData"`
	UnfilteredHits   int                `json:"unfilteredHits"`
}
type ServiceDefinition struct {
	ID                int    `json:"id"`
	ServiceDefinition string `json:"serviceDefinition"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
}
type Provider struct {
	ID                 int    `json:"id"`
	SystemName         string `json:"systemName"`
	Address            string `json:"address"`
	Port               int    `json:"port"`
	AuthenticationInfo string `json:"authenticationInfo"`
	CreatedAt          string `json:"createdAt"`
	UpdatedAt          string `json:"updatedAt"`
}
type Interfaces struct {
	ID            int    `json:"id"`
	InterfaceName string `json:"interfaceName"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}
type ServiceQueryData struct {
	ID                int               `json:"id"`
	ServiceDefinition ServiceDefinition `json:"serviceDefinition"`
	Provider          Provider          `json:"provider"`
	ServiceURI        string            `json:"serviceUri"`
	EndOfValidity     string            `json:"endOfValidity"`
	Secure            string            `json:"secure"`
	Metadata          []string          `json:"metadata"`
	Version           int               `json:"version"`
	Interfaces        []Interfaces      `json:"interfaces"`
	CreatedAt         string            `json:"createdAt"`
	UpdatedAt         string            `json:"updatedAt"`
}

type ServiceQueryForm struct {
	ServiceDefinitionRequirement string   `json:"serviceDefinitionRequirement"`
	InterfaceRequirements        []string `json:"interfaceRequirements"`
	SecurityRequirements         []string `json:"securityRequirements"`
	MetadataRequirements         []string `json:"metadataRequirements"`
	VersionRequirement           int      `json:"versionRequirement"`
	MaxVersionRequirement        int      `json:"maxVersionRequirement"`
	MinVersionRequirement        int      `json:"minVersionRequirement"`
	PingProviders                bool     `json:"pingProviders"`
}

/*************************************************************************************/

var count int

func newServiceQueryList(size int) *ServiceQueryList {

	sql := new(ServiceQueryList)

	for i := 0; i < size; i++ {
		sqd := new(ServiceQueryData)
		sql.ServiceQueryData = append(sql.ServiceQueryData, *sqd)
	}

	fmt.Println()
	fmt.Println("---NewServiceQueryList---")
	fmt.Println()
	fmt.Println("ServiceQueryData Size: ", len(sql.ServiceQueryData))
	fmt.Println()

	return sql
}

func listen(rw http.ResponseWriter, req *http.Request) {

	fmt.Println("Received a request")

	count++

	//Unmarshal the request

	var sqf = parseMSG(req)

	fmt.Println("SQF:\n ", sqf)

	//create response

	var response = constructReturnMessage(sqf)

	//Metadata

	var metaList []string
	metaList = append(metaList, "META")
	response.ServiceQueryData[0].Metadata = metaList
	response.ServiceQueryData[1].Metadata = metaList

	var metaList2 []string
	metaList2 = append(metaList2, "META", "0", "%3")
	response.ServiceQueryData[2].Metadata = metaList2

	metaList = append(metaList, "0")
	response.ServiceQueryData[3].Metadata = metaList
	response.ServiceQueryData[4].Metadata = metaList

	//Respond

	fmt.Println("Sending ServiceQueryList...")

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(response)

}

func parseMSG(req *http.Request) ServiceQueryForm {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Body:\n ", string(body))

	var request ServiceQueryForm

	if err := json.Unmarshal(body, &request); err != nil {
		panic(err)
	}

	var sqf ServiceQueryForm
	sqf = request

	return sqf
}

func constructReturnMessage(sqf ServiceQueryForm) *ServiceQueryList {

	response := newServiceQueryList(5)

	response.ServiceQueryData[0].Provider.SystemName = "THIS IS A TEST"
	response.ServiceQueryData[0].Provider.ID = 404
	response.ServiceQueryData[0].ServiceDefinition.ServiceDefinition = "THIS WORKS"

	response.ServiceQueryData[0].Version = count

	return response
}

func main() {

	count = 0

	fmt.Println("Starting up...")

	http.HandleFunc("/serviceRegistry", listen)
	http.ListenAndServe(":8000", nil)
}
