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
go run main.go

3️⃣ The app will be available at http://localhost:8080
```

---

### 🏗 Project Structure
```
.
├── assets/             # Frontend files (HTML, CSS, JS)
├── engine/             # Core business logic (War/Grab rules, Anti-cheat)
├── model/              # Data structures and entities (User, Leaderboard)
├── repository/         # Data access layer (SQLite queries, Transactions)
├── util/               # Helper functions (Config, Logger, Time formatting)
├── main.go             # Application entry point & router initialization
├── Dockerfile          # Optimized multi-stage build
├── .dockerignore       # Excluding local DB and temp files
└── fly.toml            # Fly.io infrastructure configuration
```

---

### 🤝 Contributing
This is an experimental project. Feel free to open an issue or submit a Pull Request!
