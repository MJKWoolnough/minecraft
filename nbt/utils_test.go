package nbt

import (
	"bytes"
	"testing"
)

func TestCounts(t *testing.T) {
	var (
		buf   []byte
		total int64
		n     int
	)
	b := new(bytes.Buffer)
	for i := 0; i < 1500; i++ {
		cR := &countReader{Reader: b}
		cW := &countWriter{Writer: b}
		buf = make([]byte, i)
		n, _ = cW.Write(buf)
		if n != i {
			t.Errorf("failed to write %d bytes to buffer, wrote %d", i, n)
		} else if cW.BytesWritten(&total); total != int64(i) {
			t.Errorf("written bytes returned (%d) didn't match expected (%d)", total, i)
		}
		n, _ = cR.Read(buf)
		if n != i {
			t.Errorf("failed to read %d bytes from buffer, read %d", i, n)
		} else if cR.BytesRead(&total); total != int64(i) {
			t.Errorf("read bytes returned (%d) didn't match expected (%d)", total, i)
		}
	}
}

func TestTagData(t *testing.T) {
	var (
		err error
		d   Data
		nId TagId
	)
	for id, name := range tagIdNames {
		d, err = newFromTag(TagId(id))
		if id == 0 {
			if err == nil {
				t.Errorf("tag end (0) should return an error")
			}
		} else if err != nil {
			t.Errorf("failed to get new tag data for %q (%d), error %q", name, id, err)
		} else if nId, err = idFromData(d); err != nil {
			t.Errorf("failed to get id for %q (%d) from data, error %q", name, id, err)
		} else if nId != TagId(id) {
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
