package ginq

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

func (c collection[T]) Select() []T {
	return c.items
}
