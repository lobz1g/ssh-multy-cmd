# ssh-multy-cmd

Small utility for sending some commands to different hosts and receive response from this one

## Configuration

In config/config.json placed fields for connection to hosts

```json
[
  {
    "host": "host1:port_for_host1",
    "user": "user_for_host1",
    "password": "pass_for_host1",
    "cmd": "command_for_host1"
  },
  {
    "user": "user",
    "password": "pass",
    "host": "127.0.0.1:22",
    "cmd": "ls -l"
  }
]
```

## Output

There are two outputs channels. Console and file. 
* In the console will be symbol `X` for bad response or some errors and symbol `V` for good result
* In the file will be detail information about error/response

File can be named:
* `host.log` for response from the host. e.g `127_0_0_1.log`
* `localhost.log` for errors
