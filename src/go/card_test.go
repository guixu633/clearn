package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComparePlayers(t *testing.T) {
	p1 := []string{"3d", "3h", "5d", "Jc", "7c", "Td", "Kh"}
	rank, score, err := Calculate(p1)
	assert.NoError(t, err)
	fmt.Println(rank, score)
}
