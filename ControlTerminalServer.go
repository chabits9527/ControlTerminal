/**
 *go build -ldflags -H=windowsgui .
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/cmd", cmdHandler)
	err := http.ListenAndServe("localhost:18973", nil)
	if err != nil {
		os.Exit(2)
	}
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(rw, "pong")
	if err != nil {
		return
	}
}

func cmdHandler(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		return
	}

	var body struct {
		Username string   `json:"username"`
		Password string   `json:"password"`
		Cmd      string   `json:"cmd"`
		Args     []string `json:"args"`
	}

	err = json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		return
	}
	fmt.Printf("%+v", body)

	if body.Username == "chabits" && body.Password == "q**897377595" {
		command := exec.Command(body.Cmd, body.Args...)
		command.Stdout = rw
		fmt.Printf("run cmd path ==> %v\n", command.Path)
		fmt.Printf("run cmd args ==> %v\n", command.Args)
		err := command.Run()
		if err != nil {
			fmt.Fprintf(rw, "error ==> %v", err)
		}
	}

}
