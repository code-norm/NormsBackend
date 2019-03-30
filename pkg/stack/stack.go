package stack

//ItemStack ...
type ItemStack struct {
	head *Item
	size int
}

//Item ...
type Item struct {
	text string
	next *Item
}

//Empty ...
func Empty(s *ItemStack) bool {
	return (s.size == 0)
}

//Size ...
func Size(s *ItemStack) int {
	return s.size
}
