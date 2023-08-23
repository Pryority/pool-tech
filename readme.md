# Pool Tech  

Friend Tech but part of the chat earnings go to funding a Rocket Pool and then the entire chat can split the reward ETH.

## Todo

- Either remove Tailwind Builder and just have CDN or fix tailwind builder so it works with all styles
- Add an embedded wallet (Privy) like Friend.Tech...?
- Literally all the hard smart contract stuff T_T (Rocket Pool, shared earnings, etc. etc.)
- Remove unneccessary DB stuff and other miscellaneous stuff from gowind
- Update readme

## Technology

- [**Go**](https://go.dev/doc/): The backend of the application is written in Go, a powerful and efficient programming language. Go provides excellent support for concurrency and has a strong standard library, making it suitable for building high-performance applications.

- [**Air**](https://github.com/cosmtrek/air): A development tool used to live reload Golang applications.

- [**Fiber**](https://docs.gofiber.io/): Fiber is a web framework for Go that is designed to be fast and efficient. It provides a simple and intuitive API for building web applications and APIs. Fiber is used in this application to handle HTTP requests and routing.

- [**HTML Templates**](https://docs.gofiber.io/guide/templates/): The application uses HTML templates for generating dynamic HTML content. The `github.com/gofiber/template/html/v2` package is used for rendering HTML templates with Fiber.

- [**Tailwind CSS**](https://tailwindcss.com/): Tailwind CSS is a utility-first CSS framework that provides a set of pre-defined utility classes for building responsive and modern user interfaces. It is used in this application for styling the frontend components.

- [**Docker**](https://www.docker.com/): Docker is a containerization platform that allows applications to be packaged into containers, providing a consistent and isolated environment for running the application. Docker is used in this application to build and run the application in a containerized environment.

- [**Node.js and npm**](https://nodejs.org/en): Node.js is a JavaScript runtime that allows executing JavaScript outside of a web browser. npm (Node Package Manager) is a package manager for JavaScript libraries and modules. They are used in this application for managing frontend dependencies and running build scripts.

## Building for production

To build the application for production, you can run the following command:

```bash
npm run build
```

This command will trigger the build process, which includes generating the production-ready CSS file using Tailwind CSS. The compiled CSS file will be saved as `./static/main.css`.

Please note that during the build process, you may see warnings related to the JIT (Just-in-Time) engine of Tailwind CSS, as it is currently in preview. These warnings inform you about the nature of preview features and their potential impact on your codebase.

After the build process is complete, the application will be ready for deployment in a production environment.

## Makefile

The project includes a Makefile that provides convenient commands for common tasks. Here are some available commands:

- **install**: This command will navigate to the `web` directory and run `npm install` to install the frontend dependencies.

- **dev**: This command will run the `air` command, which facilitates live reloading of the backend code during development.
