package models

type HelpMap struct {
	Key   int
	Value int
}

type HelpMapDouble struct {
	Key   int
	Value float64
}
type HelpMapDoubleList []HelpMapDouble
type HelpMapList []HelpMap

func (h HelpMapList) Less(i, j int) bool {
	return h[i].Value > h[j].Value
}
func (h HelpMapList) Len() int {
	return len(h)
}
func (h HelpMapList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h HelpMapDoubleList) Less(i, j int) bool {
	return h[i].Value > h[j].Value
}
func (h HelpMapDoubleList) Len() int {
	return len(h)
}
func (h HelpMapDoubleList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
