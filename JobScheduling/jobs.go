package main

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

const (
	FIFO          int = 1
	LIFO          int = 2
	PriorityQueue int = 3
	TimeBased     int = 4
	TokenBased    int = 5
)

var JobList chan Task

type Jobs []Task

func init() {
	JobList = make(chan Task, 50)
}

func (jobs Jobs) ScheduleUsing(algo int, routines int) error {
	switch algo {
	case FIFO:
		sort.Sort(ByArrivalTime(jobs))
		for _, job := range jobs {
			JobList <- job
		}
		close(JobList)
	case LIFO:
		sort.Sort(ByArrivalTime(jobs))
		for i := len(jobs) - 1; i >= 0; i-- {
			JobList <- jobs[i]
		}
		close(JobList)
	case PriorityQueue:
		sort.Sort(ByPriority(jobs))
		for _, job := range jobs {
			JobList <- job
		}
		close(JobList)
	default:
		return errors.New("invalid algorithm option")
	}
	// for j := range JobList {
	// 	fmt.Println(j)
	// }
	Execute(JobList, routines)
	return nil
}

func Execute(jobList <-chan Task, routines int) {
	var w sync.WaitGroup
	w.Add(routines)
	for i := 1; i <= routines; i++ {
		go func(i int, ci <-chan Task) {
			for v := range ci {
				time.Sleep(time.Millisecond)
				fmt.Printf("%d executing %v job\n", i, v)
			}
			w.Done()
		}(i, jobList)
	}
	w.Wait()

}
