package jobs

/*
Processing the validity of the
incoming file as well as decoding it
into struct Job.
*/

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Job struct {
	Arg1, Arg2 int
}

func GetJobs(path string) ([] Job, error)  {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	} else {
		if json.Valid(data) {
			var jobs [] Job
			err = json.Unmarshal(data, &jobs)
			return jobs, err
		} else {
			err = errors.New("not valid json in file")
			return nil, err
		}
	}
}