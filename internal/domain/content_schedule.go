package domain

import (
	"errors"
	"time"
)

// ContentSchedule 도메인 모델
type ContentSchedule struct {
    ID        uint      `json:"id"`
    ContentID uint      `json:"contentId"`
    StartTime time.Time `json:"startTime"`
    EndTime   time.Time `json:"endTime"`
    AdultCount int       `json:"adultCount"`
    ChildCount int       `json:"childCount"`
}

// ContentSchedule 생성자 함수
func NewContentSchedule(contentID uint, startTime, endTime time.Time, adultCount int, childCount int) (*ContentSchedule, error) {
    if startTime.IsZero() || endTime.IsZero() {
        return nil, errors.New("start time and end time cannot be empty")
    }
    if endTime.Before(startTime) {
        return nil, errors.New("end time cannot be before start time")
    }

    return &ContentSchedule{
        ContentID: contentID,
        StartTime: startTime,
        EndTime:   endTime,
        AdultCount: adultCount,
        ChildCount: childCount,
    }, nil
}

// TableName 메서드 오버라이드
func (ContentSchedule) TableName() string {
    return "CONTENT_SCHEDULES" 
}

// ContentSchedule 도메인 메서드

// 일정이 겹치는지 확인
func (cs *ContentSchedule) IsOverlapping(other *ContentSchedule) bool {
    return cs.StartTime.Before(other.EndTime) && cs.EndTime.After(other.StartTime)
}

func (s *ContentSchedule) GetSeatCount (adultCount int, childCount int) int {
    return adultCount+childCount
}



