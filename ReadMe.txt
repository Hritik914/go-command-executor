Go Command Executor
This Go program implements a simple HTTP server that allows you to execute shell commands remotely. It provides both GET and POST endpoints to execute commands and retrieve the output.
Usage
Make sure you have Go installed on your system.
Clone or download the project files to your local machine.
Open a terminal and navigate to the project directory.
Build and run the program using the following command:
	Go run main.go
This will start the HTTP server on port 8080.
Access the API endpoints using a tool like Postman, Thunder client or a web browser:
To execute a command via GET request, use the following URL format:
http://localhost:8080/api/cmd?command=<your-command-here>
Replace <your-command-here> with the actual command you want to execute. Make sure to URL encode any special characters.
To execute a command via POST request, send a JSON payload to the following URL:
http://localhost:8080/api/cmd
The JSON payload should have the following structure:
{
  "command": "<your-command-here>"
}
Replace <your-command-here> with the actual command you want to execute.
The server will return a JSON response containing the output of the executed command. The response will have the following format:
{
  "output": "<command-output>"
}

The <command-output> field will contain the output of the executed command.
Note
This program uses the os/exec package to execute shell commands. It runs the commands using sh -c <command>.
If the specified command is not found, the server will return a 404 Not Found response.
In case of any other errors during command execution, the server will return a 500 Internal Server Error response.
