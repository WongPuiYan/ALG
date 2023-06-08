# leetcode 104. Maximum Depth of Binary Tree
from typing import Optional


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution_1:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        res = 0
        depth = 0

        def traverse(root: Optional[TreeNode]) -> None:
            nonlocal res, depth
            if root is None:
                return

            depth += 1

            
            if root.left is None and root.right is None:
                res = max(res, depth)
            
            traverse(root.left)
            traverse(root.right)

            depth -= 1
        
        traverse(root)
        return res


class Solution_2:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0
        
        left_max = self.maxDepth(root.left)
        rigth_max = self.maxDepth(root.right)

        return max(left_max, rigth_max) + 1

