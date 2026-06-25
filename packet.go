package core

import (
	"encoding/json"
	"io"
)

type Packet struct {
	// Identity
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	Type string `json:"type"`

	FromID    int    `json:"fromid,omitempty"`
	From      string `json:"from,omitempty"`
	ToID      int    `json:"toid,omitempty"`
	To        string `json:"to,omitempty"`
	Scope     string `json:"scope,omitempty"`
	Direction string `json:"direction,omitempty"`
	Content   string `json:"content,omitempty"`
	Time      string `json:"time,omitempty"`
}

func ReadPacket(r io.Reader) (*Packet, error) {
	var p Packet

	if err := json.NewDecoder(r).Decode(&p); err != nil {
		return nil, err
	}

	return &p, nil
}
