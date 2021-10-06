/*
 * Generated file models/mynamespace/v4/mymodule/mymodule_model.go.
 *
 * Product version: 1.0.0-SNAPSHOT
 *
 * Part of the Sample API project
 *
 * (c) 2021 Nutanix Inc.  All rights reserved
 *
 */
package mymodule
import (
  "encoding/json"
  "errors"
  "fmt"
  import1 "models/mynamespace/v4/error"
)

type ApiLink struct {
  ObjectType_ *string `json:"$objectType,omitempty"`
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  Href *string `json:"href,omitempty"`
  Rel *string `json:"rel,omitempty"`
}

func NewApiLink() *ApiLink {
  p := new(ApiLink)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "mynamespace.v4.mymodule.ApiLink"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "mynamespace.v4.r0.a1.mymodule.ApiLink"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




type ApiResponseMetadata struct {
  ObjectType_ *string `json:"$objectType,omitempty"`
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  Links []ApiLink `json:"links,omitempty"`
}

func NewApiResponseMetadata() *ApiResponseMetadata {
  p := new(ApiResponseMetadata)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "mynamespace.v4.mymodule.ApiResponseMetadata"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "mynamespace.v4.r0.a1.mymodule.ApiResponseMetadata"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




type Hello struct {
  ObjectType_ *string `json:"$objectType,omitempty"`
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  Value *string `json:"value,omitempty"`
}

func NewHello() *Hello {
  p := new(Hello)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "mynamespace.v4.mymodule.Hello"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "mynamespace.v4.r0.a1.mymodule.Hello"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




type HelloApiResponse struct {
  ObjectType_ *string `json:"$objectType,omitempty"`
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  Data *OneOfHelloApiResponseData `json:"data,omitempty"`
  Metadata *ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewHelloApiResponse() *HelloApiResponse {
  p := new(HelloApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "mynamespace.v4.mymodule.HelloApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "mynamespace.v4.r0.a1.mymodule.HelloApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *HelloApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *HelloApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfHelloApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



type Message struct {
  ObjectType_ *string `json:"$objectType,omitempty"`
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  Body *string `json:"body,omitempty"`
  Severity *string `json:"severity,omitempty"`
}

func NewMessage() *Message {
  p := new(Message)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "mynamespace.v4.mymodule.Message"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "mynamespace.v4.r0.a1.mymodule.Message"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



type OneOfHelloApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType1005 *Hello `json:"-"`
}

func NewOneOfHelloApiResponseData() *OneOfHelloApiResponseData {
  p := new(OneOfHelloApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfHelloApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfHelloApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case Hello:
      if nil == p.oneOfType1005 {p.oneOfType1005 = new(Hello)}
      *p.oneOfType1005 = v.(Hello)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1005.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1005.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfHelloApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType1005 != nil && *p.oneOfType1005.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1005
  }
  return nil
}

func (p *OneOfHelloApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "mynamespace.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType1005 := new(Hello)
  if err := json.Unmarshal(b, vOneOfType1005); err == nil {
    if "mynamespace.v4.mymodule.Hello" == *vOneOfType1005.ObjectType_ {
          if nil == p.oneOfType1005 {p.oneOfType1005 = new(Hello)}
      *p.oneOfType1005 = *vOneOfType1005
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1005.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1005.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfHelloApiResponseData"))
}

func (p *OneOfHelloApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType1005 != nil && *p.oneOfType1005.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1005)
  }
  return nil, errors.New("No value to marshal for OneOfHelloApiResponseData")
}

