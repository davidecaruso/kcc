# **K**ey**C**hain**C**lipboard

```
Store your services' credentials and load passwords in clipboard when you need them:
   
   kcc add -s facebook.com -u john@doe.com -p secret
   kcc get -s facebook.com -u john@doe.com
   
   Usage:
     kcc [command]
   
   Available Commands:
     add         Store service credentials
     del         Delete service credentials
     get         Get password
     help        Help about any command
     version     Show current version
   
   Flags:
         --config string   config file (default is $HOME/.kcc.yaml)
     -h, --help            help for kcc
     -t, --toggle          Help message for toggle
   
   Use "kcc [command] --help" for more information about a command.
```
