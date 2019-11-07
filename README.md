# hello-go
Design, develop and implement a simple network system in Golang that creates a master-slave model with the following functionalities: 
Master Node: Master should be able to receive a WASM binary from slave and execute the arithmetic expressions based on the inputs given by slave and send the response back to the corresponding slave node/program.
Slave Node: Slave should be able to send a WASM binary along with the required inputs dynamically and wait for the response from the master node/program. We expect you to design the network and the user interaction at the command-line interface for the slave program.

Solution: 
The program "serverwasm.go" is the Master node which imports "github.com/wasmerio/go-ext-wasm/wasmer" and calls the eval() function as required.

The program "serverwasm.go" is the Master node, it is a basic TCP server which is used to check if the file eval_wasm exists in the "master" directory and then requests the client to transfer the file. 

There is a simple messaging mechanism used between slave and master:
File - File transfer is initiated from client  
Input - Input strings are sent from client
STOP - Close the slave connection, sent from client 
OK - Master indicates that the slave can start file transfer

Master accepts multiple connections from slave tasks.

Solution:
The program "client.go" is the client node, it initiates a file transfer of eval.wasm on receiving the message "OK" from master. Sequentially starts requesting input infix expressions from user and sends it to the Master and displays the answer to the user. The infix expressions are returned in int64 format.


Folders
-------

master

	//This is the server file
	** serverwasm.go

	//Screen shot with output
	** Capture.png

slave

	//This is the client file
	** client.go

	//The wasm file eval.wasm
	**eval.wasm
