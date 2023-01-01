package uuid

import (
  "crypto/rand"
  "fmt"
)

type UUID [16]byte // UUIDs are 128 bit (16 bytes)

func GenUUID() (uuid UUID, uuidStr string, err error) {
  _, err = rand.Read(uuid[:])
  if err != nil {
    return [16]byte{}, "", err
  }
  
  uuid[6] = (uuid[6] & 0x0f) | 0x40
  uuid[8] = (uuid[8] & 0x3f) | 0x80

  return [16]byte{}, fmt.Sprintf("%x-%x-%x-%x-%x", uuid[:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16]), nil
}
