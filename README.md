# conn-check

Checks the connectivity of HSN devices in a cluster over the IP layer.

## Build

Building the `main` binary:

```bash
mkdir -p build/
go build -o ./build/main ./cmd/main.go
sudo setcap cap_net_raw=+ep ./build/main
```

## Usage

Running the `main` binary:

```bash
./build/main
```

## Troubleshooting

Permission denied on socket usage:

```console
panic: socket: permission denied
```

By default, Linux has a restricted range for ping ranges:

```console
➜  ~ sudo sysctl -a | grep net.ipv4.ping_group_range
net.ipv4.ping_group_range = 1	0
```

The binary which uses this library needs to have Effective and Permitted capabilities added for `cap_net_raw`,
for using RAW and PACKET sockets:
https://man7.org/linux/man-pages/man7/cap_text_formats.7.html
https://man7.org/linux/man-pages/man7/capabilities.7.html

```bash
sudo setcap cap_net_raw=+ep <path_to_binary>
```
