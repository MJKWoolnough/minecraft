package nbt

import "testing"

func TestIndent(t *testing.T) {
	data := [...][2]string{
		{"", ""},
		{"\n", "\n	"},
		{"abc\ndef", "abc\n	def"},
		{"abc\ndef\n	ghi\n", "abc\n	def\n		ghi\n	"},
	}
	for _, d := range data {
		if out := indent(d[0]); out != d[1] {
			t.Errorf("ident mismatch with %q, expected %q, got %q", d[0], d[1], out)
		}
	}
}
