package forms

type Discover struct {
	ServiceQueryForm ServiceQueryForm
	ServiceQueryList ServiceQueryList
}

type ServiceQueryForm struct {
	ServiceDefinitionRequirement string   `json:"serviceDefinitionRequirement"`
	InterfaceRequirements        []string `json:"interfaceRequirements"`
	SecurityRequirements         []string `json:"securityRequirements"`
	MetadataRequirements         struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"metadataRequirements"`
	VersionRequirement    int  `json:"versionRequirement"`
	MaxVersionRequirement int  `json:"maxVersionRequirement"`
	MinVersionRequirement int  `json:"minVersionRequirement"`
	PingProviders         bool `json:"pingProviders"`
}

type ServiceQueryList struct {
	ServiceQueryData []struct {
		ID                int `json:"id"`
		ServiceDefinition struct {
			ID                int    `json:"id"`
			ServiceDefinition string `json:"serviceDefinition"`
			CreatedAt         string `json:"createdAt"`
			UpdatedAt         string `json:"updatedAt"`
		} `json:"serviceDefinition"`
		Provider struct {
			ID                 int    `json:"id"`
			SystemName         string `json:"systemName"`
			Address            string `json:"address"`
			Port               int    `json:"port"`
			AuthenticationInfo string `json:"authenticationInfo"`
			CreatedAt          string `json:"createdAt"`
			UpdatedAt          string `json:"updatedAt"`
		} `json:"provider"`
		ServiceURI    string `json:"serviceUri"`
		EndOfValidity string `json:"endOfValidity"`
		Secure        string `json:"secure"`
		Metadata      struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"metadata"`
		Version    int `json:"version"`
		Interfaces []struct {
			ID            int    `json:"id"`
			InterfaceName string `json:"interfaceName"`
			CreatedAt     string `json:"createdAt"`
			UpdatedAt     string `json:"updatedAt"`
		} `json:"interfaces"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	} `json:"serviceQueryData"`
	UnfilteredHits int `json:"unfilteredHits"`
}
