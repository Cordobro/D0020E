package Orchestrator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Listen(rw http.ResponseWriter, req *http.Request) {

	fmt.Println("Listening...")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	t := Spawn(body)

	/* var t interface{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	} */

	Respond(rw, t)
}

func Respond(rw http.ResponseWriter, responseStruct interface{}) {

	fmt.Println("Responding: ")
	fmt.Println(responseStruct)

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(responseStruct)
}

func SetupServer(port string, funcHandle string) {
	http.HandleFunc(funcHandle, Listen)
	http.ListenAndServe(port, nil)

	fmt.Println("Setup Complete")
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
*/

/*
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
