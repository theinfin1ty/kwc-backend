package validations

import "time"

type SeasonInput struct {
	AirDate   time.Time `json:"airDate"`
	Title     string    `json:"title" binding:"required"`
	Subtitle  string    `json:"subtitle"`
	Theme     string    `json:"theme"`
	Thumbnail string    `json:"thumbnail"`
	Url       string    `json:"url"`
}

type EpisodeInput struct {
	Title         string    `json:"title" binding:"required"`
	AirDate       time.Time `json:"airDate"`
	Subtitle      string    `json:"subtitle"`
	Thumbnail     string    `json:"thumbnail"`
	Url           string    `json:"url"`
	SeasonId      string    `json:"seasonId" binding:"required"`
	ContestantIds []string  `json:"contestantIds"`
	WinnerIds     []string  `json:"winnerIds"`
}

type QuestionInput struct {
	Question  string   `json:"question" binding:"required"`
	Answers   []string `json:"answers"`
	Images    []string `json:"images"`
	Hints     []string `json:"hints"`
	EpisodeId string   `json:"episodeId" binding:"required"`
}

type SocialInput struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ContestantInput struct {
	Name    string        `json:"name" binding:"required"`
	Image   string        `json:"image"`
	Socials []SocialInput `json:"socials"`
}
