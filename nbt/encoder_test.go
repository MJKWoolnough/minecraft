package nbt

import (
	"bytes"
	"encoding/base64"
	"testing"
)

/*func TestSize(t *testing.T) {
	for i := TagID(1); i < 12; i++ {
		o := new(bytes.Buffer)
		tag, _ := defaultConfig.newFromTag(i)
		d := NewTag("test", tag)
		n, _ := d.WriteTo(o)
		_, m, _ := ReadNBTFrom(o)
		if n != m {
			t.Errorf("written and read sizes for %s do not match, written %d, read %d", i, n, m)
		}
	}
}*/

func TestSmall(t *testing.T) { //test.nbt
	testNBT(`CgALaGVsbG8gd29ybGQIAARuYW1lAAlCYW5hbnJhbWEA`,
		NewTag("hello world", Compound{
			NewTag("name", String("Bananrama")),
		}), t)
}

func TestLarge(t *testing.T) { //bigtest.nbt
	data := make([]int8, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = int8((i*i*255 + i*7) % 100)
	}
	testNBT(`CgAFTGV2ZWwEAAhsb25nVGVzdH//////////AgAJc2hvcnRUZXN0f/8IAApzdHJpbmdUZXN0AClI`+
		`RUxMTyBXT1JMRCBUSElTIElTIEEgVEVTVCBTVFJJTkcgw4XDhMOWIQUACWZsb2F0VGVzdD7/GDID`+
		`AAdpbnRUZXN0f////woAFG5lc3RlZCBjb21wb3VuZCB0ZXN0CgADaGFtCAAEbmFtZQAGSGFtcHVz`+
		`BQAFdmFsdWU/QAAAAAoAA2VnZwgABG5hbWUAB0VnZ2JlcnQFAAV2YWx1ZT8AAAAAAAkAD2xpc3RU`+
		`ZXN0IChsb25nKQQAAAAFAAAAAAAAAAsAAAAAAAAADAAAAAAAAAANAAAAAAAAAA4AAAAAAAAADwkA`+
		`E2xpc3RUZXN0IChjb21wb3VuZCkKAAAAAggABG5hbWUAD0NvbXBvdW5kIHRhZyAjMAQACmNyZWF0`+
		`ZWQtb24AAAEmUjfVjQAIAARuYW1lAA9Db21wb3VuZCB0YWcgIzEEAApjcmVhdGVkLW9uAAABJlI3`+
		`1Y0AAQAIYnl0ZVRlc3R/BwBlYnl0ZUFycmF5VGVzdCAodGhlIGZpcnN0IDEwMDAgdmFsdWVzIG9m`+
		`IChuKm4qMjU1K24qNyklMTAwLCBzdGFydGluZyB3aXRoIG49MCAoMCwgNjIsIDM0LCAxNiwgOCwg`+
		`Li4uKSkAAAPoAD4iEAgKFixMEkYgBFZOUFwOLlgoAko4MDI+VBA6CkgsGhIUIDZWHFAqDmBYWgIY`+
		`OGIyDFRCOjxIXhpEFFI2JBweKkBgJlo0GAZiAAwiQgg8Fl5MREZSBCROHlxALiYoNEoGMAA+IhAI`+
		`ChYsTBJGIARWTlBcDi5YKAJKODAyPlQQOgpILBoSFCA2VhxQKg5gWFoCGDhiMgxUQjo8SF4aRBRS`+
		`NiQcHipAYCZaNBgGYgAMIkIIPBZeTERGUgQkTh5cQC4mKDRKBjAAPiIQCAoWLEwSRiAEVk5QXA4u`+
		`WCgCSjgwMj5UEDoKSCwaEhQgNlYcUCoOYFhaAhg4YjIMVEI6PEheGkQUUjYkHB4qQGAmWjQYBmIA`+
		`DCJCCDwWXkxERlIEJE4eXEAuJig0SgYwAD4iEAgKFixMEkYgBFZOUFwOLlgoAko4MDI+VBA6Ckgs`+
		`GhIUIDZWHFAqDmBYWgIYOGIyDFRCOjxIXhpEFFI2JBweKkBgJlo0GAZiAAwiQgg8Fl5MREZSBCRO`+
		`HlxALiYoNEoGMAA+IhAIChYsTBJGIARWTlBcDi5YKAJKODAyPlQQOgpILBoSFCA2VhxQKg5gWFoC`+
		`GDhiMgxUQjo8SF4aRBRSNiQcHipAYCZaNBgGYgAMIkIIPBZeTERGUgQkTh5cQC4mKDRKBjAAPiIQ`+
		`CAoWLEwSRiAEVk5QXA4uWCgCSjgwMj5UEDoKSCwaEhQgNlYcUCoOYFhaAhg4YjIMVEI6PEheGkQU`+
		`UjYkHB4qQGAmWjQYBmIADCJCCDwWXkxERlIEJE4eXEAuJig0SgYwAD4iEAgKFixMEkYgBFZOUFwO`+
		`LlgoAko4MDI+VBA6CkgsGhIUIDZWHFAqDmBYWgIYOGIyDFRCOjxIXhpEFFI2JBweKkBgJlo0GAZi`+
		`AAwiQgg8Fl5MREZSBCROHlxALiYoNEoGMAA+IhAIChYsTBJGIARWTlBcDi5YKAJKODAyPlQQOgpI`+
		`LBoSFCA2VhxQKg5gWFoCGDhiMgxUQjo8SF4aRBRSNiQcHipAYCZaNBgGYgAMIkIIPBZeTERGUgQk`+
		`Th5cQC4mKDRKBjAAPiIQCAoWLEwSRiAEVk5QXA4uWCgCSjgwMj5UEDoKSCwaEhQgNlYcUCoOYFha`+
		`Ahg4YjIMVEI6PEheGkQUUjYkHB4qQGAmWjQYBmIADCJCCDwWXkxERlIEJE4eXEAuJig0SgYwAD4i`+
		`EAgKFixMEkYgBFZOUFwOLlgoAko4MDI+VBA6CkgsGhIUIDZWHFAqDmBYWgIYOGIyDFRCOjxIXhpE`+
		`FFI2JBweKkBgJlo0GAZiAAwiQgg8Fl5MREZSBCROHlxALiYoNEoGMAYACmRvdWJsZVRlc3Q/349r`+
		`u/9qXgA=`,
		NewTag("Level", Compound{
			NewTag("longTest", Long(9223372036854775807)),
			NewTag("shortTest", Short(32767)),
			NewTag("stringTest", String("HELLO WORLD THIS IS A TEST STRING ÅÄÖ!")),
			NewTag("floatTest", Float(0.49823147)),
			NewTag("intTest", Int(2147483647)),
			NewTag("nested compound test", Compound{
				NewTag("ham", Compound{
					NewTag("name", String("Hampus")),
					NewTag("value", Float(0.75)),
				}),
				NewTag("egg", Compound{
					NewTag("name", String("Eggbert")),
					NewTag("value", Float(0.5)),
				}),
			}),
			NewTag("listTest (long)", NewList([]Data{
				Long(11),
				Long(12),
				Long(13),
				Long(14),
				Long(15),
			})),
			NewTag("listTest (compound)", NewList([]Data{
				Compound{
					NewTag("name", String("Compound tag #0")),
					NewTag("created-on", Long(1264099775885)),
				},
				Compound{
					NewTag("name", String("Compound tag #1")),
					NewTag("created-on", Long(1264099775885)),
				},
			})),
			NewTag("byteTest", Byte(127)),
			NewTag("byteArrayTest (the first 1000 values of (n*n*255+n*7)%100, starting with n=0 (0, 62, 34, 16, 8, ...))", ByteArray(data)),
			NewTag("doubleTest", Double(0.4931287132182315)),
		}), t)
}

func testNBT(input string, middle Tag, t *testing.T) {
	tag, err := NewDecoder(base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(input))).DecodeTag()
	if err != nil {
		t.Errorf("error encountered while reading nbt data: %q", err)
		return
	}
	if !tag.Equal(middle) {
		t.Error("parsed nbt data did not match given nbt structure")
		return
	}
	o := new(bytes.Buffer)
	b := base64.NewEncoder(base64.StdEncoding, o)
	err = NewEncoder(b).EncodeTag(tag)
	b.Close()
	if err != nil {
		t.Errorf("error encountered while writing nbt data: %q", err)
		return
	}
	if o.String() != input {
		t.Error("input and output do not match")
	}
}
