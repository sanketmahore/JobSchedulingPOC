package main

func main() {
	jobs := Jobs{{Name: "3", ArrivalTime: 2, Priority: 1},
		{Name: "2", ArrivalTime: 1, Priority: 2},
		{Name: "1", ArrivalTime: 0, Priority: 3},
		{Name: "4", ArrivalTime: 3, Priority: 4},
		{Name: "5", ArrivalTime: 4, Priority: 5},
		{Name: "6", ArrivalTime: 5, Priority: 6}}
	jobs.ScheduleUsing(FIFO, 1)
}
