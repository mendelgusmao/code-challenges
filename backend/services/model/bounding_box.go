package model

type BoundingBox struct {
	MinLon float64 `yaml:"min_lon"`
	MinLat float64 `yaml:"min_lat"`
	MaxLon float64 `yaml:"max_lon"`
	MaxLat float64 `yaml:"max_lat"`
}
