## 1094. Car Pooling

[code](../../weekly_contest/car_pooling.go)

```
You are driving a vehicle that has capacity empty seats initially available for passengers.  The vehicle only drives east (ie. it cannot turn around and drive west.)

Given a list of trips, trip[i] = [num_passengers, start_location, end_location] contains information about the i-th trip: the number of passengers that must be picked up, and the locations to pick them up and drop them off.  The locations are given as the number of kilometers due east from your vehicle's initial location.

Return true if and only if it is possible to pick up and drop off all passengers for all the given trips.


Example 1:

Input: trips = [[2,1,5],[3,3,7]], capacity = 4
Output: false

Example 2:

Input: trips = [[2,1,5],[3,3,7]], capacity = 5
Output: true

Example 3:

Input: trips = [[2,1,5],[3,5,7]], capacity = 3
Output: true

Example 4:

Input: trips = [[3,2,7],[3,7,9],[8,3,9]], capacity = 11
Output: true

Constraints:

trips.length <= 1000
trips[i].length == 3
1 <= trips[i][0] <= 100
0 <= trips[i][1] < trips[i][2] <= 1000
1 <= capacity <= 100000
```

### 解题思路：

#### 解法一

本题虽然看似一个旅行问题，但是实际上是一个计算问题，并不是要去看每一段拼车并且维护中间状态。

这道题的本质是，求出整个旅程从起点到终点上每个点上，如果完成所有拼车任务，在该点上会有多少乘客，如果发现有某个点的乘客数量超过承载能力，那么，就认为是无法完成拼车的，否则，拼车任务可以完成。

这样，就得出了一种最直观的解题思路，先求出旅程的起点和终点，然后两次for循环，第一次从起点开始，到终点结束，每次循环，遍历trips中的每一次拼单，算出当前index的点上有多少乘客，如果大于capacity，就返回false；遍历完所有情况，返回true。

第一种方式需要O(mn)的时间复杂度，m是路程总长度，n是拼车订单数量。

#### 解法二

m*n的时间复杂度，其实就是一种类似n^2的时间复杂度，思考能不能进一步简化，可以参考限制条件，因为旅程的总长度只有1000，可以考虑用一种空间换时间的计算方式，减少循环的次数。

首先初始化一个容量1000的数组，遍历每一单拼成中的所有点，然后给数组中对应位置加上那段拼车的人数，最后得到一个数组，存储了旅程中每一点最大需要同时乘坐的乘客数。只需要逐个检查是否有某个点的乘车人数比capacity大，就可以知道是否能完成拼车。

如果继续考虑权衡取舍的话，可以再多做一次trips上的for循环，先求出路程的起点终点；性能上略有出入，各自再不同的纬度做了权衡取舍。

#### 解法三

在解法二的基础上，还可以更进一步进行优化，把对每次拼车从起点到终点的一次二层嵌套循环去掉。遍历trips的时候，只在start和end位置，记录下那段拼车之后的人数；

然后遍历位置数组，每个位置的乘客数量，都是之前记录的数量，加上上一个位置的数量，这样就不用为了探求拼车起点和终点之间的位置的信息，去多次一层for循环。

这种方法，并不是很直观，需要进行一定抽象，并且看到对起点和终点之间位置做的for循环，是可以通过其他信息替代的，因为那个循环，并没有提供更多信息。

### 总结

在面临一个可以用数组存储的问题，需要根据数据给出某种结论的时候，毫无疑问，第一个思路就是怎么样做for循环，遍历数组，提取原始信息，变成作出决策有价值的，可以显著缩短读取时间的中间数据，然后在计算出的数据上逐层做计算。最后作出决策。
