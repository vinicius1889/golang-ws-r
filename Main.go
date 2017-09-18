package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"os"
	"os/exec"

)

type Script struct {
	Name string `json:"name"`
	Message string `json:"message"`
}

func runScriptInTerminal(scriptName string) Script {
	home := os.Getenv("HOME")+"/rscripts";
	out,err := exec.Command("bash", "-c", "Rscript "+home+"/"+scriptName).Output()

	if err!=nil{
		fmt.Printf("%s",err)
	}

	fmt.Printf("%s",out)

	script := Script{ Name:scriptName, Message:string(out[:])}
	return script
}

func executeScriptHandler( w http.ResponseWriter, r *http.Request){
	scriptName:=mux.Vars(r)
	fmt.Println("Execute Script Handler")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(runScriptInTerminal(scriptName["scriptName"]))
}


func test( w http.ResponseWriter, r *http.Request){
	fmt.Println("Execute Script Handler")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(Script{ Name:"teste", Message:"teste sadsadsada" })
}


func handleRequests() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/execute/{scriptName}", executeScriptHandler)
	myRouter.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":6565", myRouter))
}

func main() {
	handleRequests()
}