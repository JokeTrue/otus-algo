package arrays

type factorArray[T any] struct {
	factor int
	BaseArray[T]
}

func NewFactorArray[T any](factor int) *factorArray[T] {
	return &factorArray[T]{
		factor: factor,
		BaseArray: BaseArray[T]{
			array: make([]T, 0),
		},
	}
}

func (fa *factorArray[T]) Add(item T) {
	fa.BaseArray.Add(fa.getNewCapacity(), item)
}

func (fa *factorArray[T]) Put(index int, item T) {
	fa.BaseArray.Put(fa.getNewCapacity(), index, item)
}

func (fa *factorArray[T]) getNewCapacity() int {
	if fa.Len() < fa.Cap() {
		return fa.Cap()
	}
	return fa.Cap()*fa.factor + 1
}
