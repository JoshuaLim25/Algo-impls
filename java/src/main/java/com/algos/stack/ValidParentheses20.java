/*
20. Valid Parentheses - Easy
Topics
conpanies iconCompanies
Hint

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.

 

Example 1:

Input: s = "()"

Output: true

Example 2:

Input: s = "()[]{}"

Output: true

Example 3:

Input: s = "(]"

Output: false

Example 4:

Input: s = "([])"

Output: true
 */

import java.util.Map;
import java.util.Stack;
import java.util.HashMap;

public class ValidParentheses20 {
    public static boolean isValid(String s) {
        // approach: init a lookup table, then range through s
        // @ each step, check whether it's an opening or closing brace
        // NOTE: the KEYS are CLOSING braces
        Map<Character, Character> lookup = new HashMap<>(Map.of(
            ')', '(',
            '}', '{',
            ']', '['
        ));
        // "()[]{}"
        Stack<Character> stack = new Stack<>();
        for (char c : s.toCharArray()) {
            // the second part is why we made lookup "reversed"
            if (lookup.containsKey(c)) {
                System.out.println(String.format("c is %c, s.top is %c", c, stack.peek()));
                if (!stack.isEmpty() && lookup.get(c) == stack.peek()) {
                    return false;
                } else {
                    stack.pop();
                }
            } 
            else {
                stack.push(c);
            }
        }

        return stack.isEmpty();
    }

    public static void main(String[] args) {

        String s1 = "()[]{}";
        System.out.println(String.format("true == %b", ValidParentheses20.isValid(s1)));


        String s2 = "()[]{}";
        System.out.println(String.format("false == %b", ValidParentheses20.isValid(s2)));

        String s3 = "()";
        System.out.println(String.format("true == %b", ValidParentheses20.isValid(s3)));



    }
}
