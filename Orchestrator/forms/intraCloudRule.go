package forms

import "fmt"

/********************************* IntraCloudRule ******************************************/

type IntraCloudRule struct {
	Consumer struct {
		Address            string `json:"address"`
		AuthenticationInfo string `json:"authenticationInfo"`
		Port               int    `json:"port"`
		SystemName         string `json:"systemName"`
	} `json:"consumer"`
	ProviderIdsWithInterfaceIds []struct {
		ID     int   `json:"id"`
		IDList []int `json:"idList"`
	} `json:"providerIdsWithInterfaceIds"`
	ServiceDefinitionID int `json:"serviceDefinitionId"`
}

type ProviderIdsWithInterfaceIds struct {
	ID     int   `json:"id"`
	IDList []int `json:"idList"`
}

//Only constructs for first element in SRF
func ConstructIntraCloudRule(srf *ServiceRequestForm, discover *Discover, intraCloudRule *IntraCloudRule) {

	fmt.Println("---Inside ConstructIntraCloudRule ---")
	fmt.Println("")

	var requesterSystem = srf.RequesterSystem

	var sql = discover.ServiceQueryList

	fmt.Println("--- sql len ---")
	fmt.Println(len(sql.ServiceQueryData))
	fmt.Println("")

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

	var amountOfAvaliableServices = len(sql.ServiceQueryData) //May be wrong loop variable

	fmt.Println("--- before loop ---")
	fmt.Println(intraCloudRule)
	fmt.Println("")

	for i := 0; i < amountOfAvaliableServices; i++ {

		providerIdsWithInterfaceIds := new(ProviderIdsWithInterfaceIds)

		//ID
		providerIdsWithInterfaceIds.ID = sql.ServiceQueryData[i].Provider.ID

		var interfaceLength = len(sql.ServiceQueryData[i].Interfaces)

		fmt.Println("--- interfaceLength ---")
		fmt.Println(interfaceLength)
		fmt.Println("")

		//IDList
		var interfaceIDList []int

		for j := 0; j < interfaceLength; j++ {
			var interfaceID = sql.ServiceQueryData[i].Interfaces[j].ID

			interfaceIDList = append(interfaceIDList, interfaceID)
		}

		providerIdsWithInterfaceIds.IDList = append(providerIdsWithInterfaceIds.IDList, interfaceIDList...)

		intraCloudRule.ProviderIdsWithInterfaceIds = append(intraCloudRule.ProviderIdsWithInterfaceIds, *providerIdsWithInterfaceIds)

	}

	//ServiceDefinitionID
	intraCloudRule.ServiceDefinitionID = sql.ServiceQueryData[0].ServiceDefinition.ID

	fmt.Println("---IntraCloudRule---")
	fmt.Println("")
	fmt.Println(intraCloudRule)
	fmt.Println("")
}
