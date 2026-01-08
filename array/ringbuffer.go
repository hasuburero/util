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

func (self *RingBuffer) Push(content BufferContent) BufferContent {
	self.Content[self.Tail] = content
	self.Tail = (self.Tail + 1) % (self.Size + 1)
	if self.Tail == self.Head {
		buf := self.Content[self.Head]
		self.Head = (self.Head + 1) % (self.Size + 1)
		return buf
	} else {
		self.Length += 1
		return nil
	}
}

func (self *RingBuffer) Get() []BufferContent {
	var bc []BufferContent
	for i := self.Head; i != self.Tail; i = (i + 1) % (self.Size + 1) {
		//if i == self.Tail {
		//	break
		//}
		bc = append(bc, self.Content[i])
	}

	return bc
}
