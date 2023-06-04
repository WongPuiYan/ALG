from typing import List

"""
剑指 Offer II 078. 合并排序链表
给定一个链表数组，每个链表都已经按升序排列。

请将所有链表合并到一个升序链表中，返回合并后的链表。
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    # 分治合并
    def mergeKLists(self, lists: List[ListNode]) -> ListNode:
        if not lists:
            return
        
        n = len(lists)
        return self.merge(lists, 0, n - 1)
    
    def merge(self, lists, l, r):
        if l == r:
            return lists[l]

        mid = l + (r - l) // 2
        l1 = self.merge(lists, l, mid)
        l2 = self.merge(lists, mid + 1, r)

        return self.mergeTwoLists(l1, l2)
    
    def mergeTwoLists(self, l1, l2):
        if not l1:
            return l2
        if not l2:
            return l1
        
        if l1.val < l2.val:
            l1.next = self.mergeTwoLists(l1.next, l2)
            return l1
        else:
            l2.next = self.mergeTwoLists(l1, l2.next)
            return l2