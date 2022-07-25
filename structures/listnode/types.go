package listnode

type I[T comparable] interface {
	Value() T
	Next() I[T]
}

// G is Generic Linked List Node
type G[T comparable] struct {
	I[T]
	V T
	N *G[T]
}

func NewG[T comparable](values ...T) (head *G[T]) {
	if len(values) == 0 {
		return nil
	}
	head = &G[T]{
		V: values[0],
		N: nil,
	}
	current := head
	for i := 1; i < len(values); i++ {
		current.N = &G[T]{
			V: values[i],
			N: nil,
		}
		current = current.N
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
