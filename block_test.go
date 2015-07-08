package minecraft

import (
	"testing"
)

func TestEquality(t *testing.T) {
	testData := []Block{
		Block{BlockID: 14},
		Block{BlockID: 214},
		Block{BlockID: 792},
	}
	for i, aBlock := range testData {
		for j, bBlock := range testData {
			match := aBlock.EqualBlock(bBlock)
			sameBlock := (i == j)
			if sameBlock != aBlock.EqualBlock(bBlock) {
				if match {
					t.Errorf("Block %d matched block %d, expecting non-match", i, j)
				} else {
					t.Errorf("Block %d didn't match block %d, expecting match", i, j)
				}
			}
		}
	}
}
