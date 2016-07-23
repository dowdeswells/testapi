package domain

import (
    "time"
)

type AddScheduledAmountCmd struct {
    EndDate             time.Time         `json:"endDate"`
    UsageAmount         int64             `json:"usageAmount"`
}

type RemoveScheduledAmountCmd struct {
    EndDate             time.Time         `json:"endDate"`
}

