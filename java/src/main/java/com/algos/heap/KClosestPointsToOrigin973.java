/*
973. K Closest Points to Origin - Medium

Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane and an integer k, return the k closest points to the origin (0, 0).

The distance between two points on the X-Y plane is the Euclidean distance (i.e., √(x1 - x2)2 + (y1 - y2)2).

You may return the answer in **any order**. The answer is guaranteed to be *unique* (except for the order that it is in).

Example 1:
Input: points = [[1,3],[-2,2]], k = 1
Output: [[-2,2]]
Explanation:
The distance between (1, 3) and the origin is sqrt(10).
The distance between (-2, 2) and the origin is sqrt(8).
Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
We only want the closest k = 1 points from the origin, so the answer is just [[-2,2]].

Example 2:
Input: points = [[3,3],[5,-1],[-2,4]], k = 2
Output: [[3,3],[-2,4]]
Explanation: The answer [[-2,4],[3,3]] would also be accepted.

Constraints:
    1 <= k <= points.length <= 104
    -104 <= xi, yi <= 104
*/


/*
 * Idea:
 * For each point in points, calc the distance from origin.
     * Add this distance to a minHeap
     * have a map from { distances: Point }
     * TODO: if you have two equal distances, how to handle collisions? Guaranteed unique
*/

package com.algos.heap;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.PriorityQueue;
import java.util.Queue;

record Point(int x, int y) {};

public class KClosestPointsToOrigin973 {
    public static double calculateDistanceFromOrigin(Point p) {
        double expr = p.x()*p.x() + p.y()*p.y();
        return Math.sqrt(expr);
    }
    public static Point[] kClosestPoints(Point[] points, int k) {
        List<Point> res = new ArrayList<>();
        Map<Double, Point> lookup = new HashMap<>();
        Queue<Double> minHeap = new PriorityQueue<>(Collections.reverseOrder());

        for (Point p : points) {
            double distance = calculateDistanceFromOrigin(p);
            minHeap.offer(distance);
            if (minHeap.size() > k) {
                minHeap.poll();
            }
            lookup.put(distance, p);
        }

        int n = minHeap.size();  // NOTE: need, since n changes w/ calls to q.poll()
        for (int i = 0; i < n; ++i) {
            double dist = minHeap.poll();
            Point p = lookup.get(dist);
            res.add(p);
        }

        return res.toArray(new Point[0]);
    }

    public static int[][] kClosestFLAWED(int[][] points, int k) {
        List<int[]> res = new ArrayList<>();
        Map<Double, List<int[]>> lookup = new HashMap<>();
        Queue<Double> minHeap = new PriorityQueue<>(Collections.reverseOrder());

        for (int[] point : points) {
            double distance = Math.sqrt(point[0]*point[0] + point[1] * point[1]);
            minHeap.offer(distance);
            if (minHeap.size() > k) {
                minHeap.poll();
            }
            // if you have multiple distinct points mapping onto the same distance,
            // need to do linear chaining
            if (!lookup.containsKey(distance)) {
                List<int[]> tmp = new ArrayList<>();
                tmp.add(point);
                lookup.put(distance, tmp);
            } else {
                lookup.get(distance).add(point);
            }
        }

        int n = minHeap.size();  // NOTE: need, since n changes w/ calls to q.poll()
        for (int i = 0; i < n; ++i) {
            double dist = minHeap.poll();
            List<int[]> pointsList = lookup.get(dist);
            for (var p : pointsList) {
                res.add(p);
            }
        }

        return res.toArray(new int[0][]);
    }

    public static int[][] kClosest(int[][] points, int k) {
        List<int[]> res = new ArrayList<>();
        Queue<int[]> maxHeap = new PriorityQueue<>(
            (a, b) -> (b[0] * b[0] + b[1] * b[1]) - (a[0] * a[0] + a[1] * a[1])
        );

        // populate maxHeap with k closest points
        for (int[] point : points) {
            maxHeap.offer(point);
            if (maxHeap.size() > k) {
                maxHeap.poll();
            }
        }

        int n = maxHeap.size();
        for (int i = 0; i < n; ++i) {
            res.add(maxHeap.poll());
        }

        return res.toArray(new int[0][]);
    }


    public static void main(String[] args) {
        int[][] points1 = {{1,3},{-2,2}};
        int k = 1;
        System.out.println(String.format("[[-2,2]] == %s", Arrays.deepToString(KClosestPointsToOrigin973.kClosest(points1, k))));

        Point[] pointsArray1 = {new Point(1,3),new Point(-2,2)};
        Point[] ans1 = KClosestPointsToOrigin973.kClosestPoints(pointsArray1, k);
        System.out.println(String.format("[[-2,2]] == %s", Arrays.deepToString(ans1)));

        System.out.println();
        System.out.println("================================");
        System.out.println();



        int[][] points2 = {{3,3},{5,-1},{-2,4}};
        int k2 = 2;
        System.out.println(String.format("[[3,3],[-2,4]] == %s", Arrays.deepToString(KClosestPointsToOrigin973.kClosest(points2, k2))));

        Point[] pointsArray2 = {new Point(3,3), new Point(5,-1), new Point(-2,4)};
        Point[] ans2 = KClosestPointsToOrigin973.kClosestPoints(pointsArray2, k2);
        System.out.println(String.format("[[3,3],[-2,4]] == %s", Arrays.deepToString(ans2)));
//[[0,1],[1,0]]
    }
}

