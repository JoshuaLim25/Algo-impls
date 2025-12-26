package com.algos.greedy;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;

class JumpGame55Test {
    @Test
    void testImpossibleInput() {
        int[] nums = {3,2,1,0,4};
        boolean expected = false;
        boolean actual = JumpGame55.canJump(nums);
        assertEquals(expected, actual);
    }
    @Test
    void testTrickyInput() {
        int[] nums = {2,5,0,0};
        boolean expected = true;
        boolean actual = JumpGame55.canJump(nums);
        assertEquals(expected, actual);
    }
}
