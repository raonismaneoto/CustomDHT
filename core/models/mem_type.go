package models

type MemType int32

const (
	Mem MemType = iota
	Disk
	NotSupported
)

func GetMemTypeFromString(s string) MemType {
	switch s {
	case "Mem":
		return Mem
	case "Disk":
		return Disk
	default:
		return NotSupported
	}
}
