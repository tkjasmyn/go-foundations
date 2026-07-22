package algorithms

import "testing"

func TestMinStack(t *testing.T) {
    ms := &MinStack{}
    ms.Push(-2)
    ms.Push(0)
    ms.Push(-3)
    
    if got := ms.GetMin(); got != -3 {
        t.Errorf("GetMin() = %d, want -3", got)
    }
    
    ms.Pop()
    
    if got := ms.Top(); got != 0 {
        t.Errorf("Top() = %d, want 0", got)
    }
    
    if got := ms.GetMin(); got != -2 {
        t.Errorf("GetMin() = %d, want -2", got)
    }
}