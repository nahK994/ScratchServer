## TCPickle - A Lightweight TCP Server

Welcome to the **TCPickle** repository! This project is a work in progress and aims to implement a basic TCP server from scratch using Go.

### 🚧 Project Status

This project is still in development. Contributions, ideas, and feedback are welcome!

### 🛠️ Features

- Accepts multiple client connections.
- Handles basic HTTP methods such as GET and POST.
- Simple error handling and logging.

### 🚀 Getting Started

#### Prerequisites

To run this project, you'll need:

- **Go** (version 1.20 or later)

#### How to Use

You can easily add this package to your Go project by running the following command:

```go
go get github.com/nahK994/TCPickle
```

After installing the package, you can start the HTTP server locally in your main.go file with the following code:

```go
srv := server.InitiateHttp("127.0.0.1:8000")
srv.RequestHandler("/post", "POST", func(r models.HttpRequest, w *models.HttpResponse) {
    fmt.Println("TEST ===>", r.Body)
    w.StatusCode = 201
    w.Body = r.Body
})

srv.RequestHandler("/get", "GET", func(r models.HttpRequest, w *models.HttpResponse) {
    fmt.Println("TEST ===>", r.Body)
    w.StatusCode = 200
    w.Body = "Hello World!!"
})
log.Fatal(srv.Start())
```

You can also run RESP server by running following codes:

```go
srv := server.InitiateResp("127.0.0.1:8000")
srv.RequestHandler(func(request models.RespRequest, response *models.RespResponse) {
    response.Response = "+OK\r\n"
})
log.Fatal(srv.Start())
```

You can also check out the complete example in `cmd/main.go`.

### 📂 Project Structure

- **`server/`**: Contains the server logic and connection handling.
- **`handlers/`**: Contains command parsing and handling logic.
- **`models/`**: Contains structure of data.
- **`utils`**: Containes the utility varialbles, constants.

### 📈 Roadmap

- Handle multiple client requests.
- Handle basic HTTP methods such as GET and POST.
- Improve error handling.
- Implement logging enhancements.
- Add support for secure connections (TLS).

### 🤝 Contributing

We welcome contributions of all kinds! Whether you're fixing bugs, adding new features, or improving documentation, your input is appreciated.

#### How to Contribute

1. **Fork the repository** to your GitHub account.
2. **Clone your forked repository** to your local machine:

   ```bash
   git clone https://github.com/nahK994/TCPickle.git
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
