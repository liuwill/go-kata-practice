# Dungeon Game

[Code](../dungeon_game)

## 解题思路1：
1 首先地图的规则决定了不会产生回路，所以可以放心的进行遍历
2 另外，毫无疑问只有遍历完所有可能的路径之后，才有可能得出最后的结论
3 遍历过程中可以存储一个当前的最小血值，如果有更小的就更新
4 作为一种优化，可以记录每个位置已经探索过的，到达它的未知需要的血值，如果有新的路径到达，比较一下，如果需要更大血值，就可以淘汰当前方案

[Code](../dungeon_game/knight.go)

## 解题思路2：

通过解决思路1之后发现，根据骑士的行走规则，其实有高效的算法，不需要取模拟所有的行走动作，只需要按照特定的顺序遍历一遍所有格子，就可以完成目标。

[Code](../dungeon_game/computer.go)


## 问题：
```
The demons had captured the princess (P) and imprisoned her in the bottom-right corner of a dungeon. The dungeon consists of M x N rooms laid out in a 2D grid. Our valiant knight (K) was initially positioned in the top-left room and must fight his way through the dungeon to rescue the princess.

The knight has an initial health point represented by a positive integer. If at any point his health point drops to 0 or below, he dies immediately.

Some of the rooms are guarded by demons, so the knight loses health (negative integers) upon entering these rooms; other rooms are either empty (0's) or contain magic orbs that increase the knight's health (positive integers).

In order to reach the princess as quickly as possible, the knight decides to move only rightward or downward in each step.



Write a function to determine the knight's minimum initial health so that he is able to rescue the princess.

For example, given the dungeon below, the initial health of the knight must be at least 7 if he follows the optimal path RIGHT-> RIGHT -> DOWN -> DOWN.

-2 (K)	-3	3
-5	-10	1
10	30	-5 (P)


Note:

The knight's health has no upper bound.
Any room can contain threats or power-ups, even the first room the knight enters and the bottom-right room where the princess is imprisoned.
```
