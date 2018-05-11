package main

import (
	"net/http"
	)

type person struct {
	fName string
}

func (p *person) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	w.Write([]byte("First Name: " + p.fName))
}

func main()  {
	personOne := &person{fName:"Nishi"}
	http.ListenAndServe(":8080", personOne)

	/*myMux := http.NewServeMux()
	myMux.HandleFunc("/", someFunc)
	http.ListenAndServe(":8080", myMux)*/

}

/*func someFunc(w http.ResponseWriter, req *http.Request)  {
	w.Write([]byte("Hello Amsterdam."))
}*/
