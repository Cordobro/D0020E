package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ServiceRequestForm struct {
	RequesterSystem struct {
		SystemName         string `json:"systemName"`
		Address            string `json:"address"`
		Port               int    `json:"port"`
		AuthenticationInfo string `json:"authenticationInfo"`
	} `json:"requesterSystem"`
	RequestedService struct {
		ServiceDefinitionRequirement string   `json:"serviceDefinitionRequirement"`
		InterfaceRequirements        []string `json:"interfaceRequirements"`
		SecurityRequirements         []string `json:"securityRequirements"`
		MetadataRequirements         []string `json:"metadataRequirements"`
		VersionRequirement           int      `json:"versionRequirement"`
		MaxVersionRequirement        int      `json:"maxVersionRequirement"`
		MinVersionRequirement        int      `json:"minVersionRequirement"`
	} `json:"requestedService"`
	PreferredProviders []struct {
		ProviderCloud struct {
			Operator string `json:"operator"`
			Name     string `json:"name"`
		} `json:"providerCloud"`
		ProviderSystem struct {
			SystemName string `json:"systemName"`
			Address    string `json:"address"`
			Port       int    `json:"port"`
		} `json:"providerSystem"`
	} `json:"preferredProviders"`
	OrchestrationFlags struct {
		AdditionalProp1 bool `json:"additionalProp1"`
		AdditionalProp2 bool `json:"additionalProp2"`
		AdditionalProp3 bool `json:"additionalProp3"`
	} `json:"orchestrationFlags"`
}

type OrchestrationResponse struct {
	Response []Response `json:"response"`
}
type provider struct {
	ID                 int    `json:"id"`
	SystemName         string `json:"systemName"`
	Address            string `json:"address"`
	Port               int    `json:"port"`
	AuthenticationInfo string `json:"authenticationInfo"`
	CreatedAt          string `json:"createdAt"`
	UpdatedAt          string `json:"updatedAt"`
}
type service struct {
	ID                int    `json:"id"`
	ServiceDefinition string `json:"serviceDefinition"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
}

type interfaces struct {
	ID            int    `json:"id"`
	CreatedAt     string `json:"createdAt"`
	InterfaceName string `json:"interfaceName"`
	UpdatedAt     string `json:"updatedAt"`
}
type authorizationTokens struct {
	InterfaceName1 string `json:"interfaceName1"`
	InterfaceName2 string `json:"interfaceName2"`
}
type Response struct {
	Provider            provider            `json:"provider"`
	Service             service             `json:"service"`
	ServiceURI          string              `json:"serviceUri"`
	Secure              string              `json:"secure"`
	Metadata            []string            `json:"metadata"`
	Interfaces          []interfaces        `json:"interfaces"`
	Version             int                 `json:"version"`
	AuthorizationTokens authorizationTokens `json:"authorizationTokens"`
	Warnings            []string            `json:"warnings"`
}

var count int

//Client struct
type client struct {
	httpAdrs string
}

//Init Client
func NewClient(httpAdrs string) *client {
	c := client{httpAdrs: httpAdrs}
	return &c
}

func main() {

	count = 0

	// Taking input from user
	fmt.Println("How many forms do you want to send?: ")
	var amount int
	fmt.Scanln(&amount)

	c := NewClient("http://localhost:4245/Orc")

	var responseList []OrchestrationResponse

	for i := 0; i < amount; i++ {

		srf := new(ServiceRequestForm)
		srf.RequesterSystem.SystemName = "LOOK AT PORT DIFFERENCE"
		srf.RequesterSystem.Port = i

		var metaList []string

		counts := strconv.Itoa(count)

		metaList = append(metaList, "META", counts)

		if i%3 == 0 {
			metaList = append(metaList, "%3")
		}

		srf.RequestedService.MetadataRequirements = append(srf.RequestedService.MetadataRequirements, metaList...)

		result := ExchangeJson(c, srf)

		var response OrchestrationResponse

		if err := json.Unmarshal(result, &response); err != nil {
			panic(err)
		}

		responseList = append(responseList, response)

		count++
	}

	//Print responseList

	for i := 0; i < len(responseList); i++ {
		fmt.Println("Response #", i+1)
		fmt.Print(responseList[i])
		fmt.Println("____________________")
		fmt.Println("")
	}

	fmt.Println("Response FIRST ARRIVED")
	fmt.Print(responseList[0])
	fmt.Println("____________________")
	fmt.Println("")
}

func ExchangeJson(c *client, srf *ServiceRequestForm) []byte {

	jsonStr, err := json.Marshal(srf)
	errorHandler(err)

	resp, err := http.Post(c.httpAdrs, "application/json", bytes.NewBuffer(jsonStr))
	errorHandler(err)
	defer resp.Body.Close()

	jsonBody, err := ioutil.ReadAll(resp.Body)
	errorHandler(err)

	return jsonBody
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
