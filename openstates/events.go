package openstates

import (
	"github.com/sunlightlabs/go-sunlight/internal"
)

type EventParticipant struct {
	Chamber         string `json:"chamber"`
	Id              string `json:"id"`
	Participant     string `json:"participant"`
	ParticipantType string `json:"participant_type"`
	Type            string `json:"type"`
}

type Event struct {
	Timestamps
	Sources

	Description  string             `json:"description"`
	Id           string             `json:"id"`
	When         internal.Time      `json:"when"`
	Participants []EventParticipant `json:"participants"`
	Session      string             `json:"session"`
	State        string             `json:"state"`
	Timezone     string             `json:"timezone"`
	Type         string             `json:"type"`
}

func GetEvent(bigId string) (*Event, error) {
	t := &Event{}
	err := internal.GetURL(t, openstatesRoot, map[string]string{}, "events", bigId)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func GetEvents(criteria map[string]string) (*[]Event, error) {
	t := &[]Event{}
	err := internal.GetURL(t, openstatesRoot, criteria, "events")
	if err != nil {
		return nil, err
	}
	return t, nil
}
