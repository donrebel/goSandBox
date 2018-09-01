package datastructures

//StreamAttributes ...
type StreamAttributes struct {
	Cnt    int
	Sum    int
	Max    int
	Min    int
	Avg    int
	Range  int
	Mode   int
	Median int
}

//ModCount ...
type ModCount struct {
	m    map[int]int
	mode int
}

func newModCount() *ModCount {
	mc := &ModCount{}
	mc.m = make(map[int]int)
	return mc
}

func (mc *ModCount) insert(x int) {
	(mc.m[x])++
}

//Insert ...
func (sa *StreamAttributes) Insert(x int) {
	sa.Cnt++
	sa.Sum += x
	sa.Avg = sa.Sum / sa.Cnt
	if x > sa.Max {
		sa.Max = x
	}
	if x < sa.Min {
		sa.Min = x
	}
	sa.Range = sa.Max - sa.Min

}
