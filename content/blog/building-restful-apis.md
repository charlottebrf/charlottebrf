# Building RESTful APIs: Best Practices

*Published: March 8, 2025*  
*Category: Web Development*

Creating well-designed APIs is crucial for modern web applications. Here are some best practices I've learned for building robust, scalable RESTful APIs.

## What is a RESTful API?

REST (Representational State Transfer) is an architectural style for designing networked applications. A RESTful API follows REST principles and uses HTTP methods to perform operations on resources.

## Core Principles

### 1. Use HTTP Methods Correctly

- **GET**: Retrieve data (idempotent)
- **POST**: Create new resources
- **PUT**: Update entire resources (idempotent)
- **PATCH**: Partial updates
- **DELETE**: Remove resources (idempotent)

```go
// Example in Go using gorilla/mux
r.HandleFunc("/api/users", getUsersHandler).Methods("GET")
r.HandleFunc("/api/users", createUserHandler).Methods("POST")
r.HandleFunc("/api/users/{id}", getUserHandler).Methods("GET")
r.HandleFunc("/api/users/{id}", updateUserHandler).Methods("PUT")
r.HandleFunc("/api/users/{id}", deleteUserHandler).Methods("DELETE")
```

### 2. Use Meaningful Resource Names

Resources should be nouns, not verbs:

✅ **Good**:
- `GET /api/users`
- `POST /api/users`
- `GET /api/users/123`

❌ **Bad**:
- `GET /api/getUsers`
- `POST /api/createUser`
- `GET /api/getUserById/123`

### 3. Use HTTP Status Codes Appropriately

Common status codes and their meanings:

- **200 OK**: Successful GET, PUT, PATCH
- **201 Created**: Successful POST
- **204 No Content**: Successful DELETE
- **400 Bad Request**: Invalid request data
- **401 Unauthorized**: Authentication required
- **403 Forbidden**: Access denied
- **404 Not Found**: Resource doesn't exist
- **500 Internal Server Error**: Server error

## API Design Best Practices

### 1. Consistent Response Format

Use a consistent JSON structure for all responses:

```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  },
  "message": "User retrieved successfully"
}
```

For errors:

```json
{
  "success": false,
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Email is required",
    "details": {
      "field": "email"
    }
  }
}
```

### 2. Input Validation

Always validate input data:

```go
type CreateUserRequest struct {
    Name  string `json:"name" validate:"required,min=2,max=50"`
    Email string `json:"email" validate:"required,email"`
    Age   int    `json:"age" validate:"required,min=18,max=120"`
}

func validateInput(req CreateUserRequest) error {
    validate := validator.New()
    return validate.Struct(req)
}
```

### 3. Pagination

For endpoints that return lists, implement pagination:

```go
type PaginatedResponse struct {
    Data       []User `json:"data"`
    Page       int    `json:"page"`
    Limit      int    `json:"limit"`
    Total      int    `json:"total"`
    TotalPages int    `json:"total_pages"`
}

// GET /api/users?page=2&limit=10
```

### 4. Filtering and Sorting

Allow clients to filter and sort data:

```
GET /api/users?status=active&sort=created_at&order=desc&name=john
```

## Security Best Practices

### 1. Authentication and Authorization

```go
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        
        // Validate token
        claims, err := validateJWT(token)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
        
        // Add user context
        ctx := context.WithValue(r.Context(), "user", claims)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

### 2. Rate Limiting

Implement rate limiting to prevent abuse:

```go
func rateLimitMiddleware(limit int) func(http.Handler) http.Handler {
    limiter := make(map[string]*rate.Limiter)
    
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ip := r.RemoteAddr
            
            if _, exists := limiter[ip]; !exists {
                limiter[ip] = rate.NewLimiter(rate.Every(time.Minute), limit)
            }
            
            if !limiter[ip].Allow() {
                http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
                return
            }
            
            next.ServeHTTP(w, r)
        })
    }
}
```

### 3. Input Sanitization

Always sanitize and validate input to prevent injection attacks:

```go
func sanitizeInput(input string) string {
    // Remove potentially dangerous characters
    return html.EscapeString(strings.TrimSpace(input))
}
```

## Error Handling

Implement comprehensive error handling:

```go
type APIError struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Details interface{} `json:"details,omitempty"`
}

func handleError(w http.ResponseWriter, err error, statusCode int) {
    apiErr := APIError{
        Code:    "INTERNAL_ERROR",
        Message: err.Error(),
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "success": false,
        "error":   apiErr,
    })
}
```

## Documentation

Good API documentation is crucial:

1. **Use OpenAPI/Swagger** for interactive documentation
2. **Provide examples** for each endpoint
3. **Document error responses**
4. **Include rate limits** and authentication requirements
5. **Keep documentation up-to-date**

## Testing

Write comprehensive tests for your API:

```go
func TestCreateUser(t *testing.T) {
    // Setup
    router := setupRouter()
    
    // Test data
    user := CreateUserRequest{
        Name:  "John Doe",
        Email: "john@example.com",
        Age:   30,
    }
    
    body, _ := json.Marshal(user)
    req := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    
    // Execute
    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)
    
    // Assert
    assert.Equal(t, http.StatusCreated, rr.Code)
    
    var response map[string]interface{}
    json.Unmarshal(rr.Body.Bytes(), &response)
    assert.True(t, response["success"].(bool))
}
```

## Conclusion

Building good APIs takes practice and attention to detail. Focus on:

- **Consistency** in design and responses
- **Security** at every level
- **Performance** and scalability
- **Documentation** for developers
- **Testing** to ensure reliability

Remember: A well-designed API is not just functional—it's a joy to use!

---

*What are your favorite API design patterns? Share your thoughts in the comments below!*