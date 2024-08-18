package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
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
	if viper.GetBool("feature.active") {
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
		err = os.Remove(source)
		if err != nil {
			return "", err
		}
		return uploadTool.UploadImage(result)
	} else {
		return uploadTool.UploadImage(source)
	}
}

func (ft *FeatureTool) Segment(source string) (string, error) {
	uploadTool := UploadTool{}
	if viper.GetBool("feature.active") {
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
		err = os.Remove(source)
		if err != nil {
			return "", err
		}
		return uploadTool.UploadImage(result)
	} else {
		return uploadTool.UploadImage(source)
	}
}

func (ft *FeatureTool) Track(source string) (string, error) {
	uploadTool := UploadTool{}
	if viper.GetBool("feature.active") {
		cmd := exec.Command(
			viper.GetString("feature.pythonPath"),
			viper.GetString("feature.track.trackPath"),
			"--source", source,
			"--yolo-weights", viper.GetString("feature.track.yolo-weights"),
			"--device", strconv.Itoa(viper.GetInt("feature.track.device")),
			"--config-strongsort", viper.GetString("feature.track.config-strongsort"),
			"--save-vid")
		cmd.Stdout = nil
		cmd.Stderr = nil
		err := cmd.Run()
		if err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "跟踪失败",
			}).Error("跟踪失败")
			return "", err
		}
		println("1")
		result, err := findLatestExpPngPath(viper.GetString("feature.track.result"), "mp4")
		if result == "" {
			return "", err
		}
		println("2")
		//outPut, err := generateCopyPath(result)
		err = os.Remove(source)
		if err != nil {
			return "", err
		}
		return uploadTool.UploadVideo(result)
	} else {
		return uploadTool.UploadVideo(source)
	}
}

func findLatestExpPngPath(basePath, types string) (string, error) {
	entries, err := ioutil.ReadDir(basePath)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "读取目录内容失败",
		}).Error("读取目录内容失败")
		return "", err
	}
	var expFolders []string
	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), "exp") {
			expFolders = append(expFolders, entry.Name())
		}
	}
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
	if len(expFolders) == 0 {
		Log.WithFields(logrus.Fields{
			"error_message": "没有找到exp文件夹",
		}).Error("没有找到exp文件夹")
		return "", nil
	}
	latestExpFolder := filepath.Join(basePath, expFolders[0])
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
