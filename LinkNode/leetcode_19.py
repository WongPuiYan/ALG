from typing import Optional

"""
19. 删除链表的倒数第 N 个结点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy = ListNode(-1, head)
        slow = dummy
        fast = head

        while fast:
            if n:
                n -= 1
            else:
                slow = slow.next
            fast = fast.next

        slow.next = slow.next.next
        return dummy.next        
