package SwordToOffer

type CQueue struct {
	left  []int
	right []int
}

func Constructor() CQueue {
	return CQueue{
		left:  make([]int, 0),
		right: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.left = append(this.left, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.right) == 0 {
		if len(this.left) == 0 {
			return -1
		} else {
			for len(this.left) > 0 {
				this.right = append(this.right, this.left[len(this.left)-1])
				this.left = this.left[:len(this.left)-1]
			}
		}
	}
	val := this.right[len(this.right)-1]
	this.right = this.right[:len(this.right)-1]
	return val
}
