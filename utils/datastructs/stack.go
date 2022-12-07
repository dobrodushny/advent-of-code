package datastructs

type Stack[V any] []V

func (stack *Stack[V]) Push(val V) *Stack[V] {
	*stack = append(*stack, val)

	return stack
}

func (stack *Stack[V]) Pop() V {
	val := (*stack)[len((*stack))-1]
	*stack = (*stack)[:len((*stack))-1]

	return val
}

func (stack *Stack[V]) RPush(val V) *Stack[V] {
	*stack = append([]V{val}, *stack...)

	return stack
}

func (stack *Stack[V]) RPop() V {
	val := (*stack)[0]
	*stack = (*stack)[1:]

	return val
}

func (stack *Stack[V]) NRPush(val []V) *Stack[V] {
	*stack = append(val, *stack...)

	return stack
}

func (stack *Stack[V]) NRPop(n int) []V {
	val := append([]V{}, (*stack)[:n]...)
	*stack = (*stack)[n:]

	return val
}
