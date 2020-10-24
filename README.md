# **K**ey**C**hain**C**lipboard
> CLI command which stores your credentials and load passwords in clipboard when you need them

```
Store your services' credentials and load passwords in clipboard when you need them:

kcc add -s facebook.com -u john@doe.com
kcc get -s facebook.com -u john@doe.com

Usage:
  kcc [command]

Available Commands:
  add         store service credentials
  del         delete service credentials
  get         get password
  help        Help about any command
  version     show current version

Flags:
      --config string   config file (default is $HOME/.kcc.yaml)
  -h, --help            help for kcc
  -v, --verbose         verbose output

Use "kcc [command] --help" for more information about a command.
```

## How it works
Credentials are saved in a file with `root` permissions so you must enter your system password.
Passwords are copied in clipboard when you need them.

## Custom storage file
Default storage file is located in `$HOME/.kcc.storage`, but if you want to customize it you have to follow these steps:
- create a config file – default is `$HOME/.kcc.yaml`
- set the *storage* key in the yaml file with an absolute path to the storage file as value. Example:
```yaml
storage: /absolute/path/to/storage/file
```
- if config file location is not equal to the default, run `kcc` with the ` --config /path/to/config.yaml ` flag

## Author
[Davide Caruso](https://about.me/davidecaruso)

## License
Licensed under [MIT](LICENSE).
