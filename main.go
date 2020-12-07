package main

import (
	"fmt"
	"TestProjGo/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

)

func main() {

	res := jsonToMap("employee.json")

	resArray:=res["employees"].([]interface{})

	// fmt.Println(resArray)
	 var emps models.EmployeeDetails

	b,_ := json.Marshal(resArray)
	err := json.Unmarshal(b, &emps.Employees)
	
	
	 if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(emps)

}


func jsonToMap(fileName string) map[string]interface{} {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened ", fileName)

	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(b), &result)

	return result
}