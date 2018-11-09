//
//  libuuid for golang by cgo
//
//  library install for debian
//      $ sudo apt install uuid-dev
//  By default uuid_generate_time_safe is used.
//  Also uuid_generate can be enabled, which is slower, by uncommenting #define USEDEVREND 1
package uuid
/*
#cgo CFLAGS: -I/usr/include/uuid -pipe -Wall -O3 -fomit-frame-pointer -march=native -fopenmp -D_FILE_OFFSET_BITS=64
#cgo LDFLAGS: -luuid
#include <uuid.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>

////#define USEDEVREND 1

uuid_t uuid;
uuid_t uuid2;

static void getuuidlower(char* uuid_str) {
#ifdef USEDEVREND
	uuid_generate(uuid);
#else
	uuid_generate_time_safe(uuid);
#endif
	uuid_unparse_lower(uuid, uuid_str);
}

static void getuuidupper(char* uuid_str) {
#ifdef USEDEVREND
	uuid_generate(uuid);
#else
	uuid_generate_time_safe(uuid);
#endif
	uuid_unparse_upper(uuid, uuid_str);
}

static void getuuidmd5upper(char* uuid_str, char *name) {
	uuid_parse(uuid_str, uuid2);

	uuid_generate_md5(uuid, uuid2, name, strlen(name));

	uuid_unparse_upper(uuid, uuid_str);
}

static void getuuidmd5lower(char* uuid_str, char *name) {
	uuid_parse(uuid_str, uuid2);

	uuid_generate_md5(uuid, uuid2, name, strlen(name));

	uuid_unparse_lower(uuid, uuid_str);
}

static void getuuidsha1upper(char* uuid_str, char *name) {
	uuid_parse(uuid_str, uuid2);

	uuid_generate_sha1(uuid, uuid2, name, strlen(name));

	uuid_unparse_upper(uuid, uuid_str);
}

static void getuuidsha1lower(char* uuid_str, char *name) {
	uuid_parse(uuid_str, uuid2);

	uuid_generate_sha1(uuid, uuid2, name, strlen(name));

	uuid_unparse_lower(uuid, uuid_str);
}

static int compareuuid(char* uuid_str1, char* uuid_str2){
	uuid_parse(uuid_str1, uuid);
	uuid_parse(uuid_str2, uuid2);
	return uuid_compare(uuid, uuid2);
}

*/
import "C"
import "unsafe"
//import "log"
var uuid_byte [37]byte

func GetUUIDUpper() string{
	bs := string(uuid_byte[:37])
	uuid_str := C.CString(bs)
	defer C.free(unsafe.Pointer(uuid_str))
	C.getuuidupper(uuid_str)
	return C.GoString(uuid_str)
}

func GetUUIDLower() string{
	bs := string(uuid_byte[:37])
	uuid_str := C.CString(bs)
	defer C.free(unsafe.Pointer(uuid_str))
	C.getuuidlower(uuid_str)
	//cchar := (*C.char)(&(uuid_byte))
	return C.GoString(uuid_str)
}

func UUIDCompare(uuid_str1, uuid_str2 string) int{
	ustr1 := C.CString(uuid_str1)
  defer C.free(unsafe.Pointer(ustr1))
	ustr2 := C.CString(uuid_str2)
  defer C.free(unsafe.Pointer(ustr2))
	return int(C.compareuuid(ustr1, ustr2))
}

func GetUUIDMd5Upper(uuid, name string) string {
  uuid_str := C.CString(uuid)
	defer C.free(unsafe.Pointer(uuid_str))
	ustrn := C.CString(name)
  defer C.free(unsafe.Pointer(ustrn))
	C.getuuidmd5upper(uuid_str, ustrn)
	return C.GoString(uuid_str)
}

func GetUUIDMd5Lower(uuid, name string) string {
  uuid_str := C.CString(uuid)
	defer C.free(unsafe.Pointer(uuid_str))
	ustrn := C.CString(name)
  defer C.free(unsafe.Pointer(ustrn))
	C.getuuidmd5lower(uuid_str, ustrn)
	return C.GoString(uuid_str)
}

/*
 * Generate a SHA1 hashed (predictable) UUID based on a well-known UUID
 * providing the namespace and an arbitrary binary string.
 */
func GetUUIDSha1Lower(uuid, name string) string {
  uuid_str := C.CString(uuid)
	defer C.free(unsafe.Pointer(uuid_str))
	ustrn := C.CString(name)
  defer C.free(unsafe.Pointer(ustrn))
	C.getuuidsha1lower(uuid_str, ustrn)
	return C.GoString(uuid_str)
}

/*
 * Generate a SHA1 hashed (predictable) UUID based on a well-known UUID
 * providing the namespace and an arbitrary binary string.
 */
func GetUUIDSha1Upper(uuid, name string) string {
  uuid_str := C.CString(uuid)
	defer C.free(unsafe.Pointer(uuid_str))
	ustrn := C.CString(name)
  defer C.free(unsafe.Pointer(ustrn))
	C.getuuidsha1upper(uuid_str, ustrn)
	return C.GoString(uuid_str)
}
