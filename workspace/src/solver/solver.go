package main
/*
This is the main package of the program,
for its correct operation it is necessary to
redefine const path.
Please note that syscall is used in the package math
and you should also set the correct version math.dll
according to your system.
 */

import "C"
import (
	"./jobs"
	"./math"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

const path = "./jobs.json" // Define the path to the file with tasks
const numCPU = 4

type Result struct {
	Job, Output int
}

func solve(jb [] jobs.Job) [] Result {

	var ch = make(chan int, numCPU)
	var result [] Result

	for i := 0; i < len(jb); i++ {
		if i + numCPU <= len(jb) {
			for j := i; j < i + numCPU; j++ {
				go math.Div(jb[j].Arg1, jb[j].Arg2, j, ch) // j - id current job
			}
			for range [numCPU]int{} {
				result = append(result, Result{<-ch, <-ch})
			}
			i += numCPU
		} else {
			go math.Div(jb[i].Arg1, jb[i].Arg2, i, ch)
			result = append(result, Result{<-ch, <-ch})
		}
	}
	return result
}

func main() {
	jb, err := jobs.GetJobs(path)
	if err != nil {
		fmt.Println(err)
	}
	var result = solve(jb)
	var buf = new(bytes.Buffer)

	enc := json.NewEncoder(buf)
	err = enc.Encode(result)
	if err != nil {
		fmt.Println(err)
	}
	f, err := os.Create("solver_answer.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	io.Copy(f, buf)
}
