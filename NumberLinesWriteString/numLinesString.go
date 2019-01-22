package main

type LineString struct {
	widths []int
	s      string
}

func InitLineString(width []int, s string) *LineString {
	return &LineString{
		widths: width,
		s:      s,
	}
}

func (ls *LineString) Calculate() (int, int) {
	var totallength int

	for _, char := range ls.s {
		totallength += ls.widths[int(char)-97]
	}
	return (totallength / 100) + 1, totallength % 100
}
