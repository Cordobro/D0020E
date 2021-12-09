package orchestrator

import (
	"encoding/json"
	"fmt"
	forms "orchestrator/forms"
)

func test() {
	var serviceRequestForm forms.ServiceRequestForm

	serviceRequestForm.RequesterSystem.SystemName = "Jacks dator"
}

func constructMessage() {

}

func parseMessage(serviceRequestForm forms.ServiceRequestForm) {
	res, _ := json.Marshal(serviceRequestForm)
	fmt.Println(res)
}
