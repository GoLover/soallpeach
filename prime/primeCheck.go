package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type PrimeGenerator struct {
	primeChan  chan uint32
	primes     map[uint32]bool
	lastNumber uint32
}

func NewPrimeGenerator() *PrimeGenerator {
	primeGenerator := PrimeGenerator{
		primeChan:  make(chan uint32),
		primes:     make(map[uint32]bool),
		lastNumber: 1,
	}

	go primeGenerator.start()

	return &primeGenerator
}

func (p *PrimeGenerator) start() {
	multiples := make(map[uint32]uint32)
	wheelCycle := []uint32{10, 2, 4, 2, 4, 6, 2, 6, 4, 2, 4, 6, 6, 2, 6, 4, 2, 6, 4, 6, 8, 4, 2, 4, 2, 4, 8, 6, 4, 6, 2, 4, 6, 2, 6, 6, 4, 2, 4, 6, 2, 6, 4, 2, 4, 2, 10, 2}
	wheelCycleIndices := []uint32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 0, 0, 3, 0, 4, 0, 0, 0, 5, 0, 0, 0, 0, 0, 6, 0, 7, 0, 0, 0, 0, 0, 8, 0, 0, 0, 9, 0, 10, 0, 0, 0, 11, 0, 0, 0, 0, 0, 12, 0, 0, 0, 0, 0, 13, 0, 14, 0, 0, 0, 0, 0, 15, 0, 0, 0, 16, 0, 17, 0, 0, 0, 0, 0, 18, 0, 0, 0, 19, 0, 0, 0, 0, 0, 20, 0, 0, 0, 0, 0, 0, 0, 21, 0, 0, 0, 22, 0, 23, 0, 0, 0, 24, 0, 25, 0, 0, 0, 26, 0, 0, 0, 0, 0, 0, 0, 27, 0, 0, 0, 0, 0, 28, 0, 0, 0, 29, 0, 0, 0, 0, 0, 30, 0, 31, 0, 0, 0, 32, 0, 0, 0, 0, 0, 33, 0, 34, 0, 0, 0, 0, 0, 35, 0, 0, 0, 0, 0, 36, 0, 0, 0, 37, 0, 38, 0, 0, 0, 39, 0, 0, 0, 0, 0, 40, 0, 41, 0, 0, 0, 0, 0, 42, 0, 0, 0, 43, 0, 44, 0, 0, 0, 45, 0, 46, 0, 0, 0, 0, 0, 0, 0, 0, 0, 47}
	p.primeChan <- 2
	p.primeChan <- 3
	p.primeChan <- 5
	p.primeChan <- 7

	num := uint32(1)
	for i := uint32(0); ; i = (i + 1) % 48 {
		num += wheelCycle[i]

		factor, hasFactor := multiples[num]

		var j uint32
		if hasFactor {
			delete(multiples, num)
			j = wheelCycleIndices[(num/factor)%210]
		} else {
			factor = num
		}

		for newNum := num + factor*wheelCycle[j]; ; newNum += factor * wheelCycle[j] {
			_, hasNewFactor := multiples[newNum]
			if !hasNewFactor {
				multiples[newNum] = factor

				break
			}
			j = (j + 1) % 48
		}

		if !hasFactor {
			p.primeChan <- num
		}
	}
}

func (p *PrimeGenerator) next() uint32 {
	prime := <-p.primeChan
	p.primes[prime] = true
	return prime
}
func (p *PrimeGenerator) IsPrime(number uint32) bool {
	if p.lastNumber < number {
		for i := p.lastNumber; i <= number; i++ {
			p.lastNumber = p.next()
		}
		p.lastNumber = number
	}
	return p.primes[number]
}
func main() {
	primeGenerator := NewPrimeGenerator()
	var filePath string
	_, _ = fmt.Scanf("%s", &filePath)
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
