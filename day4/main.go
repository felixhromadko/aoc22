package main

import (
	"fmt"
	set "github.com/deckarep/golang-set/v2"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type PairOfJob struct {
	JobA Job
	JobB Job
}

type Job struct {
	TaskIDs set.Set[int]
}

func main() {
	inputRaw, err := os.ReadFile("./inf.txt")
	checkErr(err)

	var jobPairs []PairOfJob
	for _, s := range strings.Split(string(inputRaw), "\n") {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}

		// each line has two jobs - split by ,
		jobs := strings.Split(s, ",")
		jobA := ParseJob(jobs[0])
		jobB := ParseJob(jobs[1])

		jobPairs = append(jobPairs, PairOfJob{
			jobA, jobB,
		})
	}

	// find number of jobPairs where one job is fully contained in the other
	var subsetCount int
	for _, pair := range jobPairs {
		if pair.JobA.TaskIDs.IsSubset(pair.JobB.TaskIDs) ||
			pair.JobB.TaskIDs.IsSubset(pair.JobA.TaskIDs) {
			subsetCount++
		}
	}

	// find number of jobPairs where one job overlaps with the other
	var overlapCount int
	for _, pair := range jobPairs {
		if pair.JobA.TaskIDs.Intersect(pair.JobB.TaskIDs).Cardinality() != 0 {
			overlapCount++
		}
	}

	fmt.Printf("Found %d subset jobs\n", subsetCount)
	fmt.Printf("Found %d overlapping jobs\n", overlapCount)
}

func ParseJob(jobS string) Job {
	taskIDs := set.NewSet[int]()
	// each job has two numbers: lower - upper bound
	bounds := strings.Split(jobS, "-")
	lower, err := strconv.Atoi(bounds[0])
	checkErr(err)

	upper, err := strconv.Atoi(bounds[1])
	checkErr(err)

	for i := lower; i <= upper; i++ {
		taskIDs.Add(i)
	}

	return Job{taskIDs}
}
