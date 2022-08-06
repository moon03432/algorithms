# https://leetcode.com/problems/longest-valid-parentheses/
def longest_valid_parentheses(s):
    """
    :type s: str
    :rtype: int
    """
    l,r = 0,0
    max_ = 0
    for i in range(len(s)):
        if s[i] == '(':
            l +=1
        else:
            r +=1
        if l == r:
            max_ = max(max_, 2 * r)
        else:
            if r >= l:
                l = r = 0
    l,r = 0,0
    for i in range(len(s)-1,-1,-1):
        if s[i] == '(':
            l += 1
        else:
            r += 1
        if l == r:
            max_ = max(max_, 2 * l)
        else:
            if l >= r:
                l = r = 0
    return max_


if __name__ == '__main__':
    print(longest_valid_parentheses("(()"))
    print(longest_valid_parentheses(")()())"))
    print(longest_valid_parentheses(""))