# Procgo

## Install

Download latest [release](https://github.com/nigimaxx/procgo/releases/latest)

```sh
# Unpack tar
tar -xzf procgo_v0.0.0_darwin_x86_64.tar.gz

# Move to location in PATH
mv procgo /usr/local/bin/
mv procgo-daemon /usr/local/bin/

# Create required daemon log file
sudo mkdir /var/log/procgo
sudo touch /var/log/procgo/daemon.log
sudo chmod 666 /var/log/procgo/daemon.log
```

### Completion

```sh
procgo completion

# Move to some location in FPATH
mv _procgo /usr/local/share/zsh/site-functions
```

## Usage

```
procgo is tool to run local services concurrently.
It is similar to foreman. The main difference is
that it consists of a client daemon architecture
which allows it to start/stop/restart services independently.

Usage:
  procgo [command]

Available Commands:
  completion  generates zsh completions
  help        Help about any command
  kill        kills the daemon
  logs        prints the logs of the services
  restart     restarts the provided services
  start       starts the provided services
  status      lists all running services
  stop        stops the provided services

Flags:
  -h, --help              help for procgo
  -j, --procfile string   procfile (default "Procfile")
  -v, --version           version for procgo

Use "procgo [command] --help" for more information about a command.
```
