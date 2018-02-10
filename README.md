# Conky-VLC

Bridging the gap between playing something in cvlc (or vlc) and seeing what you are playing in conky.
This code is dead simple and will be very easy to change to suit your own needs.

## Getting Started

Clone this repository.

cd conky-vlc

go install

### Prerequisites

Installed and working: vlc and conky.

You need to be able to compile go code.  See golang.org for more information.

### Example

In this example a streaming url is being used, but any media should be fine.
The important part is that the local http endpoint is being started with a password.

```
cvlc http://<stream-url> --extraintf http --http-password <my-password>
```
See https://wiki.videolan.org/VLC_HTTP_requests/

Edit start_example.sh, change the password value to match <my-password>, and then make sure your script is executable.

```
chmod +x ./start_example.sh
```

Edit your .conkyrc file and add the following:

```
${color #0077ff}Playing
 ${color #7f8ed3}${exec ~/go/src/github.com/4ydx/conky-vlc/start_example.sh}     
```

Finally rename and move scripts around as you like, but be sure to check paths.

### Troubleshooting

Make sure to start conky in a terminal so that you can see any error output it might produce.
This will help you solve any path related issues you might have when conky tries to run ./start_example.sh.

```
killall conky
conky
```

## Authors

* **Nathan Findley**

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details
