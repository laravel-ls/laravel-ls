package protocol

import "encoding/json"

// ProgressToken is the progress token provided by the client or server.
type ProgressToken struct {
	name   string
	number int32
}

// MarshalJSON implements json.Marshaler.
func (v *ProgressToken) MarshalJSON() ([]byte, error) {
	if v.name != "" {
		return json.Marshal(v.name)
	}

	return json.Marshal(v.number)
}

// UnmarshalJSON implements json.Unmarshaler.
func (v *ProgressToken) UnmarshalJSON(data []byte) error {
	*v = ProgressToken{}
	if err := json.Unmarshal(data, &v.number); err == nil {
		return nil
	}

	return json.Unmarshal(data, &v.name)
}

// WorkDoneProgressParams is a parameter property of report work done progress.
type WorkDoneProgressParams struct {
	// WorkDoneToken an optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}

// PartialResultParams allows for partial results in responses.
type PartialResultParams struct {
	// PartialResultToken is a token for handling partial result updates.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
}
