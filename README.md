# 🍽️ Meal Planner App

A simple full-stack web application to register users, log in securely using JWT authentication, and track meal plans. Built with Go (Gin), PostgreSQL, and a minimal HTML/CSS/JS frontend.

---

## 🚀 Features

- User Registration & Login (with JWT)
- Passwords hashed securely using bcrypt
- Protected Profile endpoint (JWT-based)
- Basic HTML/CSS frontend to test login/register/admin
- Local storage support for client-side meal plans (in progress)
- RESTful API endpoints

---

## 🧠 Tech Stack

- **Backend:** Go (Gin Web Framework), PostgreSQL, JWT
- **Frontend:** HTML, CSS, JavaScript
- **Auth:** JWT Tokens, bcrypt Password Hashing
- **Storage:** PostgreSQL (via GORM ORM)
- **Deployment-ready:** CORS enabled, .env managed

---

## 📁 Project Structure
.
├── backend
│   ├── controllers
│   │   └── user_controller.go
│   ├── database
│   │   ├── config.go
│   │   ├── db.go
│   │   └── migrations.go
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── mealplan
│   ├── middleware
│   │   └── middleware.go
│   ├── models
│   │   └── user_model.go
│   ├── routes
│   │   └── routes.go
│   └── utils
│       └── password.go
└── frontend
    ├── index.html
    ├── script.js
    └── styles.css

## 📦 API Endpoints

### 🔐 Authentication
| Method | Endpoint             | Description          |
|--------|----------------------|----------------------|
| POST   | `/users/register`    | Register a new user  |
| POST   | `/users/login`       | Login user & get JWT |

### 👤 User
| Method | Endpoint             | Description                    |
|--------|----------------------|--------------------------------|
| GET    | `/users/:id`         | Get user profile (JWT required) |

> Add `Authorization: Bearer <token>` header to access protected routes.

---

## 🔧 Setup & Run

### 🐘 PostgreSQL Setup

Make sure you have PostgreSQL running and a user+db created.

Update your `.env` file:
DB_HOST=localhost 
DB_PORT=5432 
DB_USER=youruser 
DB_PASSWORD=yourpassword 
DB_NAME=yourdbname 
JWT_SECRET=yourjwtsecret

