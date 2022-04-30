package meta

import (
	"encoding/json"
)

type Meta struct {
	ID          string
	Name        string
	Type        string
	Description string
}

func (m *Meta) UnmarshalJSON(data []byte) (err error) {
	var v struct {
		Meta struct {
			ID          string `json:"id,required"`
			Type        string `json:"type,required"`
			Name        string `json:"name,required"`
			Description string `json:"description,required"`
		} `json:"meta,required"`
	}

	if err = json.Unmarshal(data, &v); err != nil {
		return
	}

	m.ID = v.Meta.ID
	m.Type = v.Meta.Type
	m.Name = v.Meta.Name
	m.Description = v.Meta.Description

	return
}
