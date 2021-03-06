package test

import (
	"apiproject/util"
	"testing"
)

/**
测试图片处理
*/
func TestImagingImageUtil(t *testing.T) {
	util.ResizeImage("/home/wangshibiao/test/testimg/src.jpg", "/home/wangshibiao/test/testimg/out_resize.jpg", 640, 960)
	util.CropImageCenter("/home/wangshibiao/test/testimg/out_resize.jpg", "/home/wangshibiao/test/testimg/out_crop.jpg", 640, 362)

	width, height, formatName, err := util.GetImgInfo("/home/wangshibiao/test/testimg/src.jpg")
	println(width, height, formatName, err)
}
