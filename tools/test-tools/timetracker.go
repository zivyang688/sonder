package test_tools

import (
	"log"
	"time"
)

func TimeTrack(start time.Time) {
	elasped := time.Since(start)
	log.Printf("Execution time: %s", elasped)
}

func ElaspedTimeTrack(f func()) {
	defer TimeTrack(time.Now())
	f()
}
