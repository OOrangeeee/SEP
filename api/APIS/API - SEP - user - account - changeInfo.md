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

## PUT 修改用户账户信息

PUT /users/account

通过用户ID修改用户账户信息，用户id从认证码中解析，无需传递

> Body 请求参数

```yaml
user-nickname: 晋晨曦
user-password: hsjakhdka

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 是 |令牌|
|body|body|object| 否 |none|
|» user-nickname|body|string| 是 |必须传递每个值，即使用户没有更改。没更改就传递原值|
|» user-password|body|string| 是 |用户密码|

> 返回示例

> 成功

```json
{
  "success_message": "修改信息成功"
}
```

> 请求有误

```json
{
  "error_message": "参数不全"
}
```

> 没有权限

```json
{
  "error_message": "身份令牌错误或过期"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|成功|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|请求有误|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|没有权限|Inline|

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

# 数据模型

