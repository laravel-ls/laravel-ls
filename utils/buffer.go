package utils

// Buffer is just a wrapper around a byte slice
// that provides some nice helper functions.
type Buffer []byte

// Update part of buffer denoted by start, end
func (b *Buffer) Update(start, end uint, src []byte) {
	*b = append((*b)[:start], append(src, (*b)[end:]...)...)
}
