package util

import (
	"crypto/sha256"
	"encoding/binary"
	"os"
	"os/signal"
)

func OSInterrupt() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit
}

func GenHashIDUint16(key string) uint16 {
	hf := sha256.New()
	hf.Write([]byte(key))
	return binary.LittleEndian.Uint16(hf.Sum(nil))
}
