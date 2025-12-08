import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Stack;

public class NextGreaterElement {
    /*
    Given a circular integer array nums (i.e., the next element of nums[nums.length - 1] is nums[0]), return the next
    greater number for every element in nums.
    The next greater number of a number x is the first greater number to its traversing-order next in the array,
    which means you could search circularly to find its next greater number.
    If it doesn't exist, return -1 for this number.

    Example 1:
    Input: nums = [1,2,1]
    Output: [2,-1,2]
    Explanation: The first 1's next greater number is 2;
    The number 2 can't find next greater number.
    The second 1's next greater number needs to search circularly, which is also 2.
     */
    public static List<Integer> nextGreaterElement(ArrayList<Integer> nums) {
        var n = nums.size();
        List<Integer> res = new ArrayList<>(Collections.nCopies(n, -1));
        System.out.println(res);
        Stack<Integer> s = new Stack<>();  // monotonically decreasing. Tracks indices, which map onto smaller nums
        for (int i = 0; i < n * 2; ++i) {
            var num = nums.get(i % n);
            // MUST BE WHILE, to keep popping
            while (!s.isEmpty() && num > nums.get(s.peek())) {
                var lesserIdx = s.pop();   // NOTE: index
                res.set(lesserIdx, num);
            }
            s.add(i % n);
        }
        return res;
    }

    public static void main(String[] args) {
        var nums = new ArrayList<>(List.of(1, 2, 1));
        var res = NextGreaterElement.nextGreaterElement(nums);
        System.out.println("Expected [2,-1,2], got " + res);
    }

}
