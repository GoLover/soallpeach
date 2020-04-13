package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)
func TestPrimeGenerator_IsPrime(t *testing.T) {
	testTable := []struct {
		Number uint32
		Prime bool
	}{{
		Number: 10,
		Prime: false,
	},
		{
			Number: 11,
			Prime: true,
		},
		{
			Number: 3,
			Prime: true,
		},
		{
			Number: 801403,
			Prime: true,
		},}
	primeGenerator := NewPrimeGenerator()
	for _,testData := range testTable {
		if testData.Prime != primeGenerator.IsPrime(testData.Number) {
			t.Errorf("It seems you calculate wrong, check %d", testData.Number)
		}
	}
}
func BenchmarkPrimeGenerator_IsPrime(b *testing.B) {
	b.ReportAllocs()
	filePath := `numbers.txt`
	primeGenerator := NewPrimeGenerator()
	for i := 0; i < b.N; i++ {
		file, _ := os.Open(filePath)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			num, _ := strconv.Atoi(scanner.Text())
			if primeGenerator.IsPrime(uint32(num)) {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		}
	}
}