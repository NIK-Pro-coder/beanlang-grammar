import XCTest
import SwiftTreeSitter
import TreeSitterBeanlang

final class TreeSitterBeanlangTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_beanlang())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Beanlang grammar")
    }
}
