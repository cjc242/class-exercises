package main

type Process struct {
	clock []int
	pid   int
	N     int
}

func (p *Process) Start(N int, PID int) {
	p.clock = make([]int, N)
	p.N = N
	p.pid = PID
}

// increments process' own vector clock
func (p *Process) Internal() {
	p.clock[p.pid]++
}

// increments vector timestamp, then copies it and returns that copy
func (p *Process) Send() []int {
	p.clock[p.pid]++
	snapClock := make([]int, len(p.clock))
	copy(snapClock, p.clock)
	return snapClock
}

// find max between timestamp and current process for each i then puts that into p.clock[i], then we just increment pid of clock.
// so basically merging the received timestamp into the process' own clock.
func (p *Process) Receive(ts []int) {
	for i := 0; i < p.N; i++ {
		p.clock[i] = max(p.clock[i], ts[i])
	}
	p.clock[p.pid]++
}

// To compare properly we use a for loop to get each element and compare elements from each ts, then using less and greater boolean variables
// we can decide if ts1 is less than, greater than, or concurrent as a whole compared to ts2
func Compare(ts1 []int, ts2 []int) int {
	less := false
	greater := false

	for i := 0; i < len(ts1); i++ {
		if ts1[i] < ts2[i] {
			less = true
		}
		if ts1[i] > ts2[i] {
			greater = true
		}
	}
	if less == greater {
		return 0
	}
	if less && greater == false {
		return -1
	} else {
		return 1
	}
}
