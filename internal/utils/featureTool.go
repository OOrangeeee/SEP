package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type FeatureTool struct {
}

func (ft *FeatureTool) Detect(source string) (string, error) {
	uploadTool := UploadTool{}
	cmd := exec.Command(
		viper.GetString("feature.pythonPath"),
		viper.GetString("feature.detect.detectPath"),
		"--weights", viper.GetString("feature.detect.weights"),
		"--source", source)
	cmd.Stdout = nil
	cmd.Stderr = nil
	err := cmd.Run()
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "检测失败",
		}).Error("检测失败")
		return "", err
	}
	result, err := findLatestExpPngPath(viper.GetString("feature.detect.result"), "png")
	if result == "" {
		return "", err
	}
	return uploadTool.UploadImage(result)
}

func (ft *FeatureTool) Segment(source string) (string, error) {
	uploadTool := UploadTool{}
	cmd := exec.Command(
		viper.GetString("feature.pythonPath"),
		viper.GetString("feature.segment.segmentPath"),
		"--model", viper.GetString("feature.segment.model"),
		"--image", source,
		"--result", viper.GetString("feature.segment.result"))
	cmd.Stdout = nil
	cmd.Stderr = nil
	err := cmd.Run()
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "分割失败",
		}).Error("分割失败")
		return "", err
	}
	result, err := findLatestExpPngPath(viper.GetString("feature.segment.result"), "png")
	if result == "" {
		return "", err
	}
	imageTool := ImageTool{}
	err = imageTool.ChangeColorsAndOverlay(source, result)
	if err != nil {
		return "", err
	}
	return uploadTool.UploadImage(result)
}

func findLatestExpPngPath(basePath, types string) (string, error) {
	// 读取目录内容
	entries, err := ioutil.ReadDir(basePath)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "读取目录内容失败",
		}).Error("读取目录内容失败")
		return "", err
	}
	// 过滤并找到所有以exp开头的文件夹
	var expFolders []string
	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), "exp") {
			expFolders = append(expFolders, entry.Name())
		}
	}
	// 按照exp后缀的数字排序
	sort.Slice(expFolders, func(i, j int) bool {
		numI, err := strconv.Atoi(strings.TrimPrefix(expFolders[i], "exp"))
		if err != nil {
			return false
		}
		numJ, err := strconv.Atoi(strings.TrimPrefix(expFolders[j], "exp"))
		if err != nil {
			return true
		}
		return numI > numJ
	})
	// 检查是否有exp文件夹
	if len(expFolders) == 0 {
		Log.WithFields(logrus.Fields{
			"error_message": "没有找到exp文件夹",
		}).Error("没有找到exp文件夹")
		return "", nil
	}
	// 获取数字最大的exp文件夹
	latestExpFolder := filepath.Join(basePath, expFolders[0])
	// 在该文件夹下找到唯一的文件
	pngFiles, err := filepath.Glob(filepath.Join(latestExpFolder, "*."+types))
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "查找文件失败",
		}).Error("查找文件失败")
		return "", err
	}
	if len(pngFiles) != 1 {
		Log.WithFields(logrus.Fields{
			"error_message": "找到多个或没有文件",
		}).Error("找到多个或没有文件")
		return "", nil
	}
	return pngFiles[0], nil
}
