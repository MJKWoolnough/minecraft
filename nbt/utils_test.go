package nbt

import "testing"

func TestTagData(t *testing.T) {
	var (
		err error
		d   Data
	)
	for id, name := range tagIdNames {
		d, err = newFromTag(TagId(id))
		if id == 0 {
			if err == nil {
				t.Errorf("tag end (0) should return an error")
			}
		} else if err != nil {
			t.Errorf("failed to get new tag data for %q (%d), error %q", name, id, err)
		} else if nId := d.Type(); nId != TagId(id) {
			t.Errorf("id returned for %q (%d) is incorrect, got %d", name, id, nId)
		}
	}
	if d, err = newFromTag(TagId(len(tagIdNames))); err == nil {
		t.Errorf("tag id %d has no associated name", len(tagIdNames))
	}
}

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
