package validations

import "time"

type SeasonInput struct {
	AirDate     time.Time `json:"airDate" binding:"required"`
	Description string    `json:"description"`
}

type EpisodeInput struct {
	Name        string    `json:"name" binding:"required"`
	AirDate     time.Time `json:"airDate" binding:"required"`
	Description string    `json:"description"`
	SeasonId    string    `json:"seasonId" binding:"required"`
}

type QuestionInput struct {
	EpisodeId string `json:"episodeId" binding:"required"`
	Question  string `json:"question" binding:"required"`
	Answer    string `json:"answer"`
}
