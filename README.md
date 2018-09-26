# gocli_exercises
Some projects I'm working on this week need CLI tools, so I'm doing some exercises to refresh myself about CLI implementations and integrations with golang, JS, and Python. Each branch corresponds to a different exercise.

## basic_json_echo:

This branch is a minimal implementation of a json-handling golang cli. It has no program logic, so it's mostly useless.

### Try it out with Docker:
To build it with docker, run this from the project directory:
> docker build -t gocli_exercise . && docker run -it gocli_exercise
This command builds, starts an interactive session in the container, and puts you in the bash terminal as root.

From bash in the container, you can try out the basic flags and commands:
> app -e '{"foo": "bar", "baz": ["foo", 1, "bar", {"baz": true}]}'
> app -e -d '{"foo": "bar", "baz": ["foo", 1, "bar", {"baz": true}]}'
> app -f -e example.json
> app -f -e example_invalid.json

Then leave the container:
> exit

### Install it as a CLI:
Golang has streamlined installation for CLIs. From the project directory:
> go build && go install

### Uninstalling is also a breeze. From the project directory:
> go clean -i