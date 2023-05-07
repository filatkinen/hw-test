package hw10programoptimization

import (
	"bufio"
	"io"
	"log"
	"strings"
	"sync"
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
	wg := sync.WaitGroup{}
	const threads = 4
	wg.Add(threads)
	mtx := sync.Mutex{}
	lendomain := len(domain)
	domencount := func() {
		defer wg.Done()
		for user := range u {
			if user.err != nil {
				log.Printf("error unmarshaling string: %s", user.err)
				continue
			}
			if user.userRecord.Email[len(user.userRecord.Email)-lendomain:] == domain &&
				user.userRecord.Email[len(user.userRecord.Email)-lendomain-1] == '.' {
				domen := strings.ToLower(strings.SplitN(user.userRecord.Email, "@", 2)[1])
				mtx.Lock()
				result[domen]++
				mtx.Unlock()
			}
		}
	}
	for i := 0; i < threads; i++ {
		domencount()
	}
	wg.Wait()

	return result, nil
}
