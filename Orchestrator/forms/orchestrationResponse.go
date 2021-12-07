package forms

type orchestrationResponse struct {
	Response []Response `json:"response"`
}
type Provider struct {
	ID                 int    `json:"id"`
	SystemName         string `json:"systemName"`
	Address            string `json:"address"`
	Port               int    `json:"port"`
	AuthenticationInfo string `json:"authenticationInfo"`
	CreatedAt          string `json:"createdAt"`
	UpdatedAt          string `json:"updatedAt"`
}
type Service struct {
	ID                int    `json:"id"`
	ServiceDefinition string `json:"serviceDefinition"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
}
type Metadata struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}
type Interfaces struct {
	ID            int    `json:"id"`
	CreatedAt     string `json:"createdAt"`
	InterfaceName string `json:"interfaceName"`
	UpdatedAt     string `json:"updatedAt"`
}
type AuthorizationTokens struct {
	InterfaceName1 string `json:"interfaceName1"`
	InterfaceName2 string `json:"interfaceName2"`
}
type Response struct {
	Provider            Provider            `json:"provider"`
	Service             Service             `json:"service"`
	ServiceURI          string              `json:"serviceUri"`
	Secure              string              `json:"secure"`
	Metadata            Metadata            `json:"metadata"`
	Interfaces          []Interfaces        `json:"interfaces"`
	Version             int                 `json:"version"`
	AuthorizationTokens AuthorizationTokens `json:"authorizationTokens"`
	Warnings            []string            `json:"warnings"`
}
