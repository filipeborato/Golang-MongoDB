package main

import (
	"github.com/whiteshtef/clockwork"
	"sport-test/controller"
)

func main() {

	sched := clockwork.NewScheduler()

	sched.Schedule().Every(30).Minutes().Do(controller.PopulateInformation)

	sched.Run()
}
