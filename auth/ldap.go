package auth

import (
	"gopkg.in/ldap.v3"
	"log"
	"os"
)

func LdapValidate(user string, password string) bool {
	host := os.Getenv("LDAP_HOST")
	port := os.Getenv("LDAP_PORT")
	domain := os.Getenv("LDAP_DOMAIN")

	// TLS, for testing purposes disable certificate verification,
	// check https://golang.org/pkg/crypto/tls/#Config for further information.

	//tlsConfig := &tls.Config{InsecureSkipVerify: true}
	//l, err := ldap.DialTLS("tcp", "ldap.example.com:636", tlsConfig)

	// No TLS, not recommended
	l, err := ldap.Dial("tcp", host+":"+port)

	if err != nil {
		log.Println(err)
		return false
	} else {
		//Now you should have an active connection to your LDAP server.
		// Using this connection you have to execute a bind:
		//err = l.Bind("user@test.com", "password")
		err = l.Bind(user+"@"+domain, password)
		if err != nil {
			// error in ldap bind
			log.Println(err)
			return false
		}
		// successful bind
		return true
	}
}
