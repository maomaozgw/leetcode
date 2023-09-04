from typing import Optional, List
import unittest


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        if head is None:
            return False
        slow, fast = head, head
        while slow is not None and fast is not None and fast.next is not None:
            slow = slow.next
            fast = fast.next.next
            if slow == fast:
                return True
        return False


class Test(unittest.TestCase):
    def setUp(self) -> None:
        self.s = Solution()

    def get_list(self, data: List, cycle_pos: int) -> Optional[ListNode]:
        head = ListNode(0)
        current = head
        for idx in range(0, len(data) - 1):
            current.val = data[idx]
            current.next = ListNode(0)
            current = current.next
        current.val = data[len(data) - 1]
        if cycle_pos > -1:
            tmp = head
            for idx in range(0, cycle_pos):
                tmp = tmp.next
            current.next = tmp
        return head

    def test_example1(self) -> None:
        self.assertEqual(True, self.s.hasCycle(self.get_list([3, 2, 0, -4], 1)))

    def test_example2(self) -> None:
        self.assertEqual(True, self.s.hasCycle(self.get_list([1, 2], 0)))

    def test_example3(self) -> None:
        self.assertEqual(False, self.s.hasCycle(self.get_list([1], -1)))


if __name__ == "__main__":
    unittest.main("p141")
