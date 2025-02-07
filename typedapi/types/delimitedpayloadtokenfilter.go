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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/delimitedpayloadencoding"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// DelimitedPayloadTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/899364a63e7415b60033ddd49d50a30369da26d7/specification/_types/analysis/token_filters.ts#L68-L72
type DelimitedPayloadTokenFilter struct {
	Delimiter *string                                            `json:"delimiter,omitempty"`
	Encoding  *delimitedpayloadencoding.DelimitedPayloadEncoding `json:"encoding,omitempty"`
	Type      string                                             `json:"type,omitempty"`
	Version   *string                                            `json:"version,omitempty"`
}

func (s *DelimitedPayloadTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "delimiter":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Delimiter = &o

		case "encoding":
			if err := dec.Decode(&s.Encoding); err != nil {
				return err
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDelimitedPayloadTokenFilter returns a DelimitedPayloadTokenFilter.
func NewDelimitedPayloadTokenFilter() *DelimitedPayloadTokenFilter {
	r := &DelimitedPayloadTokenFilter{}

	r.Type = "delimited_payload"

	return r
}
