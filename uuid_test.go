//
//  libuuid sample program
//
//  library install for ubuntu, debian
//      $ sudo apt install uuid-dev
//
//
package uuid

import (
	"testing"
	"log"
	"github.com/jakehl/goid"
	gouuid "github.com/satori/go.uuid"
)
var(
	uuid1, uuid2, v4UUID string
	oUuid *UUID
)

func TestNewGenRandom(t *testing.T) {
	oUuid = NewGenRandom()
	defer oUuid.Close()
	if oUuid.IsNull() {
		t.Error("Error to initialize NewGenRandom")
	}
	log.Printf("NewGenRandom UnparseUpper:%s\n", oUuid.UnparseUpper())
	if len(oUuid.Val()) != 36 {
		t.Error("Error in UnparseUpper")
	}
	uuid2 = oUuid.Val()
	log.Printf("NewGenRandom UnparseLower:%s\n", oUuid.UnparseLower())
	if len(oUuid.Val()) != 36 {
		t.Error("Error in UnparseLower()")
	}
	res:=oUuid.Compare(uuid2)
	if res != 0  {
		t.Error("Error to Compare() 1 with UnparseUpper() in NewGenRandom")
	}else{
		log.Print("Compare() 1 Ok in NewGenRandom")
	}
	res = oUuid.Compare("24F4F1D8-E394-11E8-98E0-F0DEF1A82164")
	if res == 0 {
		t.Error("Error to Compare() 2 with uuid_parse in NewGenRandom")
	}else{
    log.Print("Compare() 2 Ok in NewGenRandom")
  }
	if oUuid.IsNull() {
		t.Error("Error to check IsNull() 1 in NewGenRandom")
	}else{
    log.Print("IsNull() 1 Ok in NewGenRandom")
  }
	oUuid.Clear()
	if ! oUuid.IsNull() {
		t.Error("Error to check IsNull() 2 in NewGenRandom")
	}else{
    log.Print("Clear() and IsNull() 2 Ok in NewGenRandom")
  }

}

func TestNewGen(t *testing.T) {
	goUuid := NewGen()
	defer goUuid.Close()
	if goUuid.IsNull() {
		t.Error("Error to initialize NewGen")
	}
	log.Printf("NewGen UnparseUpper:%s\n", goUuid.UnparseUpper())
}

func TestNewGenTime(t *testing.T) {
	tUuid := NewGenTime()
	defer tUuid.Close()
	if tUuid.IsNull() {
		t.Error("Error to initialize NewGenTime")
	}
	log.Printf("NewGenTime UnparseUpper:%s\n", tUuid.UnparseUpper())
}

func TestNewGenTimeSafe(t *testing.T) {
	tsUuid := NewGenTimeSafe()
	defer tsUuid.Close()
	if tsUuid.IsNull() {
		t.Error("Error to initialize NewGenTimeSafe")
	}
	log.Printf("NewGenTimeSafe UnparseUpper:%s\n", tsUuid.UnparseUpper())
}

func TestGetUUIDUpper(t *testing.T) {
	uuid1 = GetUUIDUpper()
	l := len(uuid1)
	log.Printf("GetUUIDUpper:%s\n",uuid1)
	if l != 36 {
		t.Errorf("uuid:%v:%d string is not correct",uuid1, l)
	}
}

func TestGetUUIDLower(t *testing.T) {
	uuid2 = GetUUIDLower()
	l := len(uuid2)
	log.Printf("GetUUIDLower:%s\n",uuid2)
	if l != 36 {
		t.Errorf("uuid:%v:%d string is not correct",uuid2,l)
	}
}

