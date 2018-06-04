package main

import (
    "fmt"
    "log"
    "net/http"
    "net/smtp"
    "os"
    "bufio"
    "time"
    "flag"
)

//setting up the timer
func periodicFunc(tick time.Time){
}

func main(){

  //initializing flag
  timerPtr := flag.String("timer", "once", "Write many to start ticker, or once without ticker")

  flag.Parse()

  //setting timer and start loop of checkings
  if *timerPtr == "many"{
        readDoc()
    //you can choose range of how often checking will go, range time.NewTicker(10 * time.Second, Hour, Millisecond etc.)
    for t := range time.NewTicker(10 * time.Second).C {
        readDoc()
        periodicFunc(t)
    }
  } else {
    readDoc()
  }


}

//checking status of website
func checkStatus(site string) {

              //sending get responce to the website
              resp, err := http.Get(site)
              status := "UNKNOWN"
                  if (err == nil) && (resp.StatusCode == 200) {
                        status = "\nUP"
                        } else {
                        status = "\nDOWN \n" + err.Error()
                        }
                        body := site + status
                        //decommend send(body) to start e-mail sendings
                        send(body)
                        fmt.Println(body)
}


//sending e-mail
func send(body string) {
  //from email
	from := "your e-mail"
  //from pass
	pass := "your e-mail password"
  //whome to send
	to := "whom send notifications (e-mail)"

  //creation of the message
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Website status\n\n" +
		body


	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}

//reading .txt document websites should be one website on one line
func readDoc(){

  //put file path into ("filepath")
    sites, err := os.Open("sites.txt")
        if err != nil {
            log.Fatal(err)
            }

    defer sites.Close()

    scanner := bufio.NewScanner(sites)
    //reading line by line
      for scanner.Scan() {
        checkStatus(scanner.Text())
      }

      if err := scanner.Err(); err != nil {
        log.Fatal(err)
        }

}
