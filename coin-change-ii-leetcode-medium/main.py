import collections


# https://leetcode.com/problems/coin-change-ii/description/
def change(amount, coins) -> int:
    dp = collections.Counter()
    dp[0] = 1
    for coin in coins:
        for i in range(coin, amount + 1):
            dp[i] += dp[i - coin]

    return dp[amount]


if __name__ == '__main__':
    print(change(3, [1, 2, 5]))
