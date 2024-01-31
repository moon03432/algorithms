# https://leetcode.com/problems/longest-valid-parentheses/
def longest_valid_parentheses(s):
    """
    :type s: str
    :rtype: int
    """
    if len(s) <= 1:
        return 0

    return max(longest_valid_parentheses(s[1:]), longest_length(s))


def longest_length(s):
    max_ = 0
    score = 0

    for i in range(len(s)):
        if s[i] == '(':
            score += 1
            continue

        score -= 1

        if score < 0:
            return max_

        if score == 0:
            max_ = i + 1

    return max_


if __name__ == '__main__':
    print(longest_valid_parentheses("(()"))
    print(longest_valid_parentheses(")()())"))
    print(longest_valid_parentheses("()((((()())"))
    print(longest_valid_parentheses(""))
