package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type PropertyMap map[string]interface{}

//Value - simply returns the JSON-encoded representation of the struct.
func (p PropertyMap) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

//Scan - simply decodes a JSON-encoded value into the struct fields.
func (p *PropertyMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*p, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("Type assertion .(map[string]interface{}) failed.")
	}

	return nil
}

//TransformToPropertyMap - Value v transforms to PropertyMap in des
func TransformToPropertyMap(v interface{}) (des PropertyMap, err error) {
	des = make(PropertyMap)
	j, err := json.Marshal(v)
	if err != nil {
		return
	}

	err = json.Unmarshal(j, &des)
	return
}
