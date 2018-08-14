# Super Egg Drop

[Code](../weekly_contest_97/super_egg_drop.go)

```
You are given K eggs, and you have access to a building with N floors from 1 to N.

Each egg is identical in function, and if an egg breaks, you cannot drop it again.

You know that there exists a floor F with 0 <= F <= N such that any egg dropped at a floor higher than F will break, and any egg dropped at or below floor F will not break.

Each move, you may take an egg (if you have an unbroken one) and drop it from any floor X (with 1 <= X <= N).

Your goal is to know with certainty what the value of F is.

What is the minimum number of moves that you need to know with certainty what F is, regardless of the initial value of F?



Example 1:

Input: K = 1, N = 2
Output: 2
Explanation:
Drop the egg from floor 1.  If it breaks, we know with certainty that F = 0.
Otherwise, drop the egg from floor 2.  If it breaks, we know with certainty that F = 1.
If it didn't break, then we know with certainty F = 2.
Hence, we needed 2 moves in the worst case to know what F is with certainty.
Example 2:

Input: K = 2, N = 6
Output: 3
Example 3:

Input: K = 3, N = 14
Output: 4


Note:

1 <= K <= 100
1 <= N <= 10000
```

## 解题思路：

个人觉得，核心思路是，如果只有一个鸡蛋，那么不能在找到之前损失任何鸡蛋，这个时候只能从当前还没有探索过的最第一层楼开始尝试，直到鸡蛋破裂，正好就是分界点。

然后，通过二分搜索，可以实现快速的检索，在鸡蛋数量大于一个的情况下， 可以使用二分搜索的方式，这样得到的探索次数是最少的。