func TestNewGenMd5(t *testing.T) {
	u1 :="24f4f1d8-e394-11e8-98e0-f0def1a82164"
	//u1 :="a8558cff-e3b1-4829-937d-77616766668e"
	//u1 :="24F4F1D8-E394-11E8-98E0-F0DEF1A82164"
	//u1 :="1B4E28BA-2FA1-11D2-883F-0016D3CCA427"
	s1 :="golang.org"
	//s1 :="tip.golang.org"
	//s1 :="github.org"
	mUuid := NewGenMd5(u1, s1)
	defer mUuid.Close()
	if mUuid.IsNull() {
    t.Error("Error to initialize NewGenMd5")
  }
	log.Printf("NewGenMd5: %s\n", mUuid.UnparseLower())
	log.Printf("NewGenMd5: %s\n", mUuid.UnparseUpper())
	uuid2=mUuid.Val()
	l := len(uuid2)
	if l != 36 {
		t.Errorf("NewGenMd5:%v:%d string is not correct",uuid2,l)
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

func TestNewGenSha1(t *testing.T) {
	//u1 :="24F4F1D8-E394-11E8-98E0-F0DEF1A82164"
	u1 :="1B4E28BA-2FA1-11D2-883F-0016D3CCA427"
	//s1 :="golang.org"
	s1 :="github.org"
	//oUuid = NewGen()
	sUuid := NewGenSha1(u1, s1)
	defer sUuid.Close()
	if sUuid.IsNull() {
    t.Error("Error to initialize NewGenSha1")
  }
	log.Printf("NewGenSha1:%s \n", sUuid.UnparseLower())
	log.Printf("NewGenSha1:%s \n", sUuid.UnparseUpper())
	uuid2=sUuid.Val()
	l := len(sUuid.Val())
	if l != 36 {
		t.Errorf("NewGenSha1:%v:%d string is not correct",uuid2,l)
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
	if res != 0 {
		t.Error("UUIDCompare 1 is not correct")
	}else{
		log.Printf("compare1 Ok: %d\n", res)
	}
	res = UUIDCompare(uuid1, uuid2)
	if res == 0 {
		t.Error("UUIDCompare 2 is not correct")
	}else{
		log.Printf("compare2 Ok: %d\n",res)
	}
}

const (
	newgen = iota
	newgenrandom
	newgentime
	newgentimesafe
	newgensha1
	newgenmd5
  lower
  upper
	md5l
	md5u
	sha1l
	sha1u
	gid
	guuid4
	guuid5
)

func benchmarkGetUUID(b *testing.B, mode int){
	b.ReportAllocs()
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
		switch mode {
      case newgen:
				oUuid = NewGen()
				uuid1 = oUuid.UnparseUpper()
      case newgenrandom:
				oUuid = NewGenRandom()
				uuid1 = oUuid.UnparseUpper()
      case newgentime:
				oUuid = NewGenTime()
				uuid1 = oUuid.UnparseUpper()
      case newgentimesafe:
				oUuid = NewGenTimeSafe()
				uuid1 = oUuid.UnparseUpper()
      case newgensha1:
				oUuid = NewGenSha1("1b4e28ba-2fa1-11d2-883f-0016d3cca427", "golang.org")
				uuid1 = oUuid.UnparseUpper()
      case newgenmd5:
				oUuid = NewGenMd5("c0086038-e39e-11e8-98e0-f0def1a82164", "github.org")
				uuid1 = oUuid.UnparseUpper()
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
      case guuid4:
				gouuid.NewV4()
				//uuid1 = v4guui.String()
      case guuid5:
				gouuid.NewV5(gouuid.NamespaceDNS, "www.example.com")
			}
	}
	//log.Printf("mode %d, %s", mode, uuid1)
}

func BenchmarkNewGen(b *testing.B){ benchmarkGetUUID(b, newgen)}
func BenchmarkNewGenRandom(b *testing.B){ benchmarkGetUUID(b, newgenrandom)}
func BenchmarkNewGenTime(b *testing.B){ benchmarkGetUUID(b, newgentime)}
func BenchmarkNewGenTimeSafe(b *testing.B){ benchmarkGetUUID(b, newgentimesafe)}
func BenchmarkNewGenSha1(b *testing.B){ benchmarkGetUUID(b, newgensha1)}
func BenchmarkNewGenMd5(b *testing.B){ benchmarkGetUUID(b, newgenmd5)}
func BenchmarkGetUUIDUpper(b *testing.B){ benchmarkGetUUID(b, upper)}
func BenchmarkGetUUIDLower(b *testing.B){ benchmarkGetUUID(b, lower)}
func BenchmarkGetUUIDMd5Upper(b *testing.B){ benchmarkGetUUID(b, md5u)}
func BenchmarkGetUUIDMd5Lower(b *testing.B){ benchmarkGetUUID(b, md5l)}
func BenchmarkGetUUIDSha1Upper(b *testing.B){ benchmarkGetUUID(b, sha1u)}
func BenchmarkGetUUIDSha1Lower(b *testing.B){ benchmarkGetUUID(b, sha1l)}
func BenchmarkGetGOID     (b *testing.B){ benchmarkGetUUID(b, gid)}
func BenchmarkGetGOUUID4     (b *testing.B){ benchmarkGetUUID(b, guuid4)}
func BenchmarkGetGOUUID5     (b *testing.B){ benchmarkGetUUID(b, guuid5)}
