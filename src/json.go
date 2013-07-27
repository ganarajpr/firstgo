package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"strings"
	"errors"
	"io"
	"reflect"
)

type Function struct {
	Name string
	Inputs []string
	Outputs []string
	Definition string
	function interface {}
}

type Comparer interface {
	Compare(inputValues,outputValues []string) bool
}

func (f *Function)Compare(inputValues,outputValues []string)bool{

	return false
}



func createFunctionList(){
	functionList[0].Inputs = append(functionList[0].Inputs,"int","int")
	functionList[0].Outputs = append(functionList[0].Outputs,"int")
	functionList[0].Name="sum"
	functionList[0].Definition= ""
	functionList[0].function = func (a,b int)int {
		return a+b
	}

	functionList[1].Inputs = append(functionList[1].Inputs,"float32","float32")
	functionList[1].Outputs = append(functionList[1].Outputs,"float32")
	functionList[1].Name = "mul"
	functionList[1].Definition = "return a*b"
	functionList[1].function = func (a,b float32)float32 {
		return a*b
	}
	functionList[2].Inputs = append(functionList[2].Inputs,"float64","float64")
	functionList[2].Outputs = append(functionList[2].Outputs,"float64")
	functionList[2].Name = "div"
	functionList[2].Definition = "return a/b"
	functionList[2].function = func (a,b float64)float64 {
		return a/b
	}
	functionList[3].Inputs = append(functionList[3].Inputs,"uint64","uint64")
	functionList[3].Outputs = append(functionList[3].Outputs,"int64")
	functionList[3].Name = "sub"
	functionList[3].Definition = "return a-b"
	functionList[3].function = func (a,b uint64)uint64 {
		return a-b
	}
	functionList[4].Inputs = append(functionList[4].Inputs,"int64","int64")
	functionList[4].Outputs = append(functionList[4].Outputs,"int64")
	functionList[4].Name = "power"
	functionList[4].Definition = "return a^b"
	functionList[4].function = func (a,b int64)int64{
		return a^b
	}

	fmt.Println(functionList[0])
	fmt.Println(functionList[1])
	fmt.Println(functionList[2])
	fmt.Println(functionList[3])
	fmt.Println(functionList[4])

}

var functionList [5]Function

func StrsEquals(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}


func searchFunction(inputs,outputs []string) (Function,error){
	funcListLength := len(functionList)
	fmt.Println(functionList)
	var noOfInputs int = len(inputs)
	fmt.Println(noOfInputs)
	//var doesNotMatch bool = false
	for i :=0; i<funcListLength; i++ {
		if StrsEquals(functionList[i].Inputs,inputs) &&
			StrsEquals(functionList[i].Outputs,outputs){
			return functionList[i], nil
		}
	}
	return Function{},errors.New("No match")

}

func splitter(delimitedStr string) ([]string,[]string){
	pieces := strings.Split(delimitedStr,",")
	types := make([]string,0)
	values := make([]string,0)
	for i:= 0; i < len(pieces); i++ {
		if i % 2 == 0 {
			types = append(types,pieces[i])
		}else{
			values = append(values,pieces[i])
		}
	}
	return types,values
}

func main() {
	createFunctionList()

	http.HandleFunc("/Function", func(w http.ResponseWriter, req *http.Request) {
			var inputValues string = req.FormValue("inputs")
			var outputValues string = req.FormValue("outputs")
			inputT,_ := splitter(inputValues)
			outputT,_ := splitter(outputValues)
			response,error := searchFunction(inputT,outputT)
			fmt.Println(response)
			fmt.Println(error)
			if error == nil{
				//b,_ := json.Marshal(response)
				io.WriteString(w,string(b))
			}
			if error == nil{
				b,_ := json.Marshal(response)
				io.WriteString(w,string(b))
			}
		})

	http.ListenAndServe(":8000",nil)
}
