package client

import (
	"fmt"
	"os/exec"
	"testing"
)

var client *K8sGateKeeperClient

func TestMain(m *testing.M) {
	var err error
	client, err = NewK8sGateKeeperClient(true)

	if err != nil {
		fmt.Println(err)
		return
	}

	reset()
	m.Run()
}

func reset() {
	exec.Command("kubectl", "delete", "-f", "testdata/auth.casbin.org_casbinmodels.yaml").CombinedOutput()

	exec.Command("kubectl", "delete", "-f", "testdata/auth.casbin.org_casbinpolicies.yaml").CombinedOutput()

	res, err := exec.Command("kubectl", "apply", "-f", "testdata/auth.casbin.org_casbinmodels.yaml").CombinedOutput()
	if err != nil {
		fmt.Println(string(res), err)
		return
	}
	res, err = exec.Command("kubectl", "apply", "-f", "testdata/auth.casbin.org_casbinpolicies.yaml").CombinedOutput()
	if err != nil {
		fmt.Println(string(res), err)
		return
	}
}
