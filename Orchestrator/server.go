package Orchestrator

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
)


func SetupServer(port string, funcHandle string){
    http.HandleFunc(funcHandle, Listen)
    http.ListenAndServe(port, nil)
}

func Listen(rw http.ResponseWriter, req *http.Request) {
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        panic(err)
    }
	
    var t interface{}
    err = json.Unmarshal(body, &t)
	if err != nil {
        panic(err)
    }

    Respond(rw, t)
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
*/

func RemoveHeader(fix []byte) []byte{
    if fix[0] == []byte("{")[0]{
        return fix
    }else{
        return RemoveHeader(fix[1:])
    }
}


func Respond(rw http.ResponseWriter, responseStruct interface{}){
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(responseStruct) 
}




//https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go