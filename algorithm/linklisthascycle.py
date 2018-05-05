# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution(object):
    def hasCycle(self, head):
        """
        :type head: ListNode
        :rtype: bool
        """
        if head is None:
            return False
        current_ptr = head
        quick_ptr = head.next
        while current_ptr is not None and quick_ptr is not None and quick_ptr.next is not None:
            if current_ptr == quick_ptr:
                return True
            current_ptr = current_ptr.next
            quick_ptr = quick_ptr.next.next

        return False
