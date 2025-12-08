/*
206. Reverse Linked List - Easy

Given the head of a singly linked list, reverse the list, and return the reversed list.
 */

class ListNode {
    int data;
    ListNode next;
    public ListNode(int data) {
        this.data = data;
    }
    public ListNode(int data, ListNode next) {
        this.data = data;
        this.next = next;
    }
    public String toString() {
        ListNode cur = this;
        var sb = new StringBuilder();
        while (cur.next != null) {
            sb.append(cur.data).append(" -> ");
            cur = cur.next;
        }
        sb.append(cur.data);
        return sb.toString();
    }
}

public class ReverseLL {
    public static ListNode reverse(ListNode head) {
        ListNode prev = null;
        //    1 -> 2 -> 3 -> 4
        // p  h    n           // first save n, change link direction, then advance each one by one
        //    1 <- 2 <- 3 <- 4
        while (head != null) {
            ListNode next = head.next;
            head.next = prev;
            prev = head;
            head = next;
        }
        return prev;
    }

    public static void main(String[] args) {
        ListNode n1 = new ListNode(1);
        ListNode n2 = new ListNode(2);
        ListNode n3 = new ListNode(3);
        ListNode n4 = new ListNode(4);
        ListNode n5 = new ListNode(5);
        n1.next = n2;
        n2.next = n3;
        n3.next = n4;
        n4.next = n5;
        var newHead = ReverseLL.reverse(n1);
        System.out.println(newHead);
    }
}
