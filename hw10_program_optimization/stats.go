package hw10programoptimization

import (
	"bufio"
	"io"
	"log"
	"strings"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u := getUsers(r)
	return countDomains(u, domain)
}

type usersChan struct {
	userRecord User
	err        error
}

func getUsers(r io.Reader) chan usersChan {
	scanner := bufio.NewScanner(r)
	c := make(chan usersChan)
	var user User
	go func() {
		defer close(c)
		for {
			scanner.Scan()
			if len(scanner.Bytes()) == 0 {
				break
			}
			e := user.UnmarshalJSON(scanner.Bytes())
			c <- usersChan{
				userRecord: user,
				err:        e,
			}
		}
	}()
	return c
}

func countDomains(u chan usersChan, domain string) (DomainStat, error) {
	result := make(DomainStat)
	lendomain := len(domain)
	loop := 0
	for user := range u {
		loop++
		if user.err != nil {
			log.Printf("error unmarshaling string number=%d, error=%s", loop, user.err)
		}
		lenusername := len(user.userRecord.Email)
		if lendomain >= lenusername {
			continue
		}
		matched := true
		for i := 0; i < lendomain; i++ {
			if domain[lendomain-i-1] != user.userRecord.Email[lenusername-i-1] {
				matched = false
				break
			}
		}
		if matched && user.userRecord.Email[lenusername-lendomain-1] != '.' {
			continue
		}
		if matched {
			domen := strings.ToLower(strings.SplitN(user.userRecord.Email, "@", 2)[1])
			result[domen]++
		}
	}
	return result, nil
}
