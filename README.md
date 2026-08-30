Watchdog 🐕‍🦺

A simple server monitoring tool written in Go.

Watchdog periodically sends HTTP requests to a target server and checks whether it is alive and responding correctly. The project was created as a learning DevOps project to practice Go, HTTP, Docker, and basic service monitoring concepts.

🚀 How It Works

Watchdog checks the target server every 5 seconds.

┌──────────────┐
│   Watchdog   │
└──────┬───────┘
       │
       │ HTTP GET /health
       ▼
┌──────────────┐
│ Target Server│
│   :8080      │
└──────┬───────┘
       │
       ├── 200 OK ──────► [OK] Server is healthy
       │
       └── Error ───────► [ERROR] Server is unreachable

The check is repeated continuously.

🛠️ Technologies

* Go — application logic
* HTTP — communication with the monitored server
* Docker — containerization
* Git — version control

📦 Running Locally

Make sure you have Go installed.

Clone the repository:

git clone <your-repository-url>
cd watchdog

Run the application:

go run .

Watchdog will start checking:

http://localhost:8080/health

Every 5 seconds.

🐳 Running with Docker

Build the Docker image:

docker build -t watchdog .

Run the container:

docker run --rm watchdog

Note: The current version of Watchdog expects the monitored server to be available at localhost:8080.

📋 Example Output

When the server is running:

[OK] your server is healthy
[OK] your server is healthy
[OK] your server is healthy

When the server is unavailable:

[ERROR] your server is unreachable
[ERROR] your server is unreachable

🎯 Project Goals

This project is part of my journey into DevOps and backend development.

The main goals are to learn and practice:

* Writing backend tools with Go
* Working with HTTP requests
* Handling errors
* Understanding how services communicate
* Containerizing applications with Docker
* Understanding the basics of service monitoring
* Building and running applications in isolated environments

🗺️ Future Plans

Possible future improvements:

* Monitor multiple servers
* Configurable check interval
* Configurable target URLs
* Response time measurement
* Uptime statistics
* Better logging
* Configuration using environment variables
* Docker Compose
* Prometheus metrics
* Health-check dashboard

📚 Status

Version: v0.1

This is an educational project and is actively being developed while I learn Go and DevOps.

⸻

Made with ❤️ and Go.
