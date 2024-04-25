---
title: Sagacious Eye for Polyp
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.23"

---

# Sagacious Eye for Polyp

Base URLs:

# Authentication

# 功能

## POST 进行检测

POST /detection

上传图片进行检测

> Body 请求参数

```yaml
image: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 是 |令牌|
|body|body|object| 否 |none|
|» image|body|string(binary)| 是 |图片|

> 返回示例

> 成功

```json
{
  "success_message": "检测成功，请前往记录查看结果"
}
```

> 请求有误

```json
{
  "error_message": "图片参数错误"
}
```

> 没有权限

```json
{
  "error_message": "令牌过期或错误"
}
```

> 服务器错误

```json
{
  "error_message": "服务器运行bug，请联系管理员"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|成功|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|请求有误|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|没有权限|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|服务器错误|Inline|

### 返回数据结构

状态码 **201**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» success_message|string|true|none|成功信息|成功信息|

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» error_message|string|true|none|错误信息|错误信息|

状态码 **401**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» error_message|string|true|none|错误信息|错误信息|

状态码 **500**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» error_message|string|true|none|错误信息|错误信息|

# 数据模型

