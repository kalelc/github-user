package main

type User struct {
	ID           int     `json:"id"`
	Login        string  `json:"login"`
	AvatarURL    string  `json:"avatar_url"`
	FollowersURL string  `json:"followers_url"`
	HTMLURL      string  `json:"html_url"`
	Score        float64 `json:"score"`
}
