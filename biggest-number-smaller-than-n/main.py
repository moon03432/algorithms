# This is a sample Python script.

# Press ⌃R to execute it or replace it with your code.
# Press Double ⇧ to search everywhere for classes, files, tool windows, actions, and settings.

def construct(numbers):
    n = 0
    for digit in numbers:
        n = 10*n + digit

    return n

def smaller_than_n(n, digits):
    digits.sort()
    #print(digits)
    digits_dict = {}
    for i in range(len(digits)):
        digits_dict[digits[i]] = i
    # print(digits_dict)

    if n < 10:
        for i in range(len(digits)-1, 0, -1):
            if digits[i] < n:
                return digits[i]
        return -1

    # convert n to a list of digits
    numbers = []
    while n > 0:
        n, residu = divmod(n, 10)
        numbers = numbers + [residu]

    numbers.reverse()
    #print(numbers)

    # find first digit that is not in digits
    i = 0
    while i < len(numbers):
        if numbers[i] not in digits_dict:
            break
        i += 1

    # print(i, numbers[i])
    if i == len(numbers):
        # all digits are included, iterate from right to left to find a smaller one
        i = len(numbers)-1
        while i >= 0 and digits_dict[numbers[i]] == 0:
            i -= 1
        # print(i)
        if i == -1:
            numbers[0] = 0
            for j in range(1, len(numbers)):
                numbers[j] = digits[-1]
            return construct(numbers)

        numbers[i] = digits[digits_dict[numbers[i]]-1]
        for j in range(i+1, len(numbers)):
            numbers[j] = digits[-1]

        return construct(numbers)

    if numbers[i] > digits[0]:
        j = 0
        while j+1 < len(digits) and numbers[i] > digits[j+1]:
            j += 1
        numbers[i] = digits[j]
        for j in range(i+1,len(numbers)):
            numbers[j] = digits[-1]

        return construct(numbers)

    # numbers[i] < digits[0]
    j = i-1
    while j >= 0 and digits_dict[numbers[j]] == 0:
        j -= 1

    if j == -1:
        numbers[0] = 0
        for j in range(1, len(numbers)):
            numbers[j] = digits[-1]
        return construct(numbers)

    numbers[j] = digits[digits_dict[numbers[j]]-1]
    for k in range(j+1, len(numbers)):
        numbers[k] = digits[-1]

    return construct(numbers)

# Press the green button in the gutter to run the script.
if __name__ == '__main__':
    print(smaller_than_n(23121, [2, 4, 9]))
    print(smaller_than_n(22222, [2, 4, 9]))
    print(smaller_than_n(22222, [0, 2, 4, 9]))
    print(smaller_than_n(5416, [5, 4, 8, 2]))
    print(smaller_than_n(4413, [4, 5]))


# See PyCharm help at https://www.jetbrains.com/help/pycharm/
