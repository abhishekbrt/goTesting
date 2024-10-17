package main

import (
	"goTesting/mocking"
	"os"
)

func main(){
	sleeper:=mocking.DefaultSleeper{}
	mocking.CountDown(os.Stdout,&sleeper)
}