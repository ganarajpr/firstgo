package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strings"
	"strconv"
	"io"
	"bytes"
	"github.com/sriram-srinivasan/gore/eval"

)

type Function struct {
	Name          string
	Inputs        []string
	Outputs       []string
	Definition    string
	CallProcedure string
}

func createFunctionList() {
	functionList[0].Inputs = append(functionList[0].Inputs, "int", "int")
	functionList[0].Outputs = append(functionList[0].Outputs, "int")
	functionList[0].Name = "sum"
	functionList[0].Definition = `func sum(a,b int)int {
		return a+b
	}`
	functionList[0].CallProcedure = "sum($$_1,$$_2)"

	functionList[1].Inputs = append(functionList[1].Inputs, "int", "int")
	functionList[1].Outputs = append(functionList[1].Outputs, "int")
	functionList[1].Name = "mul"
	functionList[1].Definition = `func mul(a,b int)int {
		return a*b
	}`
	functionList[1].CallProcedure = "mul($$_1,$$_2)"

	/*functionList[2].Inputs = append(functionList[2].Inputs,"float64","float64")
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
	}*/

	//fmt.Println(functionList[0])
	//fmt.Println(functionList[1])
	//fmt.Println(functionList[2])
	//fmt.Println(functionList[3])
	//fmt.Println(functionList[4])

}

var functionList [2]Function

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


func searchFunction(inputs, outputs []string) []Function {
	funcListLength := len(functionList)
	matchedFunctionList := make([]Function, 0)
	for i := 0; i < funcListLength; i++ {
		if StrsEquals(functionList[i].Inputs, inputs) &&
				StrsEquals(functionList[i].Outputs, outputs) {
			matchedFunctionList = append(matchedFunctionList, functionList[i])
		}
	}
	return matchedFunctionList

}

func splitter(delimitedStr string) ([]string, []string) {
	pieces := strings.Split(delimitedStr, ",")
	types := make([]string, 0)
	values := make([]string, 0)
	for i := 0; i < len(pieces); i++ {
		if i%2 == 0 {
			types = append(types, pieces[i])
		}else {
			values = append(values, pieces[i])
		}
	}
	return types, values
}

func main() {
	createFunctionList()

	http.HandleFunc("/Function", func(w http.ResponseWriter, req *http.Request) {
			var inputValues string = req.FormValue("inputs")
			var outputValues string = req.FormValue("outputs")
			inputT, inputV := splitter(inputValues)
			outputT, outputV := splitter(outputValues)
			response := searchFunction(inputT, outputT)
			for _, f := range response {
				callString := f.CallProcedure
				for i, v := range inputV {
					repString := "$$_" + strconv.Itoa(i + 1)
					callString = strings.Replace(callString, repString, v, -1)
				}
				evalString := f.Definition + "\nfmt.Println(" + callString + ")"
				fmt.Println(evalString)
				out, err := eval.Eval(evalString)
				fmt.Println(err)
				fmt.Println(out)
				out = strings.TrimSuffix(out,"\n")
				outByte := []byte(out)
				fmt.Println(len(outByte))
				fmt.Println(len(outputV[0]))
				fmt.Println(outputV[0])
				if bytes.Equal(outByte,[]byte(outputV[0])) {
					b, _ := json.Marshal(f)
					io.WriteString(w, string(b))
				}
			}

		})

	http.ListenAndServe(":8000", nil)
}
