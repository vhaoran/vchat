#!/bin/bash

fn(){
   
   DIR="$1"
   echo "-------->-""$DIR""----"
   cd $DIR
   git pull origin master

};




fn   "$GOPATH/src/vchat/";
fn   "$GOPATH/src/vchatintf/";
fn   "$GOPATH/src/vchatuser/";
fn   "$GOPATH/src/vchatutil/";
fn   "$GOPATH/src/vchatmsg/";
fn   "$GOPATH/src/vchatfinance/";
fn  "$GOPATH/src/vchatgw/";
fn  "$GOPATH/src/vchatws/";

fn   "$GOPATH/src/infouser/";
fn   "$GOPATH/src/infogw/";




echo "------complete------------------"
