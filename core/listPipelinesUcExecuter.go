package core

import (
  "github.com/dev-op-spec/engine/core/models"
"github.com/dev-op-spec/engine/core/ports"
)

type listPipelinesUcExecuter interface {
  Execute(
  ) (pipelines []models.PipelineView, err error)
}

func newListPipelinesUcExecuter(
fs ports.Filesys,
yml yamlCodec,
) listPipelinesUcExecuter {

  return &listPipelinesUcExecuterImpl{
    fs:fs,
    yml:yml,
  }

}

type listPipelinesUcExecuterImpl struct {
  fs  ports.Filesys
  yml yamlCodec
}

func (x listPipelinesUcExecuterImpl) Execute(
) (pipelines []models.PipelineView, err error) {

  pipelines = make([]models.PipelineView, 0)

  var pipelineDirNames []string
  pipelineDirNames, err= x.fs.ListNamesOfPipelineDirs()
  if (nil != err) {
    return
  }

  for _, pipelineDirName := range pipelineDirNames {

    var pipelineFileBytes []byte
    pipelineFileBytes, err= x.fs.ReadPipelineFile(pipelineDirName)
    if (nil != err) {
      return
    }

    pipelineFile := pipelineFile{}
    err= x.yml.fromYaml(
      pipelineFileBytes,
      &pipelineFile,
    )
    if (nil != err) {
      return
    }

    pipelineStageViews := []models.PipelineStageView{}

    for _, pipelineStage := range pipelineFile.Stages {
      pipelineStageView := models.NewPipelineStageView(pipelineStage.Name, pipelineStage.Type)
      pipelineStageViews = append(pipelineStageViews, *pipelineStageView)
    }

    pipelineView := models.NewPipelineView(pipelineFile.Description, pipelineDirName, pipelineStageViews)

    pipelines = append(pipelines, *pipelineView)

  }

  return

}
