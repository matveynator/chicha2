package encoding

// additional types
// sorted from the smallest to the largest
const (
	NullValue  byte = 0x05
	FalseValue byte = 0x10 // 0x10 - 0x00 = 0x10 = 16
	TrueValue  byte = 0x11 // 0x11 - 0x10 = 0x01 = 1
	Int64Value byte = 0x20 // 0x20 - 0x11 = 0x0f = 15
	Int32Value byte = 0x21 // 0x21 - 0x20 = 0x01 = 1
	Int16Value byte = 0x22 // 0x22 - 0x21 = 0x01 = 1
	Int8Value  byte = 0x23 // 0x23 - 0x22 = 0x01 = 1
	// until -32
	IntSmallValue byte = 0x24 // 0x24 - 0x23 = 0x01 = 1
	// until 127
	Uint8Value    byte = 0xC4 // 0xC3 - 0x23 = 0xA0 = 160
	Uint16Value   byte = 0xC5 // 0xC4 - 0xC3 = 0x01 = 1
	Uint32Value   byte = 0xC6 // 0xC5 - 0xC4 = 0x01 = 1
	Uint64Value   byte = 0xC7 // 0xC6 - 0xC5 = 0x01 = 1
	Float64Value  byte = 0xD0 // 0xD0 - 0xC6 = 0x0a = 10
	Float32Value  byte = 0xD1 // 0xD1 - 0xD0 = 0x01 = 1 | not included in keys
	TextValue     byte = 0xDA // 0xDA - 0xD1 = 0x09 = 9
	BlobValue     byte = 0xE0 // 0xE0 - 0xDA = 0x06 = 6
	ArrayValue    byte = 0xE6 // 0xE6 - 0xE0 = 0x06 = 6
	DocumentValue byte = 0xF0 // 0xF0 - 0xE6 = 0x0e = 14
)