package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "net/smtp"
    "github.com/joho/godotenv"
    "log"
    "os"
    "bufio"
)

func main() {
    err := godotenv.Load("./config/.env")
    if err != nil {
        log.Fatal(err)
    }
    inputType()
}

func inputType() {
    var input string

    fmt.Println("What do you want to do ? 'help' for help")
    fmt.Scan(& input)
    if input == "clear" {
        clearData()
    }
    if input == "write" {
        writeUserData()
    }
    if input == "send" {
        checkUserData()
    }
    if input == "help" {
        help()
    }
    if input == "exit" {
        return
    }
    inputType()
}

type userContent struct {
    Subject string `json: 'subject'`
    Content string `json: 'content'`
}

func help() {
    fmt.Println(`
    Here are the commands.

    write: write a new mail

    send: send the mail

    clear: clear the mail

    exit: stop the program
    `)
}

func writeUserData() {
    filePath := "./config/data.json"
    
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Subject: ")
    subject, _ := reader.ReadString('\n')
    fmt.Println("Content: ")
    content, _ := reader.ReadString('\n')
    
    dataToWrite := userContent{
        Subject: subject,
        Content: content,
    }
    
    jsonData, err := json.MarshalIndent(dataToWrite, "", "    ")
    if err != nil {
        fmt.Println("Error encoding JSON:", err)
        return
    }
    fmt.Println(dataToWrite)
    err = os.WriteFile(filePath, jsonData, 0777)
    if err != nil {
        fmt.Println("Error encoding JSON:", err)
        return
    }
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
    if data.Subject != "" {
        fmt.Println("All good")
        sendMail(data.Subject, data.Content)
        return;
    } else {
        fmt.Println("not good")
        writeUserData()
    }
}

func clearData() {
    filePath := "./config/data.json"
    dataToWrite := userContent{
        Subject: "",
        Content: "",
    }

    jsonData, err := json.MarshalIndent(dataToWrite, "", "    ")

    if err != nil {
        fmt.Println("error encoding JSON: ", err)
        return
    }
 
    err = ioutil.WriteFile(filePath, jsonData, 0777)
}

func sendMail(subject string, content string) {
    senderMail := os.Getenv("SENDER_EMAIL")
    targetName := os.Getenv("TARGET_EMAIL")
    appPassword := os.Getenv("APP_PASSWORD")
    fmt.Println(subject + content)
    auth := smtp.PlainAuth("", senderMail, appPassword, "smtp.gmail.com")
    to := []string{targetName}
    msg := []byte("Subject: " + subject +  "\r\n\r\n" + content)
    err := smtp.SendMail("smtp.gmail.com:587", auth, senderMail, to, msg)
    if err != nil {
        log.Fatal(err)
    }
}

//func getMails() {
//
//}
