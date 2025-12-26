package com.algos.matrix;

import org.junit.jupiter.api.Test;
import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.assertEquals;

class MaxAreaOfIsland695Test {
    @Test
    void testCorrectAreaForValidGrid() {
        int[][] grid = { { 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0 }, { 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0 },
                { 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0 }, { 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0 },
                { 0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0 }, { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0 },
                { 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0 }, { 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0 } };
        int expected = 6;
        MaxAreaOfIsland695 soln = new MaxAreaOfIsland695();
        int actual = soln.maxArea(grid);
        assertEquals(expected, actual);
    }

    @Test
    void testNoIslands() {
        int[][] grid = { { 0, 0, 0, 0, 0, 0, 0, 0 } };
        int expected = 0;
        MaxAreaOfIsland695 soln = new MaxAreaOfIsland695();
        int actual = soln.maxArea(grid);
        assertEquals(expected, actual);
    }

}
