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
    "strings"
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
    if input == "see" {
        seeMail()
    }
    if input == "help" {
        help()
    }
    if input == "exit" {
        return
    }
    inputType()
}

type mailStructure struct {
    Subject string `json: 'subject'`
    Content string `json: 'content'`
    MailAddress string `json: 'address'`
}

func help() {
    fmt.Println(`
    Here are the commands.

    write: write a new mail

    send: send the mail

    clear: clear the mail

    see: see the mail content

    exit: stop the program
    `)
}

func writeUserData() {
    filePath := "./config/data.json"
    
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Subject: ")
    subject, _ := reader.ReadString('\n')
    subject = strings.TrimRight(subject, "\n")
    fmt.Println("Content: ")
    content, _ := reader.ReadString('\n')
    content = strings.TrimRight(content, "\n")
    fmt.Println("Target address: ")
    address, _ := reader.ReadString('\n')
    address = strings.TrimRight(address, "\n")
    
    dataToWrite := mailStructure{
        Subject: subject,
        Content: content,
        MailAddress: address,
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

    var data mailStructure 

    err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
    
    fmt.Println(data)
    if data.Subject != "" {
        fmt.Println("All good")
        sendMail(data.Subject, data.Content, data.MailAddress)
        return;
    } else {
        fmt.Println("not good")
        writeUserData()
    }
}

func clearData() {
    filePath := "./config/data.json"
    dataToWrite := mailStructure{
        Subject: "",
        Content: "",
        MailAddress: "",
    }

    jsonData, err := json.MarshalIndent(dataToWrite, "", "    ")

    if err != nil {
        fmt.Println("error encoding JSON: ", err)
        return
    }
 
    err = ioutil.WriteFile(filePath, jsonData, 0777)
}

func sendMail(subject string, content string, address string) {
    senderMail := os.Getenv("SENDER_EMAIL")
    appPassword := os.Getenv("APP_PASSWORD")
    fmt.Println(subject + content)
    auth := smtp.PlainAuth("", senderMail, appPassword, "smtp.gmail.com")
    to := []string{address}
    msg := []byte("Subject: " + subject +  "\r\n\r\n" + content)
    err := smtp.SendMail("smtp.gmail.com:587", auth, senderMail, to, msg)
    if err != nil {
        log.Fatal(err)
    }
}

func seeMail() {
    filePath := "./config/data.json"

    jsonData, err := ioutil.ReadFile(filePath)

    if err != nil {
        fmt.Println("Error encoding JSON:", err)
        return
    }   
    
    var data mailStructure

    err = json.Unmarshal(jsonData, &data)
    if err != nil {
        fmt.Println("Error unmarshaling JSON:", err)
        return
    }
    
    subject := data.Subject
    content := data.Content
    target := data.MailAddress

    fmt.Println("Subject: ", subject)
    fmt.Println("Content: ", content)
    fmt.Println("Target: ", target)
}

//func getMails() {
//
//}
