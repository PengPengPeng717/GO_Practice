package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"
)

// 用户结构体
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 会话结构体
type Session struct {
	Username string
	Expires  time.Time
}

// 全局变量
var (
	sessions     = make(map[string]Session)
	sessionMutex sync.RWMutex
	users        = map[string]string{
		"admin": "123456",
		"user1": "password123",
		"user2": "password456",
	}
)

// 生成随机会话ID
func generateSessionID() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

// 检查用户凭据
func checkCredentials(username, password string) bool {
	if storedPassword, exists := users[username]; exists {
		return storedPassword == password
	}
	return false
}

// 创建新会话
func createSession(username string) string {
	sessionID := generateSessionID()
	sessionMutex.Lock()
	sessions[sessionID] = Session{
		Username: username,
		Expires:  time.Now().Add(30 * time.Minute),
	}
	sessionMutex.Unlock()
	return sessionID
}

// 验证会话
func validateSession(sessionID string) (string, bool) {
	sessionMutex.RLock()
	session, exists := sessions[sessionID]
	sessionMutex.RUnlock()

	if !exists {
		return "", false
	}

	if time.Now().After(session.Expires) {
		// 会话过期，删除
		sessionMutex.Lock()
		delete(sessions, sessionID)
		sessionMutex.Unlock()
		return "", false
	}

	return session.Username, true
}

// 登录页面处理
func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>用户登录</title>
    <style>
        body {
            font-family: 'Microsoft YaHei', Arial, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            margin: 0;
            padding: 0;
            height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
        }
        .login-container {
            background: white;
            padding: 40px;
            border-radius: 10px;
            box-shadow: 0 15px 35px rgba(0, 0, 0, 0.1);
            width: 100%;
            max-width: 400px;
        }
        .login-header {
            text-align: center;
            margin-bottom: 30px;
        }
        .login-header h1 {
            color: #333;
            margin: 0;
            font-size: 28px;
            font-weight: 300;
        }
        .form-group {
            margin-bottom: 20px;
        }
        .form-group label {
            display: block;
            margin-bottom: 8px;
            color: #555;
            font-weight: 500;
        }
        .form-group input {
            width: 100%;
            padding: 12px 15px;
            border: 2px solid #e1e5e9;
            border-radius: 5px;
            font-size: 16px;
            transition: border-color 0.3s ease;
            box-sizing: border-box;
        }
        .form-group input:focus {
            outline: none;
            border-color: #667eea;
        }
        .login-btn {
            width: 100%;
            padding: 12px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            border: none;
            border-radius: 5px;
            font-size: 16px;
            font-weight: 500;
            cursor: pointer;
            transition: transform 0.2s ease;
        }
        .login-btn:hover {
            transform: translateY(-2px);
        }
        .error-message {
            color: #e74c3c;
            text-align: center;
            margin-top: 10px;
            font-size: 14px;
        }
        .success-message {
            color: #27ae60;
            text-align: center;
            margin-top: 10px;
            font-size: 14px;
        }
        .demo-accounts {
            margin-top: 20px;
            padding: 15px;
            background: #f8f9fa;
            border-radius: 5px;
            font-size: 12px;
            color: #666;
        }
        .demo-accounts h4 {
            margin: 0 0 10px 0;
            color: #333;
        }
        .demo-accounts ul {
            margin: 0;
            padding-left: 20px;
        }
    </style>
</head>
<body>
    <div class="login-container">
        <div class="login-header">
            <h1>用户登录</h1>
        </div>
        <form method="POST" action="/login">
            <div class="form-group">
                <label for="username">用户名：</label>
                <input type="text" id="username" name="username" required>
            </div>
            <div class="form-group">
                <label for="password">密码：</label>
                <input type="password" id="password" name="password" required>
            </div>
            <button type="submit" class="login-btn">登录</button>
        </form>
        
        <div class="demo-accounts">
            <h4>演示账户：</h4>
            <ul>
                <li>用户名: admin, 密码: 123456</li>
                <li>用户名: user1, 密码: password123</li>
                <li>用户名: user2, 密码: password456</li>
            </ul>
        </div>
    </div>
