package test

import (
	"context"

	api "github.com/SAP/stewardci-core/pkg/apis/steward/v1alpha1"
	"github.com/SAP/stewardci-core/pkg/k8s"
)

// PipelineRunCheck is a check for a PipelineRun
type PipelineRunCheck func(k8s.PipelineRun) bool

// CreatePipelineRunCondition returns a WaitCondition for a pipelineRun with a dedicated PipelineCheck
func CreatePipelineRunCondition(pipelineRun *api.PipelineRun, check PipelineRunCheck) WaitCondition {
	return NewWaitCondition(func(ctx context.Context) (bool, error) {
		fetcher := k8s.NewPipelineRunFetcher(GetClientFactory(ctx))
		pipelineRun, err := fetcher.ByName(pipelineRun.GetNamespace(), pipelineRun.GetName())
		if err != nil {
			return true, err
		}
		result := check(pipelineRun)
		return result, nil
	})
}

// PipelineRunHasStateResult returns a PipelineRunCheck which checks if a pipelineRun has a dedicated result
func PipelineRunHasStateResult(result api.Result) PipelineRunCheck {
	return func(pr k8s.PipelineRun) bool {

		return pr.GetStatus().Result == result
	}
}
