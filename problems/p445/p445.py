from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        if l1 is None:
            return l2
        elif l2 is None:
            return l1
        l1 = self.reverse(l1)
        l2 = self.reverse(l2)
        flag = 0
        result = ListNode()
        current = result
        while l1 is not None and l2 is not None:
            c = l1.val + l2.val + flag
            if c > 10:
                flag = 1
                c -= 10
            current.next = ListNode(c)
            current = current.next
            l1 = l1.next
            l2 = l2.next
        while l1 is not None:
            c = l1.val = flag
            if c > 10:
                flag = 1
                c -= 10
            current.next = ListNode(c)
            current = current.next
            l1 = l1.next
        while l2 is not None:
            c = l2.val = flag
            if c > 10:
                flag = 1
                c -= 10
            current.next = ListNode(c)
            current = current.next
            l2 = l2.next
        if flag > 0:
            current.next = ListNode(1)
        return self.reverse(result.next)

    def reverse(self, head: ListNode) -> ListNode:
        prev = head
        current = head.next
        prev.next = None
        while current is not None:
            tmp = current.next
            current.next = prev
            prev = current
            current = tmp
        return head
