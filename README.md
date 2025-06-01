# go-juju

**A Go client for Juju.**

> This library is currently incomplete and under development. It only contains the necessary parts to run charm integration tests.

## Usage

```go
package integration_test

import (
	"testing"

	"github.com/ellanetworks/core-k8s/integration/juju"
)

const (
	CharmPath       = "../ella-core-k8s_amd64.charm"
	JujuModelName   = "test-model"
	CloudName       = "test-cloud"
	ApplicationName = "ella-core"
	EllaCoreImage   = "ghcr.io/ellanetworks/ella-core:v0.0.15"
)

func TestIntegration(t *testing.T) {
	jujuClient := juju.New()

	err := jujuClient.AddK8s(&juju.AddK8sOptions{
		Name:   CloudName,
		Client: true,
	})
	if err != nil {
		t.Fatalf("Failed to add k8s: %v", err)
	}

	t.Log("K8s cloud is added")

	err = jujuClient.Bootstrap(&juju.BootstrapOptions{
		CloudName: CloudName,
	})
	if err != nil {
		t.Fatalf("Failed to bootstrap: %v", err)
	}

	err = jujuClient.AddModel(&juju.AddModelOptions{
		Name: JujuModelName,
	})
	if err != nil {
		t.Fatalf("Failed to add model: %v", err)
	}

	t.Log("Model is added")

	err = jujuClient.Deploy(&juju.DeployOptions{
		Charm:           CharmPath,
		ApplicationName: ApplicationName,
		Resources: map[string]string{
			"core-image": EllaCoreImage,
		},
	})
	if err != nil {
		t.Fatalf("Failed to deploy: %v", err)
	}

	t.Log("Charm is deployed")
}
```
