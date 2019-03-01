package main

import (
	"bufio"
	"fmt"
	"hashCode/models"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	images := getImages("b_lovely_landscapes.txt")
	slides := createSlides(images)
	path := solve(slides)
	fmt.Println(len(path))
	for _, slide := range path {
		for _, index := range slide.Indices {
			fmt.Print(index)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func getImages(fileName string) []models.Image {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	images := make([]models.Image, 0)
	index := 0
	for scanner.Scan() {
		data := strings.Fields(scanner.Text())
		tagsAmount, _ := strconv.Atoi(data[1])
		tags := make([]string, 0)
		for i := 2; i < tagsAmount + 2; i++ {
			tags = append(tags, data[i])
		}
		images = append(images, models.Image{Orientation: data[0], Tags: tags, Index: index})
		index += 1
	}

	return images
}

func createSlides(images []models.Image) []models.Slide {
	slides := make([]models.Slide, 0)
	join := false
	var firstVertical models.Image
	for _, image := range images {
		switch image.Orientation {
		case "H":
			slides = append(slides, models.Slide{Tags: image.Tags, Indices: []int{image.Index}})
		case "V":
			if join {
				joinTags := append(firstVertical.Tags, image.Tags...)
				slides = append(slides, models.Slide{Tags: unique(joinTags), Indices: []int{firstVertical.Index, image.Index}})
				join = false
			} else {
				firstVertical = image
				join = true
			}
		}

	}

	return slides
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func min_diff(tags1 []string, tags2[]string) int{
	var common = 0

	for _, t1 := range tags1 {
		for _, t2 := range tags2 {
			if t1 == t2 {
				common = common + 1
			}
		}
	}
	array := []int{common, len(tags1)-common, len(tags2)-common}
	var min_value = min(array)
	return min_value
}

func min(values []int) (min int) {
	if len(values) == 0 {
		return 0
	}

	min = values[0]
	for _, v := range values {
		if (v < min) {
			min = v
		}
	}
	return min
}

func solve(edges []models.Slide) []models.Slide {
	var path []models.Slide

	maxEdgeIndex := 0
	for len(edges) > 0 {
		currentEdge := edges[maxEdgeIndex]
		edges = remove(edges, maxEdgeIndex)
		path = append(path, currentEdge)
		maxInterest := 0
		maxEdgeIndex = 0
		for index2, edge2 := range edges {
			value := min_diff(edge2.Tags, currentEdge.Tags)
			if value > maxInterest {
				maxInterest = value
				maxEdgeIndex = index2
			}
		}
	}

	return path
}

func remove(s []models.Slide, i int) []models.Slide {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
