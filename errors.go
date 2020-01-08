package common

import (
	"errors"
	"fmt"
	"runtime"
)

func RetErr(e error, p ...interface{}) error {
	if e!=nil {
		var s string
		fc:=myCaller(3)
		for _,j:=range p {
			s+=fmt.Sprint(j)+"; "
		}
		err:=errors.New(fc+"()"+" {"+s+"} -> "+e.Error())
		return err
	}
	return e
}

func myCaller(i int) string {

	// we get the callers as uintptrs - but we just need 1
	fpcs := make([]uintptr, 1)

	// skip 3 levels to get to the caller of whoever called Caller()
	n := runtime.Callers(i, fpcs)
	if n == 0 {
		return "n/a" // proper error her would be better
	}

	// get the info of the actual function that's in the pointer
	fun := runtime.FuncForPC(fpcs[0]-1)
	if fun == nil {
		return "n/a"
	}

	// return its name
	return fun.Name()
}