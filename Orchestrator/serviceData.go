package Orchestrator

import (
	forms "arrowhead/Orchestrator/forms"
	"fmt"
)

type serviceData struct {
	serviceRequestForm    forms.ServiceRequestForm
	orchestrationResponse forms.OrchestrationResponse
	discover              forms.Discover
	intraCloudResult      forms.IntraCloudResult
	intraCloudRule        forms.IntraCloudRule
}

func TestMH() {

	var serviceRequestForm forms.ServiceRequestForm

	serviceRequestForm.RequesterSystem.SystemName = "Jacks dator"
	fmt.Println(serviceRequestForm.RequesterSystem.SystemName)

}

func constructMessage(messageType string, serviceData serviceData) {

	switch messageType {

	case "serviceQueryForm":

		constructServiceQueryForm(serviceData)

	case "intraCloudRule":

		constructIntraCloudRule(serviceData)

	case "orchestrationResponse":

		orchestrationResponse := new(forms.OrchestrationResponse)

		serviceData.orchestrationResponse = *orchestrationResponse

	default:

	}

}

/********************************************************************************** ServiceQueryForm ******************************************************************************/

func constructServiceQueryForm(serviceData serviceData) {

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
func constructIntraCloudRule(serviceData serviceData) []forms.IntraCloudRule {

	var intraCloudRuleList []forms.IntraCloudRule

	var requesterSystem = serviceData.serviceRequestForm.RequesterSystem

	intraCloudRule := new(forms.IntraCloudRule)

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

		//add to main list
		intraCloudRuleList = append(intraCloudRuleList, *intraCloudRule)
	}

	return intraCloudRuleList
}

/********************************************************************************** OrchestrationResponse ******************************************************************************/

func constructOrchestrationResponse(serviceData serviceData) {

}

func parseMessage(messageType string, serviceData serviceData, byteValue []byte) {

	//res, _ := json.Unmarshal([]byte(byteValue), &serviceData.serviceRequestForm)

	//fmt.Println(string(res))
}
