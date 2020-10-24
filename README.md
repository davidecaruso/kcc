# **K**ey**C**hain**C**lipboard

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
  version     Show current version

Flags:
      --config string   config file (default is $HOME/.kcc.yaml)
  -h, --help            help for kcc
  -v, --verbose         verbose output

Use "kcc [command] --help" for more information about a command.
```

## Author
[Davide Caruso](https://about.me/davidecaruso)

## License
Licensed under [MIT](LICENSE).
