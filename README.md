# UNIQLO Akamai T-shirt Code

## Overview
A Japanese fashion brand called UNIQLO collaborated with an Internet company called Akamai to design a T-shirt.  

The design of the T-shirt is based on a Go program.  
However, it is also a piece of missing Go language code.  

So, the missing parts were filled in and the code was supplemented so that it would work.  

Then, the program for a simple server was restored.  

I would like to introduce this code, including how to use it, below.  

### T-shirt's Design
![T-Shirt_View](https://user-images.githubusercontent.com/36861752/233776272-a1f52816-824f-43df-96c3-61dfe6fa02b7.png)

### T-shirt's Code
![Code_View](https://user-images.githubusercontent.com/36861752/233776315-c24c45c7-0a59-48e5-9bc0-df95942e7a3c.jpg)

## Output Sample
Since this code is a program for a server, the code is invoked on the console on the left.  
The console on the left is a temporary server.  

The console on the right is the client side, executing curl commands.  

This program responds with ACTIVE while processing on the server and INACTIVE while waiting.  

### Standby
![Pre](https://user-images.githubusercontent.com/36861752/233776338-f00756ee-fb4c-472a-bd5a-f9179f7d9dc0.png)

### Active
![Active](https://user-images.githubusercontent.com/36861752/233776362-bd615819-8a72-4617-b91f-91982971feab.png)

### Inactive
![Inactive](https://user-images.githubusercontent.com/36861752/233776394-6b8691d7-576d-4795-8321-67bf32716ee6.png)

## Command Memo
### Server side
~ $ go build -o Akamai Akamai.go  
~ $ ./Akamai  
Target: Akamai T-Shirt A , Count: 0  
Target: Akamai T-Shirt A , Count: 1  
Target: Akamai T-Shirt A , Count: 2  
Target: Akamai T-Shirt A , Count: 3  
Target: Akamai T-Shirt B , Count: 0  
Target: Akamai T-Shirt A , Count: 4  
Target: Akamai T-Shirt B , Count: 1  

### Client side
~ $ curl http://localhost:8080/status  
INACTIVE  
~ $ curl -X POST -d "target=Akamai T-Shirt A &count=30" http://localhost:8080/admin  
Control message issued for Target Akamai T-Shirt A   
~ $ curl -X POST -d "target=Akamai T-Shirt B &count=30" http://localhost:8080/admin  
Control message issued for Target Akamai T-Shirt B   
~ $ curl -X POST -d "target=Akamai T-Shirt C &count=30" http://localhost:8080/admin  
Control message issued for Target Akamai T-Shirt C   
~ $ curl http://localhost:8080/status  
ACTIVE  
