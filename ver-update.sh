#!/bin/bash
exit

up(){
   ver="$1";
   file="$2";
   echo "------------->""$file";

   OS=$(uname -s)

   if [[ $OS = "GNU/Linux" ]]; then 
      sed -E -i "s/github\.com\/vhaoran\/vchat .*/github\.com\/vhaoran\/vchat $ver/g"  "$file";
   else
      sed -E -i ''  "s/github\.com\/vhaoran\/vchat .*/github\.com\/vhaoran\/vchat $ver/g"  "$file";
   fi; 


   echo "-------after update----------------"
   grep "vchat"  "$file"


};


VER=$(git ls-remote --heads |awk '{ print $1}')


up "$VER"  "$GOPATH/src/vchatintf/go.mod";
up "$VER"  "$GOPATH/src/vchatuser/go.mod";
up "$VER"  "$GOPATH/src/vchatutil/go.mod";
up "$VER"  "$GOPATH/src/vchatmsg/go.mod";
up "$VER"  "$GOPATH/src/vchatfinance/go.mod";
up "$VER"  "$GOPATH/src/vchatgw/go.mod";

up "$VER"  "$GOPATH/src/infouser/go.mod";
up "$VER"  "$GOPATH/src/infogw/go.mod";


echo "$VER"
