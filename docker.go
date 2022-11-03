package docker

import (
	// import built-in packages
	"os"
	"strings"
)

func hasDOckerEnv() bool {
	_, err := os.Open("/.dockerenv")
	// control flow
	if err != nil {
		// return value
		return false
	} else {
		// return value
		return true
	}
}

func hasDockerCGroup() bool {
	bytes, err := os.ReadFile("/proc/self/cgroup")
	// control flow
	if err != nil {
		// return value
		return false
	} else {
		// return value
		return strings.Contains(string(bytes), "docker")

	}
}

func IsDocker() bool {
	// control flow
	if hasDOckerEnv() {
		// return value
		return hasDOckerEnv()
	} else if hasDockerCGroup() {
		// return value
		return hasDockerCGroup()
	} else {
		// return value
		return false
	}
}
