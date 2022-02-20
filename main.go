package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/atomu21263/atomicgo"
)

// twitter archiveのデータ構造
type Twitter []struct {
	Tweet struct {
		CreatedAt        string   `json:"created_at"`
		DisplayTextRange []string `json:"display_text_range"`
		Entities         struct {
			Hashtags []struct {
				Indices []string `json:"indices"`
				Text    string   `json:"text"`
			} `json:"hashtags"`
			Media []struct {
				DisplayURL    string   `json:"display_url"`
				ExpandedURL   string   `json:"expanded_url"`
				ID            string   `json:"id"`
				IDStr         string   `json:"id_str"`
				Indices       []string `json:"indices"`
				MediaURL      string   `json:"media_url"`
				MediaURLHTTPS string   `json:"media_url_https"`
				Sizes         struct {
					Large struct {
						H      string `json:"h"`
						Resize string `json:"resize"`
						W      string `json:"w"`
					} `json:"large"`
					Medium struct {
						H      string `json:"h"`
						Resize string `json:"resize"`
						W      string `json:"w"`
					} `json:"medium"`
					Small struct {
						H      string `json:"h"`
						Resize string `json:"resize"`
						W      string `json:"w"`
					} `json:"small"`
					Thumb struct {
						H      string `json:"h"`
						Resize string `json:"resize"`
						W      string `json:"w"`
					} `json:"thumb"`
				} `json:"sizes"`
				Type string `json:"type"`
				URL  string `json:"url"`
			} `json:"media"`
			Symbols []interface{} `json:"symbols"`
			Urls    []struct {
				DisplayURL  string   `json:"display_url"`
				ExpandedURL string   `json:"expanded_url"`
				Indices     []string `json:"indices"`
				URL         string   `json:"url"`
			} `json:"urls"`
			UserMentions []struct {
				ID         string   `json:"id"`
				IDStr      string   `json:"id_str"`
				Indices    []string `json:"indices"`
				Name       string   `json:"name"`
				ScreenName string   `json:"screen_name"`
			} `json:"user_mentions"`
		} `json:"entities"`
		ExtendedEntities struct {
			Media []struct {
				AdditionalMediaInfo struct {
					Monetizable bool `json:"monetizable"`
				} `json:"additional_media_info"`
				DisplayURL    string   `json:"display_url"`
				ExpandedURL   string   `json:"expanded_url"`
				ID            string   `json:"id"`
				IDStr         string   `json:"id_str"`
				Indices       []string `json:"indices"`
				MediaURL      string   `json:"media_url"`
				MediaURLHTTPS string   `json:"media_url_https"`
				Sizes         struct {
					Large struct {
						H      string `json:"h"`
						Resize string `json:"resize"`
						W      string `json:"w"`
					} `json:"large"`
					Medium struct {
						H      string `json:"h"`
						Resize string `json:"resize"`
						W      string `json:"w"`
					} `json:"medium"`
					Small struct {
						H      string `json:"h"`
						Resize string `json:"resize"`
						W      string `json:"w"`
					} `json:"small"`
					Thumb struct {
						H      string `json:"h"`
						Resize string `json:"resize"`
						W      string `json:"w"`
					} `json:"thumb"`
				} `json:"sizes"`
				Type      string `json:"type"`
				URL       string `json:"url"`
				VideoInfo struct {
					AspectRatio    []string `json:"aspect_ratio"`
					DurationMillis string   `json:"duration_millis"`
					Variants       []struct {
						Bitrate     string `json:"bitrate"`
						ContentType string `json:"content_type"`
						URL         string `json:"url"`
					} `json:"variants"`
				} `json:"video_info"`
			} `json:"media"`
		} `json:"extended_entities"`
		FavoriteCount        string `json:"favorite_count"`
		Favorited            bool   `json:"favorited"`
		FullText             string `json:"full_text"`
		ID                   string `json:"id"`
		IDStr                string `json:"id_str"`
		InReplyToScreenName  string `json:"in_reply_to_screen_name"`
		InReplyToStatusID    string `json:"in_reply_to_status_id"`
		InReplyToStatusIDStr string `json:"in_reply_to_status_id_str"`
		InReplyToUserID      string `json:"in_reply_to_user_id"`
		InReplyToUserIDStr   string `json:"in_reply_to_user_id_str"`
		Lang                 string `json:"lang"`
		PossiblySensitive    bool   `json:"possibly_sensitive"`
		RetweetCount         string `json:"retweet_count"`
		Retweeted            bool   `json:"retweeted"`
		Source               string `json:"source"`
		Truncated            bool   `json:"truncated"`
	} `json:"tweet"`
}

