/*
 * Generated file .go.
 *
 * Product version: 1.0.0-SNAPSHOT
 *
 * Part of the Sample API project
 *
 * (c) 2021 Nutanix Inc.  All rights reserved
 *
 */

package routes

import (
  "fmt"
  "net/http"
  "helloworld/service/generated/apis/mynamespace/v4/mymodule"
)

const (
  MYN_package_identifier string = "FF0227B53B8A211A58638A75BA48A253"
)

var hello_endpointsimplWrapper = mymodule.Hello_endpointsimplWrapper{}

type implAccessor interface {
  GetServiceImplementation() interface{}
  SetServiceImplementation(impl interface{})
}

type HandlerImplType int
const (
  HELLO_ENDPOINTSIMPL_WRAPPER HandlerImplType = 0
)

var handlerImplName = map[HandlerImplType] string {
  HELLO_ENDPOINTSIMPL_WRAPPER : "hello_endpointsimplWrapper",
}

var handlerImplMap = map[HandlerImplType] implAccessor {
  HELLO_ENDPOINTSIMPL_WRAPPER : &hello_endpointsimplWrapper,
}

type RouteKey struct {
  Path string
  Method string
}

func HandlerMap() map[RouteKey] func(http.ResponseWriter, *http.Request) {
  r := make(map[RouteKey]func(http.ResponseWriter, *http.Request))
  var route RouteKey
  route = RouteKey{ Path:"/api/mynamespace/v4/mymodule/sayhello", Method:"Get" }
  r[route] = hello_endpointsimplWrapper.SayHello

  return r
}

func ValidateReadinessToRun() {
  unimplementedInterfaces := make([] string, 0)

  for k := range handlerImplMap {
    if handlerImplMap[k].GetServiceImplementation() == nil {
      unimplementedInterfaces = append(unimplementedInterfaces, handlerImplName[k])
    }
  }
  if len(unimplementedInterfaces) > 0 {
    errMsg := fmt.Sprintf("The following interfaces have not been implemented:%v", unimplementedInterfaces)
    panic(errMsg)
  }
}

func SetServiceImplementation(implType HandlerImplType, impl interface{}) {
  handlerImplMap[implType].SetServiceImplementation(impl)
}
