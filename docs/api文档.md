---
title: 默认模块
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
generator: "@tarslib/widdershins v4.0.30"

---

# 默认模块

Base URLs:

# Authentication

# 用户

## POST 注册

POST /user/login

> Body 请求参数

```json
{
  "username": "user9",
  "password": "password123"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET 获取用户信息

GET /user/info

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 否 |none|
|Authorization|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## PUT 更新照片

PUT /user/avatar/upload

> Body 请求参数

```yaml
data: file://D:\vsc-code\tik-tok\testdata\user2\image_969613135063421.png
user_id: u2045086133204815872

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 否 |none|
|Authorization|header|string| 否 |none|
|body|body|object| 是 |none|
|» data|body|string(binary)| 否 |none|
|» user_id|body|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

# 视频

## POST 发布视频

POST /video/publish

> Body 请求参数

```yaml
data: file://D:\vsc-code\tik-tok\testdata\user2\video_1776330013173_jp37bs.mp4
title: 标题
description: 描述

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 是 |none|
|» data|body|string(binary)| 否 |none|
|» title|body|string| 否 |none|
|» description|body|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "base": {
    "status_code": 200,
    "status_msg": "视频发布成功"
  },
  "video_id": "v2045089536433721344"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET 查找视频

GET /video/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 否 |none|
|page_size|query|number| 否 |none|
|page_num|query|number| 否 |none|
|Authorization|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "base": {
    "status_code": 200,
    "status_msg": "获取视频列表成功"
  },
  "videos": [
    {
      "video_id": "v2045089536433721344",
      "author_id": "u2045086133204815872",
      "video_url": "userdata/u2045086133204815872/v2045089536433721344_video_1776330013173_jp37bs.mp4",
      "title": "标题",
      "description": "描述",
      "created_at": "2026-04-17 10:38:35",
      "updated_at": "2026-04-17 10:38:35"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET 搜索视频

GET /video/search

> Body 请求参数

```yaml
{}

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|keyword|query|string| 否 |none|
|page_size|query|number| 否 |none|
|page_num|query|number| 否 |none|
|Authorization|header|string| 否 |none|
|body|body|object| 是 |none|

> 返回示例

> 200 Response

```json
{
  "base": {
    "status_code": 0,
    "status_msg": "string"
  },
  "videos": [
    {
      "video_id": "string",
      "author_id": "string",
      "video_url": "string",
      "title": "string",
      "description": "string",
      "created_at": "string",
      "updated_at": "string"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» base|object|true|none||none|
|»» status_code|integer|true|none||none|
|»» status_msg|string|true|none||none|
|» videos|[object]|true|none||none|
|»» video_id|string|true|none||none|
|»» author_id|string|true|none||none|
|»» video_url|string|true|none||none|
|»» title|string|true|none||none|
|»» description|string|true|none||none|
|»» created_at|string|true|none||none|
|»» updated_at|string|true|none||none|

## GET 获取排行榜

GET /video/popular

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page_size|query|string| 否 |none|
|page_num|query|string| 否 |none|
|Authorization|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "base": {
    "status_code": 0,
    "status_msg": "string"
  },
  "videos": [
    {
      "video_id": "string",
      "author_id": "string",
      "video_url": "string",
      "title": "string",
      "description": "string",
      "created_at": "string",
      "updated_at": "string"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» base|object|true|none||none|
|»» status_code|integer|true|none||none|
|»» status_msg|string|true|none||none|
|» videos|[object]|true|none||none|
|»» video_id|string|true|none||none|
|»» author_id|string|true|none||none|
|»» video_url|string|true|none||none|
|»» title|string|true|none||none|
|»» description|string|true|none||none|
|»» created_at|string|true|none||none|
|»» updated_at|string|true|none||none|

# 互动

## POST 点赞视频

POST /like/action

> Body 请求参数

```yaml
user_id: u2045086133204815872
video_id: v2045089536433721344
action_type: 2

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 是 |none|
|» user_id|body|string| 否 |none|
|» video_id|body|string| 否 |none|
|» action_type|body|number| 否 |none|

