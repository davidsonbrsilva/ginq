package ginq

const (
	collectionIsEmpty = "collection is empty"
)

type collection[item comparable] struct {
	items []item
}

func From[T comparable](items []T) collection[T] {
	return collection[T]{items: items}
}

func (c collection[T]) Where(condition func(T) bool) collection[T] {
	filtered := make([]T, 0)

	for _, item := range c.items {
		if condition(item) {
			filtered = append(filtered, item)
		}
	}

	return collection[T]{items: filtered}
}

func (c collection[T]) Contains(element T) bool {
	for _, item := range c.items {
		if item == element {
			return true
		}
	}

	return false
}

func (c collection[T]) First() T {
	if len(c.items) == 0 {
		_default[T]()
	}

	return c.items[0]
}

func (c collection[T]) FirstOrPanic() T {
	if len(c.items) == 0 {
		panic(collectionIsEmpty)
	}

	return c.items[0]
}

func (c collection[T]) FirstWith(condition func(T) bool) T {
	for _, item := range c.items {
		if condition(item) {
			return item
		}
	}

	return _default[T]()
}

func (c collection[T]) Last() T {
	if len(c.items) == 0 {
		return _default[T]()
	}

	return c.items[len(c.items)-1]
}

func (c collection[T]) LastOrPanic() T {
	if len(c.items) == 0 {
		panic(collectionIsEmpty)
	}

	return c.items[len(c.items)-1]
}

func (c collection[T]) LastWith(condition func(T) bool) T {
	for index := len(c.items) - 1; index >= 0; index-- {
		if condition(c.items[index]) {
			return c.items[index]
		}
	}

	return _default[T]()
}

func (c collection[T]) Select() []T {
	return c.items
}

func _default[T comparable]() T {
	var element T
	return element
}
