package main

import (
	"log"
	"strings"

	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

// p512Key is a Brainpool P-512 key.
var p512Key = strings.NewReader(`-----BEGIN PGP PRIVATE KEY BLOCK-----
Version: GnuPG v2

lgAAAQYEVkI7HxMJKyQDAwIIAQENBAMEP6LxOxGClZdyOEMZNRklliZ1wRo9ZyXF
fAFtsFkEe4eLhuOJA2q3jkVFyMJyk/7ZaXLY15WLuDxrnUDW2wF09A2JuNlEW1Oc
5kRmYkg8oTGYA2t4Qu0ELwgiad/lDWeI+1qTPD0hp6X+bSFC9h3/+PxzBJv3FGn2
etdbZYv/dO3+BwMCSmmKAxW2dp3rjJAGZkoiQsiykYMvbmPyEF1zBgb+74q60fex
mCYcuX0WbyjVaOkoOfMVHdH51WhW5wCpni2/C8NQ8iHt/BYGBB8CGInxImCCP+DJ
mbysA4lqfl+vyD0T2UcCOmU7rwYZLKIOit02tBlUZXN0IEtleSA8dGVzdEB0ZXN0
LnRlc3Q+iLkEExMKACEFAlZCOx8CGwMFCwkIBwIGFQgJCgsCBBYCAwECHgECF4AA
CgkQF0L6wsR8pVIsJgH9EXzEyfRMtUYWKibi00e8wAkyKS9/0Zxo5ExAVWv9QR4k
718z5QkMJgBxq33MeeSK+3gPdIqa3WflWNofTXxhTwH7Bbhkqlztqeifqe5/zBJ8
bZNJDSHlfedUz4RURWn8r5nc0oqlzQi8ISfo6kcbuvBDCn5rGeJjUkxkuDhzDzp1
9Z4AAAEKBFZCOx8SCSskAwMCCAEBDQQDBFVlAC1pwM1XHUqPq9GaG61JAfK2jsLM
Bmn3W7+k4uE34jkXrvTA8uaoX4IrMiKRl1Agqib9SaeaGENr3nWByxoC4OrXgrGR
ciFIcMFNeQcbUQmc/X5aKfZ9S2QVB3LCYJZb5fV59v2kAWq2tUAYEkevNd033Pc4
bKztotwTW0TJAwEKCf4HAwLmy/d3wx04huua/i2zcfQmb6d0gLctQRHKQQzyzWZ4
kju1qBE9e4HLXMlA2sWi/XkVzjW2zlhOhYETSSU/P4/VBccyJ4+wKwymdKkXKQRn
sR1loAgugPrzMWesof0AAfspkQmbI9qYBASdp8cXB8aIoQQYEwoACQUCVkI7HwIb
DAAKCRAXQvrCxHylUsdpAf9v/1yGYO6jfa+CNE9w9oZ7kRIO3Sq+6QrBgYTFZcR3
Nfs+EXgOUmZhMjEhOoHJwklW9rzn/pWvChjBH6I4lrY3AfwJUuXKxiJH7UfBWPjP
gSTjbqpwq4lPbT5VtGz+FRajwn9BlCOmnS2e4d7NOKOxpUcYu7GYK5+0W7bdiNbg
CxZM
=ECaI
-----END PGP PRIVATE KEY BLOCK-----
`)

func main() {
	b, err := armor.Decode(p512Key)
	if err != nil {
		log.Fatal(err)
	}
	// p, err := ioutil.ReadAll(b.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for i := range p {
	// 	fmt.Printf("%#02x, ", p[i])
	// }

	p, err := packet.Read(b.Body)
	if err != nil {
		log.Fatal(err)
	}
	if p, ok := p.(*packet.PrivateKey); ok {
		log.Println(p.PublicKey.PubKeyAlgo)
	}
}
