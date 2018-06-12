package minecraft_test

import (
	"fmt"

	"vimagination.zapto.org/minecraft"
)

func Example() {
	path := minecraft.NewMemPath()
	level, _ := minecraft.NewLevel(path)
	level.LevelName("TestMine")
	name := level.GetLevelName()
	fmt.Println(name)
	// Output:
	// TestMine
}
