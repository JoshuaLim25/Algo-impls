package com.algos.monotonicStack;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;

class RemoveKDigits402Test {
    @Test
    void testBasicNum() {
        String num = "1432219";
        int k = 3;
        String expected = "1219";
        String actual = RemoveKDigits402.removeK(num, k);
        assertEquals(expected, actual);
    }

    @Test
    void testLeadingZeroes() {
        String num = "10200";
        int k = 1;
        String expected = "200";
        String actual = RemoveKDigits402.removeK(num, k);
        assertEquals(expected, actual);
    }

    @Test
    void testEmptyResultShouldBeZero() {
        String num = "12";
        int k = 2;
        String expected = "0"; // NOT ""
        String actual = RemoveKDigits402.removeK(num, k);
        assertEquals(expected, actual);
    }
}
