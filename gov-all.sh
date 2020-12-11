#!/bin/bash

up(){
   ver="$1";
   file="$2";
   echo "------------->""$file";

   sed -E -i '' "s/github\.com\/vhaoran\/vchat .*/github\.com\/vhaoran\/vchat $ver/g"  "$file";
   echo "-------after update----------------"
   grep "vchat"  "$file"


};


VER=$(git ls-remote --heads |awk '{ print $1}')


# p "$VER"  "$GOPATH/src/vchat/go.mod";
# p "$VER"  "$GOPATH/src/vchatintf/go.mod";
# p "$VER"  "$GOPATH/src/vchatuser/go.mod";
#p "$VER"  "$GOPATH/src/vchatutil/go.mod";
#p "$VER"  "$GOPATH/src/vchatmsg/go.mod";
#p "$VER"  "$GOPATH/src/vchatfinance/go.mod";
#p "$VER"  "$GOPATH/src/vchatgw/go.mod";


# cd "$GOPATH/src/vchat/" && go vet -composites=false ./...
cd "$GOPATH/src/vchatintf/" && go vet -composites=false ./...
cd "$GOPATH/src/vchatuser/" && go vet -composites=false ./...
cd "$GOPATH/src/vchatutil/" && go vet -composites=false ./...
cd "$GOPATH/src/vchatmsg/" && go vet -composites=false ./...
cd "$GOPATH/src/vchatfinance/" && go vet -composites=false ./...
cd "$GOPATH/src/vchatgw/" && go vet -composites=false ./...
cd "$GOPATH/src/vchatws/" && go vet -composites=false ./...



echo "$VER"
