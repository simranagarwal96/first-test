#!/bin/bash
#
# CHANGE THIS FILE TO PUSH YOUR GO CODE TO A GO REPOSITORY
API_SERVER_SOURCE_PATH=../../../server/ntnx-api-helloworld
DTO_PATH=$API_SERVER_SOURCE_PATH/dto
 
mkdir -p $DTO_PATH
cp -r target/generated-sources/swagger/src/* $DTO_PATH/
