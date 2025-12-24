package com.algos.stack;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

class EvaluateReversePolishNotation150Test {
    @Test
    void testEmptyPostfixExpr() {
        String[] tokens = { };
        int expected = 0;
        int actual = EvaluateReversePolishNotation150.evalRPN(tokens);
        assertEquals(expected, actual);
    }

    @Test
    void testInvalidPostfixExpr() {
        String[] tokens = { "1", "2", "3", "+"};
        assertThrows(IllegalArgumentException.class, () -> EvaluateReversePolishNotation150.evalRPN(tokens));
    }

    @Test
    void testSimplePostfixExpr() {
        String[] tokens = { "2", "1", "+", "3", "*" };
        int expected = 9;
        int actual = EvaluateReversePolishNotation150.evalRPN(tokens);
        assertEquals(expected, actual);
    }
    @Test
    void testConvolutedPostfixExpr() {
        String[] tokens = { "10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+" };
        int expected = 22;
        int actual = EvaluateReversePolishNotation150.evalRPN(tokens);
        assertEquals(expected, actual);
    }
}


/*
String[] tokens1 = { "2", "1", "+", "3", "*" };
        System.out.println(String.format("=== %s ===", Arrays.toString(tokens1)));
        System.out.println(String.format("9 == %d", Solution.evalRPN(tokens1)));

        String[] tokens2 = { "4", "13", "5", "/", "+" };
        System.out.println(String.format("=== %s ===", Arrays.toString(tokens2)));
        System.out.println(String.format("6 == %d", Solution.evalRPN(tokens2)));

        String[] tokens3 = { "4", "13", "5", "/", "+" };
        System.out.println(String.format("=== %s ===", Arrays.toString(tokens3)));
        System.out.println(String.format("6 == %d", Solution.evalRPN(tokens3)));

        String[] tokens4 = { "10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+" };
        System.out.println(String.format("=== %s ===", Arrays.toString(tokens4)));
        System.out.println(String.format("22 == %d", Solution.evalRPN(tokens4)));
*/
