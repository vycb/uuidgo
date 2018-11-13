// Packge uuid implements a wrapper around libuuid for golang by cgo
// The libuuid library generates and parses 128-bit Universally Unique IDs
// (UUIDs). See RFC 4122 for more information.
//
// Usage:
// import "github.com/vycb/uuid"
// oUuid = NewGen()
// defer oUuid.Close()
// if oUuid.IsNull() {
//   t.Error("Error to initialize NewGen")
// }
// log.Printf("NewGen UnparseUpper:%s\n", oUuid.UnparseUpper()) 
// library install for ubuntu, debian
// 	$sudo apt install uuid-dev
// 	
// Copyright 2018 Way out enterprises. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in Wiki.
package uuid
/*
#cgo CFLAGS: -I/usr/include/uuid -pipe -Wall -O3 -Wint-conversion -fomit-frame-pointer -march=native -fopenmp -D_FILE_OFFSET_BITS=64
#cgo LDFLAGS: -luuid
#include <uuid.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>

////#define USEDEVREND 1

//unsigned char uuid[16];
//unsigned char uuid2[16];
uuid_t uuid;
uuid_t uuid2;

static inline void getuuidlower(char* uuid_str) {
//{{{
#ifdef ifdefUSEDEVREND
	uuid_generate(uuid);
#else
	uuid_generate_time_safe(uuid);
#endif
	uuid_unparse_lower(uuid, uuid_str);
} //}}}

static inline void getuuidupper(char* uuid_str) {
//{{{
#ifdef USEDEVREND
	uuid_generate(uuid);
#else
	uuid_generate_time_safe(uuid);
#endif
	uuid_unparse_upper(uuid, uuid_str);
} //}}}

static inline void getuuidmd5upper(char* uuid_str, char *name) {
//{{{
	uuid_t uuidl;
	uuid_t uuidl2;
	uuid_parse(uuid_str, uuidl);
	uuid_generate_md5(uuidl2, uuidl, name, strlen(name));
	uuid_unparse_upper(uuidl2, uuid_str);
	//printf("getuuidmd5upper: %s\n", uuid_str);
} //}}}

static inline void getuuidmd5lower(char* uuid_str, char *name) {
//{{{
	uuid_t uuidl;
	uuid_t uuidl2;
	uuid_parse(uuid_str, uuidl);
	uuid_generate_md5(uuidl2, uuidl, name, strlen(name));
	uuid_unparse_lower(uuidl2, uuid_str);
	//printf("getuuidmd5lower: %s\n", uuid_str);
} //}}}

static inline void genuuidmd5(uuid_t uuida, char* uuid_str, char *name) {
//{{{
	uuid_t uuidl;
	uuid_parse(uuid_str, uuidl);
	uuid_generate_md5(uuida, uuidl, name, strlen(name));
	//char uuid_stp[37];
	//uuid_unparse_upper(uuida, uuid_stp);
	//printf("genuuidmd5: %s\n", uuid_stp);
} //}}}

static void genuuidsha1(unsigned char* uuida, char* uuid_str, char *name) {
//{{{
	uuid_t uuidl;
	uuid_parse(uuid_str, uuidl);
	uuid_generate_sha1(uuida, uuidl, name, strlen(name));
	//char uuid_stp[37];
	//uuid_unparse_upper(uuida, uuid_stp);
	//printf("genuuidsha1: %s\n", uuid_stp);
} //}}}

static void getuuidsha1upper(char* uuid_str, char *name) {
//{{{
	uuid_t uuidl;
	uuid_t uuidl2;
	uuid_parse(uuid_str, uuidl);
	uuid_generate_sha1(uuidl2, uuidl, name, strlen(name));
	uuid_unparse_upper(uuidl2, uuid_str);
	//printf("getuuidsha1upper: %s\n", uuid_str);
} //}}}

static inline void getuuidsha1lower(char* uuid_str, char *name) {
//{{{
	uuid_parse(uuid_str, uuid2);
	uuid_generate_sha1(uuid, uuid2, name, strlen(name));
	uuid_unparse_lower(uuid, uuid_str);
	//printf("getuuidsha1lower: %s\n", uuid_str);
} //}}}

static inline int compareuuid(char* uuid_str1, char* uuid_str2){
//{{{
	uuid_parse(uuid_str1, uuid);
	uuid_parse(uuid_str2, uuid2);
	return uuid_compare(uuid, uuid2);
} //}}}

*/
import "C"
import "unsafe"
//import "log"
const (
	// uuid_t u_char[16]
	UTY = 16
	// UUID_STR_LEN
	USL = 37
)
var uuid_byte [USL]byte

type UUID struct{
	//uuid unsafe.Pointer
	uuidb []byte
	val string //*C.char
}

// use: defer oUuid.Close()
//
func (u *UUID) Close() {
	//C.free(unsafe.Pointer(u.uuid))
	//var s []byte
	//u.uuidb = s
}