// メイン処理部
func main() {
	// データを読み取り
	log.Println("- - - - - - - - - - Tweet Archive Loading Start - - - - - - - - - -")
	dataByte, _ := atomicgo.ReadAndCreateFileFlash("./tweet.js")
	log.Println("- - - - - - - - - - Tweet Archive Loading End - - - - - - - - - -")
	var tweets Twitter
	// []bytrを構造体に変換
	log.Println("- - - - - - - - - - Tweet Archive Converting To Struct Start - - - - - - - - - -")
	json.Unmarshal(dataByte, &tweets)
	log.Println("- - - - - - - - - - Tweet Archive Converting To Struct End - - - - - - - - - -")
	log.Println("- - - - - - - - - - Data Time Sort Start - - - - - - - - - -")
	sort.SliceStable(
		tweets,
		func(i, j int) bool {
			tweetTimeA, _ := time.Parse("Mon Jan 2 15:04:05 -0700 2006", tweets[i].Tweet.CreatedAt)
			tweetA, _ := strconv.Atoi(tweetTimeA.Format("20060102150405"))
			tweetTimeB, _ := time.Parse("Mon Jan 2 15:04:05 -0700 2006", tweets[j].Tweet.CreatedAt)
			tweetB, _ := strconv.Atoi(tweetTimeB.Format("20060102150405"))
			return tweetA < tweetB
		},
	)
	log.Println("- - - - - - - - - - Data Time Sort End - - - - - - - - - -")
	// 削除済みRT数
	count := 0

	key, _ := atomicgo.TwitterAPIkeysGet("twitterAPIKeys.json")
	api := atomicgo.TwitterAPISet(key)
	log.Println("- - - - - - - - - - Retweet Delete Start - - - - - - - - - -")
	// データを全部参照
	for _, tweet := range tweets {
		// RTかを確認
		if atomicgo.StringCheck(tweet.Tweet.FullText, "^RT @") {
			// RT時間を入手
			tweetTime, _ := time.Parse("Mon Jan 2 15:04:05 -0700 2006", tweet.Tweet.CreatedAt)
			logTime := tweetTime.Format("2006/01/02 15:04:05")
			// ログに表示
			fmt.Println("====================================================================================================")
			fmt.Printf("URL: https://twitter.com/i/status/%s Time:%s\n", tweet.Tweet.IDStr, logTime)
			fmt.Printf("%s\n", tweet.Tweet.FullText)
			// 削除
			fmt.Printf("Delete Wait...     ")
			time.Sleep(time.Second * 10)
			fmt.Printf("Deleting...     ")
			RetweetID, err := strconv.Atoi(tweet.Tweet.ID)
			atomicgo.PrintError("Failed Convert Str to Int", err)
			_, err = api.UnRetweet(int64(RetweetID), true)
			atomicgo.PrintError("Failed UnRetweet by RetweetID", err)
			//_, err = api.DeleteTweet(int64(RetweetID), true)
			//atomicgo.PrintError("Failed Delete Retweet", err)
			fmt.Printf("Deleted!\n")
			// API上限対策
			time.Sleep(time.Second * 10)
		}
	}
	log.Println("- - - - - - - - - - Retweet Delete End - - - - - - - - - -")
	fmt.Printf("			%v Retweet Deleted:", count)
}
