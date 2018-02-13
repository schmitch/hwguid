// +build linux

package hwguid

func HostSys(combineWith ...string) string {
	return GetEnv("HOST_SYS", "/sys", combineWith...)
}

func doSysctrl(mib string) ([]string, error) {
	sysctl, err := exec.LookPath("/sbin/sysctl")
	if err != nil {
		return []string{}, err
	}
	cmd := exec.Command(sysctl, "-n", mib)
	cmd.Env = getSysctrlEnv(os.Environ())
	out, err := cmd.Output()
	if err != nil {
		return []string{}, err
	}
	v := strings.Replace(string(out), "{ ", "", 1)
	v = strings.Replace(string(v), " }", "", 1)
	values := strings.Fields(string(v))

	return values, nil
}

func HardwareGuid() (string, error) {
	sysProductUUID := common.HostSys("class/dmi/id/product_uuid")
	switch {
	case common.PathExists(sysProductUUID):
		lines, err := common.ReadLines(sysProductUUID)
		if err == nil && len(lines) > 0 && lines[0] != "" {
			ret.HostID = strings.ToLower(lines[0])
			break
		}
		fallthrough
	default:
		values, err := doSysctrl("kernel.random.boot_id")
		if err == nil && len(values) == 1 && values[0] != "" {
			ret.HostID = strings.ToLower(values[0])
		}
	}

	return ret, nil
}