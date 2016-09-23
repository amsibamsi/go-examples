package main

type List struct {
	Value int
	Next  *List
}

func New(values []int, loop int) *List {
	if len(values) < 1 {
		return nil
	}
	var first, last *List
	for i, v := range values {
		if i == 0 {
			last = &List{v, nil}
			first = last
		} else {
			last.Next = &List{v, nil}
			last = last.Next
		}
	}
	last.Next = first.Get(loop)
	return first
}

func (l *List) Get(n int) *List {
	if n < 0 {
		return nil
	}
	e := l
	for i := n; i > 0; i-- {
		if e == nil {
			return nil
		}
		e = e.Next
	}
	return e
}

func (l *List) LoopCheckAll() bool {
	current := l
	count := 0
	var check *List
	for current != nil {
		check = l
		for i := count; i > 0; i-- {
			if current == check {
				return true
			}
			check = check.Next
		}
		current = current.Next
		count += 1
	}
	return false
}

func (l *List) LoopFast() bool {
	if l == nil {
		return false
	}
	slow, fast := l, l
	for {
		slow = slow.Next
		if slow == nil {
			return false
		}
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
		if fast == nil {
			return false
		}
		if fast == slow {
			return true
		}
	}
}
