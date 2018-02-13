// +build linux

package hwguid

import (
	"strings"
)

func MachineGuid() (string, error) {
	sysProductUUID := hostSys("class/dmi/id/product_uuid")
	switch {
	case pathExists(sysProductUUID):
		lines, err := readLines(sysProductUUID)
		if err == nil && len(lines) > 0 && lines[0] != "" {
			return strings.ToLower(lines[0]), nil
		}
		fallthrough
	default:
		values, err := doSysctrl("kernel.random.boot_id")
		if err == nil && len(values) == 1 && values[0] != "" {
			return strings.ToLower(values[0]), nil
		}
		return "", err
	}
}