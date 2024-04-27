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

# 用户/检测记录

## DELETE 删除使用记录

DELETE /users/records/{recordid}

通过用户ID和使用记录ID删除用户使用记录

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|recordid|path|integer| 是 |记录id|
|Authorization|header|string| 是 |令牌|

> 返回示例

> 删除成功

```json
{
  "success_message": "删除成功"
}
```

> 请求有误

```json
{
  "error_message": "记录id与用户id不匹配"
}
```

```json
{
  "error_message": "记录id不存在"
}
```

> 没有权限

```json
{
  "error_message": "身份令牌过期或错误"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|204|[No Content](https://tools.ietf.org/html/rfc7231#section-6.3.5)|删除成功|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|请求有误|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|没有权限|Inline|

### 返回数据结构

状态码 **204**

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

# 数据模型

