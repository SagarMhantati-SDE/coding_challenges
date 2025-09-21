What is port scanner?
-> Port scanner sends a network request to connect to specific TCP / UDP port on computer and records the response.
-> So what port scanner does is it sends the packet of network data to a port to check the current status. If you wanted to 
   check to see if your web server was working correctly, you would check the status of port 80 on that server to make sure port is open and listining.
-> First 1023 TCP ports are well know ports reserved for applications like FTP(21), HTTP(80), SSH(22) and Internet Assigned Autority reserves these points to keep them standerized.
-> TCP ports 1024-49151 are available for use by services or applications and you can register them with IANA so they are considered semi reserved ports. Ports 49152 and higher are free to use.

Port Scanning Basics:
Port scanner sends UDP or TCP network data packet and asks ports about their current status. There types of response are below:
1. Open, Accpeted: The computer responds and asks if there is anything it can do for you.
2. Closed, Not listning: The computer responds that "The port is currently in use and unavilable at this time"
3. Filtered, Dropped, Blocked: The computer even isn't bother to respond.