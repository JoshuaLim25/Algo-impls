package monotonicStack;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class NextGreaterElementTest {
    @Test
    public void testStandardCase() {
        ArrayList<Integer> input = new ArrayList<>(List.of(4, 1, 2));
        List<Integer> expected = List.of(-1, 2, -1);
        List<Integer> actual = NextGreaterElement.nextGreaterElement(input); // Call your method
        assertEquals(expected, actual);
    }

    public static void main(String[] args) {
        var nums = new ArrayList<>(List.of(1, 2, 1));
        var res = NextGreaterElement.nextGreaterElement(nums);
        System.out.println("Expected [2,-1,2], got " + res);
    }

    @Test
    public void testCircularCase() {
        // Test case provided in LeetCode example for NGE II: nums = [1,2,3,4,3]
        ArrayList<Integer> input = new ArrayList<>(List.of(1, 2, 3, 4, 3));
        List<Integer> expected = List.of(2, 3, 4, -1, 4);
        List<Integer> actual = NextGreaterElement.nextGreaterElement(input);
        assertEquals(expected, actual);
    }

    @Test
    public void testAllGreaterElementsExist() {
        ArrayList<Integer> input = new ArrayList<>(List.of(1, 5, 3, 6, 9));
        List<Integer> expected = List.of(5, 6, 6, 9, -1);
        List<Integer> actual = NextGreaterElement.nextGreaterElement(input);
        assertEquals(expected, actual);
    }

    @Test
    public void testMonotonicallyDecreasing() {
        // When the list is decreasing, all elements should map to -1 in a circular array
        // (unless wrapped around to the beginning, but 5 -> -1, 4 -> -1, etc.)
        ArrayList<Integer> input = new ArrayList<>(List.of(5, 4, 3, 2, 1));
        List<Integer> expected = List.of(-1, -1, -1, -1, -1);
        List<Integer> actual = NextGreaterElement.nextGreaterElement(input);
        assertEquals(expected, actual);
    }

    @Test
    public void testAllSameElements() {
        ArrayList<Integer> input = new ArrayList<>(List.of(2, 2, 2, 2));
        List<Integer> expected = List.of(-1, -1, -1, -1);
        List<Integer> actual = NextGreaterElement.nextGreaterElement(input);
        assertEquals(expected, actual);
    }

    @Test
    public void testEmptyList() {
        ArrayList<Integer> input = new ArrayList<>();
        List<Integer> expected = List.of(); // Expect an empty list back
        List<Integer> actual = NextGreaterElement.nextGreaterElement(input);
        assertEquals(expected, actual);
    }

    @Test
    public void testSingleElementList() {
        ArrayList<Integer> input = new ArrayList<>(List.of(10));
        // A single element in a circular array has no greater element than itself
        List<Integer> expected = List.of(-1);
        List<Integer> actual = NextGreaterElement.nextGreaterElement(input);
        assertEquals(expected, actual);
    }
}
