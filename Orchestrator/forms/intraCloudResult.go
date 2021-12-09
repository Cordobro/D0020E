package forms

/********************************* IntraCloudResult ******************************************/

type IntraCloudResult struct {
	AuthorizedProviderIdsWithInterfaceIds []struct {
		ID     int   `json:"id"`
		IDList []int `json:"idList"`
	} `json:"authorizedProviderIdsWithInterfaceIds"`
	Consumer struct {
		Address            string `json:"address"`
		AuthenticationInfo string `json:"authenticationInfo"`
		CreatedAt          string `json:"createdAt"`
		ID                 int    `json:"id"`
		Port               int    `json:"port"`
		SystemName         string `json:"systemName"`
		UpdatedAt          string `json:"updatedAt"`
	} `json:"consumer"`
	ServiceDefinitionID int `json:"serviceDefinitionId"`
}
