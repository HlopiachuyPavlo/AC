package main

import (
	"fmt"
	"regexp"
	"runtime"
	"time"
)

var funcTime map[string]time.Duration
var programStart time.Time

func startTrack() {
	funcTime = make(map[string]time.Duration)
	programStart = time.Now()
	fmt.Println("Strated")
}

func finishTrack() {
	fmt.Println("Finished")
	fmt.Println("Time: ", time.Since(programStart))
	fmt.Println(funcTime)
}

func FuncTimeTrack(start time.Time) {
	funcTime = make(map[string]time.Duration)
	programStart = time.Now()
	elapsed := time.Since(start)
	pc, _, _, _ := runtime.Caller(1)
	funcObj := runtime.FuncForPC(pc)
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")
	funcTime[name] += elapsed
	fmt.Println(fmt.Sprintf("%s took %s", name, elapsed))
}
