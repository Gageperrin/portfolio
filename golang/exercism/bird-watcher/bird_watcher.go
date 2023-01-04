package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	n := 0
	for _, v := range birdsPerDay {
		n += v
	}
	return n
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	n := 0
	var r int

	if len(birdsPerDay) - 1 < (week * 7 - 1) {
		r = len(birdsPerDay) - 1
	} else {
		r = (week * 7 - 1)
	}

	for i := (week * 7 - 7); i <= r; i++ {
		n += birdsPerDay[i]
	}
	return n
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}
	return birdsPerDay
}
