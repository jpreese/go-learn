package main

import "testing"

func TestSlowDownload(t *testing.T) {
	f := fakeImage{
		Filename: "a",
	}

	expected := "a"

	actual, err := SlowDownload(f)
	if err != nil {
		t.Fatalf("failed to download")
	}

	if expected != actual {
		t.Errorf("unexpected filename. expected %v actual %v", actual, expected)
	}
}

type fakeImage struct {
	Filename string
}

func (f fakeImage) Download() string {
	return f.Filename
}
