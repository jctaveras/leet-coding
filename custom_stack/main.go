package main

type stackElement struct {
	Val  int
	next *stackElement
}

type customStack struct {
	maxSize     int
	currentSize int
	top         *stackElement
}

func constructor(maxSize int) customStack {
	return customStack{maxSize: maxSize, currentSize: 0, top: nil}
}

func (cs *customStack) push(x int) {
	if cs.currentSize >= cs.maxSize {
		return
	}

	if cs.top == nil {
		cs.top = &stackElement{Val: x, next: nil}
	} else {
		cs.top = &stackElement{Val: x, next: cs.top}
	}

	cs.currentSize++
	return
}

func (cs *customStack) pop() int {
	if cs.top == nil || cs.currentSize <= 0 {
		return -1
	}

	value := cs.top.Val
	cs.top = cs.top.next
	cs.currentSize--

	return value
}

func (cs *customStack) increment(k int, val int) {
	if cs.currentSize < k && cs.top != nil {
		tempItem := cs.top
		for tempItem.next != nil {
			tempItem.Val += val
			tempItem = tempItem.next
		}

		tempItem.Val += val
	} else if cs.top != nil {
		tempItem := cs.top

		for index := k; index < cs.currentSize; index++ {
			tempItem = tempItem.next
		}

		for tempItem.next != nil {
			tempItem.Val += val
			tempItem = tempItem.next
		}

		tempItem.Val += val
	}

	return
}

func main() {
	stack := constructor(12)

	stack.push(83)
	stack.increment(2, 60)
	stack.increment(9, 61)
	stack.increment(1, 60)
	stack.push(82)
	stack.push(21)
	stack.push(58)
	stack.increment(8, 8)
	stack.push(22)
	stack.push(80)
	stack.increment(1, 64)
	stack.pop()
	stack.pop()
	stack.push(24)
}
