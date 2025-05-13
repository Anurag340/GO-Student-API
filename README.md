# 📚 Student API – Go Backend Server

A lightweight and efficient RESTful API built in **Go** for managing student records. This project showcases clean architecture, database integration, and graceful shutdown using Go’s standard library features.

---

## 🚀 Features

- 🔧 Create, Read, Update, Delete (CRUD) operations on student records
- 🧩 Modular and extensible code structure
- 🛡️ Graceful shutdown with context timeout
- 💾 MySQL database integration
- 📂 Clean routing using `http.ServeMux`
- 🪵 Structured logging with `slog`

---

## 🧪 API Endpoints

| Method | Endpoint                 | Description            |
|--------|--------------------------|------------------------|
| POST   | `/api/students`          | Create a new student   |
| GET    | `/api/students/{id}`     | Get a student by ID    |
| GET    | `/api/students`          | List all students      |
| PUT    | `/api/students/{id}`     | Update an existing student |
| DELETE | `/api/students/{id}`     | Delete a student by ID |
