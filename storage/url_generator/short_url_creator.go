package url_generator

import (
	"crypto/md5"
	"encoding/binary"
)

/*
 * Create Short Urls
 */

func createID(longUrl string) uint64 {
	hmd5 := md5.Sum([]byte(longUrl))
	id := binary.LittleEndian.Uint64(hmd5[:])
	return id
}
