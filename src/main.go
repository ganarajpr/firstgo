/**
 * Created with IntelliJ IDEA.
 * User: Ganaraj
 * Date: 30/06/13
 * Time: 22:35
 * To change this template use File | Settings | File Templates.
 */
package main

import  (
		"fmt"
		"net/http"
		"io"
		//"search"
)




func main() {
	//searchHandler := search.SearchHandler{}
	fmt.Println("Listening on 8000");
	http.HandleFunc("/",func(w http.ResponseWriter,req *http.Request){
			io.WriteString(w,"Hello " + req.FormValue("name"))
		});
	http.ListenAndServe(":8000",nil)
}
