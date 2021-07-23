// Copyright © 2021 Luke Carr <me+oss@carr.sh>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errorsx

import (
	"encoding/json"
	"net/http"
	"time"
)

// JsonErrorResponse represents a RFC 7807 compliant, JSON-based REST API error response.
type JsonErrorResponse struct {
	Timestamp string                 `json:"timestamp"`        // The timestamp of when the error occurred.
	Type      string                 `json:"type"`             // A URI reference that identifies the error's type.
	Title     string                 `json:"title"`            // A human-readable summary of the error.
	Status    int                    `json:"status"`           // The HTTP status code generated by the error's response.
	Detail    string                 `json:"detail,omitempty"` // A human-readable explanation of this specific error.
	Instance  string                 `json:"instance"`         // A URI reference that identifies this error's occurrence.
	Extra     map[string]interface{} `json:"extra,omitempty"`  // Any additional metadata to accompany the error.
}

// WriteToHttp writes a JsonErrorResponse to a http.ResponseWriter.
func (j JsonErrorResponse) WriteToHttp(w http.ResponseWriter) {
	js, err := json.Marshal(j)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/problem+json")
	w.Write(js)
}

// New creates a new Rfc 7807 compliant error, represented as a JSON response.
func New(status int, typ, title, instance string, extra map[string]interface{}, details ...string) JsonErrorResponse {
	response := JsonErrorResponse{
		Timestamp: time.Now().Format(time.RFC3339),
		Type:      typ,
		Title:     title,
		Status:    status,
		Instance:  instance,
		Extra:     extra,
	}

	if len(details) == 1 {
		response.Detail = details[0]
	}

	return response
}
