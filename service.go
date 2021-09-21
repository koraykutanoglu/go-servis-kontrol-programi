package main

import (
	"fmt"
	"os/exec"
	"time"
	//"syscall"

)

func main() {

	for  {
		time.Sleep(10 * time.Second)
		fmt.Println("10 Saniyede bir tüm servisler sorgulanıyor\n")

		cmd := exec.Command("systemctl", "check", "sshd")
		out, err := cmd.CombinedOutput()
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				fmt.Printf("hata çıktısı: %v\n", exitErr)
				cmd := exec.Command("systemctl", "start","sshd")
				err := cmd.Start()
				if err!=nil {

				}
				fmt.Printf("SSH tekrar aktif edildi\n")

			} else {
				fmt.Printf("failed to run systemctl: %v", err)
				//os.Exit(1) //bu komut programi sonlandiriyor

			}
		}
		fmt.Printf("SSH Status is: %s\n", string(out))

		cmd = exec.Command("systemctl", "check", "docker.service")
		out, err = cmd.CombinedOutput()
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				fmt.Printf("hata çıktısı: %v\n", exitErr)
				cmd := exec.Command("systemctl", "start","sshd")
				err := cmd.Start()
				if err!=nil {

				}

			} else {
				fmt.Printf("failed to run systemctl: %v", err)
				//os.Exit(1) //bu komut programi sonlandiriyor
			}
		}
		fmt.Printf("Docker.Service Status is: %s\n", string(out))

		cmd = exec.Command("systemctl", "check", "cron.service")
		out, err = cmd.CombinedOutput()
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				fmt.Printf("hata çıktısı: %v\n", exitErr)
				cmd := exec.Command("systemctl", "start","cron.service")
				err := cmd.Start()
				if err!=nil {

				}
				fmt.Printf("Cron tekrar aktif edildi\n")

			} else {
				fmt.Printf("failed to run systemctl: %v", err)
				//os.Exit(1) //bu komut programi sonlandiriyor

			}
		}
		fmt.Printf("Cron Status is: %s\n", string(out))
	}
}

