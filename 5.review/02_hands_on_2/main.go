//Serve the files in the "starting-files" folder
//Use "http.FileServer"

package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)

}
