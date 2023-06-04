from typing import Optional, List

"""
23. 合并 K 个升序链表
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        import heapq
        dummy = ListNode(-1)
        p = dummy
        min_heap = []

        for head in lists:
            while head:
                heapq.heappush(min_heap, head.val)
                head = head.next
        
        while min_heap:
            p.next = ListNode(heapq.heappop(min_heap))
            p = p.next
        
        return dummy.next
