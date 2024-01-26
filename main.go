package main

import (
	"fmt"
	"io"
	"net/http"
	// "strings"
)

const myurl = "http://localhost:3001/arduino"

func main() {

  // getResponce := PerformGetRequest(myurl)
  // fmt.Println(getResponce)
  //
  // payload := strings.NewReader(`{
  //   "id":100,
  //   "name":"bad person"
  // }`)
  // postResponce := PerformPostRequest(myurl,payload)
  // fmt.Println(postResponce)
  WebApp()
}





// ######################## HTTP requests ###########################

func PerformGetRequest(url string) string {

	// ---------get the responce-----------
	responce, getRequestErr := http.Get(url) // make get request
	if getRequestErr != nil {
		panic(getRequestErr)
	}
	defer responce.Body.Close() //close the responce in the end

  content := ReadBody(responce)
	return content
}

func PerformPostRequest(url string, body io.Reader) string {

  // --------------------make a post request -------------------------
	responce, postRequestErr := http.Post(url, "application/json", body) // make post request
	if postRequestErr != nil {
		panic(postRequestErr)
	}
	defer responce.Body.Close() //close the responce in the end
  content := ReadBody(responce)
  return content
}




// ########## small toos #################
func ReadBody(responce *http.Response) string {

	body, bodyErr := io.ReadAll(responce.Body)
	if bodyErr != nil {
		panic(bodyErr)
	}
	content := string(body)
	return content
}
