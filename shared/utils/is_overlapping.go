package utils

import "time"

func IsOverlappingInt(start1, range1, start2, range2 int) bool {
	end1 := start1 + range1
	end2 := start2 + range2
	return start1 < end2 && end1 > start2
}

func IsOverlappingDate(start1, end1, start2, end2 time.Time) bool {
	return start1.Before(end2) && end1.After(start2)
}
