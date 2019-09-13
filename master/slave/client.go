package main

 import (
         "fmt"
         "net"
	     "bufio"
         "os"
         "io"
         "strings"
 )

func CheckErrorU(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

 func sendFile(serverAddr string) {

        
		conn, err := net.Dial("tcp", serverAddr)
		CheckErrorU(err)
		defer conn.Close()
		fmt.Printf("Connection established between %s and localhost.\n", serverAddr)
         fmt.Printf("Remote Address : %s \n", conn.RemoteAddr().String())
         fmt.Printf("Local Address : %s \n", conn.LocalAddr().String()) 
		 
        //text := "File"
        fmt.Fprintf(conn, "File" + "\n")
        // listen for reply
    		message, _ := bufio.NewReader(conn).ReadString('\n')
			temp := strings.TrimSpace(string(message))
                if temp == "STOP" {
				    fmt.Print("Message from server: "+message)
                }else {
				// send to socket
        
				srcFile := "./eval.wasm" 
				// open file to upload
				fmt.Printf("Sending File from client: %s \n", conn.LocalAddr().String())
				fi, err := os.Open(srcFile)
				CheckErrorU(err)
				defer fi.Close()

				// upload
				_, err = io.Copy(conn, fi)
				CheckErrorU(err)
				
				}
       
   }


func sendRegx(serverAddr string){
        
// Wait for some data to be input

			conn, err := net.Dial("tcp", serverAddr)
			CheckErrorU(err)
			defer conn.Close()
			//text := "File"
        fmt.Fprintf(conn, "Input" + "\n")
			
          for { 
           // read in input from stdin
    		reader := bufio.NewReader(os.Stdin)
    		fmt.Print("Enter Infix Expression: ")
    		text, _ := reader.ReadString('\n')
    		// send to socket
    		fmt.Fprintf(conn, text + "\n")
			
			
    		// listen for reply
    		message, _ := bufio.NewReader(conn).ReadString('\n')
			temp := strings.TrimSpace(string(message))
                if temp == "STOP" {
				    break
                }else {
				fmt.Print("Message from server: "+message)
				}
        
           }
}


 func main() {
         
         
		servAddress := "localhost:6000"
		
		// Send the wasm file to the server
		sendFile(servAddress)
        
		// Receive inputs and send to server
        sendRegx(servAddress)
        //conn.Close()
}


 

 