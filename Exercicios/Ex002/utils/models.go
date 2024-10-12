package utils

type Author struct {
	name string
	nacionality string
}

type Publisher struct{
	name string
	nacionality string
}

type Book struct{
	name string
	year uint32
	author Author
	publisher Publisher
	gender string
	value float32
}

