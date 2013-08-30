package minecraft

import (
	"testing"
)

func TestEquality(t *testing.T) {
	testData := []Block{
		Block{BlockId: 14},
		Block{BlockId: 214},
		Block{BlockId: 792},
	}
	for i, aBlock := range testData {
		for j, bBlock := range testData {
			match := aBlock.Equal(bBlock)
			pmatch := aBlock.Equal(&bBlock)
			if match != pmatch {
				t.Errorf("Block %d didn't match with pointer and non-pointer arguments", j)
			}
			sameBlock := (i == j)
			if sameBlock != match {
				if match {
					t.Errorf("Block %d matched block %d, expecting non-match", i, j)
				} else {
					t.Errorf("Block %d didn't match block %d, expecting match", i, j)
				}
			}
		}
	}
}
