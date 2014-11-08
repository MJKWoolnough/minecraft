// Package Minecraft will be a full featured minecraft level editor/viewer.
package minecraft

type TransparentBlockList []uint16

func (t *TransparentBlockList) Add(blockId uint16) bool {
	for _, b := range *t {
		if b == blockId {
			return false
		}
	}
	*t = append(*t, blockId)
	return true
}

func (t *TransparentBlockList) Remove(blockId uint16) bool {
	for n, b := range *t {
		if b == blockId {
			lt := len(*t) - 1
			(*t)[n], (*t) = (*t)[lt], (*t)[:lt]
			return true
		}
	}
	return false
}

type LightBlockList map[uint16]uint8

func (l LightBlockList) Add(blockId uint16, light uint8) bool {
	toRet := true
	if _, ok := l[blockId]; ok {
		toRet = false
	}
	l[blockId] = light
	return toRet
}

func (l LightBlockList) Remove(blockId uint16) bool {
	if _, ok := l[blockId]; ok {
		delete(l, blockId)
		return true
	}
	return false
}

var (
	TransparentBlocks = TransparentBlockList{0, 6, 18, 20, 26, 27, 28, 29, 30, 31, 33, 34, 37, 38, 39, 40, 50, 51, 52, 54, 55, 59, 63, 64, 65, 66, 69, 70, 71, 75, 76, 77, 78, 79, 81, 83, 85, 90, 92, 93, 94, 96, 102, 106, 107, 117, 118, 119, 120, 750}
	LightBlocks       = LightBlockList{
		10:  15,
		11:  15,
		39:  1,
		50:  14,
		51:  15,
		62:  13,
		74:  13,
		76:  7,
		89:  15,
		90:  11,
		91:  15,
		94:  9,
		117: 1,
		119: 15,
		120: 1,
		122: 1,
		124: 15,
		130: 7,
		138: 15,
	}
)
