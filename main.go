package main

import (
	"fmt"
	"github.com/c-bata/go-neural-network/networks"
	"github.com/c-bata/go-neural-network/matrix"
)

func main() {
	var n networks.SimpleNet = *networks.NewSimpleNet()
	fmt.Printf("Initilized Weight: %#v\n", n.Weight)

	x := []float64{0.6, 0.9}
	p := n.Predict(x)
	fmt.Printf("Predict: %#v\n", p)

	fmt.Printf("Argmax %v\n", matrix.ArgMax(p)) // 最大値のインデックス

	t := []float64{0, 0, 1}  // 正解ラベル
	fmt.Printf("Loss %v\n", n.Loss(x, t))
}
