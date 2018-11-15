# execitu

## Installation
```
go get -u github.com/theoden9014/exercitu/...
```

## Usage
```
$ exercitu --help
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  exercitu [command]

Available Commands:
  help        Help about any command
  http        http mode

Flags:
  -c, --client string    client mode
      --config string    config file (default is $HOME/.exercitu.yaml)
      --count uint       requests per second
      --debug            debug output
  -h, --help             help for exercitu
      --logfile string   send output to a log file
  -P, --parallels int    number of parallels (default 1)
  -q, --quiet            quiet output
  -r, --rate uint        requests per second (default 50)
  -s, --server           server mode
  -t, --time uint        time in seconds
      --timeout int      time in seconds
  -v, --verbose          verbose output

Use "exercitu [command] --help" for more information about a command.
```

## Contribution
1. Fork it
2. Download your fork to your PC (git clone https://github.com/your_username/exercitu && cd exercitu)
3. Create your feature branch (git checkout -b my-new-feature)
4. Make changes and add them (git add .)
5. Commit your changes (git commit -m 'Add some feature')
6. Push to the branch (git push origin my-new-feature)
7. Create new pull request