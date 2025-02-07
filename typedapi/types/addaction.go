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

// AddAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/899364a63e7415b60033ddd49d50a30369da26d7/specification/indices/update_aliases/types.ts#L30-L44
type AddAction struct {
	Alias         *string  `json:"alias,omitempty"`
	Aliases       []string `json:"aliases,omitempty"`
	Filter        *Query   `json:"filter,omitempty"`
	Index         *string  `json:"index,omitempty"`
	IndexRouting  *string  `json:"index_routing,omitempty"`
	Indices       []string `json:"indices,omitempty"`
	IsHidden      *bool    `json:"is_hidden,omitempty"`
	IsWriteIndex  *bool    `json:"is_write_index,omitempty"`
	MustExist     *bool    `json:"must_exist,omitempty"`
	Routing       *string  `json:"routing,omitempty"`
	SearchRouting *string  `json:"search_routing,omitempty"`
}

func (s *AddAction) UnmarshalJSON(data []byte) error {

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

		case "alias":
			if err := dec.Decode(&s.Alias); err != nil {
				return err
			}

		case "aliases":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Aliases = append(s.Aliases, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Aliases); err != nil {
					return err
				}
			}

		case "filter":
			if err := dec.Decode(&s.Filter); err != nil {
				return err
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return err
			}

		case "index_routing":
			if err := dec.Decode(&s.IndexRouting); err != nil {
				return err
			}

		case "indices":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Indices = append(s.Indices, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Indices); err != nil {
					return err
				}
			}

		case "is_hidden":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IsHidden = &value
			case bool:
				s.IsHidden = &v
			}

		case "is_write_index":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IsWriteIndex = &value
			case bool:
				s.IsWriteIndex = &v
			}

		case "must_exist":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.MustExist = &value
			case bool:
				s.MustExist = &v
			}

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return err
			}

		case "search_routing":
			if err := dec.Decode(&s.SearchRouting); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewAddAction returns a AddAction.
func NewAddAction() *AddAction {
	r := &AddAction{}

	return r
}
