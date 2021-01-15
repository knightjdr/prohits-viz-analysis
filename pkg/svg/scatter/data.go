package scatter

import (
	"math"
	"strconv"

	customMath "github.com/knightjdr/prohits-viz-analysis/pkg/math"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func formatData(scatter *Scatter, axisLength float64) {
	scatter.Ticks = defineTicks(scatter.Plot, scatter.LogBase)
	scaleData(scatter, axisLength)
}

func defineTicks(plot []types.ScatterPoint, logBase string) Ticks {
	minmax := getAxisMinMax(plot)

	if logBase != "none" {
		return calculateLogTicks(logBase, minmax)
	}

	return calculateLinearTicks(minmax)
}

func getAxisMinMax(plot []types.ScatterPoint) map[string]map[string]float64 {
	minmax := map[string]map[string]float64{
		"x": {
			"max": -math.MaxFloat64,
			"min": math.MaxFloat64,
		},
		"y": {
			"max": -math.MaxFloat64,
			"min": math.MaxFloat64,
		},
	}

	for _, point := range plot {
		if point.X > minmax["x"]["max"] {
			minmax["x"]["max"] = point.X
		}
		if point.X < minmax["x"]["min"] {
			minmax["x"]["min"] = point.X
		}

		if point.Y > minmax["y"]["max"] {
			minmax["y"]["max"] = point.Y
		}
		if point.Y < minmax["y"]["min"] {
			minmax["y"]["min"] = point.Y
		}
	}

	return minmax
}

func calculateLogTicks(logBase string, minmax map[string]map[string]float64) Ticks {
	return Ticks{
		X: calculateLogTicksForAxis(logBase, minmax["x"]["max"], minmax["x"]["min"]),
		Y: calculateLogTicksForAxis(logBase, minmax["y"]["max"], minmax["y"]["min"]),
	}
}

func calculateLogTicksForAxis(logBase string, max, min float64) []float64 {
	ticks := make([]float64, 0)
	firstTick := getLowerLogTick(logBase, min)
	lastTick := getUpperLogTick(logBase, max)

	logFloat, _ := strconv.ParseFloat(logBase, 64)
	for i := firstTick; i < lastTick; i *= logFloat {
		ticks = append(ticks, i)
	}
	ticks = append(ticks, lastTick)

	return ticks
}

func getLowerLogTick(logBase string, min float64) float64 {
	if logBase == "2" {
		logValue := math.Log2(min)
		if logValue < 0.25 {
			return 0.125
		}
		if logValue < 0.5 {
			return 0.25
		}
		return 1
	}

	logValue := math.Log10(min)
	if logValue < 0.1 {
		return 0.01
	}
	if logValue < 1 {
		return 0.1
	}
	return 1
}

func getUpperLogTick(logBase string, max float64) float64 {
	if logBase == "2" {
		return math.Pow(2, math.Ceil(math.Log2(max)))
	}
	return math.Pow(10, math.Ceil(math.Log10(max)))
}

func calculateLinearTicks(minmax map[string]map[string]float64) Ticks {
	return Ticks{
		X: calculateLinearTicksForAxis(minmax["x"]["max"]),
		Y: calculateLinearTicksForAxis(minmax["y"]["max"]),
	}
}

func calculateLinearTicksForAxis(max float64) []float64 {
	ticks := make([]float64, 0)

	power := int(math.Log10(max))
	step := math.Pow10(power)

	lastTick := getUpperLinearTick(max, step)
	for i := 0.0; i <= lastTick; i += step {
		ticks = append(ticks, i)
	}

	if ticks[len(ticks)-1] != lastTick {
		ticks = append(ticks, lastTick)
	}

	return ticks
}

func getUpperLinearTick(max, exp float64) float64 {
	return math.Ceil(max/exp) * exp
}

func scaleData(scatter *Scatter, axisLength float64) {
	scaleXValue := getScaler(scatter.LogBase, axisLength, scatter.Ticks.X)
	scaleYValue := getScaler(scatter.LogBase, axisLength, scatter.Ticks.Y)

	for i := range scatter.Plot {
		scatter.Plot[i].X = scaleXValue(math.Max(scatter.Plot[i].X, scatter.Ticks.X[0]))
		scatter.Plot[i].Y = scaleYValue(math.Max(scatter.Plot[i].Y, scatter.Ticks.Y[0]))
	}

	scatter.Ticks.XLabel = make([]string, len(scatter.Ticks.X))
	for i := range scatter.Ticks.X {
		scatter.Ticks.XLabel[i] = strconv.FormatFloat(scatter.Ticks.X[i], 'f', -1, 64)
		scatter.Ticks.X[i] = scaleXValue(scatter.Ticks.X[i])
	}
	scatter.Ticks.YLabel = make([]string, len(scatter.Ticks.Y))
	for i := range scatter.Ticks.Y {
		scatter.Ticks.YLabel[i] = strconv.FormatFloat(scatter.Ticks.Y[i], 'f', -1, 64)
		scatter.Ticks.Y[i] = scaleYValue(scatter.Ticks.Y[i])
	}
}

func getScaler(logBase string, axisLength float64, ticks []float64) func(float64) float64 {
	max := ticks[len(ticks)-1]
	if logBase != "none" {
		logFunc := math.Log10
		if logBase == "2" {
			logFunc = math.Log2
		}
		min := ticks[0]
		k := axisLength / (logFunc(max) - logFunc(min))
		c := -1 * k * logFunc(min)
		return func(point float64) float64 {
			return customMath.Round(k*logFunc(point)+c, 0.01)
		}
	}
	return func(point float64) float64 {
		return customMath.Round((point/max)*axisLength, 0.01)
	}
}
