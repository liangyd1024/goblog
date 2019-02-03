//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/12/10

package admin

import (
	"goblog/src/controller"
	. "goblog/src/logs"
	"goblog/src/utils/bizerror"
	"goblog/src/utils/datetime"
	"strings"
)

type FileUploadController struct {
	controller.BaseController
}

func (fileUploadController *FileUploadController) URLMapping() {
	fileUploadController.Mapping("UploadEditorImage", fileUploadController.UploadEditorImage)
}

// @router /admin/file/upload/editor [post]
func (fileUploadController *FileUploadController) UploadEditorImage() {
	defer fileUploadController.PanicHandler()

	guid := fileUploadController.GetString("guid", datetime.ParseNowTime(datetime.FM_FULL_DATE_TIME))

	file, header, err := fileUploadController.GetFile("editormd-image-file")
	bizerror.Check(err)

	fileNames := strings.Split(header.Filename, ".")
	uploadUrl := "static/upload/editor/" + guid + "_" + fileNames[0] + "." + fileNames[1]
	Log.Info("call UploadEditorImage file name:%v,size:%v,uploadUrl:%v", header.Filename, header.Size, uploadUrl)
	defer bizerror.Check(file.Close())

	err = fileUploadController.SaveToFile("editormd-image-file", uploadUrl)
	bizerror.Check(err)

	result := make(map[string]interface{})
	result["success"] = 1
	result["message"] = "上传成功"
	result["url"] = fileUploadController.Site(uploadUrl)

	Log.Info("call UploadEditorImage result:%v", result)
	fileUploadController.Data["json"] = result
	fileUploadController.ServeJSON()
}
