package concourse

import (
	"encoding/json"
	"fmt"
)

func (c *Command) Check() error {
	setupLogging(c.stderr)

	var req CheckRequest
	decoder := json.NewDecoder(c.stdin)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		return fmt.Errorf("invalid payload: %s", err)
	}

	return nil
}
