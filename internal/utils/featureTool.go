package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type FeatureTool struct {
}

func (ft *FeatureTool) Detect(source string) (string, error) {
	uploadTool := UploadTool{}
	if viper.GetBool("feature.active") {
		uuidTool := UUIDTool{}
		uuid := uuidTool.GenerateUUID()

		// 获取原文件的扩展名
		ext := filepath.Ext(source)

		// 创建新的文件名(UUID + 原扩展名)
		newFileName := uuid + ext
		newSource := filepath.Join(filepath.Dir(source), newFileName)

		// 重命名文件
		if err := os.Rename(source, newSource); err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "重命名文件失败",
			}).Error("重命名文件失败")
			println("重命名文件失败")
			println(err.Error())
			return "", err
		}

		// 更新 source 为新的文件路径
		source = newSource

		localDir := "./temp/" + uuid
		if err := os.MkdirAll(localDir, os.ModePerm); err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "创建临时目录失败",
			}).Error("创建临时目录失败")
			println("创建临时目录失败")
			println(err.Error())
			return "", err
		}

		// 延迟删除临时目录
		defer func() {
			os.RemoveAll(localDir)
		}()

		sshSecret := viper.GetString("feature.secret")
		sshOpts := "-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null"
		sshPass := fmt.Sprintf("sshpass -p '%s' ", sshSecret)

		cmd := exec.Command("bash", "-c", sshPass+fmt.Sprintf("ssh %s -p 27075 root@connect.yza1.seetacloud.com 'mkdir -p /services/images/%s'", sshOpts, uuid))
		if err := cmd.Run(); err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "在远程服务器创建目录失败",
			}).Error("在远程服务器创建目录失败")
			return "", err
		}

		scpCmd := exec.Command("bash", "-c", sshPass+fmt.Sprintf("scp %s -P 27075 %s root@connect.yza1.seetacloud.com:/services/images/%s/", sshOpts, source, uuid))
		if err := scpCmd.Run(); err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "上传文件到远程服务器失败",
			}).Error("上传文件到远程服务器失败")
			return "", err
		}

		detectCmd := exec.Command("bash", "-c", sshPass+fmt.Sprintf("ssh %s -p 27075 root@connect.yza1.seetacloud.com '/root/miniconda3/bin/python3.9 /sep/ai/ai/Polyp_detection/detect.py --weights /sep/ai/ai/Polyp_detection/weights/best.pt --source /services/images/%s/%s'", sshOpts, uuid, filepath.Base(source)))
		if err := detectCmd.Run(); err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "执行检测命令失败",
			}).Error("执行检测命令失败")
			return "", err
		}

		downloadCmd := exec.Command("bash", "-c", sshPass+fmt.Sprintf("scp %s -r -P 27075 root@connect.yza1.seetacloud.com:/sep/ai/ai/Polyp_detection/runs/detect %s/", sshOpts, localDir))
		if err := downloadCmd.Run(); err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "下载检测结果失败",
			}).Error("下载检测结果失败")
			return "", err
		}

		// 删除原文件
		if err := os.Remove(source); err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "删除原文件失败",
			}).Error("删除原文件失败")
			println("删除原文件失败")
			println(err.Error())
			return "", err
		}

		resultsDir := filepath.Join(localDir, "detect")
		dirs, _ := ioutil.ReadDir(resultsDir)
		for _, d := range dirs {
			if d.IsDir() {
				resultPath := filepath.Join(resultsDir, d.Name(), filepath.Base(source))
				if _, err := os.Stat(resultPath); err == nil {
					// 上传图片
					uploadedURL, err := uploadTool.UploadImage(resultPath)
					if err != nil {
						Log.WithFields(logrus.Fields{
							"error":         err,
							"error_message": "上传结果图片失败",
							"result_path":   resultPath,
						}).Error("上传结果图片失败")
						return "", err
					}

					// 返回上传后的URL
					Log.WithField("uploaded_url", uploadedURL).Info("成功上传图片")
					return uploadedURL, nil
				}
			}
		}
	} else {
		return uploadTool.UploadImage(source)
	}

	Log.WithFields(logrus.Fields{
		"error":         "no result file found",
		"error_message": "未找到结果文件",
	}).Error("未找到结果文件")
	println("未找到结果文件")
	return "", fmt.Errorf("no result file found")
}

