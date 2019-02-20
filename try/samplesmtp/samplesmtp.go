package samplesmtp

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"time"

	"github.com/flashmob/go-guerrilla"
)

// SendMail to the specified receipient with the given body.
func SendMail(server, from, to, body string) bool {
	// Connect to the remote SMTP server.
	c, err := smtp.Dial(server)
	if err != nil {
		log.Fatal(err)
		return false
	}

	// Set the sender and recipient first
	if err := c.Mail(from); err != nil {
		log.Fatal(err)
		return false
	}
	if err := c.Rcpt(to); err != nil {
		log.Fatal(err)
		return false
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
		return false
	}
	_, err = fmt.Fprintf(wc, body)
	if err != nil {
		log.Fatal(err)
		return false
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
		return false
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// variables to make ExamplePlainAuth compile, without adding
// unnecessary noise there.
// var (
// 	from       = "gopher@example.net"
// 	msg        = []byte("dummy message")
// 	recipients = []string{"abhijith.da@veritas.com"}
// )

// func ExamplePlainAuth() {
// 	// hostname is used by PlainAuth to validate the TLS certificate.
// 	hostname := "mail.example.com"
// 	auth := smtp.PlainAuth("", "user@example.com", "password", hostname)

// 	err := smtp.SendMail(hostname+":25", auth, from, recipients, msg)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func ExampleSendMail() {
// 	// Set up authentication information.
// 	auth := smtp.PlainAuth("", "user@example.com", "password", "mail.example.com")

// 	// Connect to the server, authenticate, set the sender and recipient,
// 	// and send the email all in one step.
// 	to := []string{"abhijith.da@veritas.com"}
// 	msg := []byte("To: abhijith.da@veritas.com\\r\n" +
// 		"Subject: discount Gophers!\r\n" +
// 		"\r\n" +
// 		"This is the email body.\r\n")
// 	err := smtp.SendMail("127.0.0.1:2526", auth, "abhijith.da@veritas.com", to, msg)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

const logFile = "log.txt"

func init() {
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	// defer f.Close()

	log.SetOutput(f)

	d := guerrilla.Daemon{}
	_, err = d.LoadConfig("guerrillad.conf.json")
	if err != nil {
		fmt.Println("ReadConfig error", err)
	}

	err = d.Start()
	if err != nil {
		fmt.Println("server error", err)
	}
}

func main() {
	server := "127.0.0.1:2526"
	from := "hacker@veritas.com"
	to := "abhijith.da@veritas.com"
	body := "This is the email body"
	SendMail(server, from, to, body)
	time.Sleep(time.Minute * 5)
}
