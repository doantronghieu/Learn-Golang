package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

/*
Image represents the structure for storing image information.
Fields:
- Id: Unique identifier for the image.
- Url: URL of the image.
- Width: Width of the image.
- Height: Height of the image.
- CloudName: Cloud service name associated with the image (optional, not persisted in the database).
- Extension: File extension of the image (optional, not persisted in the database).

Function TableName returns the name of the database table associated with the Image model.
*/
type Image struct {
	Id        int    `json:"id" gorm:"column:id"`
	Url       string `json:"url" gorm:"column:url"`
	Width     int    `json:"width" gorm:"column:width"`
	Height    int    `json:"height" gorm:"column:height"`
	CloudName string `json:"cloud_name,omitempty" gorm:"-"`
	Extension string `json:"extension,omitempty" gorm:"-"`
}

/*
TableName returns the name of the "images" table in the database.
*/
func (Image) TableName() string { return "images" }

/*
Fullfill appends the provided domain to the image URL.
Parameters:
- domain: The domain to be appended to the image URL.
*/
func (i *Image) Fullfill(domain string) {
	i.Url = fmt.Sprintf("%s/%s", domain, i.Url)
}

/*
Implements the sql.Scanner interface for converting a database column value
to a Go type and sets it to the Image instance.
Parameters:
- value: The database column value to be scanned and converted to the Image type.
*/
func (i *Image) Scan(value interface{}) error {
	// Convert the value to a byte slice.
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal data from DB:", value))
	}

	// Unmarshal the JSON data into a temporary Image struct.
	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	// Update the current Image instance with the unmarshaled data.
	*i = img
	return nil
}


/*
Implements the driver.Valuer interface for converting the Image instance to
a value that can be stored in a database column.
Returns:
- driver.Value: The value to be stored in the database column.
- error: An error if the conversion fails.
*/
func (i *Image) Value() (driver.Value, error) {
	// If the Image instance is nil, return nil.
	if i == nil {
		return nil, nil
	}
	// Marshal the Image struct into a JSON byte slice.
	return json.Marshal(i)
}