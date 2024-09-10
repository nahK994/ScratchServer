## TinyHTTP - A Lightweight HTTP Server

Welcome to the **TinyHTTP** repository! This project is a work in progress and aims to implement a basic HTTP server from scratch using Go.

### 🚧 Project Status

This project is still in development. Contributions, ideas, and feedback are welcome!

### 🚀 Getting Started

#### Prerequisites

To run this project, you'll need:

- **Go** (version 1.20 or later)

#### How to Use

You can easily add this package to your Go project by running the following command:

```go
go get github.com/nahK994/SimpleServer
```

After installing the package, you can start the HTTP server locally in your main.go file with the following code:

```go
srv := server.Initiate("127.0.0.1:8000")
srv.RequestHandler("/post", http.MethodPost, func(r models.Request, w *models.Response) {
    fmt.Println("TEST ===>", r.Body)
    w.StatusCode = http.StatusCreated
    w.Body = r.Body
})

srv.RequestHandler("/post", http.MethodGet, func(r models.Request, w *models.Response) {
    fmt.Println("TEST ===>", r.Body)
    w.StatusCode = http.StatusOK
    w.Body = "Hello World!!"
})
srv.RequestHandler("/get", http.MethodGet, func(r models.Request, w *models.Response) {
    fmt.Println("TEST ===>", r.Body)
    w.StatusCode = http.StatusOK
    w.Body = "Get request -> Hello World!!"
})
log.Fatal(srv.Start())
```

You can also check out the complete example in `cmd/main.go`.

### Roadmap: TinyHTTP

#### Milestone 1: Basic HTTP Server (In Progress)
- [x] Implement basic HTTP server handling requests (e.g., GET, POST, DELETE, PUT, PATCH).
- [x] Set up routing system to handle different endpoints (e.g., `/`, `/about`, etc.).
- [x] Ensure basic error handling (e.g., 404 Not Found, 500 Internal Server Error).
- [ ] Support static file serving (e.g., HTML, CSS, JavaScript).
- [ ] Add logging for incoming requests and server responses.
- [ ] Write unit tests for core functionalities.

#### Milestone 2: Advanced Features (Planned)
- [ ] Implement middleware support (e.g., for authentication, logging, request modification).
- [ ] Add support for query parameters and URL parameters.
- [ ] Implement session management and cookie handling.
- [ ] Add HTTPS support for secure communication.
- [ ] Support CORS (Cross-Origin Resource Sharing) for API-based communication.
- [ ] Provide rate limiting to control the number of requests handled.

#### Milestone 3: Optimization and Deployment (Planned)
- [ ] Optimize server performance and resource usage.
- [ ] Add support for WebSocket connections.
- [ ] Add detailed logging and monitoring tools (e.g., request times, error rates).
- [ ] Create a deployment guide with instructions for Docker and cloud services (e.g., AWS, Heroku).
- [ ] Provide documentation and examples for users and contributors.
- [ ] Implement load balancing and horizontal scaling for high traffic environments.


### 🤝 Contributing

We welcome contributions of all kinds! Whether you're fixing bugs, adding new features, or improving documentation, your input is appreciated.

#### How to Contribute

1. **Fork the repository** to your GitHub account.
2. **Clone your forked repository** to your local machine:

   ```bash
   git clone https://github.com/nahK994/TinyHTTP.git
   ```
   
4. **Create a new branch** for your feature or fix:
    ```bash
    git checkout -b feature/YourFeature
    ```
5. **Make your changes** and commit them with descriptive messages:
    ```bash
    git commit -m "Add new feature"
    ```
6. **Push your changes** to your forked repository:
    ```bash
    git push origin feature/YourFeature
    ```
7. **Open a pull request** to merge your changes into the main repository.



### 📝 License

This project is open-source and available under the MIT License.


#### Feel free to reach out if you have any questions or need help!
