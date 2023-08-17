package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
	"strconv"
)

const (
	populationSize = 50
	mutationRate   = 0.2
	generations    = 100
)

func main() {
	rand.Seed(time.Now().UnixNano())
	filenames := []string{"randomGraph5120.txt",
	"randomGraph2560.txt",
	"randomGraph1280.txt",
	"randomGraph640.txt",
	"randomGraph320.txt",
	"randomGraph160.txt",
	"randomGraph80.txt",
	"randomGraph40.txt",
	"randomGraph20.txt",
	"randomGraph10.txt"}
		
	for _, filename := range filenames {
		start := time.Now()
		matrix := readMatrixFromFile(filename)
		population := createInitialPopulation(len(matrix))
		for i := 0; i < generations; i++ {
			newPopulation := make([][]int, 0)
			for j := 0; j < populationSize/2; j++ {
				parent1 := selectParents(population, matrix)
				parent2 := selectParents(population, matrix)
				child1 := crossover(parent1, parent2)
				child2 := crossover(parent2, parent1)
				mutate(child1)
				mutate(child2)
				newPopulation = append(newPopulation, child1, child2)
			}
			population = newPopulation
		}
		bestIndividual := findBestIndividual(population, matrix)
		bestDistance := fitness(bestIndividual, matrix)
		fmt.Printf("Best route: %v\n", bestIndividual)
		fmt.Printf("Distance: %v\n", bestDistance)
		end := time.Now()
		duration := end.Sub(start).Microseconds()
		fmt.Printf("Time taken by Quick Sort: %d microseconds\n", duration)
	}
}

func createInitialPopulation(size int) [][]int {
	population := make([][]int, populationSize)
	for i := range population {
		population[i] = rand.Perm(size)
	}
	return population
}

func readMatrixFromFile(filename string) [][]int {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	var matrix [][]int
	for _, line := range lines {
		if line == "" {
			continue
		}
		numStrs := strings.Split(line, ",")
		var numArray []int
		for _, numStr := range numStrs {
			num, err := strconv.Atoi(strings.TrimSpace(numStr))
			if err != nil {
				return nil
			}
			numArray = append(numArray, num)
		}
		matrix = append(matrix, numArray)
	}
	return matrix
}


func fitness(individual []int, matrix [][]int) int {
	totalDistance := 0
	for i := 0; i < len(individual)-1; i++ {
		totalDistance += matrix[individual[i]][individual[i+1]]
	}
	totalDistance += matrix[individual[len(individual)-1]][individual[0]]
	return totalDistance
}

func selectParents(population [][]int, matrix [][]int) []int {
	parent1Index := rand.Intn(len(population))
	parent2Index := rand.Intn(len(population))
	parent1Fitness := fitness(population[parent1Index], matrix)
	parent2Fitness := fitness(population[parent2Index], matrix)
	if parent1Fitness < parent2Fitness {
		return population[parent1Index]
	}
	return population[parent2Index]
}

func crossover(parent1 []int, parent2 []int) []int {
	crossoverPoint := rand.Intn(len(parent1) - 1) + 1
	child := append([]int{}, parent1[:crossoverPoint]...)
	for _, gene := range parent2 {
		if !contains(child, gene) {
			child = append(child, gene)
		}
	}
	return child
}

func mutate(individual []int) {
	if rand.Float64() < mutationRate {
		idx1 := rand.Intn(len(individual))
		idx2 := rand.Intn(len(individual))
		individual[idx1], individual[idx2] = individual[idx2], individual[idx1]
	}
}

func findBestIndividual(population [][]int, matrix [][]int) []int {
	bestIndividual := population[0]
	bestFitness := fitness(bestIndividual, matrix)
	for _, individual := range population {
		if fitness(individual, matrix) < bestFitness {
			bestIndividual = individual
			bestFitness = fitness(individual, matrix)
		}
	}
	return bestIndividual
}

func parseInt(s string) int {
	i := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			i = i*10 + int(c-'0')
		}
	}
	return i
}

func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
