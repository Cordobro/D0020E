package forms

import "fmt"

type Discover struct {
	ServiceQueryForm ServiceQueryForm
	ServiceQueryList ServiceQueryList
}

type ServiceQueryForm struct {
	ServiceDefinitionRequirement string   `json:"serviceDefinitionRequirement"`
	InterfaceRequirements        []string `json:"interfaceRequirements"`
	SecurityRequirements         []string `json:"securityRequirements"`
	MetadataRequirements         []string `json:"metadataRequirements"`
	VersionRequirement           int      `json:"versionRequirement"`
	MaxVersionRequirement        int      `json:"maxVersionRequirement"`
	MinVersionRequirement        int      `json:"minVersionRequirement"`
	PingProviders                bool     `json:"pingProviders"`
}

/* type MetadataRequirements struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
} */

/********************* ServiceQueryList ************************/

type ServiceQueryList struct {
	ServiceQueryData []ServiceQueryData `json:"serviceQueryData"`
	UnfilteredHits   int                `json:"unfilteredHits"`
}
type ServiceDefinition struct {
	ID                int    `json:"id"`
	ServiceDefinition string `json:"serviceDefinition"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
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

/* type Metadata struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
} */

type Interfaces struct {
	ID            int    `json:"id"`
	InterfaceName string `json:"interfaceName"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}
type ServiceQueryData struct {
	ID                int               `json:"id"`
	ServiceDefinition ServiceDefinition `json:"serviceDefinition"`
	Provider          Provider          `json:"provider"`
	ServiceURI        string            `json:"serviceUri"`
	EndOfValidity     string            `json:"endOfValidity"`
	Secure            string            `json:"secure"`
	Metadata          []string          `json:"metadata"`
	Version           int               `json:"version"`
	Interfaces        []Interfaces      `json:"interfaces"`
	CreatedAt         string            `json:"createdAt"`
	UpdatedAt         string            `json:"updatedAt"`
}

func ConstructServiceQueryForm(srf *ServiceRequestForm, serviceQueryForm *ServiceQueryForm) {

	var requestedService = srf.RequestedService

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

	fmt.Println("---Construct ServiceQueryForm---")
	fmt.Println("")
	fmt.Println(serviceQueryForm)
	fmt.Println("")

}

/********* Testing ************/

func TestNewServiceQueryList(size int) *ServiceQueryList {

	sql := new(ServiceQueryList)

	for i := 0; i < size; i++ {
		sqd := new(ServiceQueryData)
		sql.ServiceQueryData = append(sql.ServiceQueryData, *sqd)

		for j := 0; j < size; j++ {

			in := new(Interfaces)
			sql.ServiceQueryData[i].Interfaces = append(sql.ServiceQueryData[i].Interfaces, *in)
		}

	}

	fmt.Println()
	fmt.Println("--- Inside NewServiceQueryList ---")
	fmt.Println()
	fmt.Println("ServiceQueryData Size: ", len(sql.ServiceQueryData))
	fmt.Println()

	return sql
}
