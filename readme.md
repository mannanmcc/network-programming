Until now we have a way one client can connect and send text message.

To test it, open up another terminal window and test using netcat command
$nc localhost 4242 
This line will open a connection with server, Then type something.

To install netcat follow following two command on mac
1. ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)" < /dev/null 2> /dev/null
2. brew install netcat