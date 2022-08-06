# [9,-1,-5,0,5,1,0,-4]
# ->
# [-1,-5,-4,0,0,9,5,1]
# (negatives，0，positives)
def partition(numbers):
    index_p = len(numbers)-1
    while index_p >= 0 and numbers[index_p] > 0:
        index_p -= 1

    index_z = index_p
    while index_z >= 0 and numbers[index_z] == 0:
        index_z -= 1

    index_n = 0
    while index_n <= index_z:
        if numbers[index_n] < 0:
            index_n += 1
            continue

        if numbers[index_n] == 0:
            numbers[index_n], numbers[index_z] = numbers[index_z], 0
            index_z -= 1
            continue

        numbers[index_n], numbers[index_z], numbers[index_p] = numbers[index_z], 0, numbers[index_n]
        index_z -= 1
        index_p -= 1

    return numbers


if __name__ == '__main__':
    print(partition([9,-1,-5,0,5,1,0,-4]))