/**
 * @file A lil grammar for a lil language
 * @author Team Bean
 * @license MIT
 */

/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

module.exports = grammar({
  name: "beanlang",

  rules: {
    // TODO: add the actual grammar rules
    source_file: $ => "hello"
  }
});
