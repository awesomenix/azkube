package runtime

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/luis/runtime"

// APIError error information returned by the API
type APIError struct {
	// StatusCode - HTTP Status code
	StatusCode *string `json:"statusCode,omitempty"`
	// Message - Cause of the error.
	Message *string `json:"message,omitempty"`
}

// CompositeChildModel child entity in a LUIS Composite Entity.
type CompositeChildModel struct {
	// Type - Type of child entity.
	Type *string `json:"type,omitempty"`
	// Value - Value extracted by LUIS.
	Value *string `json:"value,omitempty"`
}

// CompositeEntityModel LUIS Composite Entity.
type CompositeEntityModel struct {
	// ParentType - Type/name of parent entity.
	ParentType *string `json:"parentType,omitempty"`
	// Value - Value for composite entity extracted by LUIS.
	Value *string `json:"value,omitempty"`
	// Children - Child entities.
	Children *[]CompositeChildModel `json:"children,omitempty"`
}

// EntityModel an entity extracted from the utterance.
type EntityModel struct {
	// AdditionalProperties - Unmatched properties from the message are deserialized this collection
	AdditionalProperties map[string]interface{} `json:""`
	// Entity - Name of the entity, as defined in LUIS.
	Entity *string `json:"entity,omitempty"`
	// Type - Type of the entity, as defined in LUIS.
	Type *string `json:"type,omitempty"`
	// StartIndex - The position of the first character of the matched entity within the utterance.
	StartIndex *int32 `json:"startIndex,omitempty"`
	// EndIndex - The position of the last character of the matched entity within the utterance.
	EndIndex *int32 `json:"endIndex,omitempty"`
}

// MarshalJSON is the custom marshaler for EntityModel.
func (em EntityModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if em.Entity != nil {
		objectMap["entity"] = em.Entity
	}
	if em.Type != nil {
		objectMap["type"] = em.Type
	}
	if em.StartIndex != nil {
		objectMap["startIndex"] = em.StartIndex
	}
	if em.EndIndex != nil {
		objectMap["endIndex"] = em.EndIndex
	}
	for k, v := range em.AdditionalProperties {
		objectMap[k] = v
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for EntityModel struct.
func (em *EntityModel) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		default:
			if v != nil {
				var additionalProperties interface{}
				err = json.Unmarshal(*v, &additionalProperties)
				if err != nil {
					return err
				}
				if em.AdditionalProperties == nil {
					em.AdditionalProperties = make(map[string]interface{})
				}
				em.AdditionalProperties[k] = additionalProperties
			}
		case "entity":
			if v != nil {
				var entity string
				err = json.Unmarshal(*v, &entity)
				if err != nil {
					return err
				}
				em.Entity = &entity
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				em.Type = &typeVar
			}
		case "startIndex":
			if v != nil {
				var startIndex int32
				err = json.Unmarshal(*v, &startIndex)
				if err != nil {
					return err
				}
				em.StartIndex = &startIndex
			}
		case "endIndex":
			if v != nil {
				var endIndex int32
				err = json.Unmarshal(*v, &endIndex)
				if err != nil {
					return err
				}
				em.EndIndex = &endIndex
			}
		}
	}

	return nil
}

// EntityWithResolution ...
type EntityWithResolution struct {
	// Resolution - Resolution values for pre-built LUIS entities.
	Resolution interface{} `json:"resolution,omitempty"`
	// AdditionalProperties - Unmatched properties from the message are deserialized this collection
	AdditionalProperties map[string]interface{} `json:""`
	// Entity - Name of the entity, as defined in LUIS.
	Entity *string `json:"entity,omitempty"`
	// Type - Type of the entity, as defined in LUIS.
	Type *string `json:"type,omitempty"`
	// StartIndex - The position of the first character of the matched entity within the utterance.
	StartIndex *int32 `json:"startIndex,omitempty"`
	// EndIndex - The position of the last character of the matched entity within the utterance.
	EndIndex *int32 `json:"endIndex,omitempty"`
}

// MarshalJSON is the custom marshaler for EntityWithResolution.
func (ewr EntityWithResolution) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ewr.Resolution != nil {
		objectMap["resolution"] = ewr.Resolution
	}
	if ewr.Entity != nil {
		objectMap["entity"] = ewr.Entity
	}
	if ewr.Type != nil {
		objectMap["type"] = ewr.Type
	}
	if ewr.StartIndex != nil {
		objectMap["startIndex"] = ewr.StartIndex
	}
	if ewr.EndIndex != nil {
		objectMap["endIndex"] = ewr.EndIndex
	}
	for k, v := range ewr.AdditionalProperties {
		objectMap[k] = v
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for EntityWithResolution struct.
func (ewr *EntityWithResolution) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "resolution":
			if v != nil {
				var resolution interface{}
				err = json.Unmarshal(*v, &resolution)
				if err != nil {
					return err
				}
				ewr.Resolution = resolution
			}
		default:
			if v != nil {
				var additionalProperties interface{}
				err = json.Unmarshal(*v, &additionalProperties)
				if err != nil {
					return err
				}
				if ewr.AdditionalProperties == nil {
					ewr.AdditionalProperties = make(map[string]interface{})
				}
				ewr.AdditionalProperties[k] = additionalProperties
			}
		case "entity":
			if v != nil {
				var entity string
				err = json.Unmarshal(*v, &entity)
				if err != nil {
					return err
				}
				ewr.Entity = &entity
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ewr.Type = &typeVar
			}
		case "startIndex":
			if v != nil {
				var startIndex int32
				err = json.Unmarshal(*v, &startIndex)
				if err != nil {
					return err
				}
				ewr.StartIndex = &startIndex
			}
		case "endIndex":
			if v != nil {
				var endIndex int32
				err = json.Unmarshal(*v, &endIndex)
				if err != nil {
					return err
				}
				ewr.EndIndex = &endIndex
			}
		}
	}

	return nil
}

