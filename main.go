package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func openBrowserWithURLs(browser, profile string, urls []string, batchSize int) {
	for i := 0; i < len(urls); i += batchSize {
		end := i + batchSize
		if end > len(urls) {
			end = len(urls)
		}

		batch := urls[i:end]

		var args []string
		switch browser {
		case "firefox":
			args = append([]string{"-P", profile}, batch...)
			fmt.Printf("%s.exe -P %s %s\n", browser, profile, strings.Join(batch, " "))
			cmd := exec.Command(browser+".exe", args...)
			cmd.Env = append(os.Environ(), "MOZ_DISABLE_NONLOCAL_CONNECTIONS=1")
			if err := cmd.Run(); err != nil {
				log.Fatalf("Error opening Firefox: %v", err)
			}
		case "chrome":
			args = append([]string{"--ignore-certificate-errors"}, batch...)
			fmt.Printf("%s.exe --ignore-certificate-errors %s\n", browser, strings.Join(batch, " "))
			if err := exec.Command(browser+".exe", args...).Run(); err != nil {
				log.Fatalf("Error opening Chrome: %v", err)
			}
		default:
			log.Fatalf("Unsupported browser: %s", browser)
		}

		fmt.Println("Press enter for new urls")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}

func main() {
	filepath := flag.String("file", "webapps.txt", "Path to the file containing URLs")
	browser := flag.String("browser", "firefox", "Browser to use (chrome or firefox)")
	profile := flag.String("profile", "Insecure", "Firefox profile to use")
	batchSize := flag.Int("batch", 50, "Number of URLs to open at once")
	flag.Parse()

	if *browser != "chrome" && *browser != "firefox" {
		log.Fatalf("Unsupported browser. Supported browsers are: chrome, firefox")
	}

	file, err := os.Open(*filepath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			urls = append(urls, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	openBrowserWithURLs(*browser, *profile, urls, *batchSize)
}
