//go:build opennet
// +build opennet

package networkpolicy

import (
	"testing"

	f "github.com/SAP/stewardci-core/test/framework"
)

func Test_PipelineRunOpenNetworkPolicy(t *testing.T) {
	npTest := f.TestPlan{TestBuilder: PipelineRunNetworkOpenPolicy,
		Count: 1,
	}
	f.ExecutePipelineRunTests(t, npTest)
}
