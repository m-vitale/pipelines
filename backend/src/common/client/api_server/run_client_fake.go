package api_server

import (
	"fmt"

	workflowapi "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/go-openapi/strfmt"
	runparams "github.com/googleprivate/ml/backend/api/go_http_client/run_client/run_service"
	runmodel "github.com/googleprivate/ml/backend/api/go_http_client/run_model"
)

const (
	RunForDefaultTest     = "RUN_DEFAULT"
	RunForClientErrorTest = "RUN_CLIENT_ERROR"
)

func getDefaultRun(id string, name string) *runmodel.APIRunDetail {
	return &runmodel.APIRunDetail{
		Workflow: getDefaultWorkflowAsString(),
		Run: &runmodel.APIRun{
			CreatedAt: strfmt.NewDateTime(),
			ID:        id,
			Name:      name,
		},
	}
}

type RunClientFake struct{}

func NewRunClientFake() *RunClientFake {
	return &RunClientFake{}
}

func (c *RunClientFake) Get(params *runparams.GetRunV2Params) (*runmodel.APIRunDetail,
	*workflowapi.Workflow, error) {
	switch params.RunID {
	case RunForClientErrorTest:
		return nil, nil, fmt.Errorf(ClientErrorString)
	case RunForDefaultTest:
		return getDefaultRun(params.RunID, "RUN_NAME"), getDefaultWorkflow(), nil
	default:
		return nil, nil, fmt.Errorf(InvalidFakeRequest)
	}
}

func (c *RunClientFake) List(params *runparams.ListRunsParams) (
	[]*runmodel.APIRun, string, error) {
	const (
		FirstToken  = ""
		SecondToken = "SECOND_TOKEN"
		FinalToken  = ""
	)

	token := ""
	if params.PageToken != nil {
		token = *params.PageToken
	}

	switch token {
	case FirstToken:
		return []*runmodel.APIRun{
			getDefaultRun("100", "MY_FIRST_RUN").Run,
			getDefaultRun("101", "MY_SECOND_RUN").Run,
		}, SecondToken, nil
	case SecondToken:
		return []*runmodel.APIRun{
			getDefaultRun("102", "MY_THIRD_RUN").Run,
		}, FinalToken, nil
	default:
		return nil, "", fmt.Errorf(InvalidFakeRequest)
	}
}

func (c *RunClientFake) ListAll(params *runparams.ListRunsParams, maxResultSize int) (
	[]*runmodel.APIRun, error) {
	return listAllForRun(c, params, maxResultSize)
}
