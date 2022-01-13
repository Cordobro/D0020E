package Orchestrator

import (
	"net"
)

const (

	SERVICE_ADDRESS = ""
	AUTH_ADDRESS = ""
    CONN_PORT = ""
)




func Spawn(conn net.Conn, data interface{}){
	
	serviceData := NewServiceData(data)
	r := NewRequest(*serviceData, SERVICE_ADDRESS, AUTH_ADDRESS)
	//Checks what services
	intraCloudRule := sendServiceRequest(r)

	//Checks if requestor has authentication
	

}