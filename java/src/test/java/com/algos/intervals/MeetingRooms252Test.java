package com.algos.intervals;

import org.junit.jupiter.api.Test;

import com.algos.interval.MeetingRooms252;

import static org.junit.jupiter.api.Assertions.assertEquals;

class MeetingRooms252Test {
    @Test
    void testEmptyIntervals() {
        int[][] intervals = {};
        boolean expected = true;
        assertEquals(expected, MeetingRooms252.canAttendMeetings1(intervals));
    }

    @Test
    void testTwoIntervalsWithSameStartTime() {
        int[][] intervals = {{1,2},{1, 4}};
        boolean expected = false;
        assertEquals(expected, MeetingRooms252.canAttendMeetings1(intervals));
    }

}
