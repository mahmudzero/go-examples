// Package foo provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package foo

import (
	"encoding/json"
	"fmt"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// CopilotUsageMetrics Summary of Copilot usage.
type CopilotUsageMetrics struct {
	// Breakdown Breakdown of Copilot code completions usage by language and editor
	Breakdown *[]CopilotUsageMetrics_Breakdown_Item `json:"breakdown"`

	// Day The date for which the usage metrics are reported, in `YYYY-MM-DD` format.
	Day openapi_types.Date `json:"day"`

	// TotalAcceptancesCount The total number of Copilot code completion suggestions accepted by users.
	TotalAcceptancesCount *int `json:"total_acceptances_count,omitempty"`

	// TotalActiveChatUsers The total number of users who interacted with Copilot Chat in the IDE during the day specified.
	TotalActiveChatUsers *int `json:"total_active_chat_users,omitempty"`

	// TotalActiveUsers The total number of users who were shown Copilot code completion suggestions during the day specified.
	TotalActiveUsers *int `json:"total_active_users,omitempty"`

	// TotalChatAcceptances The total instances of users who accepted code suggested by Copilot Chat in the IDE (panel and inline).
	TotalChatAcceptances *int `json:"total_chat_acceptances,omitempty"`

	// TotalChatTurns The total number of chat turns (prompt and response pairs) sent between users and Copilot Chat in the IDE.
	TotalChatTurns *int `json:"total_chat_turns,omitempty"`

	// TotalLinesAccepted The total number of lines of code completions accepted by users.
	TotalLinesAccepted *int `json:"total_lines_accepted,omitempty"`

	// TotalLinesSuggested The total number of lines of code completions suggested by Copilot.
	TotalLinesSuggested *int `json:"total_lines_suggested,omitempty"`

	// TotalSuggestionsCount The total number of Copilot code completion suggestions shown to users.
	TotalSuggestionsCount *int `json:"total_suggestions_count,omitempty"`
}

// CopilotUsageMetrics_Breakdown_Item Breakdown of Copilot usage by editor for this language
type CopilotUsageMetrics_Breakdown_Item struct {
	// AcceptancesCount The number of Copilot suggestions accepted by users in the editor specified during the day specified.
	AcceptancesCount *int `json:"acceptances_count,omitempty"`

	// ActiveUsers The number of users who were shown Copilot completion suggestions in the editor specified during the day specified.
	ActiveUsers *int `json:"active_users,omitempty"`

	// Editor The editor in which Copilot suggestions were shown to users for the specified language.
	Editor *string `json:"editor,omitempty"`

	// Language The language in which Copilot suggestions were shown to users in the specified editor.
	Language *string `json:"language,omitempty"`

	// LinesAccepted The number of lines of code accepted by users in the editor specified during the day specified.
	LinesAccepted *int `json:"lines_accepted,omitempty"`

	// LinesSuggested The number of lines of code suggested by Copilot in the editor specified during the day specified.
	LinesSuggested *int `json:"lines_suggested,omitempty"`

	// SuggestionsCount The number of Copilot suggestions shown to users in the editor specified during the day specified.
	SuggestionsCount     *int                   `json:"suggestions_count,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// Getter for additional properties for CopilotUsageMetrics_Breakdown_Item. Returns the specified
// element and whether it was found
func (a CopilotUsageMetrics_Breakdown_Item) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for CopilotUsageMetrics_Breakdown_Item
func (a *CopilotUsageMetrics_Breakdown_Item) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for CopilotUsageMetrics_Breakdown_Item to handle AdditionalProperties
func (a *CopilotUsageMetrics_Breakdown_Item) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["acceptances_count"]; found {
		err = json.Unmarshal(raw, &a.AcceptancesCount)
		if err != nil {
			return fmt.Errorf("error reading 'acceptances_count': %w", err)
		}
		delete(object, "acceptances_count")
	}

	if raw, found := object["active_users"]; found {
		err = json.Unmarshal(raw, &a.ActiveUsers)
		if err != nil {
			return fmt.Errorf("error reading 'active_users': %w", err)
		}
		delete(object, "active_users")
	}

	if raw, found := object["editor"]; found {
		err = json.Unmarshal(raw, &a.Editor)
		if err != nil {
			return fmt.Errorf("error reading 'editor': %w", err)
		}
		delete(object, "editor")
	}

	if raw, found := object["language"]; found {
		err = json.Unmarshal(raw, &a.Language)
		if err != nil {
			return fmt.Errorf("error reading 'language': %w", err)
		}
		delete(object, "language")
	}

	if raw, found := object["lines_accepted"]; found {
		err = json.Unmarshal(raw, &a.LinesAccepted)
		if err != nil {
			return fmt.Errorf("error reading 'lines_accepted': %w", err)
		}
		delete(object, "lines_accepted")
	}

	if raw, found := object["lines_suggested"]; found {
		err = json.Unmarshal(raw, &a.LinesSuggested)
		if err != nil {
			return fmt.Errorf("error reading 'lines_suggested': %w", err)
		}
		delete(object, "lines_suggested")
	}

	if raw, found := object["suggestions_count"]; found {
		err = json.Unmarshal(raw, &a.SuggestionsCount)
		if err != nil {
			return fmt.Errorf("error reading 'suggestions_count': %w", err)
		}
		delete(object, "suggestions_count")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for CopilotUsageMetrics_Breakdown_Item to handle AdditionalProperties
func (a CopilotUsageMetrics_Breakdown_Item) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.AcceptancesCount != nil {
		object["acceptances_count"], err = json.Marshal(a.AcceptancesCount)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'acceptances_count': %w", err)
		}
	}

	if a.ActiveUsers != nil {
		object["active_users"], err = json.Marshal(a.ActiveUsers)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'active_users': %w", err)
		}
	}

	if a.Editor != nil {
		object["editor"], err = json.Marshal(a.Editor)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'editor': %w", err)
		}
	}

	if a.Language != nil {
		object["language"], err = json.Marshal(a.Language)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'language': %w", err)
		}
	}

	if a.LinesAccepted != nil {
		object["lines_accepted"], err = json.Marshal(a.LinesAccepted)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'lines_accepted': %w", err)
		}
	}

	if a.LinesSuggested != nil {
		object["lines_suggested"], err = json.Marshal(a.LinesSuggested)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'lines_suggested': %w", err)
		}
	}

	if a.SuggestionsCount != nil {
		object["suggestions_count"], err = json.Marshal(a.SuggestionsCount)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'suggestions_count': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}
