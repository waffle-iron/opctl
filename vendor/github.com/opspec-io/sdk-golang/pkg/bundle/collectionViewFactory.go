package bundle

//go:generate counterfeiter -o ./fakeCollectionViewFactory.go --fake-name fakeCollectionViewFactory ./ collectionViewFactory

import (
	"github.com/opspec-io/sdk-golang/pkg/model"
	"github.com/opspec-io/sdk-golang/util/format"
	"github.com/opspec-io/sdk-golang/util/fs"
	"path"
)

type collectionViewFactory interface {
	Construct(
		collectionBundlePath string,
	) (
		collectionView model.CollectionView,
		err error,
	)
}

func newCollectionViewFactory(
	fileSystem fs.FileSystem,
	opViewFactory opViewFactory,
	yaml format.Format,
) collectionViewFactory {

	return &_collectionViewFactory{
		fileSystem:    fileSystem,
		opViewFactory: opViewFactory,
		yaml:          yaml,
	}

}

type _collectionViewFactory struct {
	fileSystem    fs.FileSystem
	opViewFactory opViewFactory
	yaml          format.Format
}

func (this _collectionViewFactory) Construct(
	collectionBundlePath string,
) (
	collectionView model.CollectionView,
	err error,
) {

	collectionManifestPath := path.Join(collectionBundlePath, NameOfCollectionManifestFile)

	collectionManifestBytes, err := this.fileSystem.GetBytesOfFile(
		collectionManifestPath,
	)
	if nil != err {
		return
	}

	collectionManifest := model.CollectionManifest{}
	err = this.yaml.To(
		collectionManifestBytes,
		&collectionManifest,
	)
	if nil != err {
		return
	}

	var opViews []model.OpView
	childFileInfos, err := this.fileSystem.ListChildFileInfosOfDir(collectionBundlePath)
	if nil != err {
		return
	}

	for _, childFileInfo := range childFileInfos {
		opView, err := this.opViewFactory.Construct(
			path.Join(collectionBundlePath, childFileInfo.Name()),
		)
		if nil == err {
			opViews = append(opViews, opView)
		}
	}

	collectionView = model.CollectionView{
		Description: collectionManifest.Description,
		Name:        collectionManifest.Name,
		Ops:         opViews,
	}

	return

}
