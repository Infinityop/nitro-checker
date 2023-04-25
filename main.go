package main

import (

	"fmt"	"io/ioutil"

	"net/http"

	"os"

	"strings"

	"sync"

	"time"

)

func checkToken(token string, wg *sync.WaitGroup, successFile *os.File, invalidFile *os.File) {

	defer wg.Done()

	headers := http.Header{}

	headers.Set("Authorization", token)

	resp, err := http.Get("https://discord.com/api/v8/users/@me")

	if err != nil {

		fmt.Fprintf(invalidFile, "%s\n", token)

		return

	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		fmt.Fprintf(invalidFile, "%s\n", token)

		return

	}

	if resp.StatusCode == http.StatusOK {

		if strings.Contains(string(body), `"premium_type":1`) {

			fmt.Fprintf(successFile, "%s\n", token)

			fmt.Printf("\033[32m[Success] Nitro Classic on: %s\033[0m\n", token)

		} else if strings.Contains(string(body), `"premium_type":2`) {

			fmt.Fprintf(successFile, "%s\n", token)

			fmt.Printf("\033[32m[Success] Nitro Boost on: %s\033[0m\n", token)

		} else {

			fmt.Fprintf(invalidFile, "%s\n", token)

			fmt.Printf("\033[31m[No Nitro] %s\033[0m\n", token)

		}

	} else {

		fmt.Fprintf(invalidFile, "%s\n", token)

		fmt.Printf("\033[33m[Failed] To get data: %s\033[0m\n", token)

	}

}

func main() {

	var wg sync.WaitGroup

	successFile, err := os.Create("nitro.txt")

	if err != nil {

		panic(err)

	}

	defer successFile.Close()

	invalidFile, err := os.Create("invalid.txt")

	if err != nil {

		panic(err)

	}

	defer invalidFile.Close()

	numThreads := 10

	fmt.Print("Enter number of threads to use (1-100): ")

	fmt.Scanln(&numThreads)

	if numThreads < 1 || numThreads > 100 {

		fmt.Println("Invalid number of threads. Using default value of 10.")

		numThreads = 10

	}

	tokens, err := ioutil.ReadFile("tokens.txt")

	if err != nil {

		panic(err)

	}

	tokenList := strings.Split(string(tokens), "\n")

	startTime := time.Now()

	for _, token := range tokenList {

		if token == "" {

			continue

		}

		wg.Add(1)

		go checkToken(token, &wg, successFile, invalidFile)

	}

	wg.Wait()

	fmt.Printf("Finished checking %d tokens in %.2f seconds.\n", len(tokenList), time.Since(startTime).Seconds())

}
