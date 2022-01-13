package forms

type OrchestrationResponse struct {
	Response []struct {
		Provider struct {
			ID                 int    `json:"id"`
			SystemName         string `json:"systemName"`
			Address            string `json:"address"`
			Port               int    `json:"port"`
			AuthenticationInfo string `json:"authenticationInfo"`
			CreatedAt          string `json:"createdAt"`
			UpdatedAt          string `json:"updatedAt"`
		} `json:"provider"`
		Service struct {
			ID                int    `json:"id"`
			ServiceDefinition string `json:"serviceDefinition"`
			CreatedAt         string `json:"createdAt"`
			UpdatedAt         string `json:"updatedAt"`
		} `json:"service"`
		ServiceURI string `json:"serviceUri"`
		Secure     string `json:"secure"`
		Metadata   struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"metadata"`
		Interfaces []struct {
			ID            int    `json:"id"`
			CreatedAt     string `json:"createdAt"`
			InterfaceName string `json:"interfaceName"`
			UpdatedAt     string `json:"updatedAt"`
		} `json:"interfaces"`
		Version             int `json:"version"`
		AuthorizationTokens struct {
			InterfaceName1 string `json:"interfaceName1"`
			InterfaceName2 string `json:"interfaceName2"`
		} `json:"authorizationTokens"`
		Warnings []string `json:"warnings"`
	} `json:"response"`
}

//This is assuming we have already run the matchmaker and will only return one service to the consumer (the best one)

func ConstructOrchestrationResponse(sql *ServiceQueryList) *OrchestrationResponse {

	var oResponse OrchestrationResponse

	var serviceQueryData = sql.ServiceQueryData[0]
	var response = oResponse.Response[0]

	var provider = response.Provider
	var service = response.Service
	var metadata = response.Metadata

	//Provider
	provider.ID = serviceQueryData.Provider.ID
	provider.SystemName = serviceQueryData.Provider.SystemName
	provider.Address = serviceQueryData.Provider.Address
	provider.Port = serviceQueryData.Provider.Port
	provider.AuthenticationInfo = serviceQueryData.Provider.AuthenticationInfo
	provider.CreatedAt = serviceQueryData.Provider.CreatedAt
	provider.UpdatedAt = serviceQueryData.Provider.UpdatedAt

	//Service
	service.ID = serviceQueryData.ServiceDefinition.ID
	service.ServiceDefinition = serviceQueryData.ServiceDefinition.ServiceDefinition
	service.CreatedAt = serviceQueryData.ServiceDefinition.CreatedAt
	service.UpdatedAt = serviceQueryData.ServiceDefinition.UpdatedAt

	//serviceURI
	response.ServiceURI = serviceQueryData.ServiceURI

	//secure
	response.Secure = serviceQueryData.Secure

	//metadata
	metadata.AdditionalProp1 = serviceQueryData.Metadata.AdditionalProp1
	metadata.AdditionalProp2 = serviceQueryData.Metadata.AdditionalProp2
	metadata.AdditionalProp3 = serviceQueryData.Metadata.AdditionalProp3

	//interfaces
	for i := 0; i < len(serviceQueryData.Interfaces); i++ {
		response.Interfaces[i].ID = serviceQueryData.Interfaces[i].ID
		response.Interfaces[i].InterfaceName = serviceQueryData.Interfaces[i].InterfaceName
		response.Interfaces[i].CreatedAt = serviceQueryData.Interfaces[i].CreatedAt
		response.Interfaces[i].UpdatedAt = serviceQueryData.Interfaces[i].UpdatedAt
	}

	//version
	response.Version = serviceQueryData.Version

	//autorizationTokens
	response.AuthorizationTokens.InterfaceName1 = serviceQueryData.Interfaces[0].InterfaceName
	response.AuthorizationTokens.InterfaceName2 = serviceQueryData.Interfaces[1].InterfaceName

	//warnings

	oResponse.Response[0] = response

	return &oResponse

}
