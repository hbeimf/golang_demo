#! /bin/bash

thrift -out .. --gen go example.thrift

thrift -out .. --gen go call.thrift
