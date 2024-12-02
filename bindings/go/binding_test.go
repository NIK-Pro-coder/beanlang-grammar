package tree_sitter_beanlang_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_beanlang "github.com/nik-pro-coder/beanlang-grammar/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_beanlang.Language())
	if language == nil {
		t.Errorf("Error loading Beanlang grammar")
	}
}
