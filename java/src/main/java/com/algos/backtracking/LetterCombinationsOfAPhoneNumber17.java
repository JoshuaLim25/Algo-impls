/*
17. Letter Combinations of a Phone Number - Medium

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example 1:
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Example 2:
Input: digits = "2"
Output: ["a","b","c"]

Constraints:
    1 <= digits.length <= 4
    digits[i] is a digit in the range ['2', '9'].
*/

package com.algos.backtracking;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class LetterCombinationsOfAPhoneNumber17 {
    List<String> res = new ArrayList<>();
    Map<Character, String> map;
    String digits;
    StringBuilder sb = new StringBuilder();
    public List<String> letterCombinations(String digits) {
        this.digits = digits;
        map = new HashMap<>(Map.of(
            '2', "abc",
            '3', "def",
            '4', "ghi",
            '5', "jkl",
            '6', "mno",
            '7', "pqrs",
            '8', "tuv",
            '9', "wxyz"
        ));

        backtrack(0);

        return res;
    }
    private void backtrack(int curDigit) {
        // "23"
        if (curDigit == digits.length()) {
            res.add(sb.toString());
            return;
        }

        String letters = map.get(digits.charAt(curDigit)); // "abc"

        for (int i = 0; i < letters.length(); ++i) {
            sb.append(letters.charAt(i));
            backtrack(curDigit+1);
            sb.deleteCharAt(sb.length()-1);
        }
    }

    public static void main(String[] args) {
        String digits = "23";
        List<String> expected = List.of("ad","ae","af","bd","be","bf","cd","ce","cf");
        LetterCombinationsOfAPhoneNumber17 soln = new LetterCombinationsOfAPhoneNumber17();
        System.out.println(String.format("%s == %s", expected, soln.letterCombinations(digits)));


        String digits2 = "2";
        List<String> expected2 = List.of("a","b","c");
        System.out.println(String.format("%s == %s", expected2, soln.letterCombinations(digits2)));
    }
}
