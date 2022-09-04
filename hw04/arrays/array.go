package arrays

type Array[T any] interface {
	Add(item T)
	Get(index int) T
	Remove(index int) T
	Put(index int, item T)
}

type BaseArray[T any] struct {
	array []T
}

func (ba *BaseArray[T]) Len() int {
	return len(ba.array)
}

func (ba *BaseArray[T]) Cap() int {
	return cap(ba.array)
}

func (ba *BaseArray[T]) Get(index int) T {
	return ba.array[index]
}

func (ba *BaseArray[T]) Add(newCapacity int, item T) {
	if ba.Len() < ba.Cap() {
		ba.array = append(ba.array, item)
		return
	}

	newArray := make([]T, 0, newCapacity)
	newArray = append(newArray, ba.array...) // Copy all elements of existing array
	newArray = append(newArray, item)        // Add new element to the end of array

	ba.array = newArray
}

func (ba *BaseArray[T]) Put(newCapacity int, index int, item T) {
	if ba.Len() < ba.Cap() {
		ba.array = append(ba.array[:index+1], ba.array[index:]...)
		ba.array[index] = item
		return
	}

	newArray := make([]T, 0, newCapacity)
	newArray = append(newArray, ba.array[:index]...) // Copy all elements before index
	newArray = append(newArray, item)                // Add new element at index
	newArray = append(newArray, ba.array[index:]...) // Copy all elements after index

	ba.array = newArray
}

func (ba *BaseArray[T]) Remove(index int) T {
	removedValue := ba.Get(index)

	newArray := make([]T, 0, ba.Cap()-1)
	newArray = append(newArray, ba.array[:index]...)   // Copy all elements before index
	newArray = append(newArray, ba.array[index+1:]...) // Copy all elements after index (without element)

	ba.array = newArray
	return removedValue
}
