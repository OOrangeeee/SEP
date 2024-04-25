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

## GET 获取用户账户信息

GET /users/account

通过用户ID获取用户账户信息

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 是 |用户验证身份的令牌|

> 返回示例

> 成功

```json
{
  "success_message": "成功获取用户信息",
  "UserId": 430000200210273860,
  "UserName": "2w4R5z",
  "UserEmail": "g.jttxfhhr@kflqosk.dz"
}
```

> 401 Response

```json
{
  "error_message": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|没有权限|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» success_message|string|true|none|成功信息|成功信息|
|» UserId|integer|true|none|用户id|每个用户对应的用户id|
|» UserName|string|true|none|用户名|唯一的用户名，类似于账号|
|» UserEmail|string|true|none|用户邮箱|用户邮箱，用于验证注册|
|» UserNickName|string|true|none|用户昵称|一个随意好听的名字|

状态码 **401**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» error_message|string|true|none|错误信息|用户令牌过期或错误|

# 数据模型

<h2 id="tocS_UserInfo">UserInfo</h2>

<a id="schemauserinfo"></a>
<a id="schema_UserInfo"></a>
<a id="tocSuserinfo"></a>
<a id="tocsuserinfo"></a>

```json
{
  "UserId": 0,
  "UserName": "string",
  "UserEmail": "string",
  "UserNickName": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|UserId|integer|true|none|用户id|每个用户对应的用户id|
|UserName|string|true|none|用户名|唯一的用户名，类似于账号|
|UserEmail|string|true|none|用户邮箱|用户邮箱，用于验证注册|
|UserNickName|string|true|none|用户昵称|一个随意好听的名字|

