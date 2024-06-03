package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	p := plot.New()

	p.Title.Text = "Exponential Distribution"
	p.X.Label.Text = "Variable (x)"
	p.Y.Label.Text = "Num of Appearence (P) "

	if mean, err := strconv.ParseFloat(os.Args[1], 32); err == nil {
		lambda := 1 / mean
		var dotsNum = 10000000

		plotutil.AddLinePoints(p,
			fmt.Sprintf("λ: %.2f", 1.0000), getExpDistribution(dotsNum, 1.0000),
			fmt.Sprintf("λ: %.2f", lambda), getExpDistribution(dotsNum, lambda))

		p.Save(4*vg.Inch, 4*vg.Inch, "result.png")
	}
}

func getExpArr(_mean float64, _length int) []float64 {
	var arr []float64

	for i := 0; i < int(_length); i++ {
		arr = append(arr, _mean*rand.ExpFloat64())
	}
	return arr
}

func getExpDistribution(_dotsNum int, _lambda float64) plotter.XYs {
	intervals := 5
	points := make(plotter.XYs, intervals+1)
	_mean := 1 / _lambda

	for i := 0; i <= intervals; i++ {
		points[i].X = float64(i) * _mean
		points[i].Y = float64(0)
	}

	var randArr []float64 = getExpArr(_mean, _dotsNum)

	for i := 0; i < _dotsNum; i++ {
		randNum := randArr[i]
		index := int(randNum / _mean)

		// 	大於 max interval * spacing
		// 	--> 歸類在 interval+1 的 session
		if index >= intervals {
			points[intervals].Y++
			continue
		}
		points[index].Y++
	}

	for i := 0; i <= intervals; i++ {
		points[i].Y = points[i].Y / float64(_dotsNum)
	}

	var x_Mean = getXMean(randArr)
	var x2_Mean = getX2Mean(randArr)

	fmt.Printf("[ lambda: %.4f ]\n", _lambda)
	fmt.Printf("mean: %.4f \n", x_Mean)
	fmt.Printf("x2mean: %.4f \n", x2_Mean)

	var var1 = x2_Mean - (x_Mean * x_Mean)

	if var1 < 0 {
		var1 = -var1
	}
	fmt.Printf("variance (using square delta): %.4f \n", var1)

	var var2 = getDelX2Mean(randArr, x_Mean)
	fmt.Printf("variance (using delta square): %.4f \n\n", var2)

	return points
}

func getXMean(_arr []float64) float64 {
	var sum = 0.0
	for _, n := range _arr {
		sum += n
	}
	return sum / float64(len(_arr))
}

func getX2Mean(_arr []float64) float64 {
	var sum = 0.0
	for _, n := range _arr {
		sum += (n * n)
	}
	return sum / float64(len(_arr))
}

func getDelX2Mean(_arr []float64, mean float64) float64 {
	var sum = 0.0
	for _, n := range _arr {
		sum += (n - mean) * (n - mean)
	}
	return sum / float64(len(_arr))
}
