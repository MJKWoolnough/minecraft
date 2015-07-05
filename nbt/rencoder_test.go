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
	testRNBT(smallTestData, new(smallTest), &expecting, t)
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
	testRNBT(largeTestData, new(largeTest), &expecting, t)
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
