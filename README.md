# 🚀 Montink ERP Integration

> A backend service built with **Go (Golang)** to manage, synchronize, and analyze orders from the Montink Print-on-Demand platform.

![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-blue.svg)
![Status](https://img.shields.io/badge/Status-In%20Development-yellow)

## 📖 About The Project

This project serves as an **ERP (Enterprise Resource Planning)** backend specifically designed for stores operating on Montink. It solves common challenges when integrating with the Montink API, such as inconsistent data types (strings representing floats) and complex nested JSON structures.

**Key Features:**
* **🔄 Order Synchronization:** Fetches orders from Montink and persists them locally.
* **📦 Data Normalization:** Handles complex JSON responses, standardizing currency values and date formats.
* **🗄️ SQLite Persistence:** Lightweight local database storage for offline access and history.
* **🏗️ Clean Architecture:** Follows the Standard Go Project Layout, separating Domain, Infrastructure, and Application layers.

---

## 🛠️ Tech Stack

* **Language:** [Go](https://go.dev/) (Golang)
* **Database:** SQLite (via `modernc.org/sqlite` - pure Go driver)
* **Architecture:** Domain-Driven Design (DDD) principles
* **Environment Management:** `godotenv`

---

## 📂 Project Structure

This project follows the **Standard Go Layout**:

```text
montink_erp/
├── cmd/
│   └── api/
│       └── main.go       # Application entry point
├── internal/
│   ├── domain/           # Business logic & Entity definitions (Structs)
│   ├── database/         # SQLite repository implementation
│   └── montink/          # HTTP Client to communicate with external API
├── orders.db             # Local SQLite database
└── go.mod                # Dependency management