package models

type ScheduleParam struct {
	Model

	StartTime int `json:"startTime"`
	EndTime   int `json:"endTime"`

	Crontab string `json:"crontab"`
}
