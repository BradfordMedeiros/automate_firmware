Simple tcp client that talks to fs_mount_mqtt daemon.  Can override tcp port to connect to.  

This initiates a tcp connection with the server to issue a command.   
Uses custom application protocol where commands are json formatted.  For details look in the request.go code.

Refer to the top level read me for specific commands.
