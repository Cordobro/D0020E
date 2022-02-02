package Orchestrator

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var SERVICE_ADDRESS string
var AUTH_ADDRESS string

func readFile(searched string) string {

	var a string

	content, err := ioutil.ReadFile("Orchestrator/Config.txt")

	if err != nil {
		log.Fatal(err)
	}

	s := strings.Fields(string(content))

	for i, j := range s {
		if searched == j {
			a = strings.TrimSpace(s[i+1])
			fmt.Println(a)
			return a
		}
	}
	return "error, no result"
}

func InitOrchestrator() {
	SERVICE_ADDRESS = readFile("SERVICE_ADDRESS")
	AUTH_ADDRESS = readFile("AUTH_ADDRESS")
}

func spawn(data []byte) interface{} {
	serviceData := newServiceData(data)
	r := newRequest(*serviceData, SERVICE_ADDRESS, AUTH_ADDRESS)
	sendServiceRequest(r)
	//sendAuthQuery(r)
	//sendTokenQuery(r)
	matchmaker(r)

	return r.data.OrchestrationResponse
}
