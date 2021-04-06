package services

import (
	"fmt"
	"github.com/eshaanmangal/Go-Project-Structure/repository"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func HandlePushElement(w http.ResponseWriter, r *http.Request) {
	stackElement := mux.Vars(r)["element"]
	fmt.Println(mux.Vars(r))
	var element repository.Element
	element.StackElement = stackElement
	pushingElement, err := repository.Create(&element)
	if err!=nil {
		log.Fatalln(err)
	}
	fmt.Println(pushingElement)
}
