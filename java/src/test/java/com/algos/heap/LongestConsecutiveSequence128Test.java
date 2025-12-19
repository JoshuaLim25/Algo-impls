package com.algos.heap;
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;

class LongestConsecutiveSequence128Test {
    // public static void main(String[] args) {
    //     int[] nums = {100,4,200,1,3,2};
    //     int soln = Solution.longestConsecutive(nums);
    //     System.out.println(String.format("4 == %d", soln));
    //
    //     int[] nums2 = {0,3,7,2,5,8,4,6,0,1};
    //     int soln2 = Solution.longestConsecutive(nums2);
    //     System.out.println(String.format("9 == %d", soln2));
    //
    //
    //     int[] nums3 = {1,0,1,2};
    //     int soln3 = Solution.longestConsecutive(nums3);
    //     System.out.println(String.format("3 == %d", soln3));
    // }
    @Test
    void testBasicUnsortedInput() {
        int[] nums = {100,4,200,1,3,2};
        int expected = 4;
        int actual = LongestConsecutiveSequence128.longestConsecutive(nums);
        assertEquals(expected, actual);
    }
    @Test
    void emptyInputReturnsZero() {
        int[] nums = {};
        int expected = 0;
        int actual = LongestConsecutiveSequence128.longestConsecutive(nums);
        assertEquals(expected, actual);
    }

    @Test
    void testDuplicatesInInput() {
        int[] nums = {1,0,1,2};
        int expected = 3;
        int actual = LongestConsecutiveSequence128.longestConsecutive(nums);
        assertEquals(expected, actual);
    }

}
