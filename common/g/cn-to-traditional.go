package g

import "github.com/teamlint/opencc"

func cn2tradiation(s string)(string,error){
	cc,err := opencc.New("s2t")
	if err != nil{
		return "",err
	}
	return cc.Convert(s)
}
