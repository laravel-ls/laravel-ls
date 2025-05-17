package protocol

import "encoding/json"

// A token that can be used to report work done progress.
// Can be a string or a number.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#progress
type ProgressToken struct {
	name   string
	number int32
}

func (v *ProgressToken) MarshalJSON() ([]byte, error) {
	if v.name != "" {
		return json.Marshal(v.name)
	}

	return json.Marshal(v.number)
}

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
