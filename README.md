# uuidgo
libuuid for golang by cgo

library install for debian  
$ sudo apt install uuid-dev

 Usage:
	import "github.com/vycb/uuid"	 
  oUuid = NewGen()  
  defer oUuid.Close()  
  
  if oUuid.IsNull() {  
    t.Error("Error to initialize NewGen")  
  }  
  log.Printf("NewGen UnparseUpper:%s\n", oUuid.UnparseUpper())  


  
