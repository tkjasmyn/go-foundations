package algorithms

type MinStack struct {
	stack []int
	minstack []int
}

func (m *MinStack) Push(val int)  {
	m.stack = append(m.stack, val)

	if len(m.minstack) == 0 || val <= m.minstack[len(m.minstack)-1] {
		m.minstack = append(m.minstack, val)
	}
}

func (m *MinStack) Pop()  {
	if m.stack[len(m.stack)-1] == m.minstack[len(m.minstack)-1] {
		m.stack = m.stack[:len(m.stack)-1]
		m.minstack = m.minstack[:len(m.minstack)-1]
	} else {
		m.stack = m.stack[:len(m.stack)-1]
	}
}

func (m *MinStack) Top() int {
	return m.stack[len(m.stack)-1]
}

func (m *MinStack) GetMin() int {
	return m.minstack[len(m.minstack)-1]
}