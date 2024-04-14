import unittest
from typing import Optional
import json
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = int(val)
        self.left = left
        self.right = right
    
    def __new__(cls, *args, **kwargs):
        instance = super().__new__(cls)
        return instance

    @staticmethod
    def new_from_leetcode(leet_code_str: str) -> Optional["TreeNode"]:
        args = leet_code_str.split(",")
        null_str = 'null'
        if len(args) == 0:
            return None
        root_val = int(args[0])
        root_node = TreeNode(val=root_val)
        nodes = [root_node]
        for i in range(1,len(args),2):
            p_node = nodes.pop(0)
            if args[i] != null_str:
                left_node = TreeNode(val=args[i])
                p_node.left = left_node
                nodes.append(left_node)
            if len(args) - i > 1 and args[i] != null_str:
                right_node = TreeNode(val=args[i+1])
                p_node.right = right_node
                nodes.append(right_node)
        return root_node

class Solution:
    def sumOfLeftLeaves(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0
        return self.sum_left(root.left) + self.sum_right(root.right)

    def sum_left(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0
        if root.left is None and root.right is None:
            return root.val
        return self.sumOfLeftLeaves(root)
   
    def sum_right(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0
        if root.left is None and root.right is None:
            return 0
        return self.sumOfLeftLeaves(root) 

class Test(unittest.TestCase):
    def setUp(self) -> None:
        self.s = Solution()
        return super().setUp()
    
    def test_example_1(self) -> None:
        self.assertEqual(
            24,
            self.s.sumOfLeftLeaves(TreeNode.new_from_leetcode('3,9,20,null,null,15,7'))
        )

    def test_exmple_2(self) -> None:
        self.assertEqual(
            0, 
            self.s.sumOfLeftLeaves(TreeNode.new_from_leetcode('1'))
        )


if __name__ == "__main__":
    unittest.main("p404")