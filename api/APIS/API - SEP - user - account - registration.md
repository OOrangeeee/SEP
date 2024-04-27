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

# 用户/账户

## POST 注册用户账户

POST /users/account

创建一个新的用户账户

> Body 请求参数

```yaml
user-name: orange
user-password: "123456"
user-email: 123@qq.com
user-nickname: 晋晨曦

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» user-name|body|string| 是 |用户名|
|» user-password|body|string| 是 |用户密码，密码判断在前端做|
|» user-email|body|string| 是 |用户邮箱|
|» user-nickname|body|string| 是 |用户昵称|

> 返回示例

> 成功

```json
{
  "success_message": "注册成功，请前往邮箱激活"
}
```

> 请求有误

```json
{
  "error_message": "用户名已被注册"
}
```

```json
{
  "error_message": "用户邮箱错误"
}
```

> 禁止访问

```json
{
  "error_message": "发送邮件太频繁，等待五分钟再试"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|成功|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|请求有误|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|禁止访问|Inline|

### 返回数据结构

状态码 **201**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» success_message|string|true|none|成功信息|成功信息|

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» error_message|string|true|none|错误信息|错误信息|

状态码 **403**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» error_message|string|true|none|错误信息|错误信息|

## PUT 激活用户账户

PUT /users/account/activation/{token}

通过激活码激活用户账户

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|path|string| 是 |激活验证的唯一标识符|

> 返回示例

> 成功

```json
{
  "success_message": "账户激活成功"
}
```

> 请求有误

```json
{
  "error_message": "验证令牌错误"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|成功|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|请求有误|Inline|

### 返回数据结构

状态码 **201**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» success_message|string|true|none|成功信息|成功信息|

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» error_message|string|true|none|错误信息|错误信息|

# 数据模型

