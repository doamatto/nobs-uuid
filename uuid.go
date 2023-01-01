package uuid

import (
  "crypto/rand"
)

type UUID [16]byte // UUIDs are 128 bit (16 bytes)

func genUUID() (uuid UUID, err error) {
  _, err = rand.Read(uuid[:])
  if err != nil {
    return
  }
  
  uuid[6] = (uuid[6] & 0x0f) | 0x40
  return
}
