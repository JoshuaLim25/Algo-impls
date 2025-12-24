/*
150. Evaluate Reverse Polish Notation - Medium

You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation (i.e., postfix).

Evaluate the expression. Return an integer that represents the value of the expression.

Note that:
    The valid operators are '+', '-', '*', and '/'.
    Each operand may be an integer or another expression.
    The division between two integers always truncates toward zero.
    There will not be any division by zero.
    The input represents a valid arithmetic expression in a reverse polish notation.
    The answer and all the intermediate calculations can be represented in a 32-bit integer.

Example 1:
Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9

Example 2:
Input: tokens = ["4","13","5","/","+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6

Example 3:
Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
Output: 22
Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22

Constraints:
    1 <= tokens.length <= 104
    tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the range [-200, 200].
*/

package com.algos.stack;

import java.util.ArrayDeque;
import java.util.Arrays;
import java.util.Deque;

public class EvaluateReversePolishNotation150 {
    // NOTE: for the life of me couldn't figure out why I was getting bogus, and
    // it was because of fallthrough behavior from this piece of shit switch stmt
    // use the modernized ver.
    public static int evalRPN_OLD(String[] tokens) {
        Deque<String> stack = new ArrayDeque<>();
        for (String token : tokens) {
            // this token is either a number or an operand
            // Case 1: operator
            // 1 2 3 + /
            // 1 / (2 + 3)
            if (token.equals("*") || token.equals("/") || token.equals("+") || token.equals("-")) {
                // either we have AT LEAST two numbers in stack
                // or something has gone wrong (eg 1 + ?)
                int op2 = Integer.parseInt(stack.pop());
                int op1 = Integer.parseInt(stack.pop());
                System.out.println(String.format("%d %s %d", op1, token, op2));
                switch (token) {
                    case "*":
                        stack.push(Integer.toString(op1 * op2));
                        break;
                    case "/":
                        stack.push(Integer.toString(op1 / op2));
                        break;
                    case "+":
                        stack.push(Integer.toString(op1 + op2));
                        break;
                    case "-":
                        stack.push(Integer.toString(op1 - op2));
                        break;
                }
                // Case 2: must be a number/operand
            } else {
                stack.push(token);
            }
        }

        return Integer.parseInt(stack.pop());
    }

    public static int evalRPN(String[] tokens) {
        if (tokens == null || tokens.length == 0) return 0;
        Deque<String> stack = new ArrayDeque<>();
        for (String token : tokens) {
            switch (token) {
                case "+" -> {
                    int op2 = Integer.parseInt(stack.pop());
                    int op1 = Integer.parseInt(stack.pop());
                    stack.push(Integer.toString(op1 + op2));
                }
                case "-" -> {
                    int op2 = Integer.parseInt(stack.pop());
                    int op1 = Integer.parseInt(stack.pop());
                    stack.push(Integer.toString(op1 - op2));
                }
                case "/" -> {
                    int op2 = Integer.parseInt(stack.pop());
                    int op1 = Integer.parseInt(stack.pop());
                    stack.push(Integer.toString(op1 / op2));
                }
                case "*" -> {
                    int op2 = Integer.parseInt(stack.pop());
                    int op1 = Integer.parseInt(stack.pop());
                    stack.push(Integer.toString(op1 * op2));
                }
                default -> stack.push(token);
            }
            ;
        }

        if (stack.size() > 1) throw new IllegalArgumentException("Final stack size > 1 implies bogus input");

        return Integer.parseInt(stack.pop());
    }

    public static void main(String[] args) {
        String[] tokens1 = { "2", "1", "+", "3", "*" };
        System.out.println(String.format("=== %s ===", Arrays.toString(tokens1)));
        System.out.println(String.format("9 == %d", EvaluateReversePolishNotation150.evalRPN(tokens1)));

        String[] tokens2 = { "4", "13", "5", "/", "+" };
        System.out.println(String.format("=== %s ===", Arrays.toString(tokens2)));
        System.out.println(String.format("6 == %d", EvaluateReversePolishNotation150.evalRPN(tokens2)));

        String[] tokens3 = { "4", "13", "5", "/", "+" };
        System.out.println(String.format("=== %s ===", Arrays.toString(tokens3)));
        System.out.println(String.format("6 == %d", EvaluateReversePolishNotation150.evalRPN(tokens3)));

        String[] tokens4 = { "10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+" };
        System.out.println(String.format("=== %s ===", Arrays.toString(tokens4)));
        System.out.println(String.format("22 == %d", EvaluateReversePolishNotation150.evalRPN(tokens4)));
    }
}
