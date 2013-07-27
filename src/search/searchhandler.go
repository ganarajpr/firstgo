/**
 * Created with IntelliJ IDEA.
 * User: Ganaraj
 * Date: 15/07/13
 * Time: 19:59
 * To change this template use File | Settings | File Templates.
 */
package search

import (
	"net/http"
	"io"
	"fmt"

)

type SearchHandler struct {}



func (c *SearchHandler) ServeHTTP(w http.ResponseWriter, request *http.Request){
	io.WriteString(w,"Hello " + request.FormValue("name"))
}

func PrintMe(){
	fmt.Println("Hello")
}

