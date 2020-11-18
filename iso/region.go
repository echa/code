// Copyright (c) 2013-2020 KIDTSUNAMI
// Author: alex@kidtsunami.com

package iso

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

// ISO 3166-2 Country subdivision code
type Region string

const (
	RegionUndefined Region = ""
)

func ParseRegion(c string) Region {
	if c == "" {
		return RegionUndefined
	}
	ff := strings.SplitN(c, "-", 2)
	if len(ff) != 2 {
		return RegionUndefined
	}
	if !ParseCountry(ff[0]).IsValid() {
		return RegionUndefined
	}
	return Region(c)
}

func (r Region) IsValid() bool {
	return r != RegionUndefined
}

// Text/JSON conversion
func (r Region) MarshalText() ([]byte, error) {
	return []byte(r), nil
}

func (r *Region) UnmarshalText(data []byte) error {
	rr := ParseRegion(string(data))
	if !rr.IsValid() {
		return fmt.Errorf("iso: invalid ISO region code '%s'", string(data))
	}
	*r = rr
	return nil
}

// SQL conversion
func (r *Region) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		*r = ParseRegion(v)
	case []byte:
		*r = ParseRegion(string(v))
	}
	if !(*r).IsValid() {
		return fmt.Errorf("iso: invalid ISO region code '%v'", value)
	}
	return nil
}

func (r Region) Value() (driver.Value, error) {
	return string(r), nil
}

func (c Region) String() string {
	// TODO
	// if n, ok := region_names[string(c)]; ok {
	// 	return n
	// }
	return string(c)
}

func (r Region) Country() Country {
	if !r.IsValid() {
		return Country("")
	}
	ff := strings.SplitN(string(r), "-", 2)
	return Country(ff[0])
}
