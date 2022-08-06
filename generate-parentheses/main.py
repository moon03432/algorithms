# https://leetcode.com/problems/generate-parentheses/

def generate(n):
    if n == 0:
        yield ""
        return
    if n == 1:
        yield "()"
        return

    for i in range(n):
        for s1 in generate(i):
            for s2 in generate(n - 1 - i):
                yield s1 + "(" + s2 + ")"


def generate_parenthesis(n: int):
    return list(generate(n))


if __name__ == '__main__':
    print(generate_parenthesis(3))
