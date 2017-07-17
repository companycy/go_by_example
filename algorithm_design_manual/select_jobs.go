package main

import (
	"fmt"
)

var maxFee int

func calc_jobs() {
	total_fee, begin, end := jobs[0].fee, jobs[0].begin, jobs[0].end	
	for i := 0; i < len(jobs)-1; i++ {
		next_job := jobs[i+1]
		if end <= next_job.begin {
			end = next_job.end
			total_fee += next_job.fee
		}
	}
	return total_fee
}

func select_jobs() {
	if len(jobs) == 1 {
		return current_fee + jobs.fee
	}

	for i := 0; i < len(jobs); i++ {
		total_fee := calc_jobs(jobs[i:])
		if maxFee < newFee {
			maxFee = newFee
		}
	}
}

func main() {
	
}
