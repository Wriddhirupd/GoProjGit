package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func JsonToMap(fileName string) map[string]interface{} {
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

// func mapToStruct(mapName map) (emps []models.EmployeeDetail) {
// 	for k, v := range mapName {
//         switch val := v.(type) {
//         case string:
//                 fmt.Println(k, "is string", val)
//         case int:
//                 fmt.Println(k, "is int", val)
//         case []interface{}:
// 				fmt.Println(k, "is an array")
// 				var emp models.EmployeeDetail
//                 for key, value := range val {
// 						// fmt.Println(i, v)
// 						mapstructure.Decode(value, &emp)
// 						emps = append(emps, emp)
// 				}
				
//         default:
//                 fmt.Println(k, "is unknown type")
// 		}
		
// 		return emps
// 	}
// }
