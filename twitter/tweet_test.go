package twitter

import "testing"

func Test_TwitterEasy(t *testing.T) {
	user1 := 10
	user2 := 11
	user3 := 12
	user4 := 13
	obj := Constructor()
	obj.PostTweet(user1, 1)
	obj.PostTweet(user3, 21)
	obj.Follow(user2, user1)
	obj.Follow(user2, user3)
	obj.PostTweet(user1, 2)

	obj.PostTweet(user1, 3)
	message1 := obj.GetNewsFeed(user1)
	message2 := obj.GetNewsFeed(user2)
	message3 := obj.GetNewsFeed(user3)

	if len(message1) != 3 || len(message2) != 4 || len(message3) != 1 {
		t.Error("Translate test_TwitterEasy Step1", message1, message2, message3)
	}
	// obj.Unfollow(followerId, followeeId)

	obj.Unfollow(user2, user1)
	message1 = obj.GetNewsFeed(user1)
	message2 = obj.GetNewsFeed(user2)
	message3 = obj.GetNewsFeed(user3)

	if len(message1) != 3 || len(message2) != 1 || len(message3) != 1 {
		t.Error("Translate test_TwitterEasy Step2", message1, message2, message3)
	}

	message4 := obj.GetNewsFeed(user4)
	if len(message4) != 0 {
		t.Error("Translate test_TwitterEasy Step3", message1, message2, message3)
	}

	extramUser1 := 21
	extramUser2 := 22
	obj.Follow(extramUser1, extramUser2)

	obj.PostTweet(extramUser2, 200)
	eMessage1 := obj.GetNewsFeed(extramUser1)
	eMessage2 := obj.GetNewsFeed(extramUser2)
	obj.Follow(extramUser1, extramUser2)
	if len(eMessage1) != 1 || len(eMessage2) != 1 {
		t.Error("Translate test_TwitterEasy Step4", eMessage1, eMessage2)
	}

	obj.Follow(user1, extramUser2)
	message1 = obj.GetNewsFeed(user1)
	eMessage2 = obj.GetNewsFeed(extramUser2)
	if len(message1) != 4 {
		t.Error("Translate test_TwitterEasy Step5", message1, eMessage2)
	}

	obj.Follow(user1, user3)
	message1 = obj.GetNewsFeed(user1)
	message3 = obj.GetNewsFeed(user3)
	if len(message1) != 5 {
		t.Error("Translate test_TwitterEasy Step6", message1, message3)
	}

	noneUser1 := 41
	noneUser2 := 42
	// noneUser3 := 43
	obj.Unfollow(noneUser1, noneUser2)
	obj.Unfollow(noneUser1, extramUser2)
	obj.Follow(noneUser2, noneUser1)
	obj.Unfollow(noneUser2, noneUser1)

	obj.Follow(user1, extramUser1)
	message1 = obj.GetNewsFeed(user1)
	if len(message1) != 5 {
		t.Error("Translate test_TwitterEasy Step7", message1)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
