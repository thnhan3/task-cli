package cmd

import (
	"encoding/json"
	"time"
)

type Task struct {
	ID        int         `json:"id"`
	Content   string      `json:"content"`
	Status    ETaskStatus `json:"status"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

type ETaskStatus string

const (
	TODO       ETaskStatus = "TODO"
	DONE       ETaskStatus = "DONE"
	IN_PROCESS ETaskStatus = "IN_PROCESS"
)

func (t Task) MarshalJSON() ([]byte, error) {
	type Alias Task
	return json.Marshal(&struct {
		Alias
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	}{
		Alias:     (Alias)(t),
		CreatedAt: t.CreatedAt.Format("02/01/2006 15:04"),
		UpdatedAt: t.UpdatedAt.Format("02/01/2006 15:04"),
	})
}

// UnmarshalJSON tùy chỉnh giải mã JSON của struct Task
func (t *Task) UnmarshalJSON(data []byte) error {
	type Alias Task
	aux := &struct {
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var err error
	t.CreatedAt, err = time.Parse("02/01/2006 15:04", aux.CreatedAt)
	if err != nil {
		return err
	}

	t.UpdatedAt, err = time.Parse("02/01/2006 15:04", aux.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
