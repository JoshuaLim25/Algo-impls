/*
402. Remove K Digits - Medium

Given string num representing a non-negative integer num, and an integer k, return the smallest possible integer after removing k digits from num.

Example 1:
Input: num = "1432219", k = 3
Output: "1219"
Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.

Example 2:
Input: num = "10200", k = 1
Output: "200"
Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.

Example 3:
Input: num = "10", k = 2
Output: "0"
Explanation: Remove all the digits from the number and it is left with nothing which is 0.

Constraints:
    1 <= k <= num.length <= 105
    num consists of only digits.
    num does not have any leading zeros except for the zero itself.
*/

/*
Want to get rid of the biggest (in value), most significant digits. s.size() == n-k
 */

package com.algos.monotonicStack;

import java.util.ArrayDeque;
import java.util.Deque;

public class RemoveKDigits402 {
    public static String removeK(String num, int k) {
        Deque<Character> stack = new ArrayDeque<>();  // monotonically INCREASING
        // "1432219", k = 3. Output "1219".
        for (char n : num.toCharArray()) {
            // NOTE: Don't need to do this, since they're ints anyway
            // int cur = Character.getNumericValue(n);
            // NOTE: Don't need <= here, since nums compared have the same value anyway
            while (!stack.isEmpty() && k > 0 && n < stack.peek()) {
                stack.pop();  // just get rid of the bigger, more signif. num
                --k;  // NOTE: ACTUALLY REMOVE IT
            }
            stack.push(n);
        }

        StringBuilder sb = new StringBuilder();

        while (k > 0 && !stack.isEmpty()) {
            // consider [12], k = 2.
            // while above doesnt fire both times, since first [], then 2 >= 1,
            // meaning k is still 2 here. We never used our removals
            stack.pop();
            --k;
        }

        while (!stack.isEmpty()) {
            // sb.insert(0, stack.pop());  // slower to prepend
            sb.append(stack.pop());
        }
        sb.reverse();

        // deal w/potentially many leading zeroes
        while (sb.length() > 1 && sb.charAt(0) == '0') {
            sb.deleteCharAt(0);
        }

        // deal w/case where num == [12], k = 2, end up with "" but want "0"
        return (sb.length() == 0) ? "0" : sb.toString();
    }
}

