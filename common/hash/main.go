package main

import (
	"fmt"
	"hash/fnv"
)

func main() {

	// Tenant, Seq, CIDR
	fmt.Printf("%d", GetHash("T.aweoivhaewoivhaweoi"))

}

func GetHash(key string) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return hash.Sum32()
}
