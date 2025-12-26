package com.algos.backtracking;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;

class WordSearch79Test {
    @Test
    void testValidPath() {
        char[][] board = { { 'A', 'B', 'C', 'E' }, { 'S', 'F', 'C', 'S' }, { 'A', 'D', 'E', 'E' } };
        String word = "ABCCED";
        boolean expected = true;
        WordSearch79 soln = new WordSearch79();
        boolean actual = soln.exist(board, word);
        assertEquals(expected, actual);
    }

    @Test
    void testInvalidPath() {
        char[][] board = { { 'A', 'B', 'C', 'E' }, { 'S', 'F', 'C', 'S' }, { 'A', 'D', 'E', 'E' } };
        String word = "ABCB";
        boolean expected = false;
        WordSearch79 soln = new WordSearch79();
        boolean actual = soln.exist(board, word);
        assertEquals(expected, actual);
    }
}
