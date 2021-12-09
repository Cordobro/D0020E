package Orchestrator
import (
	forms "arrowhead/Orchestrator/forms"
)


type serviceData struct {
	serviceRequestForm    forms.ServiceRequestForm
	orchestrationResponse forms.OrchestrationResponse
	discover              forms.Discover
	intraCloudResult 	  forms.IntraCloudResult
	intraCloudRule        forms.IntraCloudRule
}
