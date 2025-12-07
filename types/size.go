package types

type (
	SizeB  int64
	SizeKB int64
	SizeMB int64
	SizeGB int64
	SizeTB int64
	SizePB int64
)

func (s SizeB) ToKB() SizeKB {
	return SizeKB(s / 1024)
}

func (s SizeB) ToMB() SizeMB {
	return SizeMB(s / 1024 / 1024)
}

func (s SizeB) ToGB() SizeGB {
	return SizeGB(s / 1024 / 1024 / 1024)
}

func (s SizeB) ToTB() SizeTB {
	return SizeTB(s / 1024 / 1024 / 1024 / 1024)
}

func (s SizeB) ToPB() SizePB {
	return SizePB(s / 1024 / 1024 / 1024 / 1024 / 1024)
}

// SizeKB 转换为其他单位
func (s SizeKB) ToB() SizeB {
	return SizeB(s * 1024)
}

func (s SizeKB) ToMB() SizeMB {
	return SizeMB(s / 1024)
}

func (s SizeKB) ToGB() SizeGB {
	return SizeGB(s / 1024 / 1024)
}

func (s SizeKB) ToTB() SizeTB {
	return SizeTB(s / 1024 / 1024 / 1024)
}

func (s SizeKB) ToPB() SizePB {
	return SizePB(s / 1024 / 1024 / 1024 / 1024)
}

// SizeMB 转换为其他单位
func (s SizeMB) ToB() SizeB {
	return SizeB(s * 1024 * 1024)
}

func (s SizeMB) ToKB() SizeKB {
	return SizeKB(s * 1024)
}

func (s SizeMB) ToGB() SizeGB {
	return SizeGB(s / 1024)
}

func (s SizeMB) ToTB() SizeTB {
	return SizeTB(s / 1024 / 1024)
}

func (s SizeMB) ToPB() SizePB {
	return SizePB(s / 1024 / 1024 / 1024)
}

// SizeGB 转换为其他单位
func (s SizeGB) ToB() SizeB {
	return SizeB(s * 1024 * 1024 * 1024)
}

func (s SizeGB) ToKB() SizeKB {
	return SizeKB(s * 1024 * 1024)
}

func (s SizeGB) ToMB() SizeMB {
	return SizeMB(s * 1024)
}

func (s SizeGB) ToTB() SizeTB {
	return SizeTB(s / 1024)
}

func (s SizeGB) ToPB() SizePB {
	return SizePB(s / 1024 / 1024)
}

// SizeTB 转换为其他单位
func (s SizeTB) ToB() SizeB {
	return SizeB(s * 1024 * 1024 * 1024 * 1024)
}

func (s SizeTB) ToKB() SizeKB {
	return SizeKB(s * 1024 * 1024 * 1024)
}

func (s SizeTB) ToMB() SizeMB {
	return SizeMB(s * 1024 * 1024)
}

func (s SizeTB) ToGB() SizeGB {
	return SizeGB(s * 1024)
}

func (s SizeTB) ToPB() SizePB {
	return SizePB(s / 1024)
}

// SizePB 转换为其他单位
func (s SizePB) ToB() SizeB {
	return SizeB(s * 1024 * 1024 * 1024 * 1024 * 1024)
}

func (s SizePB) ToKB() SizeKB {
	return SizeKB(s * 1024 * 1024 * 1024 * 1024)
}

func (s SizePB) ToMB() SizeMB {
	return SizeMB(s * 1024 * 1024 * 1024)
}

func (s SizePB) ToGB() SizeGB {
	return SizeGB(s * 1024 * 1024)
}

func (s SizePB) ToTB() SizeTB {
	return SizeTB(s * 1024)
}