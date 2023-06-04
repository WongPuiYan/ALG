"""
剑指 Offer 52. 两个链表的第一个公共节点
输入两个链表，找出它们的第一个公共节点。
"""


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> ListNode:
        p_a, p_b = headA, headB

        while p_a != p_b:
            p_a = p_a.next if p_a else headB
            p_b = p_b.next if p_b else headA
        
        return p_a
