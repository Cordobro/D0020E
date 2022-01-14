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

<<<<<<< HEAD
func ConstructIntraCloudRule(srf *ServiceRequestForm, sql *ServiceQueryList) *IntraCloudRule {
=======
func ConstructIntraCloudRule(srf *ServiceRequestForm, discover *Discover) *IntraCloudRule {

	fmt.Println("---Inside ConstructIntraCloudRule ---")
	fmt.Println("")
>>>>>>> fb29c1863f8ec0251b6bfaf75eff812182345dea

	var requesterSystem = srf.RequesterSystem

	var sql = discover.ServiceQueryList

	fmt.Println("--- sql len ---")
	fmt.Println(len(sql.ServiceQueryData))
	fmt.Println("")

	var intraCloudRule IntraCloudRule

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

	var amountOfAvaliableServices = len(sql.ServiceQueryData)

	fmt.Println("--- before loop ---")
	fmt.Println(intraCloudRule)
	fmt.Println("")

	for i := 0; i < amountOfAvaliableServices; i++ {

		//ID
		var providerID = sql.ServiceQueryData[i].Provider.ID

		var interfaceLength = len(sql.ServiceQueryData[i].Interfaces)

		fmt.Println("--- provider ID ---")
		fmt.Println(providerID)
		fmt.Println("")

		fmt.Println("--- interfaceLength ---")
		fmt.Println(interfaceLength)
		fmt.Println("")

		//IDList
		var interfaceIDList []int

		for j := 0; j < interfaceLength; j++ {
			var interfaceID = sql.ServiceQueryData[i].Interfaces[j].ID

			interfaceIDList = append(interfaceIDList, interfaceID)
		}

		intraCloudRule.ProviderIdsWithInterfaceIds[i].ID = providerID
		intraCloudRule.ProviderIdsWithInterfaceIds[i].IDList = interfaceIDList[:]

		//ServiceDefinitionID
		intraCloudRule.ServiceDefinitionID = sql.ServiceQueryData[i].ID

	}

	fmt.Println("---IntraCloudRule---")
	fmt.Println("")
	fmt.Println(intraCloudRule)
	fmt.Println("")

	return &intraCloudRule
}
