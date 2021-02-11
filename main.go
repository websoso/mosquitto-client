package main

import "github.com/onesosoo/mosquitto-client/mc"

func main() {

	o := mc.NewOperation("sample_name")
	o.Execute("sample_content")

}
