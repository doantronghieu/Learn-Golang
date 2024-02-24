package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

// ItemStatus represents the status of a todo item.
type ItemStatus int

const (
	ItemStatusDoing   ItemStatus = iota // Enum value for 'Doing'
	ItemStatusDone                       // Enum value for 'Done'
	ItemStatusDeleted                    // Enum value for 'Deleted'
)

var allItemStatuses = [3]string{"Doing", "Done", "Deleted"}

// String returns the string representation of the ItemStatus.
func (item *ItemStatus) String() string {
	return allItemStatuses[*item]
}

// parseStr2ItemStatus converts a string to the corresponding ItemStatus enum.
func parseStr2ItemStatus(s string) (ItemStatus, error) {
	for i := range allItemStatuses {
		if allItemStatuses[i] == s {
			return ItemStatus(i), nil
		}
	}

	return ItemStatus(0), errors.New("invalid status string")
}

// Scan converts the SQL value to ItemStatus during database scanning.
func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan data from SQL: %s", value)
	}

	v, err := parseStr2ItemStatus(string(bytes))
	if err != nil {
		return fmt.Errorf("failed to scan data from SQL: %s", value)
	}
	*item = v

	return nil
}

// Value converts ItemStatus to the SQL value during database operations.
func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}

	return item.String(), nil
}

// MarshalJSON converts ItemStatus to JSON format.
func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

// UnmarshalJSON converts JSON data to ItemStatus.
func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	itemValue, err := parseStr2ItemStatus(str)
	if err != nil {
		return err
	}

	*item = itemValue

	return nil
}
