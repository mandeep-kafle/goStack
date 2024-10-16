package stack

type Stack struct {
	arr []interface{}
}

func (s *Stack) Len() interface{}     { return len(s.arr) }
func (s *Stack) Push(val interface{}) { s.arr = append(s.arr, val) }

func (s *Stack) Pop() {
	s.arr = s.arr[:len(s.arr)-1]
	return
}
func (s *Stack) Top() interface{} { return s.arr[len(s.arr)-1] }
func NewStack() *Stack {
	return &Stack{arr: make([]interface{}, 0)}
}
