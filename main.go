package main

import "github.com/websoso/mosquitto-client/mc"

func main() {

	o := mc.NewOperation("sample_name")
	o.Execute("sample_content")

}
