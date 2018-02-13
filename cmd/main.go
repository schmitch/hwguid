package main

import (
	"github.com/schmitch/hwguid"
	"fmt"
)

func main() {
	uuid, _ := hwguid.MachineGuid()
	fmt.Printf("Hardware UUID: %s\n", uuid)
}
