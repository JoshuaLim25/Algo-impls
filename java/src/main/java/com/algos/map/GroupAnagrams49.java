/*
49. Group Anagrams - Medium

- Given an array of strings `strs`, group the anagrams together. You can return the answer in any order.
    - DEF: an **anagram** is a word or phrase formed by rearranging the letters of a different word or phrase, using all the original letters exactly once.

Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Explanation:
    There is no string in strs that can be rearranged to form "bat".
    The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
    The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.

Example 2:

Input: strs = [""]

Output: [[""]]

Example 3:

Input: strs = ["a"]

Output: [["a"]]
*/

/*
 * Approach:
 * Literally just sort each word alphabetically
 * "eat" -> "aet"
 * "tea" -> "aet"
     * Then for every string in sorted strs,
*/

package com.algos.map;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class GroupAnagrams49 {
    // dont want set, bc "aab" and "baa"
    public static List<List<String>> groupAnagrams(String[] strs) {
        if (strs == null || strs.length == 0) return new ArrayList<>();
        // associates { sortedWord: list<original words> }
        Map<String, List<String>> groupsMap = new HashMap<>();
        List<List<String>> res = new ArrayList<>();

        for (String word : strs) {
            char[] wordArray = word.toCharArray();  // [e a t]
            Arrays.sort(wordArray); // [ a e t ]
            String sortedWord = Arrays.toString(wordArray);
            // value is a list of the original words
            if (!groupsMap.containsKey(sortedWord)) {
                groupsMap.put(sortedWord, new ArrayList<>());
            }
            groupsMap.get(sortedWord).add(word);
        }

        groupsMap.forEach((k, v) -> res.add(v));
        return res;
    }

    public static void main(String[] args) {
        String[] strs1 = {"eat","tea","tan","ate","nat","bat"};
        System.out.println(String.format("[['bat'],['nat','tan'],['ate','eat','tea']] == %s", GroupAnagrams49.groupAnagrams(strs1)));
    }
}

