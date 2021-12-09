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