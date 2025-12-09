/*
875. Koko Eating Bananas
Solved
Medium
Topics
conpanies iconCompanies

Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.

Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return the minimum integer k such that she can eat all the bananas within h hours.


Example 1:

Input: piles = [3,6,7,11], h = 8
Output: 4

Example 2:

Input: piles = [30,11,23,4,20], h = 5
Output: 30

Example 3:

Input: piles = [30,11,23,4,20], h = 6
Output: 23
*/

import java.util.Arrays;

public class KokoEatingBananas875 {
    public static int minEatingSpeed(int[] piles, int h) {
        // straightforward: you want to minimize k, which is the rate at which koko eats bananas. You want the lowest val s.t. koko can still finish eating all the banananas <= h hours, which is when the guards get back.
        int k = 1, eatingSpeed = 1, biggestPile = Arrays.stream(piles).max().getAsInt();  // NOTE: cannot have a rate < 1
        // the idea is to run a binary search looking for a "target" eat speed.
        // NOTE: there's no real point in sorting. that property doesn't help
        while (k <= biggestPile) {
            int mid = k + (biggestPile - k) / 2;
            int hours = hoursSpent(piles, mid);

            if (hours <= h) {
                // we may be able to minimize eat speed further
                // NOTE: we only update eatingSpeed if there's room for improvement
                eatingSpeed = mid;
                biggestPile = mid - 1;
            } else {
                // it took too long to eat, so need a higher eat speed
                k = mid + 1;
            }
        }

        return eatingSpeed;
    }
    private static int hoursSpent(int[] piles, int curEatSpeed) {
        int hours = 0;
        for (int i = 0; i < piles.length; ++i) {
            // need to update hours taken. no real need to check if piles[i] > k
            hours += piles[i] / curEatSpeed;
            if (piles[i] % curEatSpeed != 0) ++hours;
        }
        return hours;
    }


    public static void main(String[] args) {
        int[] piles = {3,6,7,11};
        int h = 8;
        System.out.println(String.format("4 == %d", KokoEatingBananas875.minEatingSpeed(piles, h)));


        int[] piles2 = {30,11,23,4,20};
        int h2 = 5;
        System.out.println(String.format("30 == %d", KokoEatingBananas875.minEatingSpeed(piles2, h2)));


        int[] piles3 = {30,11,23,4,20};
        int h3 = 6;
        System.out.println(String.format("23 == %d", KokoEatingBananas875.minEatingSpeed(piles3, h3)));
    }
}
