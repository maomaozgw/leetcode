package listnode

type I[T comparable] interface {
	Value() T
	Next() I[T]
}

// G is Generic Linked List Node
type G[T comparable] struct {
	Val  T
	Next *G[T]
}

func NewG[T comparable](values ...T) (head *G[T]) {
	if len(values) == 0 {
		return nil
	}
	head = &G[T]{
		Val:  values[0],
		Next: nil,
	}
	current := head
	for i := 1; i < len(values); i++ {
		current.Next = &G[T]{
			Val:  values[i],
			Next: nil,
		}
		current = current.Next
	}
	return head
}

// D is Double Linked List Node
type D[T comparable] struct {
	V T
	N *D[T]
	P *D[T]
}

func NewD[T comparable](values ...T) (head *D[T]) {
	if len(values) == 0 {
		return nil
	}
	head = &D[T]{
		V: values[0],
		P: nil,
	}
	current := head
	for i := 1; i < len(values); i++ {
		current.N = &D[T]{
			V: values[0],
			P: current,
		}
		current = current.N
	}
	return head
}
