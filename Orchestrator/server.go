package Orchestrator

import (
	"fmt"
	"log"
	"net"

	//"os"

	//"strconv"
	"encoding/json"
)

type server struct{	
    CONN_PORT string
}

func NewServer(CONN_PORT string) *server{
	s := server{CONN_PORT: CONN_PORT}
	return &s
}

const (
    CONN_HOST = "localhost"
    CONN_TYPE = "tcp"
)

func Listen(s *server) {

    // Listen for incoming connections.
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+s.CONN_PORT)
    errorHandler(err)

    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        errorHandler(err)
        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {

  // Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
  // Read the incoming connection into the buffer.

	_, err := conn.Read(buf)
	errorHandler(err)

	//data := []byte(strconv.Itoa(reqLen))

    var serviceRequestForm interface{}


    fmt.Println(string(buf))

    buf = RemoveHeader(buf)
    fmt.Println(string(buf))

	Spawn(conn, serviceRequestForm)
}


func RemoveHeader(fix []byte) []byte{
    if fix[0] == []byte("{")[0]{
        return fix
    }else{
        return RemoveHeader(fix[1:])
    }
}


func Respond(conn net.Conn, responseStruct interface{}){
    jsonData, err := json.Marshal(responseStruct)
    errorHandler(err)

    // Send a response back to person contacting us.
	conn.Write(jsonData)
	// Close the connection when you're done with it.
	conn.Close()
}




//https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go