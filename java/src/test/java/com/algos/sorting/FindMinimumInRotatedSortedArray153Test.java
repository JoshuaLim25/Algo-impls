package com.algos.sorting;

import org.junit.jupiter.api.Test;

import com.algos.search.FindMinimumInRotatedSortedArray153;

import static org.junit.jupiter.api.Assertions.assertEquals;

class FindMinimumInRotatedSortedArray153Test {
    @Test
    void testEmptyInput() {
        int[] nums = null;
        int expected = -1;
        int actual = FindMinimumInRotatedSortedArray153.findMin(nums);
        assertEquals(expected, actual);
    }

    @Test
    void testDuplicatesInInput() {
        int[] nums = {5, 5, 5, 5, 6, 6, 7, 7, 0, 1, 2, 2};
        int expected = 0;
        int actual = FindMinimumInRotatedSortedArray153.findMin(nums);
        assertEquals(expected, actual);
    }
}
