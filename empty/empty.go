package empty

type Shan []int

func (s *Shan) Push(v int) {
	*s = append(*s,v)
}
