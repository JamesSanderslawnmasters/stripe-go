//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// SigmaScheduledQueryRunStatus is the possible values for status for a scheduled query run.
type SigmaScheduledQueryRunStatus string

// List of values that SigmaScheduledQueryRunStatus can take.
const (
	SigmaScheduledQueryRunStatusCanceled  SigmaScheduledQueryRunStatus = "canceled"
	SigmaScheduledQueryRunStatusCompleted SigmaScheduledQueryRunStatus = "completed"
	SigmaScheduledQueryRunStatusFailed    SigmaScheduledQueryRunStatus = "failed"
	SigmaScheduledQueryRunStatusTimedOut  SigmaScheduledQueryRunStatus = "timed_out"
)

// SigmaScheduledQueryRunParams is the set of parameters that can be used when updating a scheduled query run.
type SigmaScheduledQueryRunParams struct {
	Params `form:"*"`
}

// SigmaScheduledQueryRunListParams is the set of parameters that can be used when listing scheduled query runs.
type SigmaScheduledQueryRunListParams struct {
	ListParams `form:"*"`
}
type SigmaScheduledQueryRunError struct {
	Message string `json:"message"`
}

// SigmaScheduledQueryRun is the resource representing a scheduled query run.
type SigmaScheduledQueryRun struct {
	APIResource
	Created              int64                        `json:"created"`
	DataLoadTime         int64                        `json:"data_load_time"`
	Error                *SigmaScheduledQueryRunError `json:"error"`
	File                 *File                        `json:"file"`
	ID                   string                       `json:"id"`
	Livemode             bool                         `json:"livemode"`
	Object               string                       `json:"object"`
	Query                string                       `json:"query"`
	ResultAvailableUntil int64                        `json:"result_available_until"`
	SQL                  string                       `json:"sql"`
	Status               SigmaScheduledQueryRunStatus `json:"status"`
	Title                string                       `json:"title"`
}

// SigmaScheduledQueryRunList is a list of scheduled query runs as retrieved from a list endpoint.
type SigmaScheduledQueryRunList struct {
	APIResource
	ListMeta
	Data []*SigmaScheduledQueryRun `json:"data"`
}

// UnmarshalJSON handles deserialization of a SigmaScheduledQueryRun.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *SigmaScheduledQueryRun) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type sigmaScheduledQueryRun SigmaScheduledQueryRun
	var v sigmaScheduledQueryRun
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = SigmaScheduledQueryRun(v)
	return nil
}
