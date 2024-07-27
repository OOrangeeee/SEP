# Sagacious Eye for Polyp

> A project based on multimodal high-precision artificial intelligence for the recognition of intestinal polyps.
>
> 一个基于多模态高精度人工智能的智能肠道息肉识别项目。

## 功能

1. 对肠道图片进行AI检测，识别出肠道息肉。
2. 对肠道图片进行息肉AI分割，便于医生诊断。
3. 对肠道视频进行追踪，实时检测息肉。
4. 支持基本用户功能，如登录、注册、修改密码等。
5. 支持管理员功能，如用户管理、数据管理等。
6. 支持诊断报告生成，支持诊断记录存储。

前端部分仍在开发中，后端部分已经完成。

## AI模型

AI模型使用了多种开源模型，包括：

- StrongSORT++：用于视频追踪：https://github.com/dyhBUPT/StrongSORT https://github.com/bharath5673/StrongSORT-YOLO
- YOLOv5：用于息肉检测：https://github.com/ultralytics/yolov5
- Polyp-PVT：用于息肉分割：https://github.com/DengPingFan/Polyp-PVT.git

## 使用

可由Docker-Compose一键部署，或者手动部署。

//TODO：前端开发完成后，补充部署说明。
