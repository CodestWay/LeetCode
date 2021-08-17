package SwordToOffer

type MinStack struct {
	left  []int
	right []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		left:  make([]int, 0),
		right: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.left = append(this.left, x)
	if len(this.right) == 0 {
		this.right = append(this.right, x)
	} else {
		if x <= this.right[len(this.right)-1] {
			this.right = append(this.right, x)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.left) == 0 {
		return
	}
	pop := this.left[len(this.left)-1]
	this.left = this.left[:len(this.left)-1]
	if pop == this.right[len(this.right)-1] {
		this.right = this.right[:len(this.right)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.left) == 0 {
		return -1
	}
	return this.left[len(this.left)-1]
}

func (this *MinStack) Min() int {
	if len(this.right) == 0 {
		return -1
	}
	return this.right[len(this.right)-1]
}
