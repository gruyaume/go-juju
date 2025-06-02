# go-juju

**A Go client for Juju.**

> This library is currently incomplete and under development. It only contains the necessary parts to run charm integration tests.

## Usage

```go
package integration_test

import (
	"testing"

	"github.com/gruyaume/go-juju/juju"
)

const (
	JujuModelName   = "test-model"
	ApplicationName = "ella-core"
	EllaCoreImage   = "ghcr.io/ellanetworks/ella-core:v0.0.15"
	CharmPath       = "./ella-core-k8s_amd64.charm"
)

func TestCharmIntegration(t *testing.T) {
	jujuClient := juju.New()

	err := jujuClient.AddModel(&juju.AddModelOptions{
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
