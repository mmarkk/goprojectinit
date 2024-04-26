# goprojectinit

## Description

This is a Go application. Replace this line with a brief description of your application.

-   Prompts the user for the application name.
-   Creates a base directory with the given application name.
-   Establishes a set of common directories like cmd, internal, pkg, etc., which are typical in a structured Go application.

This CLI tool simplifies the setup of a new Go project by automatically generating a clean, organized directory structure, following common conventions in the Go community. Adjust the folder names and structure as needed based on your specific project requirements or organizational standards

## Directory Structure

-   `cmd`:
-   `cmd/<appName>`: <appname> will be replaced with the application name you enter
-   `configs`:
-   `internal`:
-   `pkg`:
-   `scripts`:
-   `vendor`:
-   `api`:
-   `ui`:
-   `deployments`:

## How to Run

To run this application, follow these steps:

1. Make sure you have Go installed on your machine. If not, you can download it from the official Go website: [https://golang.org/](https://golang.org/).

2. Clone the repository to your local machine:

    ```
    git clone https://github.com/mmarkk/goprojectinit.git
    ```

3. Change to the project directory:

    ```
    cd goprojectinit
    ```

4. Build the application:

    ```
    go build ./cmd/goprojectinit/
    ```

5. Run the application:

    ```
    ./goprojectinit
    ```

6. You should now see the application running in your terminal.

If you encounter any issues or have any questions, please feel free to reach out.

## Contributing

Thank you for your interest in contributing to this application! Here are some guidelines on how you can contribute:

1. Fork the repository and clone it to your local machine.
2. Create a new branch for your contribution: `git checkout -b feature/your-feature-name`.
3. Make your changes and commit them: `git commit -m "Add your commit message"`.
4. Push your changes to your forked repository: `git push origin feature/your-feature-name`.
5. Open a pull request in the original repository and provide a clear description of your changes.

We appreciate your contributions and will review your pull request as soon as possible. Thank you for helping to improve this application!

## License

This code is released under the [MIT License](https://opensource.org/licenses/MIT). You are free to use, modify, and distribute this code as long as you include a reference to the original source.
