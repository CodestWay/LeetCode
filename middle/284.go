package middle

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
}

type PeekingIterator struct {
	index int
	data  []int
}

func Constructor(iter *Iterator) *PeekingIterator {
	p := &PeekingIterator{index: 0}
	for iter.hasNext() {
		p.data = append(p.data, iter.next())
	}
	return p
}

func (this *PeekingIterator) hasNext() bool {
	return len(this.data) > this.index
}

func (this *PeekingIterator) next() int {
	if this.index > len(this.data)-1 {
		return 0
	}
	val := this.data[this.index]
	this.index++
	return val
}

func (this *PeekingIterator) peek() int {
	return this.data[this.index]
}
