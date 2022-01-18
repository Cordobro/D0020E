package Orchestrator

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var	SERVICE_ADDRESS string
var	AUTH_ADDRESS string

func ReadFile(searched string) string {

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

func InitOrchestrator(){
	SERVICE_ADDRESS = ReadFile("SERVICE_ADDRESS")
	AUTH_ADDRESS = ReadFile("AUTH_ADDRESS")
<<<<<<< HEAD
	CONN_PORT = ReadFile("CONN_PORT")

	//fmt.Println(CONN_PORT)

}

func Spawn(conn net.Conn, data interface{}){
	
	serviceData := NewServiceData(data)
	r := NewRequest(*serviceData, SERVICE_ADDRESS, AUTH_ADDRESS)
	//Checks what services
	//intraCloudRule := sendServiceRequest(r)

	//Checks if requestor has authentication
	

=======
}

func Spawn(data []byte) interface{}{
	serviceData := NewServiceData(data)
	r := NewRequest(*serviceData, SERVICE_ADDRESS, AUTH_ADDRESS)
	sendServiceRequest(r)
	//sendAuthQuery(r)
	//sendTokenQuery(r)
	Matchmaker(r)
	
	return r.data.OrchestrationResponse
>>>>>>> 58dd4d2bcfd6d6a09cc13104f7ee865b69efb275
}
