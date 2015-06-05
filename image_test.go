package image

import "testing"

func TestExtension(t *testing.T) {
	curLen := 110
	if len(Extensions) != curLen {
		t.Fatalf("Length doesn't match. Expected %d, Got %d", curLen, len(Extensions))
	}
}

func TestIs(t *testing.T) {
	if Is("a/b/c/bar.css") {
		t.Fatal("Wrong detection. Must not be image")
	}

	if !Is("a/b/c/bar.jpg") {
		t.Fatal("Wrong detection. Must be image")
	}

	if !Is("a/b/c/bar.JPG") {
		t.Fatal("Wrong detection. Must be image")
	}

	if Is("a/b/c/barjpg") {
		t.Fatal("Wrong detection. Must not be image")
	}
}
