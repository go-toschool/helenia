package helenia

import (
	"time"
)

// Assist represents a relation between talks - speaker(user) - assistant(user)
type Assist struct {
	ID          string `json:"id" db:"id"`
	TalkID      string `json:"talk_id" db:"talk_id"`
	SpeakerID   string `json:"speaker_id" db:"speaker_id"`
	AssistantID string `json:"assistant_id" db:"assistant_id"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"-" db:"deleted_at"`
}

// AssistQuery represents a possible assits query fields
type AssistQuery struct {
	AssistID    string
	TalkID      string
	SpeakerID   string
	AssistantID string
}

// Assists represents the interface exposed to consume and interact with
// talks assitance
type Assists interface {
	Add(*AssistQuery) (*Assists, error)
	Get(*AssistQuery) (*AssistQuery, error)
	Update(*AssistQuery) (*Assists, error)
	Delete(*AssistQuery) (*Assists, error)
}
