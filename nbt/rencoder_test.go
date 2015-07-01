package nbt

import (
	"bytes"
	"encoding/base64"
	"reflect"
	"testing"
)

type smallTest struct {
	Name string `nbt:"name"`
}

func TestRSmall(t *testing.T) {
	var expecting smallTest
	expecting.Name = "Bananrama"
	testRNBT(`CgALaGVsbG8gd29ybGQIAARuYW1lAAlCYW5hbnJhbWEA`, new(smallTest), &expecting, t)
}

type largeTest struct {
	LongTest   int64   `nbt:"longTest"`
	ShortTest  int16   `nbt:"shortTest"`
	StringTest string  `nbt:"stringTest"`
	FloatTest  float32 `nbt:"floatTest"`
	IntTest    int32   `nbt:"intTest"`
	Nested     struct {
		Ham nv `nbt:"ham"`
		Egg nv `nbt:"egg"`
	} `nbt:"nested compound test"`
	ListTestLong     []int64 `nbt:"listTest (long)"`
	ListTestCompound [2]struct {
		Name    string `nbt:"name"`
		Created int64  `nbt:"created-on"`
	} `nbt:"listTest (compound)"`
	ByteTest      int8    `nbt:"byteTest"`
	ByteArrayTest []int8  `nbt:"byteArrayTest (the first 1000 values of (n*n*255+n*7)%100, starting with n=0 (0, 62, 34, 16, 8, ...))"`
	DoubleTest    float64 `nbt:"doubleTest"`
}

type nv struct {
	Name  string  `nbt:"name"`
	Value float32 `nbt:"value"`
}

func TestRLarge(t *testing.T) {
	var expecting largeTest
	expecting.LongTest = 9223372036854775807
	expecting.ShortTest = 32767
	expecting.StringTest = "HELLO WORLD THIS IS A TEST STRING ÅÄÖ!"
	expecting.FloatTest = 0.49823147
	expecting.IntTest = 2147483647
	expecting.Nested.Ham.Name = "Hampus"
	expecting.Nested.Ham.Value = 0.75
	expecting.Nested.Egg.Name = "Eggbert"
	expecting.Nested.Egg.Value = 0.5
	expecting.ListTestLong = []int64{11, 12, 13, 14, 15}
	expecting.ListTestCompound[0].Name = "Compound tag #0"
	expecting.ListTestCompound[0].Created = 1264099775885
	expecting.ListTestCompound[1].Name = "Compound tag #1"
	expecting.ListTestCompound[1].Created = 1264099775885
	expecting.ByteTest = 127
	expecting.ByteArrayTest = byteArrayTestData()
	expecting.DoubleTest = 0.4931287132182315
	testRNBT(`CgAFTGV2ZWwEAAhsb25nVGVzdH//////////AgAJc2hvcnRUZXN0f/8IAApzdHJpbmdUZXN0AClI`+
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
		`u/9qXgA=`, new(largeTest), &expecting, t)
}

func testRNBT(input string, into, match interface{}, t *testing.T) {
	name, err := RDecode(base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(input)), into)
	if err != nil {
		t.Errorf("error encountered while reading nbt data: %q", err)
		return
	}
	if !reflect.DeepEqual(into, match) {
		t.Error("parsed nbt data did not match given nbt structure")
		return
	}
	o := new(bytes.Buffer)
	b := base64.NewEncoder(base64.StdEncoding, o)
	err = REncode(b, name, into)
	b.Close()
	if err != nil {
		t.Errorf("error encountered while writing nbt data: %q", err)
		return
	}
	if o.String() != input {
		t.Error("input and output do not match")
	}
}
