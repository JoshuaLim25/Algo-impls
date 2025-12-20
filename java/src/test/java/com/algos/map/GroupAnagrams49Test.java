package com.algos.map;

import org.junit.jupiter.api.Test;
import static org.assertj.core.api.Assertions.assertThat;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class GroupAnagrams49Test {
    @Test
    void testManyGroupsOfAnagrams() {
        String[] anagrams = {"eat","tea","tan","ate","nat","bat"};
        List<List<String>> expected = Arrays.asList(
            List.of("eat", "tea", "ate"),
            List.of("tan", "nat"),
            List.of("bat")
        );
        List<List<String>> actual = GroupAnagrams49.groupAnagrams(anagrams);
        assertThat(expected).containsExactlyInAnyOrderElementsOf(actual);
    }

    @Test
    void testEmptyInput() {
        String[] strs = {};
        List<List<String>> expected = new ArrayList<>();
        List<List<String>> actual = GroupAnagrams49.groupAnagrams(strs);
        assertThat(expected).containsExactlyInAnyOrderElementsOf(actual);
    }
}
