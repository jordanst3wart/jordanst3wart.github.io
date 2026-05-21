package main

import "testing"

func TestParseMarkdownWithMetaData(t *testing.T) {

	str := `---
title: Hello, world!
tags:
  - foo
  - bar
---
### hi

what is up?

	`
	source := []byte(str)
	_, metadata, err := ParseMarkdownWriting(source)
	if err != nil {
		t.Errorf("failed with error: %v", err)
	}
	if metadata.Title != "Hello, world!" {
		t.Error("title failed with")
	}
	//if got != want {
	//	t.Errorf("got %d, want %d", got, want)
	//}
}
