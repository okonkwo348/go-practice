package main

import (
	"strconv"
)

func add(a,b string)(string){
	c,err1:=strconv.Atoi(a)
	d,err2:=strconv.Atoi(b)
	e:=c+d
	if err1!=nil{
		return "wrong first value"
	}
	if err2!=nil{
		return "wrong second value"
	}
	return strconv.Itoa(e)
	
}
func sub(a,b string)(string){
	c,err1:=strconv.Atoi(a)
	d,err2:=strconv.Atoi(b)
	e:=c-d
	if err1!=nil{
		return "wrong first value"
	}
	if err2!=nil{
		return "wrong second value"
	}
	return strconv.Itoa(e)
	
}

func mul(a,b string)(string){
	c,err1:=strconv.Atoi(a)
	d,err2:=strconv.Atoi(b)
	
	e:=c*d
	if err1!=nil{
		return "wrong first value"
	}
	if err2!=nil{
		return "wrong second value"
	}
	if c==0 || d==0{
		return "0"
	}
	return strconv.Itoa(e)
}

func div(a,b string)(string){
	c,err1:=strconv.Atoi(a)
	d,err2:=strconv.Atoi(b)
	
	e:=c/d
	if err1!=nil{
		return "wrong first value"
	}
	if err2!=nil{
		return "wrong second value"
	}
	if d==0{
		return "can't"
	}
	if c==0{
		return "0"
	}
	
	return strconv.Itoa(e)
}
