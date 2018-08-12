# Advantage Shuffle

[Code](../weekly_contest_93/advantage_shuffle.go)

```
Given two arrays A and B of equal size, the advantage of A with respect to B is the number of indices i for which A[i] > B[i].

Return any permutation of A that maximizes its advantage with respect to B.



Example 1:

Input: A = [2,7,11,15], B = [1,10,4,11]
Output: [2,11,7,15]
Example 2:

Input: A = [12,24,8,32], B = [13,25,32,11]
Output: [24,32,8,12]


Note:

1 <= A.length = B.length <= 10000
0 <= A[i] <= 10^9
0 <= B[i] <= 10^9

```

## 解题思路

本题的思路是必须对两个从小到大排序的数组进行筛选，先要对a和b分别排序。然后按照田忌赛马的方式去对局，A的最优策略是尽可能用最小的代价，打败B中最弱小的个体，放弃B中较大的数字，这样可以获得最大收益，核心就是要先排序，然后通过标记的方式，防止重复比较，如果碰到无法战胜的对手，就选择己方最弱的数字来比较.

在排序过程中，必须存储a和b中元素原来的位置，才能将选出的A中元素，放到正确的位置上，这个时候不能丢失原始的位置信息。
