// Package minecraft is a level viewer/editor for the popular creative game minecraft
package minecraft // import "vimagination.zapto.org/minecraft"

// TransparentBlockList is a slice of the block ids that are transparent.
type TransparentBlockList []uint16

// Add is a convenience method for the transparent block list. It adds a new
// block id to the list, making sure to not add duplicates
func (t *TransparentBlockList) Add(blockID uint16) bool {
	for _, b := range *t {
		if b == blockID {
			return false
		}
	}
	*t = append(*t, blockID)
	return true
}

// Remove is a convenience method to remove a block id from the transparent
// block list
func (t *TransparentBlockList) Remove(blockID uint16) bool {
	for n, b := range *t {
		if b == blockID {
			lt := len(*t) - 1
			(*t)[n], (*t) = (*t)[lt], (*t)[:lt]
			return true
		}
	}
	return false
}

// LightBlockList is a map of block ids to the amount of light they give off
type LightBlockList map[uint16]uint8

// Add is a convenience method for the light block list. It adds a new block id
// to the list with its corresponding light level
func (l LightBlockList) Add(blockID uint16, light uint8) bool {
	toRet := true
	if _, ok := l[blockID]; ok {
		toRet = false
	}
	l[blockID] = light
	return toRet
}

// Remove is a convenience method to remove a block id from the light block
// list
func (l LightBlockList) Remove(blockID uint16) bool {
	if _, ok := l[blockID]; ok {
		delete(l, blockID)
		return true
	}
	return false
}

var (
	// TransparentBlocks is a slice of the block ids that are transparent.
	// This is used in lighting calculations and is user overrideable for custom
	// blocks
	TransparentBlocks = TransparentBlockList{0, 6, 18, 20, 26, 27, 28, 29, 30, 31, 33, 34, 37, 38, 39, 40, 50, 51, 52, 54, 55, 59, 63, 64, 65, 66, 69, 70, 71, 75, 76, 77, 78, 79, 81, 83, 85, 90, 92, 93, 94, 96, 102, 106, 107, 117, 118, 119, 120, 750}
	// LightBlocks is a map of block ids to the amount of light they give off
	LightBlocks = LightBlockList{
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
