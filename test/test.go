package main

import (
	"SEP/internal/utils"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	source := "./uploads/1.png"
	uuidTool := utils.UUIDTool{}
	uuid := uuidTool.GenerateUUID()
	ext := filepath.Ext(source)

	// 创建新的文件名(UUID + 原扩展名)
	newFileName := uuid + ext
	newSource := filepath.Join(filepath.Dir(source), newFileName)

	// 重命名文件
	if err := os.Rename(source, newSource); err != nil {
		println(0)
	}

	// 更新 source 为新的文件路径
	source = newSource
	localDir := "./temp/" + uuid
	if err := os.MkdirAll(localDir, os.ModePerm); err != nil {
		println("1")
	}

	sshPass := "sshpass -p 'WZ+34ybGhwC5' "
	cmd := sshPass + fmt.Sprintf("ssh -p 27075 root@connect.yza1.seetacloud.com 'mkdir -p /services/images/%s'", uuid)
	if err := exec.Command("bash", "-c", cmd).Run(); err != nil {
		println("3")
		println(err.Error())
	}

	scpCmd := sshPass + fmt.Sprintf("scp -P 27075 %s root@connect.yza1.seetacloud.com:/services/images/%s/", source, uuid)
	if err := exec.Command("bash", "-c", scpCmd).Run(); err != nil {
		println("4")
		println(err.Error())
	}

	detectCmd := sshPass + fmt.Sprintf("ssh -p 27075 root@connect.yza1.seetacloud.com '/root/miniconda3/bin/python3.9 /sep/ai/ai/Polyp_detection/detect.py --weights /sep/ai/ai/Polyp_detection/weights/best.pt --source /services/images/%s/%s'", uuid, filepath.Base(source))
	if err := exec.Command("bash", "-c", detectCmd).Run(); err != nil {
		println("5")
		println(err.Error())
		println(detectCmd)
	}

	downloadCmd := sshPass + fmt.Sprintf("scp -rP 27075 root@connect.yza1.seetacloud.com:/sep/ai/ai/Polyp_detection/runs/detect %s/", localDir)
	if err := exec.Command("bash", "-c", downloadCmd).Run(); err != nil {
		println("6")
		println(err.Error())
	}
	if err := os.Remove(source); err != nil {
		println("2")
		println(err.Error())
	}
	resultsDir := filepath.Join(localDir, "detect")
	dirs, _ := ioutil.ReadDir(resultsDir)
	for _, d := range dirs {
		if d.IsDir() {
			resultPath := filepath.Join(resultsDir, d.Name(), filepath.Base(source))
			if _, err := os.Stat(resultPath); err == nil {
				cleanupCmd := sshPass + fmt.Sprintf("ssh -p 27075 root@connect.yza1.seetacloud.com 'rm -rf /services/images/%s'", uuid)
				if err := exec.Command("bash", "-c", cleanupCmd).Run(); err != nil {
					println("7")
				}
				println(resultPath)
			}
		}
	}
}
