package arrays

type singleArray[T any] struct {
	BaseArray[T]
}

func NewSingleArray[T any]() *singleArray[T] {
	return &singleArray[T]{
		BaseArray: BaseArray[T]{
			array: make([]T, 0),
		},
	}
}

func (sa *singleArray[T]) Add(item T) {
	sa.BaseArray.Add(sa.getNewCapacity(), item)
}

func (sa *singleArray[T]) Put(index int, item T) {
	sa.BaseArray.Put(sa.getNewCapacity(), index, item)
}

func (sa *singleArray[T]) getNewCapacity() int {
	return sa.Cap() + 1
}
