package Orchestrator

import "net"

type request struct {
	data ServiceData
	conn net.Conn

	serviceRegAdrs string
	authAdrs       string
}

func NewRequest(data ServiceData, conn net.Conn, serviceRegAdrs string, authAdrs string) *request {
	r := request{data: data,
		conn:           conn,
		serviceRegAdrs: serviceRegAdrs,
		authAdrs:       authAdrs}
	return &r
}

func DoRequest() {

}

func Matchmaker() {}
