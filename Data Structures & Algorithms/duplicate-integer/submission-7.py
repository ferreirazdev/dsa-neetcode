class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        prevList = {}
        for n in nums:
            if n in prevList:
                return True
            
            prevList[n] = True

        return False
        