> 返回示例

> 200 Response

```json
"你还没有点赞过"
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET 查看喜欢视频列表

GET /like/list

> Body 请求参数

```yaml
{}

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 否 |none|
|page_number|query|number| 否 |none|
|page_size|query|number| 否 |none|
|Authorization|header|string| 否 |none|
|body|body|object| 是 |none|

> 返回示例

> 200 Response

```json
{
  "base_response": {
    "status_code": 0,
    "status_msg": "string"
  },
  "video_list": [
    {
      "video_id": "string",
      "author_id": "string",
      "video_url": "string",
      "title": "string",
      "description": "string",
      "created_at": "string",
      "updated_at": "string"
    }
  ],
  "like_count": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» base_response|object|true|none||none|
|»» status_code|integer|true|none||none|
|»» status_msg|string|true|none||none|
|» video_list|[object]|true|none||none|
|»» video_id|string|false|none||none|
|»» author_id|string|false|none||none|
|»» video_url|string|false|none||none|
|»» title|string|false|none||none|
|»» description|string|false|none||none|
|»» created_at|string|false|none||none|
|»» updated_at|string|false|none||none|
|» like_count|integer|true|none||none|

## POST 评论

POST /comment/pubilsh

> Body 请求参数

```yaml
comment_id: u2045086133204815872
video_id: v2045089536433721344
comment_text: 我的诗山代码哈哈哈哈

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 是 |none|
|» comment_id|body|string| 否 |none|
|» video_id|body|string| 否 |none|
|» comment_text|body|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "base_response": {
    "status_code": 200,
    "status_msg": "success"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET 获取视频评论

GET /comment/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|video_id|query|string| 否 |none|
|page_size|query|string| 否 |none|
|page_number|query|string| 否 |none|
|Authorization|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "comment_list": [
    {
      "id": "c2045135617045565440",
      "user_id": "u2045086133204815872",
      "video_id": "v2045089536433721344",
      "content": "我的诗山代码哈哈哈哈",
      "created_at": "2026-04-17 13:41:41",
      "updated_at": "2026-04-17 13:41:41"
    },
    {
      "id": "c2045135618740064256",
      "user_id": "u2045086133204815872",
      "video_id": "v2045089536433721344",
      "content": "我的诗山代码哈哈哈哈",
      "created_at": "2026-04-17 13:41:42",
      "updated_at": "2026-04-17 13:41:42"
    },
    {
      "id": "c2045135620417785856",
      "user_id": "u2045086133204815872",
      "video_id": "v2045089536433721344",
      "content": "我的诗山代码哈哈哈哈",
      "created_at": "2026-04-17 13:41:42",
      "updated_at": "2026-04-17 13:41:42"
    },
    {
      "id": "c2045135622233919488",
      "user_id": "u2045086133204815872",
      "video_id": "v2045089536433721344",
      "content": "我的诗山代码哈哈哈哈",
      "created_at": "2026-04-17 13:41:43",
      "updated_at": "2026-04-17 13:41:43"
    },
    {
      "id": "c2045135623668371456",
      "user_id": "u2045086133204815872",
      "video_id": "v2045089536433721344",
      "content": "我的诗山代码哈哈哈哈",
      "created_at": "2026-04-17 13:41:43",
      "updated_at": "2026-04-17 13:41:43"
    },
    {
      "id": "c2045135625304150016",
      "user_id": "u2045086133204815872",
      "video_id": "v2045089536433721344",
      "content": "我的诗山代码哈哈哈哈",
      "created_at": "2026-04-17 13:41:43",
      "updated_at": "2026-04-17 13:41:43"
    },
    {
      "id": "c2045135627082534912",
      "user_id": "u2045086133204815872",
      "video_id": "v2045089536433721344",
      "content": "我的诗山代码哈哈哈哈",
      "created_at": "2026-04-17 13:41:44",
      "updated_at": "2026-04-17 13:41:44"
    },
    {
      "id": "c2045135628370186240",
      "user_id": "u2045086133204815872",
      "video_id": "v2045089536433721344",
      "content": "我的诗山代码哈哈哈哈",
      "created_at": "2026-04-17 13:41:44",
      "updated_at": "2026-04-17 13:41:44"
    },
    {
      "id": "c2045135630089850880",
      "user_id": "u2045086133204815872",
      "video_id": "v2045089536433721344",
      "content": "我的诗山代码哈哈哈哈",
      "created_at": "2026-04-17 13:41:44",
      "updated_at": "2026-04-17 13:41:44"
    },
    {
      "id": "c2045135632639987712",
      "user_id": "u2045086133204815872",
      "video_id": "v2045089536433721344",
      "content": "我的诗山代码哈哈哈哈",
      "created_at": "2026-04-17 13:41:45",
      "updated_at": "2026-04-17 13:41:45"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 删除评论

POST /comment/delete

> Body 请求参数

```yaml
comment_id: c2045143046974083072

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 是 |none|
|» comment_id|body|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "base_response": {
    "status_code": 0,
    "status_msg": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» base_response|object|true|none||none|
|»» status_code|integer|true|none||none|
|»» status_msg|string|true|none||none|

# 社交

## POST 关注用户

POST /relation/action

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 否 |none|
|to_user_id|query|string| 否 |none|
|action_type|query|integer| 否 |none|
|Authorization|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "base_response": {
    "status_code": 0,
    "status_msg": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» base_response|object|true|none||none|
|»» status_code|integer|true|none||none|
|»» status_msg|string|true|none||none|

## GET 查询关注列表

GET /following/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 否 |none|
|page_number|query|number| 否 |none|
|page_size|query|number| 否 |none|
|Authorization|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "base_response": {
    "status_code": 0,
    "status_msg": "string"
  },
  "subscriber_list": [
    {
      "id": "string",
      "name": "string"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» base_response|object|true|none||none|
|»» status_code|integer|true|none||none|
|»» status_msg|string|true|none||none|
|» subscriber_list|[object]|true|none||none|
|»» id|string|true|none||none|
|»» name|string|true|none||none|

## GET 查询粉丝列表

GET /follower/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 否 |none|
|page_number|query|string| 否 |none|
|page_size|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "base_response": {
    "status_code": 0,
    "status_msg": "string"
  },
  "fans_list": [
    {
      "id": "string",
      "name": "string"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» base_response|object|true|none||none|
|»» status_code|integer|true|none||none|
|»» status_msg|string|true|none||none|
|» fans_list|[object]|true|none||none|
|»» id|string|false|none||none|
|»» name|string|false|none||none|

## GET 查看我的好友

GET /friend/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 否 |none|
|page_number|query|string| 否 |none|
|page_size|query|string| 否 |none|
|Authorization|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "base_response": {
    "status_code": 0,
    "status_msg": "string"
  },
  "fans_list": [
    {
      "id": "string",
      "name": "string"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» base_response|object|true|none||none|
|»» status_code|integer|true|none||none|
|»» status_msg|string|true|none||none|
|» fans_list|[object]|true|none||none|
|»» id|string|true|none||none|
|»» name|string|true|none||none|

# 数据模型

<h2 id="tocS_Comments">Comments</h2>

<a id="schemacomments"></a>
<a id="schema_Comments"></a>
<a id="tocScomments"></a>
<a id="tocscomments"></a>

```json
{
  "id": "string",
  "user_id": "string",
  "video_id": "string",
  "content": "string",
  "created_at": "string",
  "updated_at": "string",
  "deleted_at": "string",
  "parent_comment_id": "string",
  "like_count": -2147483648
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none||none|
|user_id|string¦null|false|none||none|
|video_id|string¦null|false|none||none|
|content|string¦null|false|none||none|
|created_at|string¦null|false|none||none|
|updated_at|string¦null|false|none||none|
|deleted_at|string¦null|false|none||none|
|parent_comment_id|string¦null|false|none||none|
|like_count|integer¦null|false|none||none|

<h2 id="tocS_Followers">Followers</h2>

<a id="schemafollowers"></a>
<a id="schema_Followers"></a>
<a id="tocSfollowers"></a>
<a id="tocsfollowers"></a>

```json
{
  "create_time": "string",
  "status": -2147483648,
  "following_user_id": "string",
  "following_user_name": "string",
  "following_user_avatar_url": "string",
  "follower_user_id": "string",
  "follower_user_name": "string",
  "follower_user_avatar_url": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|create_time|string¦null|false|none||none|
|status|integer¦null|false|none||none|
|following_user_id|string¦null|false|none||none|
|following_user_name|string¦null|false|none||none|
|following_user_avatar_url|string¦null|false|none||none|
|follower_user_id|string¦null|false|none||none|
|follower_user_name|string¦null|false|none||none|
|follower_user_avatar_url|string¦null|false|none||none|

<h2 id="tocS_Users">Users</h2>

<a id="schemausers"></a>
<a id="schema_Users"></a>
<a id="tocSusers"></a>
<a id="tocsusers"></a>

```json
{
  "id": "string",
  "username": "string",
  "password": "string",
  "photo_url": "string",
  "created_at": "string",
  "updated_at": "string",
  "deleted_at": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|string|true|none||none|
|username|string|true|none||none|
|password|string|true|none||none|
|photo_url|string¦null|false|none||none|
|created_at|string¦null|false|none||none|
|updated_at|string¦null|false|none||none|
|deleted_at|string¦null|false|none||none|

<h2 id="tocS_VideoLikers">VideoLikers</h2>

<a id="schemavideolikers"></a>
<a id="schema_VideoLikers"></a>
<a id="tocSvideolikers"></a>
<a id="tocsvideolikers"></a>

```json
{
  "id": 1,
  "user_id": "string",
  "video_id": "string",
  "author_id": "string",
  "video_url": "string",
  "cover_url": "string",
  "title": "string",
  "description": "string",
  "visit_count": -2147483648,
  "like_count": -2147483648,
  "comment_count": -2147483648,
  "created_at": "string",
  "updated_at": "string",
  "deleted_at": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer|true|none||none|
|user_id|string¦null|false|none||none|
|video_id|string|true|none||none|
|author_id|string¦null|false|none||none|
|video_url|string¦null|false|none||none|
|cover_url|string¦null|false|none||none|
|title|string¦null|false|none||none|
|description|string¦null|false|none||none|
|visit_count|integer¦null|false|none||none|
|like_count|integer¦null|false|none||none|
|comment_count|integer¦null|false|none||none|
|created_at|string¦null|false|none||none|
|updated_at|string¦null|false|none||none|
|deleted_at|string¦null|false|none||none|

<h2 id="tocS_Videos">Videos</h2>

<a id="schemavideos"></a>
<a id="schema_Videos"></a>
<a id="tocSvideos"></a>
<a id="tocsvideos"></a>

```json
{
  "video_id": "string",
  "author_id": "string",
  "video_url": "string",
  "cover_url": "string",
  "title": "string",
  "description": "string",
  "visit_count": -2147483648,
  "like_count": -2147483648,
  "comment_count": -2147483648,
  "created_at": "string",
  "updated_at": "string",
  "deleted_at": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|video_id|string|true|none||none|
|author_id|string¦null|false|none||none|
|video_url|string¦null|false|none||none|
|cover_url|string¦null|false|none||none|
|title|string¦null|false|none||none|
|description|string¦null|false|none||none|
|visit_count|integer¦null|false|none||none|
|like_count|integer¦null|false|none||none|
|comment_count|integer¦null|false|none||none|
|created_at|string¦null|false|none||none|
|updated_at|string¦null|false|none||none|
|deleted_at|string¦null|false|none||none|

