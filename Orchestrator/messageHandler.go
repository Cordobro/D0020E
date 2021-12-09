package Orchestrator

import (
	"encoding/json"
	"fmt"
	forms "arrowhead/Orchestrator/forms"
)

func Test() {
	/*
	var serviceRequestForm forms.ServiceRequestForm

	serviceRequestForm.RequesterSystem.SystemName = "Jacks dator"
	fmt.Println(serviceRequestForm.RequesterSystem.SystemName)
	*/
	println("hello")
}

func constructMessage() {

}

func parseMessage(serviceRequestForm forms.ServiceRequestForm) {
	res, _ := json.Marshal(serviceRequestForm)
	fmt.Println(res)
}
