package main

import (
	"encoding/pem"
	"fmt"
)

//var privPEM = `-----BEGIN RSA PRIVATE KEY-----
//MIICXQIBAAKBgQCqmc6dxR+tsJsGPF00Wr3qmV/3ApwE1DEarZ9BSiTYTQVUDy+W
//uZ+niQFi6C6V5RceEKAP1h7/Z1LTWPkaJ8wyJzFip0hD9iEei4SI6NPQZLDPmPYY
//sXKBVUFpnfqQ8rIH+kv2d2OG2BXDp16LGYjHTdr8JJEm1aQ+AjWZh9pzbQIDAQAB
//AoGANS/eirF6PtxgeIE5Tak8rHdEw+28VoURChA3JlPHSOg6UQqq+4LDk6fTFtLs
//My9JFcZ5IHbREy9TUzDZ+J2Pu1BnjUi+G2m54sdJLglc5sD2zqpJmfxGSFGEBKqt
//If2G22Fizo2UD8pfmK7tPvy1pYpHt2J7fmeR1uQzYq4zXMECQQDiAZnrAjJITLmu
//kq9IWmwfRfiiPYwGYaidcrHzrfCwYTItFX9UMhyLRZm7zWy1Kpos6KZnQUUOYP7D
//enaQSBP3AkEAwT3YkaGht/y7c4xYfg6KL8wSfVVDbFC9J+qBWrCzxYM98VHuNjR0
//iyq84hIMBkoZ4m8072S4PiNkQZv6QHeSuwJBAKG2PZjPSIU9CPtlj6/4qzaxTVdh
//LIkAZbLK95OBmR/LXCiwIhxvgscQdRDQywDSS+DoUvC83hmMw53BSYaxXD8CQAl/
//VaaKsB0P0dKzAiJn6ojA2ePJDgBD05gjoWnop108vw2ePjvxxgyU9CWUR30DpVQI
//rSxa4edD7AiBdwI2HkMCQQDKb+XNhoLhK8/RMylbDIs6V9Ushmic1OYBV2uNqYJF
//yJoQR1+12jndEMAa8y0uVFxxCxBlw1a/CYVsBIwYpuhu
//-----END RSA PRIVATE KEY-----`




func main() {
	block, _ := pem.Decode([]byte(privPEM))

	fmt.Printf("type: %s\n", block.Type)
	fmt.Printf("key: \n%x\n", block.Bytes)
}
