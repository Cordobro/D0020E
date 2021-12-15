package Orchestrator

import (
	forms "arrowhead/Orchestrator/forms"
	"encoding/json"
	"fmt"
)

type serviceData struct {
	serviceRequestForm    forms.ServiceRequestForm
	orchestrationResponse forms.OrchestrationResponse
	discover              forms.Discover
	intraCloudResult      forms.IntraCloudResult
	intraCloudRule        forms.IntraCloudRule
	tokenRule             forms.TokenRule
	tokenResult           forms.TokenData
}

func NewServiceData(request forms.ServiceRequestForm) *serviceData {

	s := serviceData{}
	s.serviceRequestForm = request

	return &s
}

func TestMH() {

	var serviceRequestForm forms.ServiceRequestForm

	serviceRequestForm.RequesterSystem.SystemName = "Jacks dator"

	test, _ := json.Marshal(serviceRequestForm)

	ParseMessage(test)

}

/********************************************************************************** ServiceQueryForm ******************************************************************************/

func ConstructServiceQueryForm(serviceData serviceData) {

	var requestedService = serviceData.serviceRequestForm.RequestedService

	discover := new(forms.Discover)
	var serviceQueryForm = discover.ServiceQueryForm

	//ServiceDefinitionRequirement
	serviceQueryForm.ServiceDefinitionRequirement = requestedService.ServiceDefinitionRequirement

	//InterfaceRequirements
	serviceQueryForm.InterfaceRequirements = requestedService.InterfaceRequirements

	//SecurityRequirements
	serviceQueryForm.SecurityRequirements = requestedService.SecurityRequirements

	//MetadataRequirements
	serviceQueryForm.MetadataRequirements = requestedService.MetadataRequirements

	//VersionRequirement
	serviceQueryForm.VersionRequirement = requestedService.VersionRequirement

	//MaxVersionRequirement
	serviceQueryForm.MaxVersionRequirement = requestedService.MaxVersionRequirement

	//MinVersionRequirement
	serviceQueryForm.MinVersionRequirement = requestedService.MinVersionRequirement

	//PingProviders
	serviceQueryForm.PingProviders = true

	serviceData.discover.ServiceQueryForm = serviceQueryForm
}

/********************************************************************************** IntraCloudRule ******************************************************************************/

//might not work
func ConstructIntraCloudRule(serviceData serviceData) {

	var requesterSystem = serviceData.serviceRequestForm.RequesterSystem

	var intraCloudRule forms.IntraCloudRule

	/***Consumer***/

	//Address
	intraCloudRule.Consumer.Address = requesterSystem.Address

	//AuthenticationInfo
	intraCloudRule.Consumer.AuthenticationInfo = requesterSystem.AuthenticationInfo

	//Port
	intraCloudRule.Consumer.Port = requesterSystem.Port

	//SystemName
	intraCloudRule.Consumer.SystemName = requesterSystem.SystemName

	/***ProviderIdsWithInterfaceIds***/

	var amountOfAvaliableServices = len(serviceData.discover.ServiceQueryList.ServiceQueryData)

	for i := 0; i < amountOfAvaliableServices; i++ {

		//ID
		var providerID = serviceData.discover.ServiceQueryList.ServiceQueryData[i].Provider.ID

		var interfaceLength = len(serviceData.discover.ServiceQueryList.ServiceQueryData[i].Interfaces)

		//IDList
		var interfaceIDList []int

		for j := 0; j < interfaceLength; j++ {
			var interfaceID = serviceData.discover.ServiceQueryList.ServiceQueryData[i].Interfaces[j].ID

			interfaceIDList = append(interfaceIDList, interfaceID)
		}

		intraCloudRule.ProviderIdsWithInterfaceIds[i].ID = providerID
		intraCloudRule.ProviderIdsWithInterfaceIds[i].IDList = interfaceIDList[:]

		//ServiceDefinitionID
		intraCloudRule.ServiceDefinitionID = serviceData.discover.ServiceQueryList.ServiceQueryData[i].ID
	}

	serviceData.intraCloudRule = intraCloudRule
}

/********************************************************************************** TokenRule ******************************************************************************/

func ConstructTokenRule(serviceData serviceData) {

	tokenRule := new(forms.TokenRule)

	//Consumer
	tokenRule.Consumer.Address = serviceData.serviceRequestForm.RequesterSystem.Address
	tokenRule.Consumer.AuthenticationInfo = serviceData.serviceRequestForm.RequesterSystem.AuthenticationInfo
	tokenRule.Consumer.Port = serviceData.serviceRequestForm.RequesterSystem.Port
	tokenRule.Consumer.SystemName = serviceData.serviceRequestForm.RequesterSystem.SystemName

	//ConsumerCloud
	//Where do we get these?
	/*
		tokenRule.ConsumerCloud.AuthenticationInfo =
		tokenRule.ConsumerCloud.GatekeeperRelayIds =
		tokenRule.ConsumerCloud.GatewayRelayIds =
		tokenRule.ConsumerCloud.Name =
		tokenRule.ConsumerCloud.Neighbor =
		tokenRule.ConsumerCloud.Operator =
		tokenRule.ConsumerCloud.Secure =
	*/

	//Duration
	//tokenRule.Duration =

	//Providers
	//provider

	//Service
	//tokenRule.Service

}

/********************************************************************************** OrchestrationResponse ******************************************************************************/

func ConstructOrchestrationResponse(serviceData serviceData, serviceQueryList forms.ServiceQueryList) {

	var serviceQueryData = serviceQueryList.ServiceQueryData[0]

	var response = serviceData.orchestrationResponse.Response[0]

	var provider = response.Provider
	var service = response.Service
	var metadata = response.Metadata
	//var authorizationTokens = response.AuthorizationTokens

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

	//warnings

	/********************* OBS!! DONT HAVE YET **************************/

	serviceData.orchestrationResponse.Response[0] = response

}

func ParseMessage(byteValue []byte) {

	//var dat map[string]interface{}

	var requestForm forms.ServiceRequestForm

	var _ = json.Unmarshal(byteValue, &requestForm)

	fmt.Println("___INSIDE ParseMessage___")
	fmt.Println(requestForm)
}

func ComposeMessage() {

}
