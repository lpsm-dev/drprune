## Concepts

### Cobra

Cobra is a CLI framework for Golang. Using it you can speed up your development and creating a powerful and modern CLI application. Cobra is built on a structure of commands, arguments and flags:

- Commands represent actions.
- Args are things.
- Flags are modifiers for those actions.

The best applications will read like sentences when used. Users will know how to use the application because they will natively understand how to use it. This pattern is: `APPCLI VERB NOUN --ADJECTIVE` or `APPCLI COMMAND ARG --FLAG`. A few good real world examples may better illustrate this point:

```bash
git clone URL --bare
```

or 

```
hugo server --port=1313
```

#### Commands

Command is the central point of the application. Each interaction thar the application supports will be contained in a command. We can create commands with children commands and optionally run an `action`. In the example above, `server` is the command.

#### Flags

A flag is a way to modify the behavior of a command. Cobra supports fully POSIX-compliant flags as well the Go flag package. A Cobra command can define flags that persist through to children commands and flags that are only available to that command. In the example above, `port` is the flag.
