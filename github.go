package main

const usersURL = "https://api.github.com/search/users"

type UsersSearchResult struct {
	TotalCount int     `json:"total_count"`
	Users      []*User `json:"items"`
}

type User struct {
	ID               int     `json:"id"`
	Login            string  `json:"login"`
	AvatarURL        string  `json:"avatar_url"`
	SubscriptionsURL string  `json:"subscriptions_url"`
	OrganizationsURL string  `json:"organizations_url"`
	Score            float64 `json:"score"`
}
