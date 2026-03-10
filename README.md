# 🧸 Grab Your Labubu!

Grab Your Labubu is a simulation app designed to handle high-traffic "grab" events. It is engineered to manage thousands of simultaneous clicks using Go's sync.Map for thread-safe concurrency and SQLite WAL Mode for fast, reliable data persistence.

### 🚀 [Live here](https://grab-your-labubu.fly.dev)

---

### 🛡️ Anti-Cheat & Security Features
To ensure a fair "War" experience, the following server-side protections are implemented:
* **Rate Limiting**: Prevents automated script/bot spamming by limiting requests per IP address using a sliding window or cooldown period.
* **Request Validation**: Each "Grab" request is validated against a server-side timestamp to prevent replay attacks.
* **Idempotency Checks**: Ensures that a single winning event cannot be processed multiple times for the same user session.
* **Concurrency Guard**: Uses sync.Once or atomic counters to ensure only one winner is declared in a microsecond race condition.

---

### 🚀 Features
* **High Concurrency**: Optimized locking mechanism in Go to handle thousands of simultaneous users.
* **Persistent Leaderboard**: All winners are stored permanently in SQLite.
* **Production Ready**: Multi-stage Docker build resulting in a lightweight image (< 20MB).
* **Cloud Native**: Designed for Fly.io with persistent volume support.

---

### 🛠 Tech Stack
* **Language**: Go (Golang)
* **Database**: SQLite (with WAL Mode & Busy Timeout)
* **Infrastructure**: Docker & Fly.io

---

### 📦 Local Development
```
1️⃣ Clone the repository
git clone https://github.com/deng37/grab-your-labubu.git
cd grab-your-labubu

2️⃣ Run the application
go run cmd/api/main.go

3️⃣ The app will be available at http://localhost:8080
```

---

### 🏗 Project Structure
```
.
├── assets/             # Static files (Images, CSS, Client-side JS)
├── cmd/
│   └── api/
│       └── main.go     # Application entry point & router initialization
├── internal/           # Private application and library code
│   ├── engine/         # Core business logic (War/Grab rules, Concurrency)
│   ├── middleware/     # Custom HTTP Middlewares (Auth, CORS, Logging)
│   ├── model/          # Data structures and entities
│   ├── repository/     # Data access layer (SQLite/DB operations)
│   └── util/           # Shared helpers (IP Tracking, Headers, Config)
├── index.html          # Main landing page for the "Labubu War" arena
├── Dockerfile          # Optimized multi-stage build (Alpine-based)
├── fly.toml            # Fly.io infrastructure & auto-scaling config
├── Makefile            # Automation shortcuts (run, build, deploy)
└── go.mod              # Go module dependency management
```

---

### ✋🏻 Copyright Info
This project is for personal, non-commercial experimentation only.
Labubu is a character owned by Pop Mart / Kasing Lung.
No copyright infringement intended.
