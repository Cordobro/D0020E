package forms

import "fmt"

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

/* type metadata struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
} */

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

//This is assuming we have already run the matchmaker and will only return one service to the consumer (the best one)

func ConstructOrchestrationResponse(sql *ServiceQueryList, oResponse *OrchestrationResponse) {

	fmt.Println("---Inside OrchestrationResponse ---")

	var serviceQueryData = sql.ServiceQueryData

	fmt.Println("serviceQueryData length: ", len(serviceQueryData))

	var response Response

	//Make sures that serviceQueryData isn't empty

	if len(serviceQueryData) > 0 {

		var provider = response.Provider
		var service = response.Service
		var metadata = response.Metadata

		//Provider
		provider.ID = serviceQueryData[0].Provider.ID
		provider.SystemName = serviceQueryData[0].Provider.SystemName
		provider.Address = serviceQueryData[0].Provider.Address
		provider.Port = serviceQueryData[0].Provider.Port
		provider.AuthenticationInfo = serviceQueryData[0].Provider.AuthenticationInfo
		provider.CreatedAt = serviceQueryData[0].Provider.CreatedAt
		provider.UpdatedAt = serviceQueryData[0].Provider.UpdatedAt

		response.Provider = provider

		//Service
		service.ID = serviceQueryData[0].ServiceDefinition.ID
		service.ServiceDefinition = serviceQueryData[0].ServiceDefinition.ServiceDefinition
		service.CreatedAt = serviceQueryData[0].ServiceDefinition.CreatedAt
		service.UpdatedAt = serviceQueryData[0].ServiceDefinition.UpdatedAt

		response.Service = service

		//serviceURI
		response.ServiceURI = serviceQueryData[0].ServiceURI

		//secure
		response.Secure = serviceQueryData[0].Secure

		//metadata
		metadata = serviceQueryData[0].Metadata

		response.Metadata = metadata

		fmt.Println("Response")
		fmt.Println(response)
		fmt.Println("")

		//interfaces

		var ifList []interfaces

		for i := 0; i < len(serviceQueryData[0].Interfaces); i++ {

			var interf interfaces

			interf.ID = serviceQueryData[0].Interfaces[i].ID
			interf.InterfaceName = serviceQueryData[0].Interfaces[i].InterfaceName
			interf.CreatedAt = serviceQueryData[0].Interfaces[i].CreatedAt
			interf.UpdatedAt = serviceQueryData[0].Interfaces[i].UpdatedAt

			ifList = append(ifList, interf)
		}

		response.Interfaces = ifList

		//version
		response.Version = serviceQueryData[0].Version

		//autorizationTokens
		//Where do we get these?
		//response.AuthorizationTokens.InterfaceName1 = serviceQueryData[0].Interfaces[0].InterfaceName
		//response.AuthorizationTokens.InterfaceName2 = serviceQueryData[0].Interfaces[1].InterfaceName

		//warnings
		//no warnings yet

	}

	oResponse.Response = append(oResponse.Response, response)

}
