package Orchestrator

import (
<<<<<<< HEAD
	"fmt"
	//"log"
	"net"

	//"os"

	//"strconv"
=======
	"io/ioutil"
	"net/http"
>>>>>>> 9799d911f72546e0cb95c74d3d523fc8f5f6addb
	"encoding/json"
)




func Listen(rw http.ResponseWriter, req *http.Request) {
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        panic(err)
    }
	
    Spawn(body)

    var t interface{}
    err = json.Unmarshal(body, &t)
	if err != nil {
        panic(err)
    }

    Respond(rw, t)
}

func Respond(rw http.ResponseWriter, responseStruct interface{}){
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(responseStruct) 
}

func SetupServer(port string, funcHandle string){ 
    http.HandleFunc(funcHandle, Listen)
    http.ListenAndServe(port, nil)
}

/*
func NewServer(CONN_PORT string) *server{
	s := server{CONN_PORT: CONN_PORT}
	return &s
}

const (
    CONN_HOST = "localhost"
    CONN_TYPE = "tcp"
)
*/
/*
type server struct{	
    CONN_PORT string
}*/
/*
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
*/
/*
// Handles incoming requests.
func handleRequest(conn net.Conn) {

  // Make a buffer to hold incoming data.
	buf := make([]byte, 2048)
  // Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	errorHandler(err)

    var serviceRequestForm interface{}
    //fmt.Println(string(buf))

    buf = RemoveHeader(buf)
    buf = RemoveEnd(buf)

    if err := json.Unmarshal(buf, &serviceRequestForm); err != nil {
        panic(err)
    }
    //fmt.Println(serviceRequestForm) 

	Spawn(conn, serviceRequestForm)
}
<<<<<<< HEAD
=======
*/
<<<<<<< HEAD
>>>>>>> 9799d911f72546e0cb95c74d3d523fc8f5f6addb

=======
/*
>>>>>>> 58dd4d2bcfd6d6a09cc13104f7ee865b69efb275
func RemoveHeader(fix []byte) []byte{
    if fix[0] == []byte("{")[0]{
        return fix
    }else{
        return RemoveHeader(fix[1:])
    }
}

func RemoveEnd(fix []byte) []byte{
    for i := 0; i < len(fix); i++ {
        if fix[i] == 0 {
            return fix[0:i]
        }
    }
    return fix
}



*/



//https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go