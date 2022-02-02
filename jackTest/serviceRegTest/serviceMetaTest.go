package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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
var listSize int
var database []ServiceQueryData

func newServiceQueryList(size int) *ServiceQueryList {

	sql := new(ServiceQueryList)

	for i := 0; i < size; i++ {
		sqd := new(ServiceQueryData)
		sql.ServiceQueryData = append(sql.ServiceQueryData, *sqd)
	}

	return sql
}

func listen(rw http.ResponseWriter, req *http.Request) {

	fmt.Println("Received a request")

	//Unmarshal the request

	var sqf = parseMSG(req)

	fmt.Println("Received message:\n ", sqf)

	//create response

	var response = constructReturnMessage(sqf)
	response.ServiceQueryData = database

	for i := 0; i < len(response.ServiceQueryData); i++ {
		countString := strconv.Itoa(count)
		response.ServiceQueryData[i].Provider.SystemName = countString
		count++
	}

	//Respond

	fmt.Println("Sending response...\n", response)

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(response)

	fmt.Println()
}

func parseMSG(req *http.Request) ServiceQueryForm {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var request ServiceQueryForm

	if err := json.Unmarshal(body, &request); err != nil {
		panic(err)
	}

	var sqf ServiceQueryForm
	sqf = request

	return sqf
}

func constructReturnMessage(sqf ServiceQueryForm) *ServiceQueryList {

	response := newServiceQueryList(listSize)

	return response
}

func setupDB(size int) {

	fmt.Println("Setting up Database...")

	for i := 0; i < size; i++ {
		sqd := new(ServiceQueryData)
		database = append(database, *sqd)
	}

}

func fillDB() {
	for i := 0; i < len(database); i++ {
		database[i].ServiceDefinition.ServiceDefinition = "Temp"
		database[i].ServiceDefinition.ID = 0 + i
	}
}

func printDB() {
	for i := 0; i < len(database); i++ {
		fmt.Println("")
		fmt.Println("Index: ", i)
		fmt.Println("ServiceDefinition: ", database[i].ServiceDefinition.ServiceDefinition)
		fmt.Println("ID: ", database[i].ServiceDefinition.ID)
		fmt.Print("Metadata: ", database[i].Metadata)
		fmt.Println("")

	}
}

func addMetaData() {

	database[0].Metadata = append(database[0].Metadata, "Fahrenheit")
	database[0].Metadata = append(database[0].Metadata, "Sundsvall")

	database[1].Metadata = append(database[1].Metadata, "Celsius")
	database[1].Metadata = append(database[1].Metadata, "Luleå")
	database[1].Metadata = append(database[1].Metadata, "Today")

	database[2].Metadata = append(database[2].Metadata, "Celsius")
	database[2].Metadata = append(database[2].Metadata, "Today")
	database[2].Metadata = append(database[2].Metadata, "Stockholm")

}

func main() {

	count = 0
	listSize = 3

	setupDB(3)
	fillDB()
	addMetaData()
	printDB()

	fmt.Println("")
	fmt.Println("Waiting for a request...")

	http.HandleFunc("/serviceRegistry", listen)
	http.ListenAndServe(":8000", nil)
}
