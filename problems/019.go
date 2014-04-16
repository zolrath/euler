/*
Problem 19: Counting Sundays
You are given the following information, but you may prefer to do some research for yourself.
1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
Answer: 171
*/

package problems

import "time"

const ANSWER_019 = 171

func Euler019() int {
	// monthDays := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	// This method is a bit too easy/slow.
	// TODO: Actually do math to solve this.

	months := []time.Month{time.January, time.February, time.March,
		time.April, time.May, time.June, time.July, time.August,
		time.September, time.October, time.November, time.December}
	var date time.Time
	sundays := 0

	for year := 1901; year < 2001; year++ {
		for month := 0; month < 12; month++ {
			date = time.Date(year, months[month], 1, 12, 0, 0, 0, time.UTC)
			if date.Weekday() == time.Sunday {
				sundays++
			}
		}
	}
	return sundays
}
