class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        
        s_hash = {}

        for ch in s:
            s_hash[ch] = s_hash.get(ch, 0) + 1
        
        for ch in t:
            if ch not in s_hash:
                return False

            s_hash[ch]-=1

            if s_hash[ch] < 0:
                return False

        return True