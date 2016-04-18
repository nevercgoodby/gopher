package main

type Bitmap struct {
	// 保存实际的 bit 数据
	data []byte
	// 指示该 Bitmap 的 bit 容量
	bitsize uint32
	// 该 Bitmap 被设置为 1 的最大位置（方便遍历）
	maxpos uint32
}

// SetBit 将 offset 位置的 bit 置为 value (0/1)
func (this *Bitmap) Init(size uint32) bool {
	if size*8 > 0xFFFFFFFF {
		return false
	}
	this.data = make([]byte, size)
	this.bitsize = size * 8
	return true
}

// SetBit 将 offset 位置的 bit 置为 value (0/1)
func (this *Bitmap) SetBit(offset uint32, value uint8) bool {
	index, pos := offset/8, offset%8

	if this.bitsize < offset {
		return false
	}

	if value == 0 {
		// &^ 清位
		this.data[index] &^= 0x01 << pos
	} else {
		this.data[index] |= 0x01 << pos

		// 记录曾经设置为 1 的最大位置
		if this.maxpos < offset {
			this.maxpos = offset
		}
	}

	return true
}

func (this *Bitmap) GetBit(offset uint32) (uint8, bool) {
	index, pos := offset/8, offset%8

	if this.bitsize < offset {
		return 0, false
	}
	return this.data[index] & 0x01 << pos, true

}
