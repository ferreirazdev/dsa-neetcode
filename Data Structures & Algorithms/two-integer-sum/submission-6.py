class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        cp_hash = {}

        for i, num in enumerate(nums):
            diff = target - num
            if diff in cp_hash:
                return [cp_hash[diff], i]
            
            cp_hash[num] = i

        return []