/*
1046. Last Stone Weight - Easy

You are given an array of integers stones where stones[i] is the weight of the ith stone.

We are playing a game with the stones. On each turn, we choose the heaviest two stones and smash them together. Suppose the heaviest two stones have weights x and y with x <= y. The result of this smash is:
    If x == y, both stones are destroyed, and
    If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.

At the end of the game, there is at most one stone left.
Return the weight of the last remaining stone. If there are no stones left, return 0.

Example 1:
Input: stones = [2,7,4,1,8,1]
Output: 1
Explanation:
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of the last stone.
 */

import java.util.Collections;
import java.util.PriorityQueue;
import java.util.Queue;


public class LastStoneWeight1046 {
    private static Queue<Integer> maxHeap;

    public static int getWeight(int[] stones) {
        maxHeap = new PriorityQueue<>(Collections.reverseOrder());
        for (int weight : stones) {
            maxHeap.add(weight);
        }

        while (maxHeap.size() != 0 && maxHeap.size() != 1) {
            System.out.println(String.format("maxHeap size = %d", maxHeap.size()));
            // take 2 biggest stones out.
            // NOTE: order of instantiation MATTERS. x <= y.
            int y = maxHeap.poll();
            int x = maxHeap.poll();

            // If x == y, both stones are destroyed, and
            // If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
            if (x != y) {
                y = y - x ;
                maxHeap.add(y);
            }
        }
        if (maxHeap.size() == 1) {
            return maxHeap.poll();
        }
        return 0;
    }

    public static void main(String[] args) {
        int[] stones = new int[]{2,7,4,1,8,1};
        System.out.println(LastStoneWeight1046.getWeight(stones));
    }

}
