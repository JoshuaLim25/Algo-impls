/*
739. Daily Temperatures - Medium

Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.

Example 1:
Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

Example 2:
Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]
*/

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collection;
import java.util.Collections;
import java.util.List;
import java.util.Stack;
import java.util.stream.Collectors;

public class DailyTemperatures739 {
    List<Integer> foo = new ArrayList<Integer>(Collections.nCopies(10, 0));

    /*
     * 0 1 2 3 4 5 6 7
     * Input: temperatures = [73,74,75,71,69,72,76,73]
     * ^
     * i
     * []
     * ^
     * Output: [1,1,4,2,1,1,0,0]
     * res[i] == numDaysWaited, which compares temperatures[i] to
     * temperatures[stack.top]
     */
    public int[] dailyTemperatures(int[] temperatures) {
        int[] res = new int[temperatures.length];
        var stack = new Stack<Integer>();

        for (int i = 0; i < temperatures.length; ++i) {
            int tempToday = temperatures[i];
            while (stack.size() != 0 && tempToday > temperatures[stack.peek()]) {
                // top of stack was a prev day that was colder than today was
                // so caclulate span of days in between
                int colderDay = stack.pop();
                int daysWaited = i - colderDay;
                res[colderDay] = daysWaited;
            }
            stack.push(i);
        }

        return res;
    }

    public static void main(String[] args) {
        DailyTemperatures739 soln = new DailyTemperatures739();
        System.out.println(soln.foo);
        var temperatures = new int[] { 73, 74, 75, 71, 69, 72, 76, 73 };
        List<Integer> answerList = Arrays.stream(soln.dailyTemperatures(temperatures)).boxed()
                .collect(Collectors.toList());
        System.out.println(String.format("[1, 1, 4, 2, 1, 1, 0, 0] == %s", answerList));

    }
}
