# SecureBootCheck

## Description

A small program to check if Secure Boot is enabled or not. Written using Go.
Primarily because I work in Ubuntu most of the time, and I occassionally play Valorant on Windows, and they have a kernel level anti-cheat that requires Secure Boot to be enabled.
This utility is to check if Secure Boot is enabled or not without calling upon the wrath of the Vanguard gods.

## Usage

Run the executable in admin mode. If you do forget, it will prompt you to run it in admin mode.

## Building

Run `go build -o securebootcheck.exe main.go` in the root directory. It will create an executable in the root directory.