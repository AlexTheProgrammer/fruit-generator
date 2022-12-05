package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	idx := r1.Intn(len(fruit))

	fmt.Printf(strings.ToLower(fruit[idx]))
}

var fruit = []string{
	"Acerola",
	"Apple",
	"Apricots",
	"Avocado",
	"Banana",
	"Blackberries",
	"Blackcurrant",
	"Blueberries",
	"Breadfruit",
	"Cantaloupe",
	"Carambola",
	"Cherimoya",
	"Cherries",
	"Clementine",
	"CoconutMeat",
	"Cranberries",
	"DateFruit",
	"Durian",
	"Elderberries",
	"Feijoa",
	"Figs",
	"Gooseberries",
	"Grapefruit",
	"Grapes",
	"Guava",
	"Honeydew",
	"Jackfruit",
	"JujubeFruit",
	"Kiwifruit",
	"Kumquat",
	"Lemon",
	"Lime",
	"Longan",
	"Loquat",
	"Lychee",
	"Mandarin",
	"Mango",
	"Mangosteen",
	"Mulberries",
	"Nectarine",
	"Olives",
	"Orange",
	"Papaya",
	"PassionFruit",
	"Peaches",
	"Pear",
	"Persimmon",
	"Pitaya",
	"Pineapple",
	"Pitanga",
	"Plantain",
	"Plums",
	"Pomegranate",
	"PricklyPear",
	"Prunes",
	"Pummelo",
	"Quince",
	"Raspberries",
	"Rhubarb",
	"Sapodilla",
	"Sapote",
	"Soursop",
	"Strawberries",
	"Tamarind",
	"Tangerine",
	"Watermelon",
}