// uuid_is_null(uuid_t uuid)
// clear.c
func (u *UUID) Clear() {
	up := C.CBytes(u.uuidb)
  defer C.free(unsafe.Pointer(up))
	C.uuid_clear((*C.u_char)(up))
	ubyt := C.GoBytes(up, C.int(UTY))
	u.uuidb=ubyt
	//u.Close()
}

// uuid_is_null(uuid_t uuid)
// isnull.c
func (u *UUID) IsNull() int{
	up := C.CBytes(u.uuidb)
  defer C.free(unsafe.Pointer(up))
	res := int(C.uuid_is_null((*C.u_char)(up)))
	//log.Printf("IsNull:uuidb: len:%v, cap:%v", len(u.uuidb), cap(u.uuidb))

	//if len(u.uuidb) == 0{
		//return true
	//}
	return res
}

// get: Val()
//
func (u *UUID) Val() string{
	if len(u.val) == 0{
		u.val = u.UnparseUpper()
	}
	return u.val
}

// NewGenMd5
// Generate an MD5 hashed (predictable) UUID based on a well-known UUID
// providing the namespace and an arbitrary binary string.
//
func NewGenMd5(uuid, sname string) *UUID{
	//var bup [16]byte
	//bup := make([]byte, 16)
	up := C.CBytes(uuid_byte[:UTY])
	defer C.free(unsafe.Pointer(up))

  uuid_str := C.CString(uuid)
	defer C.free(unsafe.Pointer(uuid_str))
	ustrn := C.CString(sname)
	defer C.free(unsafe.Pointer(ustrn))

	C.genuuidmd5((*C.u_char)(up), uuid_str, ustrn)
	ubyt := C.GoBytes(up, C.int(UTY))

	return &UUID{uuidb: ubyt, val: ""}
}

// NewGenSha1
// Generate a SHA1 hashed (predictable) UUID based on a well-known UUID
// providing the namespace and an arbitrary binary string.
//
func NewGenSha1(uuid, sname string) *UUID{
	//var bup [16]byte
	//bup := make([]byte, 16, 16)
	//var up unsafe.Pointer
	up := C.CBytes(uuid_byte[:UTY])
	defer C.free(unsafe.Pointer(up))

	var uuid_str *C.char
	uuid_str = C.CString(uuid)
	defer C.free(unsafe.Pointer(uuid_str))

	//var ustrn *C.char
	ustrn := C.CString(sname)
	defer C.free(unsafe.Pointer(ustrn))

	//puch := (unsafe.Pointer)(C.genuuidsha1((*C.u_char)(up), uuid_str, ustrn))
	//ubyt := C.GoBytes(puch, C.int(16))
	C.genuuidsha1((*C.u_char)(up), uuid_str, ustrn)
	ubyt := C.GoBytes(up, C.int(UTY))

	//C.uuid_generate_sha1((*C.u_char)(up), (*C.u_char)(tp), ustrn, (C.ulong)(len(sname)))
	return &UUID{uuidb: ubyt, val: ""}
}

// uuid_generate_time(uuid_t out)
// gen_uuid.c
func NewGenTime() *UUID{
	//var bup [16]byte
  up := C.CBytes(uuid_byte[:UTY])
	defer C.free(unsafe.Pointer(up))

	C.uuid_generate_time((*C.u_char)(up))
	ubyt := C.GoBytes(up, C.int(UTY))

	return &UUID{uuidb: ubyt, val: ""}
}

// uuid_generate_time_safe(uuid_t out)
// gen_uuid.c
func NewGenTimeSafe() *UUID{
	//var bup [16]byte
  up := C.CBytes(uuid_byte[:UTY])
	defer C.free(unsafe.Pointer(up))

	C.uuid_generate_time_safe((*C.u_char)(up))
	ubyt := C.GoBytes(up, C.int(UTY))

	return &UUID{uuidb: ubyt, val: ""}
}

// uuid_generate(uuid_t out)
// gen_uuid.c
func NewGen() *UUID{
	//var bup [16]byte
  up := C.CBytes(uuid_byte[:UTY])
	defer C.free(unsafe.Pointer(up))

	C.uuid_generate((*C.u_char)(up))
	ubyt := C.GoBytes(up, C.int(UTY))
	//log.Printf("NewGen:ubyt: len:%v, cap:%v", len(ubyt), cap(ubyt))

	return &UUID{uuidb: ubyt, val: ""}
}

// uuid_generate_random(uuid_t out)
// gen_uuid.c
func NewGenRandom() *UUID{
	//var bup [16]byte
  up := C.CBytes(uuid_byte[:UTY])
	defer C.free(unsafe.Pointer(up))

	//defer C.free(unsafe.Pointer(up))
	C.uuid_generate_random((*C.u_char)(up))
	ubyt := C.GoBytes(up, C.int(UTY))
	//log.Printf("NewGenRandom:ubyt: len:%v, cap:%v", len(ubyt), cap(ubyt))

	return &UUID{uuidb: ubyt, val: ""}
}

