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
	FromColor string `json:"fromcolor,omitempty"`
	ToID      int    `json:"toid,omitempty"`
	To        string `json:"to,omitempty"`
	ToColor   string `json:"tocolor,omitempty"`

	Scope     string `json:"scope,omitempty"`
	Direction string `json:"direction,omitempty"`
	Content   string `json:"content,omitempty"`
	Time      string `json:"time,omitempty"`

	Num       int    `json:"num,omitempty"`
	ColorName string `json:"colorname,omitempty"`
	ColorHex  string `json:"colorhex,omitempty"`
}

func ReadPacket(r io.Reader) (*Packet, error) {
	var p Packet

	if err := json.NewDecoder(r).Decode(&p); err != nil {
		return nil, err
	}

	return &p, nil
}
