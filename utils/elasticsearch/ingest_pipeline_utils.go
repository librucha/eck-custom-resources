package elasticsearch

import (
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/xco-sk/eck-custom-resources/apis/es.eck/v1alpha1"
	"github.com/xco-sk/eck-custom-resources/utils"
	ctrl "sigs.k8s.io/controller-runtime"
)

func DeleteIngestPipeline(esClient *elasticsearch.Client, ingestPipelineId string) (ctrl.Result, error) {
	res, err := esClient.Ingest.DeletePipeline(ingestPipelineId)
	if err != nil || res.IsError() {
		return utils.GetRequeueResult(), err
	}
	return ctrl.Result{}, nil
}

func UpsertIngestPipeline(esClient *elasticsearch.Client, ingestPipeline v1alpha1.IngestPipeline) (ctrl.Result, error) {
	pipelineId := ingestPipeline.Spec.PipelineName
	if pipelineId == "" {
		pipelineId = ingestPipeline.Name
	}
	res, err := esClient.Ingest.PutPipeline(pipelineId, strings.NewReader(ingestPipeline.Spec.Body))

	if err != nil || res.IsError() {
		return utils.GetRequeueResult(), GetClientErrorOrResponseError(err, res)
	}

	return ctrl.Result{}, nil
}
