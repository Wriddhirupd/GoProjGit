package main

import (
	"fmt"
	"TestProjGo/models"
	"encoding/json"
	// "io/ioutil"
	"log"
	// "os"
	"TestProjGo/util"

)

func main() {

	res := util.JsonToMap("employee.json")

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

