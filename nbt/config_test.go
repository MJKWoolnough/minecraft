package nbt

import "testing"

func TestTagData(t *testing.T) {
	var (
		err error
		d   Data
	)
	for id, name := range tagIdNames {
		d, err = defaultConfig.newFromTag(TagID(id))
		if id == 0 {
			if err == nil {
				t.Errorf("tag end (0) should return an error")
			}
		} else if err != nil {
			t.Errorf("failed to get new tag data for %q (%d), error %q", name, id, err)
		} else if nID := d.Type(); nID != TagID(id) {
			t.Errorf("id returned for %q (%d) is incorrect, got %d", name, id, nID)
		}
	}
	if d, err = defaultConfig.newFromTag(TagID(len(tagIdNames))); err == nil {
		t.Errorf("tag id %d has no associated name", len(tagIdNames))
	}
}
