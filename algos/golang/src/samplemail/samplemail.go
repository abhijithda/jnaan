package main

import "github.com/zemirco/email"

// const smtpServer = "mtv.nbuappsmtp.engba.veritas.com"
// const smtpServer = "nbuappsmtp.engba.symantec.com"
const smtpServer = "10.182.128.220"
const smtpPort = 25
const sender = "NoReply_Veritas_Appliance@sclautoesxd04v16.engba.symantec.com"
// const sender = "uManage@veritas.com",

func main() {
  mail := email.Mail{
    From:    sender,
    To:      "abhijith.da@veritas.com",
    Subject: "Hello there",
    HTML:    "<h1>Awesome</h1>",
  }
  err := mail.Send(smtpServer, smtpPort, sender, "PASS")
  if err != nil {
    panic(err)
  }
}

