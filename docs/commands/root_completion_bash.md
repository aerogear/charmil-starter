## root completion bash

Generate command completion script for Bash shell

### Synopsis

' Generate command completion script for Bash shell. Installing on Linux: 1. Run `cli completion bash > cli_completions` to create the script file. 2. Run `sudo mv cli_completions /etc/bash_completion.d/cli` to place the script in a special Bash completions directory. Installing on macOS: 1. Run `cli completion bash > cli_completions` to create the script file. 2. Run `sudo mv cli_completions /usr/local/etc/bash_completion.d/cli` to place the script in a special Bash completions directory. '

```
root completion bash
```

### Options

```
  -h, --help   help for bash
```

### SEE ALSO

* [root completion](root_completion.md)	 - Outputs command completion for the given shell (bash, zsh, or fish)

###### Auto generated by spf13/cobra on 5-Aug-2021
