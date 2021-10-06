/*
 * Generated file models/mynamespace/v4/error/error_model.go.
 *
 * Product version: 1.0.0-SNAPSHOT
 *
 * Part of the Sample API project
 *
 * (c) 2021 Nutanix Inc.  All rights reserved
 *
 */
package error
import (
  "bytes"
  "encoding/json"
  "errors"
  "fmt"
)

type AppMessage struct {
  ObjectType_ *string `json:"$objectType,omitempty"`
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  Code *string `json:"code,omitempty"`
  Locale *string `json:"locale,omitempty"`
  Message *string `json:"message,omitempty"`
  Severity *MessageSeverity `json:"severity,omitempty"`
}

func NewAppMessage() *AppMessage {
  p := new(AppMessage)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "mynamespace.v4.error.AppMessage"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "mynamespace.v4.r0.a1.error.AppMessage"}
  p.UnknownFields_ = map[string]interface{}{}

  p.Locale = new(string)
  *p.Locale = "en_US"


  return p
}




type ErrorResponse struct {
  ObjectType_ *string `json:"$objectType,omitempty"`
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  ErrorItemDiscriminator_ *string `json:"$errorItemDiscriminator,omitempty"`
  Error *OneOfErrorResponseError `json:"error,omitempty"`
}

func NewErrorResponse() *ErrorResponse {
  p := new(ErrorResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "mynamespace.v4.error.ErrorResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "mynamespace.v4.r0.a1.error.ErrorResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ErrorResponse) GetError() interface{} {
  if nil == p.Error {
    return nil
  }
  return p.Error.GetValue()
}

func (p *ErrorResponse) SetError(v interface{}) error {
  if nil == p.Error {
    p.Error = NewOneOfErrorResponseError()
  }
  e := p.Error.SetValue(v)
  if nil == e {
    if nil == p.ErrorItemDiscriminator_ {
      p.ErrorItemDiscriminator_ = new(string)
    }
    *p.ErrorItemDiscriminator_ = *p.Error.Discriminator
  }
  return e
}



type MessageSeverity int

const(
  MESSAGESEVERITY_INFO MessageSeverity = 0
  MESSAGESEVERITY_ERROR MessageSeverity = 1
  MESSAGESEVERITY_WARNING MessageSeverity = 2
  MESSAGESEVERITY_UNKNOWN MessageSeverity = 3
)

// returns the name of the enum given an ordinal number
func (e *MessageSeverity) name(index int) string {
  names := [...]string {
    "INFO",
    "ERROR",
    "WARNING",
    "$UNKNOWN",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *MessageSeverity) index(name string) MessageSeverity {
  names := [...]string {
    "INFO",
    "ERROR",
    "WARNING",
    "$UNKNOWN",
  }
  for idx := range names {
    if names[idx] == name {
      return MessageSeverity(idx)
    }
  }
  return MESSAGESEVERITY_UNKNOWN
}

func (e *MessageSeverity) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for MessageSeverity:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *MessageSeverity) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e MessageSeverity) Ref() *MessageSeverity {
  return &e
}


type SchemaValidationError struct {
  ObjectType_ *string `json:"$objectType,omitempty"`
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  Error *string `json:"error,omitempty"`
  Path *string `json:"path,omitempty"`
  StatusCode *int `json:"statusCode,omitempty"`
  Timestamp *string `json:"timestamp,omitempty"`
  ValidationErrorMessages []SchemaValidationErrorMessage `json:"validationErrorMessages,omitempty"`
}

func NewSchemaValidationError() *SchemaValidationError {
  p := new(SchemaValidationError)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "mynamespace.v4.error.SchemaValidationError"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "mynamespace.v4.r0.a1.error.SchemaValidationError"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




type SchemaValidationErrorMessage struct {
  ObjectType_ *string `json:"$objectType,omitempty"`
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  AttributePath *string `json:"attributePath,omitempty"`
  Location *string `json:"location,omitempty"`
  Message *string `json:"message,omitempty"`
}

func NewSchemaValidationErrorMessage() *SchemaValidationErrorMessage {
  p := new(SchemaValidationErrorMessage)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "mynamespace.v4.error.SchemaValidationErrorMessage"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "mynamespace.v4.r0.a1.error.SchemaValidationErrorMessage"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



type OneOfErrorResponseError struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType202 *SchemaValidationError `json:"-"`
  oneOfType201 []AppMessage `json:"-"`
}

func NewOneOfErrorResponseError() *OneOfErrorResponseError {
  p := new(OneOfErrorResponseError)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfErrorResponseError) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfErrorResponseError is nil"))
  }
  switch v.(type) {
    case SchemaValidationError:
      if nil == p.oneOfType202 {p.oneOfType202 = new(SchemaValidationError)}
      *p.oneOfType202 = v.(SchemaValidationError)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType202.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType202.ObjectType_
    case []AppMessage:
      p.oneOfType201 = v.([]AppMessage)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<mynamespace.v4.error.AppMessage>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<mynamespace.v4.error.AppMessage>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfErrorResponseError) GetValue() interface{} {
  if p.oneOfType202 != nil && *p.oneOfType202.ObjectType_ == *p.Discriminator {
    return *p.oneOfType202
  }
  if "List<mynamespace.v4.error.AppMessage>" == *p.Discriminator {
    return p.oneOfType201
  }
  return nil
}

func (p *OneOfErrorResponseError) UnmarshalJSON(b []byte) error {
  vOneOfType202 := new(SchemaValidationError)
  if err := json.Unmarshal(b, vOneOfType202); err == nil {
    if "mynamespace.v4.error.SchemaValidationError" == *vOneOfType202.ObjectType_ {
          if nil == p.oneOfType202 {p.oneOfType202 = new(SchemaValidationError)}
      *p.oneOfType202 = *vOneOfType202
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType202.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType202.ObjectType_
      return nil
    }
    }
  vOneOfType201 := new([]AppMessage)
  if err := json.Unmarshal(b, vOneOfType201); err == nil {
    if len(*vOneOfType201) == 0 || "mynamespace.v4.error.AppMessage" == *((*vOneOfType201)[0].ObjectType_) {
      p.oneOfType201 = *vOneOfType201
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<mynamespace.v4.error.AppMessage>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<mynamespace.v4.error.AppMessage>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfErrorResponseError"))
}

func (p *OneOfErrorResponseError) MarshalJSON() ([]byte, error) {
  if p.oneOfType202 != nil && *p.oneOfType202.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType202)
  }
  if "List<mynamespace.v4.error.AppMessage>" == *p.Discriminator {
    return json.Marshal(p.oneOfType201)
  }
  return nil, errors.New("No value to marshal for OneOfErrorResponseError")
}

