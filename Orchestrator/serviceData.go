package orchestrator
import (
	forms "orchestrator/forms"
)


type serviceData struct {
	serviceRequestForm    forms.ServiceRequestForm
	orchestrationResponse forms.OrchestrationResponse
	discover              forms.Discover
	intraCloudResult 	  forms.IntraCloudResult
	intraCloudRule        forms.IntraCloudRule
}
