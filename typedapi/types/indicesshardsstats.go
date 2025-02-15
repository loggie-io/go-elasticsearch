// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/899364a63e7415b60033ddd49d50a30369da26d7

package types

import (
	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// IndicesShardsStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/899364a63e7415b60033ddd49d50a30369da26d7/specification/indices/field_usage_stats/IndicesFieldUsageStatsResponse.ts#L49-L52
type IndicesShardsStats struct {
	AllFields FieldSummary            `json:"all_fields"`
	Fields    map[string]FieldSummary `json:"fields"`
}

func (s *IndicesShardsStats) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "all_fields":
			if err := dec.Decode(&s.AllFields); err != nil {
				return err
			}

		case "fields":
			if s.Fields == nil {
				s.Fields = make(map[string]FieldSummary, 0)
			}
			if err := dec.Decode(&s.Fields); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewIndicesShardsStats returns a IndicesShardsStats.
func NewIndicesShardsStats() *IndicesShardsStats {
	r := &IndicesShardsStats{
		Fields: make(map[string]FieldSummary, 0),
	}

	return r
}
