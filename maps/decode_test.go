package maps

import (
	"compress/gzip"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	f, err := os.Open("map_0.dat")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	r, err := gzip.NewReader(f)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	i, n, err := image.Decode(r)
	r.Close()
	f.Close()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	} else if n != "minecraftmap" {
		t.Errorf("expecting image type %q, got %q", "minecraftmap", n)
		return
	}
	g, err := os.Create("t.png")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	png.Encode(g, i)
	g.Close()
}

func TestEncode(t *testing.T) {
	f, err := os.Open("go.png")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	im, _, err := image.Decode(f)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	f.Close()
	pr, pw := io.Pipe()
	go Encode(pw, im)
	im, _, err = image.Decode(pr)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	g, err := os.Create("got.png")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	png.Encode(g, im)
	g.Close()
}

func TestConfig(t *testing.T) {
	f, err := os.Open("map_0.dat")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	r, err := gzip.NewReader(f)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	c, n, err := image.DecodeConfig(r)
	r.Close()
	f.Close()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	} else if n != "minecraftmap" {
		t.Errorf("expecting image type %q, got %q", "minecraftmap", n)
		return
	}
	fmt.Println(c)
}
