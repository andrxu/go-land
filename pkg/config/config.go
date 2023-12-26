package config

import (
	"encoding/json"
	"time"
)

type KafkaConfig struct {
	Id       string        `json:"id"`
	Interval time.Duration `json:"interval"`
}

func (c *KafkaConfig) UnmarshalJSON(data []byte) error {
	type Alias KafkaConfig
	var err error

	aux := &struct {
		Interval string `json:"interval"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Do our "custom" conversion from a plain string to a Duration
	if c.Interval, err = time.ParseDuration(aux.Interval); err != nil {
		return err
	}
	return nil
}
