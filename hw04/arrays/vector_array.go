package arrays

type vectorArray[T any] struct {
	vector int
	BaseArray[T]
}

func NewVectorArray[T any](vector int) *vectorArray[T] {
	return &vectorArray[T]{
		vector: vector,
		BaseArray: BaseArray[T]{
			array: make([]T, 0),
		},
	}
}

func (va *vectorArray[T]) Add(item T) {
	va.BaseArray.Add(va.getNewCapacity(), item)
}

func (va *vectorArray[T]) Put(index int, item T) {
	va.BaseArray.Put(va.getNewCapacity(), index, item)
}

func (va *vectorArray[T]) getNewCapacity() int {
	if va.Len() < va.Cap() {
		return va.Cap()
	}
	return va.Cap() + va.vector
}
