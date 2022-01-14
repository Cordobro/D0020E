package Orchestrator

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

var	SERVICE_ADDRESS string
var	AUTH_ADDRESS string
var CONN_PORT string


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

func SetVar(){

	SERVICE_ADDRESS = ReadFile("SERVICE_ADDRESS")
	AUTH_ADDRESS = ReadFile("AUTH_ADDRESS")
	CONN_PORT = ReadFile("CONN_PORT")

	fmt.Println(CONN_PORT)

}

func Spawn(conn net.Conn, data interface{}){
/* 	
	serviceData := NewServiceData(data)
	r := NewRequest(*serviceData, SERVICE_ADDRESS, AUTH_ADDRESS)
	//Checks what services
	intraCloudRule := sendServiceRequest(r)

	//Checks if requestor has authentication
	 */

}
