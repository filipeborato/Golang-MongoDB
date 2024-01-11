package schedule

import (
	"fmt"
	"log"
	"os"
	"sport-test/controller"
)

func CronjobCommands() bool {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) >= 1 {
		commands := map[string]func(){
			"populate-news": controller.CronPopulate,
		}
		arg := os.Args[1]
		if _, found := commands[arg]; found {
			commands[arg]()
		}
		log.Print("INFO", fmt.Sprint("[schedule-job] ", arg))
		return true
	}
	return false
}
