package search;

import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.List;
import java.util.Queue;


class TreeNode {
    public int data;
    public TreeNode left;
    public TreeNode right;
    public TreeNode() {}
    public TreeNode(int data) {
        this.data = data;
    }
}

public class BreadthFirstSearch {
    public static List<List<Integer>> bfs(TreeNode root) {
        List<List<Integer>> res = new ArrayList<>();
        if (root == null) return res;
        Queue<TreeNode> queue = new ArrayDeque<>();
        queue.add(root);

        while (!queue.isEmpty()) {
            var n = queue.size();
            List<Integer> children = new ArrayList<>();
            for (int i = 0; i < n; ++i) {
                var nextUp = queue.poll();
                if (nextUp.left != null) {
                    queue.add(nextUp.left);
                }
                if (nextUp.right != null) {
                    queue.add(nextUp.right);
                }
                children.add(nextUp.data);
            }
            res.add(children);
        }

        return res;
    }
    public static void main(String[] args) {
        /*
        Input: root = [3,9,20,null,null,15,7]
        Output: [[3],[9,20],[15,7]]
         */
        TreeNode root = new TreeNode(3);
        root.left = new TreeNode(9);
        root.right = new TreeNode(20);
        root.right.left = new TreeNode(15);
        root.right.right = new TreeNode(7);

        var res = BreadthFirstSearch.bfs(root);
        System.out.printf("Wanted [[3],[9,20],[15,7]], got: %s\n", res);
    }
}
