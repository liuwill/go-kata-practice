# Koko Eating Bananas

[ Code ](../weekly_contest_94/eating_bananas.go)

```
Koko loves to eat bananas.  There are N piles of bananas, the i-th pile has piles[i] bananas.  The guards have gone and will come back in H hours.

Koko can decide her bananas-per-hour eating speed of K.  Each hour, she chooses some pile of bananas, and eats K bananas from that pile.  If the pile has less than K bananas, she eats all of them instead, and won't eat any more bananas during this hour.

Koko likes to eat slowly, but still wants to finish eating all the bananas before the guards come back.

Return the minimum integer K such that she can eat all the bananas within H hours.



Example 1:

Input: piles = [3,6,7,11], H = 8
Output: 4
Example 2:

Input: piles = [30,11,23,4,20], H = 5
Output: 30
Example 3:

Input: piles = [30,11,23,4,20], H = 6
Output: 23


Note:

1 <= piles.length <= 10^4
piles.length <= H <= 10^9
1 <= piles[i] <= 10^9

```

# 解题思路：

初一看没有头绪，怎么样才能找到这个最小值，经过思考之后，结论是没有任何捷径可以走，找到最小值的最笨办法就是一个一个的遍历所有可能的答案，然后找到最小的符合条件的值。这个值显然在一个区间内，有序区间内最快的检索就是二分查找，二分查找的时间复杂度是O(lg N)，每遍历一个元素，都需要检查这个元素是否是可行的方案，检查的方式就是遍历piles数组每一个元素，算出需要的小时数，然后看是否符合条件，这个算法需要时间复杂度O(M)，所以，最后的时间复杂度就是O(M * lg N).

那么第二步就是如何，确定最大的可行值和最小可行值，最大值是很明确的，piles中最大那个值就是需要的最多次数，按照那个值，一定可以吃完香蕉。最小值不是那么明显，piles中最小的值显然不符合条件，因为存在limit很大的情况；1是最差情况，因为一次1个显然是不能再差了，较好的改善是，用sum(piles)的和除以limit，这样是理论上最少需要的速度。

然后就是根据检查是否当前速度可行的结论，决定二分查找的方向，并且对比之后决定什么时候已经找到了最优值。

所有代码完成之后，还需要考虑只有一个元素的情况，这种情况不需要遍历，直接计算就可以。
