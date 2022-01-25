package Orchestrator

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func listen(rw http.ResponseWriter, req *http.Request) {
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        panic(err)
    }
	
    t := spawn(body)
    respond(rw, t)
}

func respond(rw http.ResponseWriter, responseStruct interface{}){
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(responseStruct)
}

func SetupServer(){ 
    port := readFile("CONN_PORT")
    funcHandle := readFile("LISTEN_HANDLE")
    http.HandleFunc(funcHandle, listen)
    http.ListenAndServe(port, nil) 
}
