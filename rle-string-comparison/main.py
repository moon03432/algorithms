# This is a sample Python script.

# Press ⌃R to execute it or replace it with your code.
# Press Double ⇧ to search everywhere for classes, files, tool windows, actions, and settings.


def parse(s, i):
    j = i
    while j < len(s) and s[j].isdigit():
        j += 1
    return int(s[i:j]), s[j], j+1


def compare(s1, s2):
    m, n = len(s1), len(s2)
    i, j = 0, 0
    x, y = 0, 0
    c1, c2 = '', ''

    while i < m and j < n:
        if x == 0:
            x, c1, i = parse(s1, i)
        if y == 0:
            y, c2, j = parse(s2, j)

        if c1 < c2:
            return -1
        if c1 > c2:
            return 1

        if x > y:
            x, y = x - y, 0
        elif x < y:
            x, y = 0, y-x
        else:
            x, y = 0, 0

    if i == m and j == n:
        if x > 0:
            return 1
        if y > 0:
            return -1
        return 0

    if i < m:
        return 1

    return -1


# Press the green button in the gutter to run the script.
if __name__ == '__main__':
    print(compare('31A2A1B', '33A1B'))
    print(compare('1A2B', '1A'))
    print(compare('1A2A1B', '2B'))

# See PyCharm help at https://www.jetbrains.com/help/pycharm/
