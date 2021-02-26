package tool

import (
	"fmt"

	"github.com/sony/sonyflake"
)

// IssueID Unique ID generated using Sony's improved twite snowflake algorithm
// https://github.com/sony/sonyflake
func IssueID() (string, error) {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", id), nil
}
