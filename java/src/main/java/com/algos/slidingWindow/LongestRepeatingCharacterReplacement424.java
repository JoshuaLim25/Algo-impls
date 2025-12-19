/*
424. Longest Repeating Character Replacement
Medium
Topics
conpanies iconCompanies

You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.

 

Example 1:

Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.

Example 2:

Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
There may exists other ways to achieve this answer too.

Constraints:
    1 <= s.length <= 105
    s consists of only uppercase English letters.
    0 <= k <= s.length
*/

import java.util.HashMap;
import java.util.Map;

public class LongestRepeatingCharacterReplacement424 {
    public static int characterReplacement(String s, int k) {
        // k = max window size
        int l = 0, count = 0, maxFreq = 0;
        // need to account for inputs like "AABABBA", where
        // the largest window may come later
        // NOTE: think in terms of the number of *replacements*, not the max freq
        // i.e., for "sssa", you care about 'a', not 's'
        // How to represent whether we can replace up to k letters to get biggest window
        // into s with all the same letter?
        Map<Character, Integer> freqMap = new HashMap<>();
        // want biggest freq, since it requires least number of changes

        for (int r = 0; r < s.length(); ++r) {
            char c = s.charAt(r);
            freqMap.put(c, freqMap.getOrDefault(c, 0) + 1);
            maxFreq = Math.max(maxFreq, freqMap.get(c));

            // while condition for window is violated
            while ((r - l + 1) - maxFreq > k) {
                // update map with s[l], since we "forget" what it used to be
                freqMap.put(s.charAt(l), freqMap.get(s.charAt(l)) - 1);
                ++l;
            }

            // if current window larger than one we found prev, update result
            count = Math.max(count, r - l + 1);
        }
        return count;
    }

    public static void main(String[] args) {
        String s1 = "ABAB";
        int k1 = 2;
        System.out.println(String.format("4 == %d", LongestRepeatingCharacterReplacement424.characterReplacement(s1, k1)));

        String s2 = "AABABBA";
        int k2 = 1;
        System.out.println(String.format("4 == %d", LongestRepeatingCharacterReplacement424.characterReplacement(s2, k2)));

    }
}