</body>
</html>`

		tmplParsed, err := template.New("login").Parse(tmpl)
		if err != nil {
			http.Error(w, "模板解析错误", http.StatusInternalServerError)
			return
		}

		tmplParsed.Execute(w, nil)
	}
}

// 登录处理
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		if checkCredentials(username, password) {
			sessionID := createSession(username)
			http.SetCookie(w, &http.Cookie{
				Name:     "session_id",
				Value:    sessionID,
				Path:     "/",
				MaxAge:   1800, // 30分钟
				HttpOnly: true,
			})

			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/login?error=invalid", http.StatusSeeOther)
		}
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// 仪表板页面
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	username, valid := validateSession(cookie.Value)
	if !valid {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tmpl := `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>用户仪表板</title>
    <style>
        body {
            font-family: 'Microsoft YaHei', Arial, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            margin: 0;
            padding: 0;
            min-height: 100vh;
        }
        .header {
            background: white;
            padding: 20px;
            box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
        }
        .header-content {
            max-width: 1200px;
            margin: 0 auto;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }
        .welcome {
            font-size: 24px;
            color: #333;
            font-weight: 300;
        }
        .logout-btn {
            padding: 10px 20px;
            background: #e74c3c;
            color: white;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            text-decoration: none;
            font-size: 14px;
        }
        .logout-btn:hover {
            background: #c0392b;
        }
        .main-content {
            max-width: 1200px;
            margin: 40px auto;
            padding: 0 20px;
        }
        .dashboard-card {
            background: white;
            border-radius: 10px;
            padding: 30px;
            box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
            margin-bottom: 20px;
        }
        .dashboard-card h2 {
            color: #333;
            margin-top: 0;
            font-weight: 300;
        }
        .stats-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
            gap: 20px;
            margin-top: 20px;
        }
        .stat-item {
            background: #f8f9fa;
            padding: 20px;
            border-radius: 8px;
            text-align: center;
        }
        .stat-number {
            font-size: 32px;
            font-weight: bold;
            color: #667eea;
            margin-bottom: 5px;
        }
        .stat-label {
            color: #666;
            font-size: 14px;
        }
    </style>
</head>
<body>
    <div class="header">
        <div class="header-content">
            <div class="welcome">欢迎，{{.Username}}！</div>
            <a href="/logout" class="logout-btn">退出登录</a>
        </div>
    </div>
    
    <div class="main-content">
        <div class="dashboard-card">
            <h2>用户仪表板</h2>
            <p>您已成功登录系统。这里是您的个人仪表板。</p>
            
            <div class="stats-grid">
                <div class="stat-item">
                    <div class="stat-number">3</div>
                    <div class="stat-label">总项目数</div>
                </div>
                <div class="stat-item">
                    <div class="stat-number">12</div>
                    <div class="stat-label">完成任务</div>
                </div>
                <div class="stat-item">
                    <div class="stat-number">85%</div>
                    <div class="stat-label">完成率</div>
                </div>
                <div class="stat-item">
                    <div class="stat-number">5</div>
                    <div class="stat-label">待处理任务</div>
                </div>
            </div>
        </div>
    </div>
</body>
</html>`

	data := struct {
		Username string
	}{
		Username: username,
	}

	tmplParsed, err := template.New("dashboard").Parse(tmpl)
	if err != nil {
		http.Error(w, "模板解析错误", http.StatusInternalServerError)
		return
	}

	tmplParsed.Execute(w, data)
}

// 退出登录
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err == nil {
		sessionMutex.Lock()
		delete(sessions, cookie.Value)
		sessionMutex.Unlock()
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// API登录接口
func apiLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "无效的JSON格式", http.StatusBadRequest)
		return
	}

	if checkCredentials(user.Username, user.Password) {
		sessionID := createSession(user.Username)
		response := map[string]interface{}{
			"success":    true,
			"message":    "登录成功",
			"session_id": sessionID,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		response := map[string]interface{}{
			"success": false,
			"message": "用户名或密码错误",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
	}
}

func main() {
	// 设置路由
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/login/submit", loginHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/api/login", apiLoginHandler)

	// 启动服务器
	fmt.Println("服务器启动在 http://localhost:8080")
	fmt.Println("请访问 http://localhost:8080/login 进行登录")
	fmt.Println("演示账户：")
	fmt.Println("  - 用户名: admin, 密码: 123456")
	fmt.Println("  - 用户名: user1, 密码: password123")
	fmt.Println("  - 用户名: user2, 密码: password456")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
