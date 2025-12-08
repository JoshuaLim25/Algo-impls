/*
125. Valid Palindrome - Easy

A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

*/

package com.algos.string;

public class ValidPalindrome125 {
    // input: "AbBa"
    public static boolean isValid(String s) {
        var sb = new StringBuilder();
        for (var c : s.toCharArray()) {
            if (Character.isLetterOrDigit(c)) {
                sb.append(c);
            }
        }
        // string stripped of non-alphanumeric chars
        var processed = sb.toString().toLowerCase();
        int l = 0, r = processed.length() - 1;
        while (l <= r) {
            if (processed.charAt(l) != processed.charAt(r)) {
                return false;
            } 
            l++;
            r--;
        }
        // if we made it here we passed the validation
        return true;
    }
    public static void main(String[] args) {
        var s = "AbBa";
        var s2 = "A man, a plan, a canal: Panama";
        System.out.println(String.format("Input %s should be true: %s", s, ValidPalindrome125.isValid(s)));
        System.out.println(String.format("Input %s should be true: %s", s2, ValidPalindrome125.isValid(s2)));
    }
}
