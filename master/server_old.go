package main

 import (
 	"fmt"
 	"log"
 	"net"
        "bufio"
        "strings"
        "os"
        "io"
		"strconv"
        "regexp"
        
 )

func CheckErrorU(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
func CalcRegex(message string) int64{
    var ans int64
	re := regexp.MustCompile(`[0-9]*`)
    //symbol := regexp.MustCompile(`[\+\-\*\/\.]`)
	
	submatchall := re.FindAllString(message, -1)
	var arr [2]int64
	index := 0
	for _, element := range submatchall {
		//fmt.Println(element)
		num, err := strconv.ParseInt(string(element), 10, 64)
	        if err == nil {
			arr[index]=num
			//fmt.Println(arr[index])
	        }
	    index++	
	}
    
	
	        flag:=strings.Contains(message,"*")
			if flag{
			ans := arr[0]*arr[1]
			//fmt.Println("Answer:",ans)
			return ans
			}
			flag1:=strings.Contains(message,"+")
			if flag1{
			ans := arr[0]+arr[1]
			//fmt.Println("Answer:",ans)
			return ans
			}
			flag2:=strings.Contains(message,"-")
			if flag2{
			ans := arr[0]-arr[1]
			//fmt.Println("Answer:",ans)
			return ans
			}
			flag3:=strings.Contains(message,"/")
			if flag3{
			ans := arr[0]/arr[1]
			//fmt.Println("Answer:",ans)
			return ans
			}
			
	
	return ans
	
}
 func handleConnection(dstFile string, c net.Conn) {

 	log.Printf("Client %v connected. handle Connection", c.RemoteAddr())

	message, _ := bufio.NewReader(c).ReadString('\n')
	
	flag := strings.Contains(string(message), "File")
	   
	
	flagin := strings.Contains(string(message), "Input")
    //fmt.Print("Message Received:", string(message))

	print("Flag",flag, flagin)
    if flag {
       // check if the file is already created
		if _, err := os.Stat(dstFile); os.IsNotExist(err) {
     		log.Printf("file does not exist")
      		 newmessage := string("OK")
			// send new string back to client
			c.Write([]byte(newmessage + "\n")) 
			// create new file
			fo, err := os.Create(dstFile)
			CheckErrorU(err)
			defer fo.Close()

			// accept file from client & write to new file
			_, err = io.Copy(fo, c)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
				os.Exit(1)
				return
			}
			log.Printf("File Received!!", c.RemoteAddr())
        } else{
			log.Printf("File Already Received!!", c.RemoteAddr())
			newmessage := string("STOP")
			// send new string back to client
			c.Write([]byte(newmessage + "\n"))
			
		} 		    
	} else {
		  if flagin{
			handleFileInputs(c)
		  }
	
	}   
        	
		log.Printf("Connection from %v closed.", c.RemoteAddr())
		c.Close()
    }
	 

 func handleFileInputs(c net.Conn) {

 	log.Printf("Client %v connected. Handle Inputs", c.RemoteAddr())
	 for {
		// will listen for message to process ending in newline (\n)
    
	    message, _ := bufio.NewReader(c).ReadString('\n')
    		// output message received
    	fmt.Print("Message Received:", string(message))
		
		
		temp := strings.TrimSpace(string(message))
                if temp == "STOP" {
				    newmessage := string("STOP")
				// send new string back to client
					c.Write([]byte(newmessage + "\n"))
                    break
                }
		// Extract numbers
		num:= CalcRegex(string(message))		
		
    	// sample process for string received
    	//newmessage := strings.ToUpper(message)
		newmessage := string("Answer: ")
		numstr := strconv.FormatInt(num, 10)
    	// send new string back to client
    	c.Write([]byte(newmessage + numstr + "\n"))
		
     }
	 //log.Printf("Connection from %v closed.", c.RemoteAddr())
	 //c.Close()
}

 func main() {
 	ln, err := net.Listen("tcp", ":6000")
        dstFile := "./eval.wasm"

 	if err != nil {
 		log.Fatal(err)
 	}
	defer ln.Close()

 	fmt.Println("Server up and listening on port 6000")

 	for {
 		conn, err := ln.Accept()
			if err != nil {
 			log.Println(err)
 			return
			}
		
		       
 		go handleConnection(dstFile, conn)
		
		//go handleFileInputs(conn)
		
 	}
 }