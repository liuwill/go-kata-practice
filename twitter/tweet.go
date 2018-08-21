package twitter

type Twitter struct {
	Fans     map[int][]int
	Followed map[int][]int
	Tweets   map[int][]*Tweet
	Count    int
}

type Tweet struct {
	Id    int
	Owner int
	Count int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		Fans:     make(map[int][]int),
		Followed: make(map[int][]int),
		Tweets:   make(map[int][]*Tweet),
		Count:    0,
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.Tweets[userId]; !ok {
		this.Init(userId)
	}

	tweet := &Tweet{
		Id:    tweetId,
		Owner: userId,
		Count: this.Count,
	}

	this.Tweets[userId] = append(this.Tweets[userId][:], tweet)
	for _, fansId := range this.Fans[userId] {
		this.Tweets[fansId] = append(this.Tweets[fansId][:], tweet)
	}

	this.Count++
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	feeds := []int{}
	if _, ok := this.Tweets[userId]; !ok {
		return feeds
	}

	for i := len(this.Tweets[userId]) - 1; i >= 0 && len(feeds) < 10; i-- {
		feeds = append(feeds[:], this.Tweets[userId][i].Id)
	}

	return feeds
}

func (this *Twitter) Init(userId int) {
	this.Tweets[userId] = []*Tweet{}
	this.Followed[userId] = []int{}
	this.Fans[userId] = []int{}
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.Tweets[followerId]; !ok {
		this.Init(followerId)
	}

	if _, ok := this.Tweets[followeeId]; !ok {
		this.Init(followeeId)
	}

	if this.hasFollowed(followerId, followeeId) {
		return
	}

	this.Fans[followeeId] = append(this.Fans[followeeId][:], followerId)
	this.Followed[followerId] = append(this.Followed[followerId][:], followeeId)

	if len(this.Tweets[followeeId]) <= 0 {
		return
	}

	last := 0
	length := len(this.Tweets[followerId])
	for _, value := range this.Tweets[followeeId] {

		if len(this.Tweets[followerId]) < 1 {
			this.Tweets[followerId] = append(this.Tweets[followerId][:], value)
			continue
		} else if value.Count > this.Tweets[followerId][length-1].Count {
			this.Tweets[followerId] = append(this.Tweets[followerId][:], value)
			continue
		}

		head := last
		tail := len(this.Tweets[followerId])
		for head < tail {
			middle := (head + tail) / 2
			if value.Count > this.Tweets[followerId][middle].Count {
				head = middle + 1
			} else {
				tail = middle
			}
		}

		if value.Id == this.Tweets[followerId][head].Id {
			continue
		}

		// fmt.Printf("%v %v %v => ", followerId, followeeId, value)
		// println(head, tail, last, len(this.Tweets[followerId]))
		tailList := append([]*Tweet{}, this.Tweets[followerId][head:]...)
		this.Tweets[followerId] = append(this.Tweets[followerId][:head], value)
		this.Tweets[followerId] = append(this.Tweets[followerId][:], tailList...)
		last = head
	}
}

func (this *Twitter) hasFollowed(followerId int, followeeId int) bool {
	if find(this.Fans[followeeId], followerId) >= 0 {
		return true
	}

	return false
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := this.Fans[followeeId]; !ok {
		return
	}

	if _, ok := this.Followed[followerId]; !ok {
		return
	}

	followedIndex := find(this.Fans[followeeId], followerId)
	if followedIndex >= 0 {
		this.Fans[followeeId] = append(this.Fans[followeeId][:followedIndex], this.Fans[followeeId][followedIndex+1:]...)
	}

	fansIndex := find(this.Followed[followerId], followeeId)
	if fansIndex >= 0 {
		this.Followed[followeeId] = append(this.Followed[followerId][:fansIndex], this.Followed[followerId][fansIndex+1:]...)
	}

	if len(this.Tweets[followerId]) <= 0 {
		return
	}
	for i := len(this.Tweets[followerId]) - 1; i >= 0; i-- {
		currentTweet := this.Tweets[followerId][i]
		if currentTweet.Owner == followeeId {
			this.Tweets[followerId] = append(this.Tweets[followerId][:i], this.Tweets[followerId][i+1:]...)
		}
	}
}

func find(list []int, item int) int {
	for index, val := range list {
		if val == item {
			return index
		}
	}
	return -1
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
