package forms

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

func ConstructIntraCloudRule(srf *ServiceRequestForm, sql ServiceQueryList) *IntraCloudRule {

	var requesterSystem = srf.RequesterSystem

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

	for i := 0; i < amountOfAvaliableServices; i++ {

		//ID
		var providerID = sql.ServiceQueryData[i].Provider.ID

		var interfaceLength = len(sql.ServiceQueryData[i].Interfaces)

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

	return &intraCloudRule
}
