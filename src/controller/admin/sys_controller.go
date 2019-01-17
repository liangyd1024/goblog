//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2019/1/6

package admin

import (
	"goblog/src/controller"
	"goblog/src/service"
)

type SysController struct {
	controller.BaseController
}

func (sysController *SysController) MappingURL() {
	sysController.Mapping("RefreshIndexer", sysController.RefreshIndexer)
}

// @router /admin/sys/refresh/indexer
func (sysController *SysController) RefreshIndexer() {
	defer sysController.PanicHandler()

	service.SearchBiz.RefreshFullTextSearcher()

	sysController.BuildSucResponse("RefreshIndexer Successful!!!")
}
