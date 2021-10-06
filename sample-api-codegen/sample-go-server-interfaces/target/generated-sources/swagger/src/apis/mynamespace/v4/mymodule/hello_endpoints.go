/*
 * Generated file Hello_endpoints.go.
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
  "net/http"
)

type Hello_endpoints interface {

    /*
      Endpoint path: /mynamespace/v4/mymodule/sayhello
      Description: Get Hello message
      HttpMethod: Get
      RETURN TYPE:

        mynamespace.v4.mymodule.HelloApiResponse - The returned response must be of this type
        Implementation class must make sure that this is the case.
    */
    SayHello(w http.ResponseWriter, r *http.Request)


}

type Hello_endpointsimplWrapper struct {
  // create a binding to the actual interface implementation in the Init method
  // where this is used.
  svcImpl Hello_endpoints
}

  // CREATE IMPLEMENTATION BINDINGS TO ALL THE METHODS IN THE INTERFACE

func (serviceWrapper Hello_endpointsimplWrapper) SayHello(w http.ResponseWriter, r *http.Request) {
  serviceWrapper.svcImpl.SayHello(w, r)
}


func (serviceWrapper *Hello_endpointsimplWrapper) SetServiceImplementation(apiImplementation interface{}) {
  val := apiImplementation.(Hello_endpoints)
  serviceWrapper.svcImpl = val
}

func (serviceWrapper Hello_endpointsimplWrapper) GetServiceImplementation() interface{} {
   return serviceWrapper.svcImpl
}


