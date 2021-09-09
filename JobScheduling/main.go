package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	jobs := Jobs{}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 50; i++ {
		jobs = append(jobs, Task{Name: fmt.Sprintf("%d", i), ArrivalTime: float64(r1.Intn(100)), Priority: r1.Intn(100)})
	}
	jobs.ScheduleUsing(LIFO, 1)
}
