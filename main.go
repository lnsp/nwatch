package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

const (
	clearCommand = "clear"
)

var (
	n = flag.String("n", "1s", "refresh interval in seconds")
)

func clearScreen() {
	clearCmd := exec.Command(clearCommand)
	clearCmd.Stdout = os.Stdout
	clearCmd.Stderr = os.Stderr
	clearCmd.Run()
}

func runInBackground(cmd string, args ...string) chan []byte {
	cmdChan := make(chan []byte)
	go func() {
		cmd := exec.Command(cmd, args...)
		out, _ := cmd.CombinedOutput()
		cmdChan <- out
	}()
	return cmdChan
}

func runInInterval(interval time.Duration, cmd string, args ...string) {
	t := time.After(interval)
	c := runInBackground(cmd, args...)
	v := <-c
	<-t
	clearScreen()
	os.Stdout.Write(v)
}

func main() {
	flag.Parse()
	a := flag.Args()
	if len(a) < 1 {
		fmt.Println("USAGE: nwatch [-n] [cmd args ...] ")
		return
	}
	t, err := time.ParseDuration(*n)
	if err != nil {
		fmt.Println("invalid interval format")
		return
	}
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-s:
			return
		default:
			runInInterval(t, a[0], a[1:]...)
		}
	}
}
