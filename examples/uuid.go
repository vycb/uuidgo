//
//  libuuid sample program
//
//  library install for debian
//      $ sudo apt-get install uuid-dev
//
//
package main

import (
	"os"
	"log"
	"github.com/jakehl/goid"
	"github.com/vycb/uuid"
)
var(
	uuid1, uuid2, v4UUID string
)
func TestGetUUIDUpper() {
	uuid1 = uuid.GetUUIDUpper()
	l := len(uuid1)
	log.Printf("%s: %d\n",uuid1, l)
	if l != 36 {
		log.Printf("uuid:%v:%d string is not correct",uuid1, l)
	}
}

func TestGetUUIDLower() {
	uuid2 = uuid.GetUUIDLower()
	l := len(uuid2)
	log.Printf("%s: %d\n",uuid2, l)
	if l != 36 {
		log.Printf("uuid:%v:%d string is not correct",uuid2,l)
	}
}

func TestUUIDCompare() {
	res := uuid.UUIDCompare(uuid1, uuid1)
	log.Printf("compare1:%d\n", res)
	if res != 0 {
		log.Println("UUIDCompare 1 is not correct")
	}
	res = uuid.UUIDCompare(uuid1, uuid2)
	log.Printf("%s:%s:compare2:%d\n",uuid1, uuid2, res)
	if res == 0 {
		log.Println("UUIDCompare 2 is not correct")
	}
}

const (
  lower = iota
  upper
	gid
)

func benchmarkGetUUID(mode int){
  for i := 0; i < 500000; i++ {
		switch mode {
      case lower:
				uuid1 = uuid.GetUUIDLower()
				//log.Println(uuid1)
      case upper:
				uuid1 = uuid.GetUUIDUpper()
				//log.Println(uuid2)
      case gid:
				v4gob := goid.NewV4UUID()
				uuid1 = v4gob.String()
				//log.Println(v4UUID)
			}
	}
	log.Printf("mode %d, %s", mode, uuid1)
}

func BenchmarkGetUUIDLower(){ benchmarkGetUUID(lower)}
func BenchmarkGetUUIDUpper(){ benchmarkGetUUID( upper)}
func BenchmarkGetGOID     (){ benchmarkGetUUID(gid)}

func main(){
	var bench = ""
	if len(os.Args) > 1 {
		bench = os.Args[1]
	}
	switch bench {
		case "lower":
			BenchmarkGetUUIDLower()
		case "upper":
			BenchmarkGetUUIDUpper()
		case "gid":
			BenchmarkGetGOID()
		default:
			log.Println(uuid.GetUUIDUpper())
	}
}
