package array

import ()

type BufferContent interface{}
type RingBuffer struct {
	Content []BufferContent
	Length  int
	Size    int
	Head    int
	Tail    int
}

func MakeRingBuffer(size int) *RingBuffer {
	rb := new(RingBuffer)
	rb.Size = size
	rb.Content = make([]BufferContent, size+1, size+1)

	return rb
}

func (self *RingBuffer) Push(content BufferContent) {
	self.Content[self.Tail] = content
	self.Tail = (self.Tail + 1) % (self.Size + 1)
	if self.Tail == self.Head {
		self.Head = (self.Head + 1) % (self.Size + 1)
	} else {
		self.Length += 1
	}
}

func (self *RingBuffer) Get() []BufferContent {
	bc := make([]BufferContent, 0, self.Size)
	for i := self.Head; i != self.Tail; i = (i + 1) % (self.Size + 1) {
		bc = append(bc, self.Content[i])
	}

	return bc
}