// uuid_compare(uuid_t uuid1, uuid_t uuid2)
// compare.c, parse.c
func (u *UUID) Compare(uid string) int{
	uuid_str := C.CString(uid)
	defer C.free(unsafe.Pointer(uuid_str))
	//var bup [16]byte
  nup := C.CBytes(uuid_byte[:UTY])
	defer C.free(unsafe.Pointer(nup))
	C.uuid_parse(uuid_str, (*C.u_char)(nup))

  up := C.CBytes(u.uuidb)
	defer C.free(unsafe.Pointer(up))
	return int(C.uuid_compare((*C.u_char)(up), (*C.u_char)(nup)))
}

// uuid_unparse_lower(uuid_t out)
// gen_uuid.c
func (u *UUID) UnparseLower() string{
	//var bup [37]byte
	bs := string(uuid_byte[:USL])
	uuid_str := C.CString(bs)
	defer C.free(unsafe.Pointer(uuid_str))

	up := C.CBytes(u.uuidb)
  defer C.free(unsafe.Pointer(up))
	C.uuid_unparse_lower((*C.u_char)(up), uuid_str)
	u.val = C.GoString(uuid_str)
	return u.val
}

// uuid_unparse_upper(uuid_t out)
// gen_uuid.c
func (u *UUID) UnparseUpper() string{
	bs := string(uuid_byte[:37])
	uuid_str := C.CString(bs)
	defer C.free(unsafe.Pointer(uuid_str))
	//up := C.CBytes(uuid_byte[:16])
	//up := string(uuid_byte[:16])
	//uuid := C.CString(up)
	//defer C.free(unsafe.Pointer(up))
	//C.uuid_generate_time_safe((*C.u_char)(up))
	//if(len(u.uuidb)){
	//}
	up := C.CBytes(u.uuidb)
	defer C.free(unsafe.Pointer(up))

	C.uuid_unparse_upper((*C.u_char)(up), uuid_str)
	u.val = C.GoString(uuid_str)
	return u.val
}

// util: GetUUIDUpper
// uuid_generate_time_safe
func GetUUIDUpper() string{
	//var bup [37]byte
	bs := string(uuid_byte[:USL])
	uuid_str := C.CString(bs)
	defer C.free(unsafe.Pointer(uuid_str))
	C.getuuidupper(uuid_str)
	return C.GoString(uuid_str)
}

// util: GetUUIDLower
// uuid_generate_time_safe
func GetUUIDLower() string{
	//var bup [37]byte
	bs := string(uuid_byte[:USL])
	uuid_str := C.CString(bs)
	defer C.free(unsafe.Pointer(uuid_str))
	C.getuuidlower(uuid_str)
	//cchar := (*C.char)(&(uuid_byte))
	return C.GoString(uuid_str)
}

// util: UUIDCompare u1, u2 string
//
func UUIDCompare(uuid_str1, uuid_str2 string) int{
	ustr1 := C.CString(uuid_str1)
  defer C.free(unsafe.Pointer(ustr1))
	ustr2 := C.CString(uuid_str2)
  defer C.free(unsafe.Pointer(ustr2))
	return int(C.compareuuid(ustr1, ustr2))
}

// util: GetUUIDMd5Upper
// Generate an MD5 hashed (predictable) UUID based on a well-known UUID
// providing the namespace and an arbitrary binary string.
//
func GetUUIDMd5Upper(uuid, name string) string {
  uuid_str := C.CString(uuid)
	defer C.free(unsafe.Pointer(uuid_str))
	ustrn := C.CString(name)
  defer C.free(unsafe.Pointer(ustrn))
	C.getuuidmd5upper(uuid_str, ustrn)
	return C.GoString(uuid_str)
}

// util: GetUUIDMd5Lower
// Generate an MD5 hashed (predictable) UUID based on a well-known UUID
// providing the namespace and an arbitrary binary string.
//
func GetUUIDMd5Lower(uuid, name string) string {
  uuid_str := C.CString(uuid)
	defer C.free(unsafe.Pointer(uuid_str))
	ustrn := C.CString(name)
  defer C.free(unsafe.Pointer(ustrn))
	C.getuuidmd5lower(uuid_str, ustrn)
	return C.GoString(uuid_str)
}

// util: GetUUIDSha1Lower
// Generate a SHA1 hashed (predictable) UUID based on a well-known UUID
// providing the namespace and an arbitrary binary string.
//
func GetUUIDSha1Lower(uuid, name string) string {
  uuid_str := C.CString(uuid)
	defer C.free(unsafe.Pointer(uuid_str))
	ustrn := C.CString(name)
  defer C.free(unsafe.Pointer(ustrn))
	C.getuuidsha1lower(uuid_str, ustrn)
	return C.GoString(uuid_str)
}

// util: GetUUIDSha1Upper
// Generate a SHA1 hashed (predictable) UUID based on a well-known UUID
// providing the namespace and an arbitrary binary string.
//
func GetUUIDSha1Upper(uuid, name string) string {
  uuid_str := C.CString(uuid)
	defer C.free(unsafe.Pointer(uuid_str))
	ustrn := C.CString(name)
  defer C.free(unsafe.Pointer(ustrn))
	C.getuuidsha1upper(uuid_str, ustrn)
	return C.GoString(uuid_str)
}
