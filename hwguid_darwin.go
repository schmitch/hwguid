// +build darwin

package hwguid

import (
	"os"
	"os/exec"
	"strings"
)

// getSysctrlEnv sets LC_ALL=C in a list of env vars for use when running
// sysctl commands (see DoSysctrl).
func getSysctrlEnv(env []string) []string {
	foundLC := false
	for i, line := range env {
		if strings.HasPrefix(line, "LC_ALL") {
			env[i] = "LC_ALL=C"
			foundLC = true
		}
	}
	if !foundLC {
		env = append(env, "LC_ALL=C")
	}
	return env
}

func doSysctrl(mib string) ([]string, error) {
	sysctl, err := exec.LookPath("/usr/sbin/sysctl")
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

func MachineGuid() (string, error) {
	values, err := doSysctrl("kern.uuid")
	if err == nil && len(values) == 1 && values[0] != "" {
		return strings.ToLower(values[0]), nil
	}

	return "", nil
}
