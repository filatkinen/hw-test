package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v any) *ListItem
	PushBack(v any) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value any
	Next  *ListItem
	Prev  *ListItem
	Key   Key
}

type list struct {
	front, back *ListItem
	len         int
}

func NewList() List {
	return new(list)
}

func (s *list) Len() int {
	return s.len
}

func (s *list) Front() *ListItem {
	return s.front
}

func (s *list) Back() *ListItem {
	return s.back
}

func (s *list) PushFront(v any) *ListItem {
	newItem := new(ListItem)
	if s.front == nil {
		s.front = newItem
		s.back = newItem
	} else {
		s.front.Prev = newItem
		newItem.Next = s.front
		s.front = newItem
	}
	newItem.Value = v
	s.len++
	return s.front
}

func (s *list) PushBack(v any) *ListItem {
	newItem := new(ListItem)
	if s.back == nil {
		s.front = newItem
		s.back = newItem
	} else {
		s.back.Next = newItem
		newItem.Prev = s.back
		s.back = newItem
	}
	newItem.Value = v
	s.len++
	return s.back
}

func (s *list) Remove(i *ListItem) {
	switch s.len {
	case 0:
		// The list is empty: doing nothing
	case 1:
		// The list has one element
		s.front = nil
		s.back = nil
		s.len--
	default:
		// The list is inside
		switch {
		case i != s.front && i != s.back:
			i.Prev.Next = i.Next
			i.Next.Prev = i.Prev
			i.Next = nil
			i.Prev = nil
		case i == s.front:
			s.front = i.Next
			i.Next.Prev = nil
			i.Next = nil
		case i == s.back:
			s.back = i.Prev
			i.Prev.Next = nil
			i.Prev = nil
		}
		s.len--
	}
}

func (s *list) MoveToFront(i *ListItem) {
	if s.len > 1 && i != s.front {
		if i == s.back {
			s.back = i.Prev
			s.back.Next = nil
		} else {
			i.Prev.Next = i.Next
			i.Next.Prev = i.Prev
		}
		i.Prev = nil
		i.Next = s.front
		s.front.Prev = i
		s.front = i
	}
}
