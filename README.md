# Chat App Backend

This is the backend for a real-time public chat application built with Go. It provides WebSocket-based real-time messaging and RESTful API services for user authentication.

## 🚀 Features
- 🔄 WebSocket-based real-time chat
- 🔐 JWT-based authentication (Login/Register)
- 📦 REST API for user management
- 🗄️ MySQL database integration
- 📊 Efficient handling of concurrent connections

## 🛠️ Tech Stack
- **Language**: Go
- **Web Framework**: net/http (No third-party frameworks)
- **Database**: MySQL
- **Authentication**: JWT
- **Real-time Communication**: WebSockets

## 📦 Installation
1. Clone this repository:
   ```sh
   git clone https://github.com/ricoerian/chat-backend.git
   cd chat-backend
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Configure environment variables in `.env`:
   ```env
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=root
   DB_PASS=password
   DB_NAME=chat_db
   ```
4. Run the server:
   ```sh
   go run main.go
   ```

## 🏗️ API Endpoints
### **Authentication**
| Method | Endpoint      | Description         |
|--------|--------------|---------------------|
| POST   | `/register`  | Register a user     |
| POST   | `/login`     | Authenticate a user |

### **Chat**
| Method | Endpoint    | Description          |
|--------|------------|----------------------|
| GET    | `/messages` | Get chat messages   |
| POST   | `/ws`     | Send a chat message |

## 🤝 Contributing
Pull requests and feature suggestions are welcome.

## 📄 License
This project is open-source under the [MIT License](LICENSE).
