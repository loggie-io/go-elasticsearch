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

	"strconv"

	"encoding/json"
)

// ReindexStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/899364a63e7415b60033ddd49d50a30369da26d7/specification/_global/reindex_rethrottle/types.ts#L37-L51
type ReindexStatus struct {
	Batches              int64    `json:"batches"`
	Created              int64    `json:"created"`
	Deleted              int64    `json:"deleted"`
	Noops                int64    `json:"noops"`
	RequestsPerSecond    float32  `json:"requests_per_second"`
	Retries              Retries  `json:"retries"`
	Throttled            Duration `json:"throttled,omitempty"`
	ThrottledMillis      int64    `json:"throttled_millis"`
	ThrottledUntil       Duration `json:"throttled_until,omitempty"`
	ThrottledUntilMillis int64    `json:"throttled_until_millis"`
	Total                int64    `json:"total"`
	Updated              int64    `json:"updated"`
	VersionConflicts     int64    `json:"version_conflicts"`
}

func (s *ReindexStatus) UnmarshalJSON(data []byte) error {

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

		case "batches":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Batches = value
			case float64:
				f := int64(v)
				s.Batches = f
			}

		case "created":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Created = value
			case float64:
				f := int64(v)
				s.Created = f
			}

		case "deleted":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Deleted = value
			case float64:
				f := int64(v)
				s.Deleted = f
			}

		case "noops":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Noops = value
			case float64:
				f := int64(v)
				s.Noops = f
			}

		case "requests_per_second":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.RequestsPerSecond = f
			case float64:
				f := float32(v)
				s.RequestsPerSecond = f
			}

		case "retries":
			if err := dec.Decode(&s.Retries); err != nil {
				return err
			}

		case "throttled":
			if err := dec.Decode(&s.Throttled); err != nil {
				return err
			}

		case "throttled_millis":
			if err := dec.Decode(&s.ThrottledMillis); err != nil {
				return err
			}

		case "throttled_until":
			if err := dec.Decode(&s.ThrottledUntil); err != nil {
				return err
			}

		case "throttled_until_millis":
			if err := dec.Decode(&s.ThrottledUntilMillis); err != nil {
				return err
			}

		case "total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Total = value
			case float64:
				f := int64(v)
				s.Total = f
			}

		case "updated":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Updated = value
			case float64:
				f := int64(v)
				s.Updated = f
			}

		case "version_conflicts":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.VersionConflicts = value
			case float64:
				f := int64(v)
				s.VersionConflicts = f
			}

		}
	}
	return nil
}

// NewReindexStatus returns a ReindexStatus.
func NewReindexStatus() *ReindexStatus {
	r := &ReindexStatus{}

	return r
}
