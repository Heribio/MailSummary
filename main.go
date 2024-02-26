package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    //"os"
)

func main() {
    checkUserData()
}

type userContent struct {
    Name string `json: 'name'`
    Code string `json: 'code'`
}

func userData() {
    filePath := "./config/data.json"
    
    var username string
    var code string
    fmt.Println("Name: ")
    fmt.Scan(& username)
    fmt.Println("Code: ")
    fmt.Scan(& code)
    
    dataToWrite := userContent{
        Name: username,
        Code: code,
    }
    
    jsonData, err := json.MarshalIndent(dataToWrite, "", "    ")

    if err != nil {
        fmt.Println("Error encoding JSON:", err)
        return
    }
    err = ioutil.WriteFile(filePath, jsonData, 0777)
}

func checkUserData() {
    filePath := "./config/data.json"
    jsonData, err := ioutil.ReadFile(filePath)
 
    if err != nil {
        fmt.Println("Error encoding JSON:", err)
        return
    }
    err = ioutil.WriteFile(filePath, jsonData, 0777)   

    var data userContent

    err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
    
    fmt.Println(data)
    if data.Name != "" {
        fmt.Println("All good")
        return;
    } else {
        fmt.Println("not good")
        userData()
    }   
}

func clearData() {
    filePath := "./config/data.json"
    dataToWrite := userContent{
        Name: "",
        Code: "",
    }

    jsonData, err := json.MarshalIndent(dataToWrite, "", "    ")

    if err != nil {
        fmt.Println("error encoding JSON: ", err)
        return
    }
 
    err = ioutil.WriteFile(filePath, jsonData, 0777)
}
