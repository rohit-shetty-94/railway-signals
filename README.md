# 🚆 **Railway Signals**

**A RESTful API for managing railway signals and tracks, built using Go, Echo framework, and PostgreSQL.**

---

## 📌 **Table of Contents**

- [🚀 Features](#-features)
- [🛠️ Tech Stack](#-tech-stack)
- [📺 Project Structure](#-project-structure)
- [⚙️ Setup & Installation](#-setup--installation)
- [📝 API Documentation](#-api-documentation)
- [👨‍💻 Author](#-author)

---

## 🚀 **Features**

✅ CRUD operations for **Tracks** and **Signals**\
✅ Relationship management between **Signals and Tracks**\
✅ **Standardized API responses** with headers and data\
✅ **PostgreSQL database integration**\
✅ **Error handling and validation**

---

## 🛠️ **Tech Stack**

- **Golang** (Backend)
- **Echo Framework** (Web Server)
- **PostgreSQL** (Database)
- **go-pg** (ORM)

---

## 📺 **Project Structure**

```
railway_signals/
│── main.go
│── go.mod
│── go.sum
│── config/
│   ├── db.go
│── entity/
│   ├── track
│       ├── track.go
│   ├── signal
│       ├── signal.go
│── repository/
│   ├── track_repository.go
│   ├── signal_repository.go
│── handlers/
│   ├── track_handler.go
│   ├── signal_handler.go
│── utils/
│   ├── response.go
│── scripts/
│   ├── load_data.go
│── README.md
```

---

## ⚙️ **Setup & Installation**

### 1️⃣ **Clone the Repository**

```sh
git clone https://github.com/rohit-shetty-94/railway-signals.git
cd railway-signals
```

### 2️⃣ **Install Dependencies**

```sh
go mod tidy
```

### 3️⃣ **Setup PostgreSQL Database**

```sh
psql -U postgres -c "CREATE DATABASE railway_signals;"
```

### 4️⃣ **Run the Application**

```sh
go run main.go
```

Server starts at `http://localhost:8080`.

---

## 📝 **API Documentation**

### 📌 **Track Endpoints**

| Method | Endpoint      | Description                      |
| ------ | ------------- | -------------------------------- |
| GET    | `/tracks`     | Retrieve all tracks with signals |
| GET    | `/tracks/:id` | Retrieve a specific track by ID  |
| POST   | `/tracks`     | Create a new track               |
| PUT    | `/tracks/:id` | Update an existing track         |
| DELETE | `/tracks/:id` | Delete a track                   |

### 📌 **Signal Endpoints**

| Method | Endpoint              | Description                                  |
| ------ | --------------------- | -------------------------------------------- |
| GET    | `/signals`            | Retrieve all signals                         |
| GET    | `/signals/:id/tracks` | Retrieve all tracks associated with a signal |

### 📌 **Signal-Track Relationship Endpoints**

| Method | Endpoint                               | Description                               |
| ------ | -------------------------------------- | ----------------------------------------- |
| POST   | `/tracks/:track_id/signals/:signal_id` | Associate a track with a signal           |
| DELETE | `/tracks/:track_id/signals/:signal_id` | Remove association between track & signal |


---

## 👨‍💻 **Author**

**[Rohit Shetty]**\
🚀 **GitHub**: [rohit-shetty-94](https://github.com/rohit-shetty-94)\
💎 **Email**: [shettyrohit61@gamil.com](mailto\:shettyrohit61@gamil.com)

---


