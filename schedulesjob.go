package main

import (
	"github.com/whiteshtef/clockwork"
)

func main() {

	sched := clockwork.NewScheduler()

	//sched.Schedule().Every(30).Minutes().Do(controller.PopulateInformation(c))

	sched.Run()
}