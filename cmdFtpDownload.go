package main

import (
	"github.com/jlaffaye/ftp"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 4 {
		print("USAGE: cmdFtpDownload <url:port> <login> <password> <filepath>\n" +
			"Example: cmdFtpDownload ftp.some.ru:21 mylogin mypassword \\foo\\bar.txt\n" +
			"Exactly 4 args expected")
	} else {
		print(argsWithoutProg)
		c, err := ftp.Dial(argsWithoutProg[0], ftp.DialWithTimeout(5*time.Second))
		if err != nil {
			log.Fatal(err)
		}

		err = c.Login(argsWithoutProg[1], argsWithoutProg[2])
		if err != nil {
			log.Fatal(err)
		}

		r, err := c.Retr(argsWithoutProg[3])
		if err != nil {
			panic(err)
		}
		defer r.Close()

		buf, err := io.ReadAll(r)
		println(string(buf))

		if err := c.Quit(); err != nil {
			log.Fatal(err)
		}
	}
}
