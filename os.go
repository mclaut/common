package common

import "runtime"

func EOL() string {
	var ret string
	os:=runtime.GOOS
	switch os {
	case "linux":
		ret="\n"
	case "windows":
		ret="\r\n"
	case "macos":
		ret="\r"
	default:
		ret="\n"
	}

	return ret
}
