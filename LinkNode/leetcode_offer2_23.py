"""
剑指 Offer II 023. 两个链表的第一个重合节点
给定两个单链表的头节点 headA 和 headB ，请找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
"""


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> ListNode:
        p_a = headA
        p_b = headB

        while p_a != p_b:
            p_a = p_a.next if p_a else headB
            p_b = p_b.next if p_b else headA

        return p_a
