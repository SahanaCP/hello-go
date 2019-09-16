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
		wasm "github.com/wasmerio/go-ext-wasm/wasmer"
        
 )

func CheckErrorU(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
func CalcRegex(message string) int32{
    
	// Reads the WebAssembly module as bytes.
	bytes, _ := wasm.ReadBytes("./eval.wasm")
	// Instantiates the WebAssembly module.
	instance, _ := wasm.NewInstance(bytes)
	defer instance.Close()
	// Gets the 'eval' exported function from the WebAssembly instance.
	eval := instance.Exports["eval"]
	re := regexp.MustCompile(`[0-9]*`)
       // symbol := regexp.MustCompile(`[\%\.]`)
	
	submatchall := re.FindAllString(message, -1)
	var arr [2]int
	index := 0
	for _, element := range submatchall {
		//fmt.Println(element)
		num, err := strconv.Atoi(string(element))
	        if err == nil {
			arr[index]=num
	        }
	    index++	
	}
    
	var temp int32 	
	        	flag5:= strings.Contains(message,"*")
			if flag5{
				result,_ := eval(arr[0], '*', arr[1])
				temp := result.ToI32()
				return temp
			}
			flag6:= strings.Contains(message, "^")
			if flag6{
				result,_ := eval(arr[0], '^', arr[1])
				temp := result.ToI32()
				return temp
			}
			flag:=strings.Contains(message,"%")
			if flag{
				result, _ := eval(arr[0], '%', arr[1])
				temp := result.ToI32()
				return temp
			}
			flag1:=strings.Contains(message,"+")
			if flag1{
				result, _ := eval(arr[0], '+', arr[1])
				temp := result.ToI32()
				return temp
			}
			flag2:=strings.Contains(message,"-")
			if flag2{
				result, _ := eval(arr[0], '-', arr[1])
				temp := result.ToI32()
				return temp
			}
			flag3:=strings.Contains(message,"/")
			if flag3{
				result, _ := eval(arr[0], '/', arr[1])
				temp := result.ToI32()
				return temp
			}
			
	
	return temp 
	
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
		if (strings.Contains(string(message), "#")){
			message = strings.Replace(message, "#", "%", -1)
		}
    		fmt.Print("Message Received:", string(message))
		
		
		temp := strings.TrimSpace(string(message))
                if temp == "STOP" {
			newmessage := string("STOP")
			// send new string back to client
			c.Write([]byte(newmessage + "\n"))
                    break
                }
		// Extract numbers
		nums := CalcRegex(string(message))		
		
		newmessage := string("Answer: ")
		numstr := strconv.Itoa(int(nums))

		
    	// send new string back to client
    	c.Write([]byte(newmessage + numstr + "\n"))
		
     }

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
			
		
 	}
 }
