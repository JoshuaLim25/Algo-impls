package sorting;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class TreeNode {
    public int data;
    public TreeNode left;
    public TreeNode right;
}

public class BinarySearch {
    public boolean binarySearch(List<Integer> nums, int target) {
        int l = 0, r = nums.size() - 1;

        while (l < r) {
            int mid = l + (r-l)/2;
            if (target == nums.get(mid)) {
                return true;
            } else if (target < nums.get(mid)) {
                r = mid - 1;
            } else if (target > nums.get(mid)) {
                l = mid + 1;
            }
        }

        return false;
    }

    public static void main(String[] args) {
        BinarySearch b = new BinarySearch();
        List<Integer> nums = new ArrayList<>(Arrays.asList(1,2,3,4,5,6,7,8,9));
        int target = 3;
        System.out.println(b.binarySearch(nums, target));
    }
}


