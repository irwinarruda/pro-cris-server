package utils

func IsOverlapping(start1, range1, start2, range2 int) bool {
	end1 := start1 + range1
	end2 := start2 + range2
	return start1 < end2 && end1 > start2
}
