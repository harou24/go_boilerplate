# Go Boilerplate


<p align="center">
<img src="https://openupthecloud.com/wp-content/uploads/2020/01/Golang.png?ezimgfmt=rs%3Adevice%2Frscb2-1"  width=35% height=35%>
</p>


This is a Go boilerplate project to quickly start building an API server. It includes the following features:
- [Cobra](https://github.com/spf13/cobra) for command-line interface
- [Fiber](https://github.com/gofiber/fiber) as the web framework
- [Cors](https://github.com/gofiber/fiber/blob/main/middleware/cors/cors.go) middleware for handling Cross-Origin Resource Sharing

## Getting Started

1. Clone or download the repository.
2. Navigate to the root of the project directory in your terminal.
3. Run `go build` to build the project.
4. Run `./go_boilerplate start` to start the API server. By default, it will listen on port 5000. You can specify a different port using the `-p` or `--port` flag.
5. You can now start building your API routes in the `api` package.

## API Routes

By default, the boilerplate includes a single route `GET /` which returns a simple "Hello World 🖖" message. You can add more routes in the `api.Serve()` function in the `api` package.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT). Please see the [LICENSE](LICENSE) file for more details.

## Contributions

Contributions are welcome! If you find any bugs or have suggestions for improvements, please open an issue or submit a pull request.
