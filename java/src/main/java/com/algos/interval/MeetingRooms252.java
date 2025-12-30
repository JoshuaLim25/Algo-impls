/*
252. Meeting Rooms - Easy

Given an array of meeting time intervals where intervals[i] = [starti, endi], determine if a person could attend all meetings.

Example 1:
Input: intervals = [[0,30],[5,10],[15,20]]
Output: false

Example 2:
Input: intervals = [[7,10],[2,4]]
Output: true

Constraints:
    0 <= intervals.length <= 104
    intervals[i].length == 2
    0 <= starti < endi <= 106
 */


package com.algos.interval;

import java.util.Arrays;

public class MeetingRooms252 {
    public static boolean canAttendMeetings1(int[][] intervals) {
        int n = intervals.length;
        if (n == 0 || intervals == null) return true;
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));

        for (int i = 0; i < n; ++i) {
            int[] curInterval = intervals[i];
            int[] nextInterval = intervals[i+1];
            if (curInterval[1] > nextInterval[0]) return false;
        }

        return true;
    }

    public static boolean canAttendMeetings2(int[][] intervals) {
        int n = intervals.length;
        if (n == 0 || intervals == null) return true;
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));

        int[] curInterval = intervals[0];
        for (int i = 1; i < n; ++i) {
            int[] nextInterval = intervals[i];
            if (curInterval[1] > nextInterval[0]) return false;
            curInterval[1] = Math.max(curInterval[1], nextInterval[1]);
        }

        return true;
    }
}
