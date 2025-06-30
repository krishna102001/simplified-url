# Simplified URL Shortener

A lightweight backend service that transforms long URLs into short, manageable links. Built with modern technologies for reliable URL shortening functionality.

## 🛠️ Tech Stack

- **Backend Framework**: Go with Gin-Gonic
- **Database**: PostgreSQL
- **ORM**: GORM
- **Deployment**: Dockerized for easy testing and deployment

## 🚀 Features

- Convert long URLs into short, easy-to-share links
- Fast redirect service to original URLs
- RESTful API design
- Database persistence with PostgreSQL
- Dockerized application for seamless deployment

## 📡 API Endpoints

### Create Short URL

**Endpoint**: `POST /create-short-url`

**URL**: `https://simplified-url.onrender.com/create-short-url`

**Request Body**:
```json
{
  "original_url": "https://github.com/krishna102001/simplified-url/new/master"
}
```

**Response**:
```json
{
  "short_url": "4fa59618"
}
```

### Redirect to Original URL

**Endpoint**: `GET /:id`

**URL**: `https://simplified-url.onrender.com/{short_url}`

**Example**: `https://simplified-url.onrender.com/4fa59618`

This endpoint automatically redirects you to the original URL.

## 🧪 Testing

The application is fully dockerized for easy testing and deployment. You can run the entire stack locally using Docker.

### Prerequisites

- Docker
- Docker Compose (optional)

### Running the Application

```bash
# Clone the repository
git clone https://github.com/krishna102001/simplified-url.git

# Navigate to project directory
cd simplified-url

# Run with Docker
docker build -t simplified-url .
docker run -p 8080:8080 simplified-url
```

## 🔧 Configuration

Make sure to configure your PostgreSQL database connection settings in your environment variables or configuration files.

## 📝 Usage Example

1. **Create a short URL**:
   ```bash
   curl -X POST https://simplified-url.onrender.com/create-short-url \
     -H "Content-Type: application/json" \
     -d '{"original_url": "https://your-long-url.com"}'
   ```

2. **Access the short URL**:
   Visit `https://simplified-url.onrender.com/{returned_short_id}` in your browser

## 🤝 Contributing

Feel free to fork this project and submit pull requests for any improvements.

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

---

Built with ❤️ using Go and Gin-Gonic
