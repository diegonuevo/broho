# What is Broho?

Broho stands for BRowser On HOst. It's a simple client/server tool that allows a client to request opening URLs on a remote system browser. The interaction between the client and server is implemented using the gRPC framework which relies on Protocol Buffers data format. The server side (`brohod`) is written in Python, whereas the client side (`broho-push`) is available in both Python and Go.

# Motivation

The initial driver that led to the creation of this simple project was a minor annoyance I face in my day to day activity at work. My setup consists in a Linux laptop (Ubuntu) and a Windows VM running on KVM/libvirt to run some programs such as the Office 365 software suite (Outlook, PowerPoint... you know). Since the hardware resources allocated to the VM are limited, I try to avoid opening a web browser in the VM, to keep memory usage low. Besides, I'd rather have a single browser instance running on the host, with all tabs grouped under a single window/session.

Taking this into account, the ideal situation for me would be having links on the VM (in Outlook, for example) open on the browser on the Linux host instead. So I decided to write this tool to implement this.

As a bonus, I took this opportunity as a challenge, to experience with a few technologies I hadn't used before:
- **gRPC and Protocol Buffers:** This is a very simple client/server application that serves as a proof of concept of this amazing and powerful framework
- **Golang:** I'm no professional developer, but I code from time to time. I'd written some code in Python, Java, C/C++, TCL, Shell Scripts and JavaScript in the past, but I'd never gone beyond the 'Hello World' in Go, so this can be considered my baptism.

# How to use

## Install server

You first need to install the server-side component on the Linux host:
1. Place the contents of the `python` subdirectory in a system-wide location (for example `/opt/broho/`)
2. Install the Python module requirements with `pip3 install -r /opt/broho/requirements.txt` as your current user (the one you intend to launch the browser sessions with)
3. Create a systemd service to run the daemon. I provided a systemd unit file (`python/brohod.service`) that you can use as an example. Just copy it to `~/.config/systemd/user/brohod.service` under your user's home directory.

*Note: Bear in mind that the service needs to be run by the actual user having the GUI session (if you try to run it as root, you won't get URLs open in the browser under your GUI session)*

4. If you don't provide a config file for the service, it will by default listen on *127.0.0.1:50051* (useful for simple testing, but useless for our use case, since we need to make it available from the VM). Copy the provided `python/broho.conf` file to either `/etc/broho/` or `~/.config/broho/` and customize the `BindAddress` parameter to suit your needs (in my case, I used the IP address of the virbr0 NAT interface that connects the host to the VM).
5. Enable and start the systemd service you just created:
```
systemctl --user enable brohod.service
systemctl --user start brohod.service
```
6. Optionally, check the logs by looking at the system journal with `journalctl -t brohod`

## Install client

Then, install the client-side component on the Windows VM:
1. Compile the Go version of the client at `go/broho-push.go` with the Windows version of Go:
```
broho\go> go build -ldflags -H=windowsgui broho-push.go
```

*Note: the flags included prevent the CLI window from showing up when launching the program in windows*

2. Place the resulting `broho-push.exe` file in `C:\Program Files\Broho\broho-push.exe`
3. Run the program. On the first run it will create a config file with the default settings at `%APPDATA%\broho\config.json`. Edit the file and replace the `bind-address` parameter with the IP address in your setup.
4. Import the contents of the `windows\set-default.reg` file on your Windows registry by double clicking on the file to set Broho as the default browser on your Windows system
5. Head over to *Settings > Apps > Default Apps* and set Broho as your default browser
6. Go try clicking a link on your system. If everything is correct, it should now open on the browser of the Linux Host


# Security concerns

The gRPC communication between the client and the server is unencrypted, and unauthenticated. This poses a security risk if you let the server listen on an IP address reachable over your network. Anybody implementing this gRPC protocol could potentially push URLs to your daemon server and flood your system with URL open requests.

In my case I configured the server to listen on an internal interface (virbr0) that is only reachable between the VM and the Linux host.

In the future, if I feel like it, I may implement encryption and authentication features.

# Roadmap

Since this tool follows a client/server architecture and uses the general purpose and extensible gRPC framework, it can be potentially extended to implement many other features requiring VM-to-host communication. In the future I may rewrite the tool to act as a BRoker On Host to do many other things, instead of just a simple BRowser.
