package com.algos.heap;

import org.junit.jupiter.api.Test;
import static org.assertj.core.api.Assertions.assertThat;

class KClosestPointsToOrigin973Test {
    @Test
    void testSimpleCoordinates() {
        int[][] points = {{3,3},{5,-1},{-2,4}};
        int k = 2;
        int[][] expected = {{-2,4},{3,3}};
        int[][] actual = KClosestPointsToOrigin973.kClosest(points, k);
        assertThat((Object[]) expected).containsExactlyInAnyOrder((Object[]) actual);
    }

    @Test
    void testDuplicateDistances() {
        int[][] points = {{-1, 0}, {0, -1}};
        int k = 2;
        int[][] expected = {{-1, 0}, {0, -1}};
        int[][] actual = KClosestPointsToOrigin973.kClosest(points, k);
        assertThat((Object[]) expected).containsExactlyInAnyOrder((Object[]) actual);
    }

}
