# Go 用户登录系统

这是一个使用 Go 语言编写的简单用户登录系统，包含完整的 Web 界面和 API 接口。

## 功能特性

- 🎨 美观的现代化登录界面
- 🔐 安全的会话管理
- 📱 响应式设计，支持移动设备
- 🚀 RESTful API 接口
- 👤 用户仪表板
- 🔒 自动会话过期（30分钟）

## 快速开始

### 1. 运行程序

```bash
go run main.go
```

### 2. 访问应用

打开浏览器访问：http://localhost:8080

### 3. 演示账户

系统预置了以下演示账户：

| 用户名 | 密码 |
|--------|------|
| admin | 123456 |
| user1 | password123 |
| user2 | password456 |

## 项目结构

```
├── main.go          # 主程序文件
├── README.md        # 项目说明文档
└── pkg/             # Go 模块依赖
```

## API 接口

### 登录接口

**POST** `/api/login`

请求体：
```json
{
    "username": "admin",
    "password": "123456"
}
```

成功响应：
```json
{
    "success": true,
    "message": "登录成功",
    "session_id": "xxx"
}
```

失败响应：
```json
{
    "success": false,
    "message": "用户名或密码错误"
}
```

## 页面路由

- `/` - 重定向到登录页面
- `/login` - 登录页面
- `/dashboard` - 用户仪表板（需要登录）
- `/logout` - 退出登录
- `/api/login` - 登录 API 接口

## 技术栈

- **后端**: Go (Golang)
- **模板引擎**: Go HTML Template
- **会话管理**: 内存存储 + Cookie
- **前端**: HTML5 + CSS3 + JavaScript
- **样式**: 现代化渐变设计

## 安全特性

- 密码验证
- 会话管理
- Cookie 安全设置
- 会话自动过期
- 并发安全的会话存储

## 开发说明

### 添加新用户

在 `main.go` 文件的 `users` 变量中添加新用户：

```go
users = map[string]string{
    "admin": "123456",
    "user1": "password123",
    "user2": "password456",
    "newuser": "newpassword", // 添加新用户
}
```

### 自定义样式

登录页面和仪表板的样式都在 `main.go` 文件中的 HTML 模板内，可以直接修改 CSS 来自定义外观。

## 部署说明

### 生产环境建议

1. 使用 HTTPS
2. 配置反向代理（如 Nginx）
3. 使用数据库存储用户信息
4. 实现密码加密
5. 添加日志记录
6. 配置环境变量

### 构建可执行文件

```bash
go build -o login-system.exe main.go
```

## 许可证

MIT License

## 贡献

欢迎提交 Issue 和 Pull Request！ 