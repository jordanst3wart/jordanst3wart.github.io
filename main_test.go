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
	post2, err := ParseMarkdownWriting(source)
	if err != nil {
		t.Errorf("failed with error: %v", err)
	}
	if post2.Metadata.Title != "Hello, world!" {
		t.Error("title failed with")
	}
	//if got != want {
	//	t.Errorf("got %d, want %d", got, want)
	//}
}

func TestGetWriting(t *testing.T) {
	content := GetWritings()
	if content[0].Content == "" {
		t.Errorf("failed to get content")
	}
}
