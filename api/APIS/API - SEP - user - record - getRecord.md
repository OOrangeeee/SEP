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

## GET 获取所有用户使用记录

GET /users/records-all

通过用户ID获取该用户的所有使用记录

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 是 |令牌|

> 返回示例

> 成功

```json
{
  "success_message": "获取全部诊断信息成功",
  "records": {
    "0": {
      "RecordId": 220000197602187680,
      "Date": "2013-01-30 20:33:32",
      "Result": "http://jamvmwxtd.tk/baj",
      "PatientName": "Nancy Robinson"
    }
  }
}
```

> 没有权限

```json
{
  "error_message": "验证令牌过期或错误"
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
|» records|any|true|none||none|

*allOf*

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|»» *anonymous*|[RecordInfo](#schemarecordinfo)|false|none||none|
|»»» RecordId|integer|true|none|记录id|每个记录对应的id|
|»»» Date|string|true|none|检测时间|检测时间|
|»»» Result|string|true|none|结果|可能是图片或者视频的链接|
|»»» PatientName|string|true|none|患者姓名|患者姓名|

*and*

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|»» *anonymous*|[RecordInfo](#schemarecordinfo)|false|none||none|
|»»» RecordId|integer|true|none|记录id|每个记录对应的id|
|»»» Date|string|true|none|检测时间|检测时间|
|»»» Result|string|true|none|结果|可能是图片或者视频的链接|
|»»» PatientName|string|true|none|患者姓名|患者姓名|

*and*

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|»» *anonymous*|[RecordInfo](#schemarecordinfo)|false|none||none|
|»»» RecordId|integer|true|none|记录id|每个记录对应的id|
|»»» Date|string|true|none|检测时间|检测时间|
|»»» Result|string|true|none|结果|可能是图片或者视频的链接|
|»»» PatientName|string|true|none|患者姓名|患者姓名|

状态码 **401**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» error_message|string|true|none|错误信息|错误信息|

## GET 获取用户使用记录

GET /users/records/{recordsid}

通过使用记录ID获取用户使用记录

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|recordsid|path|integer| 是 |使用记录id|
|Authorization|header|string| 是 |登录令牌|

> 返回示例

> 成功

```json
{
  "success_message": "获取对应记录成功",
  "RecordId": 140000201710096930,
  "Date": "1970-12-31 13:48:41",
  "Result": "http://prtvzvad.sr/wfcbi",
  "PatientName": "Paul Lee"
}
```

> 请求有误

```json
{
  "error_message": "参数错误"
}
```

> 没有权限

```json
{
  "error_message": "登录令牌错误或过期"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|请求有误|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|没有权限|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» success_message|string|true|none|成功信息|成功信息|
|» RecordId|integer|true|none|记录id|每个记录对应的id|
|» Date|string|true|none|检测时间|检测时间|
|» Result|string|true|none|结果|可能是图片或者视频的链接|
|» PatientName|string|true|none|患者姓名|患者姓名|

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» error_message|string|true|none|错误信息|错误信息|

状态码 **401**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» error_message|string|true|none|错误信息|错误信息|

# 数据模型

<h2 id="tocS_RecordInfo">RecordInfo</h2>

<a id="schemarecordinfo"></a>
<a id="schema_RecordInfo"></a>
<a id="tocSrecordinfo"></a>
<a id="tocsrecordinfo"></a>

```json
{
  "RecordId": 0,
  "Date": "string",
  "Result": "string",
  "PatientName": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|RecordId|integer|true|none|记录id|每个记录对应的id|
|Date|string|true|none|检测时间|检测时间|
|Result|string|true|none|结果|可能是图片或者视频的链接|
|PatientName|string|true|none|患者姓名|患者姓名|

