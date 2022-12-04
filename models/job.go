package model

import "github.com/msojocs/AutoTask/v1/db"

// Job 任务集合
type Job struct {
	ID          int64
	Name        string
	Description string
	AuthorId    int64
	Active      bool
}

func GetJobListByUserId(userId int64) ([]Job, error) {
	var jobs []Job
	result := db.DB.Where("author_id = ?", userId).Find(&jobs)
	return jobs, result.Error
}

func GetJobRequests() {
	var jobs []Job
	db.DB.Where("at_jobs.active = ?", 1).Joins("JOIN at_job_requests jr ON jr.job_id = at_jobs.id").Joins("JOIN at_requests r ON r.id = jr.request_id AND r.active = ?", 1).Find(&jobs).Limit(10)
}
