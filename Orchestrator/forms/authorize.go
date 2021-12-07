package forms

/********************************* Authorize ******************************************/

type Authorize struct {
	IntraCloudRule   IntraCloudRule
	IntraCloudResult IntraCloudResult
}

/********************************* IntraCloudRule ******************************************/

type IntraCloudRule struct {
	Cloud                       Cloud                         `json:"cloud"`
	ProviderIdsWithInterfaceIds []ProviderIdsWithInterfaceIds `json:"providerIdsWithInterfaceIds"`
	ServiceDefinition           string                        `json:"serviceDefinition"`
}

type Cloud struct {
	AuthenticationInfo string `json:"authenticationInfo"`
	GatekeeperRelayIds []int  `json:"gatekeeperRelayIds"`
	GatewayRelayIds    []int  `json:"gatewayRelayIds"`
	Name               string `json:"name"`
	Neighbor           bool   `json:"neighbor"`
	Operator           string `json:"operator"`
	Secure             bool   `json:"secure"`
}

type ProviderIdsWithInterfaceIds struct {
	ID     int   `json:"id"`
	IDList []int `json:"idList"`
}

/********************************* IntraCloudResult ******************************************/

type IntraCloudResult struct {
	AuthorizedProviderIdsWithInterfaceIds []AuthorizedProviderIdsWithInterfaceIds `json:"authorizedProviderIdsWithInterfaceIds"`
	Cloud                                 Cloud                                   `json:"cloud"`
	ServiceDefinition                     string                                  `json:"serviceDefinition"`
}

type AuthorizedProviderIdsWithInterfaceIds struct {
	ID     int   `json:"id"`
	IDList []int `json:"idList"`
}

type Cloud struct {
	AuthenticationInfo string `json:"authenticationInfo"`
	CreatedAt          string `json:"createdAt"`
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Neighbor           bool   `json:"neighbor"`
	Operator           string `json:"operator"`
	OwnCloud           bool   `json:"ownCloud"`
	Secure             bool   `json:"secure"`
	UpdatedAt          string `json:"updatedAt"`
}
