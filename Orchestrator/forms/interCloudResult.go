package forms

type InterCloudResult struct {
	AuthorizedProviderIdsWithInterfaceIds []struct {
		ID     int   `json:"id"`
		IDList []int `json:"idList"`
	} `json:"authorizedProviderIdsWithInterfaceIds"`

	Cloud struct {
		AuthenticationInfo string `json:"authenticationInfo"`
		CreatedAt          string `json:"createdAt"`
		ID                 int    `json:"id"`
		Name               string `json:"name"`
		Neighbor           bool   `json:"neighbor"`
		Operator           string `json:"operator"`
		OwnCloud           bool   `json:"ownCloud"`
		Secure             bool   `json:"secure"`
		UpdatedAt          string `json:"updatedAt"`
	} `json:"cloud"`

	ServiceDefinition string `json:"serviceDefinition"`
}
