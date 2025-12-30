package com.algos.backtracking;

import org.junit.jupiter.api.Test;
import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.List;

class LetterCombinationsOfAPhoneNumber17Test {
    @Test
    void testBasicDigitCombo() {
        String digits = "23";
        List<String> expected = List.of("ad","ae","af","bd","be","bf","cd","ce","cf");
        LetterCombinationsOfAPhoneNumber17 soln = new LetterCombinationsOfAPhoneNumber17();
        var actual = soln.letterCombinations(digits);
        assertEquals(expected, actual);
    }

}
