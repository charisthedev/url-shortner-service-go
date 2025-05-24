# URL Shortener Service (Go)

A simple and efficient URL shortening service built in Go using PostgreSQL for persistence and Base62 encoding for short codes.  
This project demonstrates creating short, readable URLs similar to services like bit.ly or tinyurl.com.

---

## Features

- Shorten long URLs into compact Base62 encoded short codes
- Decode short codes back to original URLs
- Store URLs and their short codes in PostgreSQL
- Transaction-safe creation with rollback on failure
- Simple RESTful API endpoints
- Modular, testable codebase with SQLC for type-safe DB queries
- Base62 encoding combined with obfuscation for security & unpredictability

---

## Technologies Used

- **Go** - API server and core logic
- **PostgreSQL** - Database storage for URLs
- **SQLC** - Generates type-safe Go database queries
- **Base62** - Encoding scheme for short codes
- **github.com/mattheath/base62** - Base62 encoding/decoding Go library
- **net/http** - HTTP server and routing

---

## Getting Started

### Prerequisites

- Go (1.19+ recommended)
- PostgreSQL database
- `sqlc` installed for generating DB query code ([install guide](https://sqlc.dev))

---

### Setup

1. **Clone the repository**

```bash
git clone https://github.com/charisthedev/url-shortner-service-go.git
cd url-shortner-service-go
```

1. **Set up your PostgreSQL database**
   Create a new database and user, then run the migration SQL scripts to create the necessary table

Configure environment variables

Create a .env file or export your database URL environment variable:

export DB_URL="postgres://user:password@localhost:5432/yourdbname?sslmode=disable"
Generate database query code

Make sure you have sqlc installed, then run:

```bash
sqlc generate
```

This will generate the Go database query types and methods under your internal/database package.

Running the Service
Build and run the Go server:

```bash
go build -o url-shortner
./url-shortner
```

The service will start listening on your configured port (default: :5000).

API Endpoints

1. Create Short URL
   POST /shorten

Request Body:

```json
{
  "url": "https://example.com/very/long/url"
}
```

Response:

```json
{
  "message": "url created",
  "data": {
    "url": "http://localhost:5000/abc123"
  }
}
```

2. Redirect to Original URL
   GET /{short_code}

Example: GET /abc123

Redirects user to the original URL associated with the short_code.

## How It Works

-The server stores the original URL in the database.
-It generates a unique numeric ID (auto-incremented).
-The numeric ID is obfuscated using a prime multiplication and modulo to avoid predictability.
-The obfuscated number is encoded to a Base62 string as the short code.
-When accessing the short URL, the short code is decoded and reversed to get the original ID.
-The original URL is fetched and the user is redirected.

```csharp
.
├── cmd/ # Main app entrypoint
├── controllers/ # HTTP handlers
├── internal/
│ ├── database/ # SQLC generated DB code
│ └── utils/ # Utility functions (encoding, responses)
├── db/ # SQL migration and queries
├── go.mod
└── main.go
```

## Utilities

utils.HashUrl(id int64) string — Obfuscates and encodes ID to Base62 shortcode

## Contributing

Feel free to open issues or submit pull requests. Please maintain consistent code style and include tests.

## 📄 License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT)  
© 2025 [CharisTheDev](https://github.com/charisthedev)

## 📬 Contact

Developed and maintained by [CharisTheDev](https://github.com/charisthedev)  
For inquiries, reach out via GitHub or open an issue.
