package com.algos.twoPointer;

import org.junit.jupiter.api.Test;
import static org.assertj.core.api.Assertions.assertThat;

import java.util.ArrayList;
import java.util.List;

class ThreeSum15Test {
    @Test
    void testBasicInput() {
        int[] nums = new int[] { -100, -70, -60, 110, 120, 130, 160 };
        List<List<Integer>> expected = List.of(
                List.of(-100, -60, 160),
                List.of(-70, -60, 130));
        List<List<Integer>> actual = ThreeSum15.threeSum(nums);
        assertThat(expected).containsExactlyInAnyOrderElementsOf(actual);
    }

    @Test
    void testInputWithDuplicates() {
        int[] nums = { 0, 1, 1 };
        List<List<Integer>> expected = new ArrayList<>();
        List<List<Integer>> actual = ThreeSum15.threeSum(nums);
        assertThat(expected).containsExactlyInAnyOrderElementsOf(actual);
    }
}
