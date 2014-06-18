package main

import (
	"math/rand"
	"testing"
	"time"
)
import "fmt"

func TestSkill(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	guy := randPerson()
	fmt.Println(guy.getWeaponSkill())
}
