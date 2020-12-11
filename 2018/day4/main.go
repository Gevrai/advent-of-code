package main

import (
	"regexp"
	"sort"
	"strings"
	"time"

	"advent-of-code-2018/utils"
)

func main() {
	utils.DownloadDayInput(2018, 4, false)
	input := utils.ReadInputFileRelativeSplitNewline()

	var events []Event
	for _, l := range input {
		events = append(events, ParseEvent(l))
	}
	sort.Slice(events, func(i, j int) bool { return events[i].timestamp.Before(events[j].timestamp) })

	guards := make(map[int][]TimeSlice)

	reg := regexp.MustCompile("Guard #([0-9]+) begins shift")

	currentGuard := 0
	currentSlice := TimeSlice{}
	for _, e := range events {
		switch {
		case strings.Contains(e.event, "begins shift"):
			m := reg.FindAllStringSubmatch(e.event, 1)
			currentGuard = utils.ParseInt(m[0][1], 10)
			currentSlice = TimeSlice{}
		case strings.Contains(e.event, "falls asleep"):
			currentSlice.start = e.timestamp
		case strings.Contains(e.event, "wakes up"):
			currentSlice.end = e.timestamp
			guards[currentGuard] = append(guards[currentGuard], currentSlice)
		}
	}

	var maxID int
	var maxAsleep time.Duration
	for id, slice := range guards {
		var asleep time.Duration
		for _, t := range slice {
			asleep += t.end.Sub(t.start)
		}
		if maxAsleep < asleep {
			maxID = id
			maxAsleep = asleep
		}
	}

	minutes := make([]int, 60)
	for _, slice := range guards[maxID] {
		for i := slice.start.Minute(); i < slice.end.Minute(); i++ {
			minutes[i]++
		}
	}
	var max int
	for i := range minutes {
		if minutes[i] > minutes[max] {
			max = i
		}
	}

	println("Part 1:", maxID*max)

	var maxMinute, maxIndex int
	for id, slices := range guards {
		minutes := make([]int, 60)
		for _, slice := range slices {
			for i := slice.start.Minute(); i < slice.end.Minute(); i++ {
				minutes[i]++
			}
		}
		for i := range minutes {
			if minutes[i] > maxMinute {
				maxID = id
				maxMinute = minutes[i]
				maxIndex = i
			}
		}
	}

	println("Part 2:", maxIndex*maxID)
}

type Event struct {
	timestamp time.Time
	event     string
}

type TimeSlice struct {
	start, end time.Time
}

func ParseEvent(input string) Event {

	parts := strings.Split(input, "]")

	const timeFormat = "2006-01-02 15:04"
	t, err := time.Parse(timeFormat, strings.Trim(parts[0], "[]"))
	utils.PanicIfError(err)

	return Event{
		event:     strings.TrimSpace(parts[1]),
		timestamp: t,
	}
}
