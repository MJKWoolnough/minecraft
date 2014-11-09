package nbt

import "testing"

func TestTagData(t *testing.T) {
	var (
		err error
		d   Data
	)
	for id, name := range tagIdNames {
		d, err = defaultConfig.newFromTag(TagId(id))
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
	if d, err = defaultConfig.newFromTag(TagId(len(tagIdNames))); err == nil {
		t.Errorf("tag id %d has no associated name", len(tagIdNames))
	}
}