// EntityWithScore ...
type EntityWithScore struct {
	// Score - Associated prediction score for the intent (float).
	Score *float64 `json:"score,omitempty"`
	// AdditionalProperties - Unmatched properties from the message are deserialized this collection
	AdditionalProperties map[string]interface{} `json:""`
	// Entity - Name of the entity, as defined in LUIS.
	Entity *string `json:"entity,omitempty"`
	// Type - Type of the entity, as defined in LUIS.
	Type *string `json:"type,omitempty"`
	// StartIndex - The position of the first character of the matched entity within the utterance.
	StartIndex *int32 `json:"startIndex,omitempty"`
	// EndIndex - The position of the last character of the matched entity within the utterance.
	EndIndex *int32 `json:"endIndex,omitempty"`
}

// MarshalJSON is the custom marshaler for EntityWithScore.
func (ews EntityWithScore) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ews.Score != nil {
		objectMap["score"] = ews.Score
	}
	if ews.Entity != nil {
		objectMap["entity"] = ews.Entity
	}
	if ews.Type != nil {
		objectMap["type"] = ews.Type
	}
	if ews.StartIndex != nil {
		objectMap["startIndex"] = ews.StartIndex
	}
	if ews.EndIndex != nil {
		objectMap["endIndex"] = ews.EndIndex
	}
	for k, v := range ews.AdditionalProperties {
		objectMap[k] = v
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for EntityWithScore struct.
func (ews *EntityWithScore) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "score":
			if v != nil {
				var score float64
				err = json.Unmarshal(*v, &score)
				if err != nil {
					return err
				}
				ews.Score = &score
			}
		default:
			if v != nil {
				var additionalProperties interface{}
				err = json.Unmarshal(*v, &additionalProperties)
				if err != nil {
					return err
				}
				if ews.AdditionalProperties == nil {
					ews.AdditionalProperties = make(map[string]interface{})
				}
				ews.AdditionalProperties[k] = additionalProperties
			}
		case "entity":
			if v != nil {
				var entity string
				err = json.Unmarshal(*v, &entity)
				if err != nil {
					return err
				}
				ews.Entity = &entity
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ews.Type = &typeVar
			}
		case "startIndex":
			if v != nil {
				var startIndex int32
				err = json.Unmarshal(*v, &startIndex)
				if err != nil {
					return err
				}
				ews.StartIndex = &startIndex
			}
		case "endIndex":
			if v != nil {
				var endIndex int32
				err = json.Unmarshal(*v, &endIndex)
				if err != nil {
					return err
				}
				ews.EndIndex = &endIndex
			}
		}
	}

	return nil
}

// IntentModel an intent detected from the utterance.
type IntentModel struct {
	// Intent - Name of the intent, as defined in LUIS.
	Intent *string `json:"intent,omitempty"`
	// Score - Associated prediction score for the intent (float).
	Score *float64 `json:"score,omitempty"`
}

// LuisResult prediction, based on the input query, containing intent(s) and entities.
type LuisResult struct {
	autorest.Response `json:"-"`
	// Query - The input utterance that was analized.
	Query *string `json:"query,omitempty"`
	// AlteredQuery - The corrected utterance (when spell checking was enabled).
	AlteredQuery     *string      `json:"alteredQuery,omitempty"`
	TopScoringIntent *IntentModel `json:"topScoringIntent,omitempty"`
	// Intents - All the intents (and their score) that were detected from utterance.
	Intents *[]IntentModel `json:"intents,omitempty"`
	// Entities - The entities extracted from the utterance.
	Entities *[]EntityModel `json:"entities,omitempty"`
	// CompositeEntities - The composite entities extracted from the utterance.
	CompositeEntities *[]CompositeEntityModel `json:"compositeEntities,omitempty"`
	SentimentAnalysis *Sentiment              `json:"sentimentAnalysis,omitempty"`
}

// Sentiment sentiment of the input utterance.
type Sentiment struct {
	// Label - The polarity of the sentiment, can be positive, neutral or negative.
	Label *string `json:"label,omitempty"`
	// Score - Score of the sentiment, ranges from 0 (most negative) to 1 (most positive).
	Score *float64 `json:"score,omitempty"`
}
