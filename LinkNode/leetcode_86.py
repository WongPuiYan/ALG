from typing import Optional
"""
86. 分隔链表
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

你应当 保留 两个分区中每个节点的初始相对位置。
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def partition(self, head: Optional[ListNode], x: int) -> Optional[ListNode]:
        dummy_small = ListNode(-1)
        dummy_large = ListNode(-1)

        p_small, p_large = dummy_small, dummy_large
        p = head

        while p:
            if p.val < x:
                p_small.next = p
                p_small = p_small.next
            else:
                p_large.next = p
                p_large = p_large.next
            tmp = p.next
            p.next = None
            p = tmp

        p_small.next = dummy_large.next
        return dummy_small.next
