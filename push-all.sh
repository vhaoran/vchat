#!/bin/bash

up(){
   
   DIR="$1"
   echo "-------->-""$DIR""----"
   cd $DIR
   git add .
   git commit -m "auto commit"
   git push
};




up   "$GOPATH/src/vchat/";
up   "$GOPATH/src/vchatintf/";
up   "$GOPATH/src/vchatuser/";
up   "$GOPATH/src/vchatutil/";
up   "$GOPATH/src/vchatmsg/";
up   "$GOPATH/src/vchatfinance/";
up   "$GOPATH/src/vchatgw/";
up   "$GOPATH/src/vchatws/";

up   "$GOPATH/src/infoser/";
up   "$GOPATH/src/infogw/";





echo "$VER"