func (ft *FeatureTool) Segment(source string) (string, error) {
	uploadTool := UploadTool{}
	if viper.GetBool("feature.active") {
		uuidTool := UUIDTool{}
		uuid := uuidTool.GenerateUUID()

		// 获取原文件的扩展名
		ext := filepath.Ext(source)

		// 创建新的文件名(UUID + 原扩展名)
		newFileName := uuid + ext
		newSource := filepath.Join(filepath.Dir(source), newFileName)

		// 重命名文件
		if err := os.Rename(source, newSource); err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "重命名文件失败",
			}).Error("重命名文件失败")
			println(err.Error())
			return "", err
		}

		// 更新 source 为新的文件路径
		source = newSource

		localDir := "./temp/" + uuid
		if err := os.MkdirAll(localDir, os.ModePerm); err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "创建临时目录失败",
			}).Error("创建临时目录失败")
			println(err.Error())
			return "", err
		}

		// 延迟删除临时目录
		defer func() {
			os.RemoveAll(localDir)
		}()

		sshSecret := viper.GetString("feature.secret")
		sshOpts := "-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null"
		sshPass := fmt.Sprintf("sshpass -p '%s' ", sshSecret)

		cmd := exec.Command("bash", "-c", sshPass+fmt.Sprintf("ssh %s -p 27075 root@connect.yza1.seetacloud.com 'mkdir -p /services/images/%s'", sshOpts, uuid))
		if err := cmd.Run(); err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "在远程服务器创建目录失败",
			}).Error("在远程服务器创建目录失败")
			return "", err
		}

		scpCmd := exec.Command("bash", "-c", sshPass+fmt.Sprintf("scp %s -P 27075 %s root@connect.yza1.seetacloud.com:/services/images/%s/", sshOpts, source, uuid))
		if err := scpCmd.Run(); err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "上传文件到远程服务器失败",
			}).Error("上传文件到远程服务器失败")
			return "", err
		}

		segmentCmd := exec.Command("bash", "-c", sshPass+fmt.Sprintf("ssh %s -p 27075 root@connect.yza1.seetacloud.com '/root/miniconda3/bin/python3.9 /sep/ai/ai/Polyp-PVT-main/run.py --model /sep/ai/ai/Polyp-PVT-main/model_pth/PolypPVT.pth --result /sep/ai/ai/Polyp-PVT-main/result/ --image /services/images/%s/%s'", sshOpts, uuid, filepath.Base(source)))
		if err := segmentCmd.Run(); err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "执行分割命令失败",
			}).Error("执行分割命令失败")
			return "", err
		}

		downloadCmd := exec.Command("bash", "-c", sshPass+fmt.Sprintf("scp %s -r -P 27075 root@connect.yza1.seetacloud.com:/sep/ai/ai/Polyp-PVT-main/result %s/", sshOpts, localDir))
		if err := downloadCmd.Run(); err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "下载分割结果失败",
			}).Error("下载分割结果失败")
			return "", err
		}

		resultsDir := filepath.Join(localDir, "result")
		dirs, _ := ioutil.ReadDir(resultsDir)
		for _, d := range dirs {
			if d.IsDir() {
				resultPath := filepath.Join(resultsDir, d.Name(), filepath.Base(source))
				if _, err := os.Stat(resultPath); err == nil {
					imageTool := ImageTool{}
					err = imageTool.ChangeColorsAndOverlay(source, resultPath)
					if err != nil {
						Log.WithFields(logrus.Fields{
							"error":         err,
							"error_message": "整合失败",
						}).Error("整合失败")
						println(err.Error())
						return "", err
					}
					// 删除原文件
					if err := os.Remove(source); err != nil {
						Log.WithFields(logrus.Fields{
							"error":         err,
							"error_message": "删除原文件失败",
						}).Error("删除原文件失败")
						println(err.Error())
						return "", err
					}
					// 上传图片
					uploadedURL, err := uploadTool.UploadImage(resultPath)
					if err != nil {
						Log.WithFields(logrus.Fields{
							"error":         err,
							"error_message": "上传结果图片失败",
							"result_path":   resultPath,
						}).Error("上传结果图片失败")
						return "", err
					}

					// 返回上传后的URL
					Log.WithField("uploaded_url", uploadedURL).Info("成功上传图片")
					return uploadedURL, nil
				}
			}
		}
	} else {
		return uploadTool.UploadImage(source)
	}

	Log.WithFields(logrus.Fields{
		"error":         "no result file found",
		"error_message": "未找到结果文件",
	}).Error("未找到结果文件")
	println("未找到结果文件")
	return "", fmt.Errorf("no result file found")
}

