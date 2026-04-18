
# TikTok 项目目录结构


```
tik-tok/
├── biz/                        
│   ├── handler/                  # HTTP 请求处理器
│   │   ├── api/                  # API 相关服务
│   │   │   └── empty_service.go
│   │   ├── communication/        # 社交功能处理器
│   │   │   └── communication_service.go
│   │   ├── interaction/          # 互动功能处理器
│   │   │   └── interaction_service.go
│   │   ├── user/                 # 用户相关处理器
│   │   │   └── user_service.go
│   │   ├── video/                # 视频相关处理器
│   │   │   └── video_service.go
│   │   └── ping.go               # 健康检查接口
│   │
│   ├── model/                    # 数据模型层（Proto Buffer 生成）
│   │   ├── api/
│   │   │   └── api.pb.go
│   │   ├── communication/
│   │   │   └── commuinication.pb.go
│   │   ├── interaction/
│   │   │   └── interaction.pb.go
│   │   ├── user/
│   │   │   └── user.pb.go
│   │   └── video/
│   │       └── video.pb.go
│   │
│   └── router/                   # 路由及其中间件配置
│       ├── api/
│       │   ├── api.go
│       │   └── middleware.go
│       ├── communication/
│       │   ├── commuinication.go
│       │   └── middleware.go
│       ├── interaction/
│       │   ├── interaction.go
│       │   └── middleware.go
│       ├── user/
│       │   ├── user.go
│       │   └── middleware.go
│       ├── video/
│       │   ├── video.go
│       │   └── middleware.go
│       └── register.go           
│
├── config/                       #全局配置文件
│   ├── config.go
│   └── config.yml
│
├── dal/                          # 数据库的初始化及其相关操作
│   ├── mysql/
│   │   ├── init.go
│   │   └── action.go
│   └── redis/
│       └── init.go
│
├── idl/                          # 数据库模型接口定义
│   ├── api.proto
│   ├── communication.proto
│   ├── interaction.proto
│   ├── user.proto
│   └── video.proto
│
├── mw/                           # 中间件层
│   └── token/
│       ├── init.go
│       ├── jwt.go
│       └── token.go
│
├── utils/                        #常用的小组件
│   ├── encryption.go             # 加密工具类
│   ├── id.go                     # ID生成工具类
│   └── time.go                   # 时间工具类
│
├── script/                      
│   └── bootstrap.sh
│
├── .gitignore
├── .hz
├── Dockerfile                   # Dockerfile 配置文件
├── compose.yml                  # Docker Compose 配置文件  
├── go.mod
├── go.sum
├── hz.yaml
├── main.go
├── router.go
└── router_gen.go
```
