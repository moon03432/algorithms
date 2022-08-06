# [a,b,c]->[[a],[b],[c],[a,b],[a,c],[b,c],[a,b,c]]

def combination_generator(elements):
    n = len(elements)
    size = 2 ** n - 1
    res = [0] * size

    def generate(k):
        lst = []
        for e in elements:
            k, r = divmod(k, 2)
            if r > 0:
                lst.append(e)
        return lst

    for i in range(size):
        res[i] = generate(i+1)

    return res


if __name__ == '__main__':
    print(combination_generator(['a','b','c']))