/*
3. Longest Substring Without Repeating Characters - Medium

Given a string s, find the length of the longest substring **without duplicate characters**.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

import java.util.HashSet;
import java.util.Set;

public class LongestSubstringWithoutRepeatingCharacters3 {
    public static int lengthOfLongestSubstring(String s) {
        // for sure need a set to keep track of encountered chars
        Set<Character> set = new HashSet<>();

        int l = 0, count = 0;
        for (int r = 0; r < s.length(); ++r) {
            char c = s.charAt(r);
            // r-l+1 is the # of elts in the window currently
            if (!set.contains(c)) {
                set.add(c);
                count = Math.max(count, r - l + 1);
                continue;
            }

            // need to deal with duplicates
            // BAD: while (set.contains(s.charAt(l))) {
            while (set.contains(c)) {
                char left = s.charAt(l);
                set.remove(left);
                ++l;
            }

            // NOTE: after you finish updating left, need to add the current char
            // bc the first if never executed
            set.add(c);

            // update count
            count = Math.max(count, r - l + 1);
        }
        return count;
    }
    public static void main(String[] args) {
        String s1 = "abcabcbb";
        System.out.println(String.format("3 == %d", LongestSubstringWithoutRepeatingCharacters3.lengthOfLongestSubstring(s1)));

        String s2 = "bbbbb";
        System.out.println(String.format("1 == %d", LongestSubstringWithoutRepeatingCharacters3.lengthOfLongestSubstring(s2)));

        String s3 = "pwwkew";
        System.out.println(String.format("3 == %d", LongestSubstringWithoutRepeatingCharacters3.lengthOfLongestSubstring(s3)));
    }

}
