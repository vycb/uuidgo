//
//  libuuid sample program
//
//  library install for debian
//      $ sudo apt-get install uuid-dev
//
//
package uuid

import (
	"testing"
	"log"
	"github.com/jakehl/goid"
)
var(
	uuid1, uuid2, v4UUID string
)
func TestGetUUIDUpper(t *testing.T) {
	uuid1 = GetUUIDUpper()
	l := len(uuid1)
	log.Printf("%s\n",uuid1)
	if l != 36 {
		t.Errorf("uuid:%v:%d string is not correct",uuid1, l)
	}
}

func TestGetUUIDLower(t *testing.T) {
	uuid2 = GetUUIDLower()
	l := len(uuid2)
	log.Printf("%s\n",uuid2)
	if l != 36 {
		t.Errorf("uuid:%v:%d string is not correct",uuid2,l)
	}
}

func TestGetUUIDMd5Upper(t *testing.T) {
	u1 :="24F4F1D8-E394-11E8-98E0-F0DEF1A82164"
	//u1 :="1b4e28ba-2fa1-11d2-883f-0016d3cca427"
	s1 :="golang.org"
	//s1 :="github.org"
	uuid2 = GetUUIDMd5Upper(u1, s1)
	l := len(uuid2)
	log.Printf("MD5:%s\n",uuid2)
	if l != 36 {
		t.Errorf("MD5:%v:%d string is not correct",uuid2,l)
	}
}

func TestGetUUIDMd5Lower(t *testing.T) {
	u1 :="24f4f1d8-e394-11e8-98e0-f0def1a82164"
	s1 :="golang.org"
	//s1 :="github.org"
	uuid2 = GetUUIDMd5Lower(u1, s1)
	l := len(uuid2)
	log.Printf("MD5:%s\n",uuid2)
	if l != 36 {
		t.Errorf("MD5:%v:%d string is not correct",uuid2,l)
	}
}

func TestGetUUIDSha1Upper(t *testing.T) {
	//u1 :="24F4F1D8-E394-11E8-98E0-F0DEF1A82164"
	u1 :="1B4E28BA-2FA1-11D2-883F-0016D3CCA427"
	//s1 :="golang.org"
	s1 :="github.org"
	uuid2 = GetUUIDSha1Upper(u1, s1)
	l := len(uuid2)
	log.Printf("SHA1:%s\n",uuid2)
	if l != 36 {
		t.Errorf("SHA1:%v:%d string is not correct",uuid2,l)
	}
}

func TestGetUUIDSha1Lower(t *testing.T) {
	//u1 :="24F4F1D8-E394-11E8-98E0-F0DEF1A82164"
	u1 :="1b4e28ba-2fa1-11d2-883f-0016d3cca427"
	//s1 :="golang.org"
	s1 :="github.org"
	uuid2 = GetUUIDSha1Lower(u1, s1)
	l := len(uuid2)
	log.Printf("SHA1:%s\n",uuid2)
	if l != 36 {
		t.Errorf("SHA1:%v:%d string is not correct",uuid2,l)
	}
}

func TestUUIDCompare(t *testing.T) {
	res := UUIDCompare(uuid1, uuid1)
	log.Printf("compare1:%d\n", res)
	if res != 0 {
		t.Error("UUIDCompare 1 is not correct")
	}
	res = UUIDCompare(uuid1, uuid2)
	log.Printf("%s:%s:compare2:%d\n",uuid1, uuid2, res)
	if res == 0 {
		t.Error("UUIDCompare 2 is not correct")
	}
}

const (
  lower = iota
  upper
	md5l
	md5u
	sha1l
	sha1u
	gid
)

func benchmarkGetUUID(b *testing.B, mode int){
	b.ReportAllocs()
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
		switch mode {
      case lower:
				uuid1 = GetUUIDLower()
      case upper:
				uuid1 = GetUUIDUpper()
      case md5u:
				uuid1 = GetUUIDMd5Upper("C0086038-E39E-11E8-98E0-F0DEF1A82164", "github.org")
      case md5l:
				uuid1 = GetUUIDMd5Lower("c0086038-e39e-11e8-98e0-f0def1a82164", "github.org")
      case sha1l:
				uuid1 = GetUUIDSha1Lower("1b4e28ba-2fa1-11d2-883f-0016d3cca427", "golang.org")
      case sha1u:
				uuid1 = GetUUIDSha1Upper("1B4E28BA-2FA1-11D2-883F-0016D3CCA427", "golang.org")
      case gid:
				v4gob := goid.NewV4UUID()
				uuid1 = v4gob.String()
			}
	}
	//log.Printf("mode %d, %s", mode, uuid1)
}

func BenchmarkGetUUIDUpper(b *testing.B){ benchmarkGetUUID(b, upper)}
func BenchmarkGetUUIDLower(b *testing.B){ benchmarkGetUUID(b, lower)}
func BenchmarkGetUUIDMd5Upper(b *testing.B){ benchmarkGetUUID(b, md5u)}
func BenchmarkGetUUIDMd5Lower(b *testing.B){ benchmarkGetUUID(b, md5l)}
func BenchmarkGetUUIDSha1Upper(b *testing.B){ benchmarkGetUUID(b, sha1u)}
func BenchmarkGetUUIDSha1Lower(b *testing.B){ benchmarkGetUUID(b, sha1l)}
func BenchmarkGetGOID     (b *testing.B){ benchmarkGetUUID(b, gid)}