func (ft *FeatureTool) Track(source string) (string, error) {
	uploadTool := UploadTool{}
	if viper.GetBool("feature.active") {
		uuidTool := UUIDTool{}
		uuid := uuidTool.GenerateUUID()

		// 获取原文件的扩展名
		ext := filepath.Ext(source)

		// 创建新的文件名(UUID + 原扩展名)
		newFileName := uuid + ext
		newSource := filepath.Join(filepath.Dir(source), newFileName)

		// 重命名文件
		if err := os.Rename(source, newSource); err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "重命名文件失败",
			}).Error("重命名文件失败")
			println(err.Error())
			return "", err
		}

		// 更新 source 为新的文件路径
		source = newSource

		localDir := "./temp/" + uuid
		if err := os.MkdirAll(localDir, os.ModePerm); err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "创建临时目录失败",
			}).Error("创建临时目录失败")
			println(err.Error())
			return "", err
		}

		// 延迟删除临时目录
		defer func() {
			os.RemoveAll(localDir)
		}()

		sshSecret := viper.GetString("feature.secret")
		sshOpts := "-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null"
		sshPass := fmt.Sprintf("sshpass -p '%s' ", sshSecret)

		cmd := exec.Command("bash", "-c", sshPass+fmt.Sprintf("ssh %s -p 27075 root@connect.yza1.seetacloud.com 'mkdir -p /services/videos/%s'", sshOpts, uuid))
		if err := cmd.Run(); err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "在远程服务器创建目录失败",
			}).Error("在远程服务器创建目录失败")
			return "", err
		}

		scpCmd := exec.Command("bash", "-c", sshPass+fmt.Sprintf("scp %s -P 27075 %s root@connect.yza1.seetacloud.com:/services/videos/%s/", sshOpts, source, uuid))
		if err := scpCmd.Run(); err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "上传文件到远程服务器失败",
			}).Error("上传文件到远程服务器失败")
			return "", err
		}

		trackCmd := exec.Command("bash", "-c", sshPass+fmt.Sprintf("ssh %s -p 27075 root@connect.yza1.seetacloud.com '/root/miniconda3/bin/python3.9 /sep/ai/ai/StrongSORT-YOLO-main/track_v5.py --yolo-weights /sep/ai/ai/StrongSORT-YOLO-main/weights/best.pt --device 0 --config-strongsort /sep/ai/ai/StrongSORT-YOLO-main/strong_sort/configs/strong_sort.yaml --save-vid --source /services/videos/%s/%s'", sshOpts, uuid, filepath.Base(source)))
		if err := trackCmd.Run(); err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "执行跟踪命令失败",
			}).Error("执行跟踪命令失败")
			return "", err
		}

		downloadCmd := exec.Command("bash", "-c", sshPass+fmt.Sprintf("scp %s -r -P 27075 root@connect.yza1.seetacloud.com:/sep/ai/ai/StrongSORT-YOLO-main/runs/track %s/", sshOpts, localDir))
		if err := downloadCmd.Run(); err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "下载跟踪结果失败",
			}).Error("下载跟踪结果失败")
			return "", err
		}

		// 删除原文件
		if err := os.Remove(source); err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "删除原文件失败",
			}).Error("删除原文件失败")
			println(err.Error())
			return "", err
		}

		resultsDir := filepath.Join(localDir, "track")
		dirs, _ := ioutil.ReadDir(resultsDir)
		for _, d := range dirs {
			if d.IsDir() {
				resultPath := filepath.Join(resultsDir, d.Name(), filepath.Base(source))
				if _, err := os.Stat(resultPath); err == nil {
					// 上传
					uploadedURL, err := uploadTool.UploadVideo(resultPath)
					if err != nil {
						Log.WithFields(logrus.Fields{
							"error":         err,
							"error_message": "上传结果视频失败",
							"result_path":   resultPath,
						}).Error("上传结果视频失败")
						return "", err
					}

					// 返回上传后的URL
					Log.WithField("uploaded_url", uploadedURL).Info("上传结果视频成功")
					return uploadedURL, nil
				}
			}
		}
	} else {
		return uploadTool.UploadImage(source)
	}

	Log.WithFields(logrus.Fields{
		"error":         "no result file found",
		"error_message": "未找到结果文件",
	}).Error("未找到结果文件")
	println("未找到结果文件")
	return "", fmt.Errorf("no result file found")
}
