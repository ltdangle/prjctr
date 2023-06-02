package main

type zoo struct {
	Name         string `json:"zooName"`
	WolfCage     *wolfCage
	FoxCage      *foxCage
	ElephantCage *elephantCage
	ZebraCage    *zebraCage
	PanteraCage  *panteraCage
	Zookeeper    *zookeeper
}
