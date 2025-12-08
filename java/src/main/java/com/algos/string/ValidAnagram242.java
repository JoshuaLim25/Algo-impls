/*
242. Valid Anagram - Easy

Given two strings `s` and `t`, return true if `t` is an ANAGRAM of `s`, and false otherwise.
 

Example 1:

Input: s = "anagram", t = "nagaram"

Output: true
*/

package com.algos.strings;

public class ValidAnagram242 {
    public static boolean isValid(String src, String target) {
        if (target.length() != src.length())
            return false;
        // two loops to populate list, or
        var letterFreq = new int[26]; // itoA == i + '0' == Adding
        for (char c : src.toCharArray()) {
            letterFreq[c - 'a']++;
        }
        for (char c : target.toCharArray()) {
            int freq = letterFreq[c - 'a']--;
            if (freq < 0) {
                return false;
            }
        }
        for (int n : letterFreq) {
            if (n != 0)
                return false;
        }

        return true;
    }

    public static void main(String[] args) {
        String s1 = "anagram";
        String t1 = "nagaram";
        System.out.println(String.format("Is \"%s\" an anagram of \"%s\"? %b", t1, s1, isValid(s1, t1)));

        String s2 = "rat";
        String t2 = "car";
        System.out.println(String.format("Is \"%s\" an anagram of \"%s\"? %b", t2, s2, isValid(s2, t2)));
    }
}
