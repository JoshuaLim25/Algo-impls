/*
703. Kth Largest Element in a Stream - Easy

You are part of a university admissions office and need to keep track of the kth highest test score from applicants in real-time. This helps to determine cut-off marks for interviews and admissions dynamically as new applicants submit their scores.

You are tasked to implement a class which, for a given integer k, maintains a stream of test scores and continuously returns the kth highest test score after a new score has been submitted. More specifically, we are looking for the kth highest score in the sorted list of all scores.

Implement the KthLargest class:
    KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of test scores nums.
    int add(int val) Adds a new test score val to the stream and returns the element representing the kth largest element in the pool of test scores so far.

 

Example 1:

Input:
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]

Output: [null, 4, 5, 5, 8, 8]

Explanation:

KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3); // return 4
kthLargest.add(5); // return 5
kthLargest.add(10); // return 5
kthLargest.add(9); // return 8
kthLargest.add(4); // return 8
*/

import java.util.Arrays;
import java.util.Collections;
import java.util.Iterator;
import java.util.List;
import java.util.PriorityQueue;
import java.util.Queue;
import java.util.stream.Collectors;

// the whole idea of this problem is to use a MIN heap to store all of the BIGGEST numbers (**minimax**)
public class KthLargestElementInStream703 {
    // Initializes the object with the integer k and the stream of test scores nums.
    Queue<Integer> pq;
    int k;

    // 1 2 3 4 5 6 7, k = 2
    public KthLargestElementInStream703(int k, int[] nums) {
        this.pq = new PriorityQueue<>(); // min heap, use Collections.reverseOrder() for max heap
        this.k = k;
        if (nums == null || nums.length == 0)
            return;
        for (int n : nums) {
            add(n);
        }
    }

    // Adds a new test score val to the stream and returns the element representing
    // the kth largest element in the pool of test scores so far.
    public int add(int val) {
        pq.add(val);
        // Case 1: pq.size() <= k, so nothing to update in pq
        if (pq.size() > k) {
            pq.poll();
        }
        return pq.peek();
    }

    public static void main(String[] args) {
        int k = 3;
        var stream = new int[] { 4, 5, 8, 2 };
        KthLargestElementInStream703 kthLargest = new KthLargestElementInStream703(k, stream);
        int a1 = kthLargest.add(3); // return 4 . 23458
        int a2 = kthLargest.add(5); // return 5
        int a3 = kthLargest.add(10); // return 5
        int a4 = kthLargest.add(9); // return 8
        int a5 = kthLargest.add(4); // return 8
        System.out.println(String.format("4 == %d", a1));
        System.out.println(String.format("5 == %d", a2));
        System.out.println(String.format("5 == %d", a3));
        System.out.println(String.format("8 == %d", a4));
        System.out.println(String.format("8 == %d", a5));
    }
}
