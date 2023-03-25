package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

const (
	markName  string = "GOLANG_CLI_REMINDER"
	markValue string = "1"
)

func main() {
	if len(os.Args) < 3 {
		log.Printf("Usage: %s <hh:mm> <text message>\n", os.Args[0])
	}

	now := time.Now()

	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	t, err := w.Parse(os.Args[1], now)

	if err != nil {
		log.Println(err.Error())
	}

	if t == nil {
		log.Println("Unable to parse time!")
		os.Exit(2)
	}

	if now.After(t.Time) {
		log.Println("Set a future time")
		os.Exit(3)
	}

	diff := t.Time.Sub(now)

	if os.Getenv(markName) == markValue {
		time.Sleep(diff)
		err := beeep.Alert("Reminder", strings.Join(os.Args[2:], " "), "")

		if err != nil {
			log.Println(err.Error())
			os.Exit(4)
		}
	} else {
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue))
		if err := cmd.Start(); err != nil {
			log.Println(err.Error())
			os.Exit(5)
		}
		log.Println("Reminder will be displayed after:", diff.Round(time.Second), "seconds.")
		os.Exit(0)
	}
}
