package main

type ByArrivalTime []Task

func (s ByArrivalTime) Len() int {
	return len(s)
}
func (s ByArrivalTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByArrivalTime) Less(i, j int) bool {
	return s[i].ArrivalTime < s[j].ArrivalTime
}

type ByPriority []Task

func (s ByPriority) Len() int {
	return len(s)
}
func (s ByPriority) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByPriority) Less(i, j int) bool {
	if s[i].ArrivalTime < s[j].ArrivalTime {
		return true
	} else if s[i].ArrivalTime == s[j].ArrivalTime {
		if float64(s[i].Priority) < float64(s[j].Priority) {
			return true
		}
	}
	return false
}
