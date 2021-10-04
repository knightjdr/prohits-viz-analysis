package scatter

import (
	"math"
	"strconv"

	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	customMath "github.com/knightjdr/prohits-viz-analysis/pkg/math"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func formatData(scatter *Scatter, axisLength float64) {
	axisBoundaries := defineAxisBoundaries(scatter.Plot, scatter.LogBase)
	scatter.Ticks = defineTicks(axisBoundaries, scatter.LogBase)
	scatter.Axes = defineAxes(scatter.Ticks)
	scaleData(scatter, axisLength)
}

func defineAxisBoundaries(plot []types.ScatterPoint, logBase string) boundaries {
	minMax := getAxisMinMax(plot)

	if logBase != "none" {
		return defineLogTickLimits(logBase, minMax)
	}

	return defineLinearTickLimits(minMax)
}

func getAxisMinMax(plot []types.ScatterPoint) boundaries {
	minmax := boundaries{
		x: boundary{
			max: -math.MaxFloat64,
			min: math.MaxFloat64,
		},
		y: boundary{
			max: -math.MaxFloat64,
			min: math.MaxFloat64,
		},
	}

	for _, point := range plot {
		if point.X > minmax.x.max {
			minmax.x.max = point.X
		}
		if point.X < minmax.x.min {
			minmax.x.min = point.X
		}

		if point.Y > minmax.y.max {
			minmax.y.max = point.Y
		}
		if point.Y < minmax.y.min {
			minmax.y.min = point.Y
		}
	}

	if minmax.x.max < 0 {
		minmax.x.max = 0
	}
	if minmax.y.max < 0 {
		minmax.y.max = 0
	}
	if minmax.x.min > 0 {
		minmax.x.min = 0
	}
	if minmax.y.min > 0 {
		minmax.y.min = 0
	}

	return minmax
}

func defineLinearTickLimits(minMax boundaries) boundaries {
	return boundaries{
		x: defineLinearTickLimitsForAxis(minMax.x),
		y: defineLinearTickLimitsForAxis(minMax.y),
	}
}

func defineLinearTickLimitsForAxis(axis boundary) boundary {
	axisMax := axis.max
	axisMin := axis.min

	if axisMax > 0 && axisMin < 0 {
		powerMax := math.Floor(math.Log10(math.Abs(axisMax)))
		powerMin := math.Floor(math.Log10(math.Abs(axisMin)))

		if powerMax < powerMin {
			axisMax = math.Pow(10, powerMin)
		}
		if powerMin < powerMax {
			axisMin = -math.Pow(10, powerMax)
		}
	}

	return boundary{
		max: defineLinearTickLimit(axisMax, math.Signbit(axis.max)),
		min: defineLinearTickLimit(axisMin, math.Signbit(axis.min)),
	}
}

func defineLinearTickLimit(boundary float64, isNegative bool) float64 {
	if boundary == 0 {
		return 0
	}

	exp := math.Pow(10, math.Floor(math.Log10(math.Abs(boundary))))
	if isNegative {
		return math.Floor(boundary/exp) * exp
	}
	return math.Ceil(boundary/exp) * exp
}

func defineLogTickLimits(base string, minMax boundaries) boundaries {
	return boundaries{
		x: defineLogTickLimitsForAxis(base, minMax.x),
		y: defineLogTickLimitsForAxis(base, minMax.y),
	}
}

func defineLogTickLimitsForAxis(base string, axis boundary) boundary {
	limits := boundary{}

	if axis.max != 0 {
		limits.max = defineUpperLogTickLimit(base, math.Abs(axis.max))
	} else {
		limits.max = -1 * defineLowerLogTickLimit(base, math.Abs(axis.max))
	}

	if axis.min != 0 {
		limits.min = -1 * defineUpperLogTickLimit(base, math.Abs(axis.min))
	} else {
		limits.min = defineLowerLogTickLimit(base, math.Abs(axis.min))
	}

	return limits
}

func defineUpperLogTickLimit(logBase string, value float64) float64 {
	if logBase == "2" {
		return math.Pow(2, math.Ceil(math.Log2(value)))
	}
	return math.Pow(10, math.Ceil(math.Log10(value)))
}

func defineLowerLogTickLimit(logBase string, value float64) float64 {
	if logBase == "2" {
		if value < 1 {
			return 0.5
		}
		return 1
	}

	if value < 1 {
		return 0.1
	}
	return 1
}

func defineTicks(axisBoundaries boundaries, logBase string) Ticks {
	if logBase != "none" {
		return calculateLogTicks(logBase, axisBoundaries)
	}

	return calculateLinearTicks(axisBoundaries)
}

func calculateLinearTicks(axisBoundaries boundaries) Ticks {
	return Ticks{
		X: calculateLinearTicksForAxis(axisBoundaries.x),
		Y: calculateLinearTicksForAxis(axisBoundaries.y),
	}
}

func calculateLinearTicksForAxis(axis boundary) []float64 {
	maxAbsoluteValue := math.Max(math.Abs(axis.max), math.Abs(axis.min))
	power := math.Floor(math.Log10(maxAbsoluteValue - 0.5))
	step := math.Pow(10, power)

	ticks := make([]float64, 0)
	for i := axis.min; i <= axis.max; i += step {
		ticks = append(ticks, i)
	}
	if ticks[len(ticks)-1] != axis.max {
		ticks = append(ticks, axis.max)
	}
	return ticks
}

func calculateLogTicks(logBase string, axisBoundaries boundaries) Ticks {
	return Ticks{
		X: calculateLogTicksForAxis(logBase, axisBoundaries.x),
		Y: calculateLogTicksForAxis(logBase, axisBoundaries.y),
	}
}

func calculateLogTicksForAxis(logBase string, axis boundary) []float64 {
	ticks := make([]float64, 0)

	logBaseAsFloat, _ := strconv.ParseFloat(logBase, 64)

	stepMultiplier := logBaseAsFloat
	if axis.min <= 0 {
		stepMultiplier = 1 / logBaseAsFloat
	}

	if axis.min < 0 && axis.max > 0 {
		end := -1 / logBaseAsFloat
		for i := axis.min; i < axis.max; i *= stepMultiplier {
			ticks = append(ticks, i)
			lastTickIndex := len(ticks) - 1
			if ticks[lastTickIndex] >= end && ticks[lastTickIndex] < 0 {
				i = -ticks[lastTickIndex] / logBaseAsFloat
				stepMultiplier = logBaseAsFloat
			}
		}
	} else {
		for i := axis.min; i < axis.max; i *= stepMultiplier {
			ticks = append(ticks, i)
		}
	}
	ticks = append(ticks, axis.max)

	return ticks
}

func defineAxes(ticks Ticks) Axes {
	defineOrigin := func(axis []float64) float64 {
		lastIndex := len(axis) - 1
		if axis[0] == 0 ||
			axis[lastIndex] == 0 ||
			(axis[0] < 0 && axis[lastIndex] > 0) {
			return 0
		}
		iMin := 0
		for i, tick := range axis {
			if math.Abs(tick) < math.Abs((axis[iMin])) {
				iMin = i
			}
		}
		return axis[iMin]
	}

	xOrigin := defineOrigin(ticks.X)
	yOrigin := defineOrigin(ticks.Y)
	return Axes{
		X: Line{
			X1: ticks.X[0],
			X2: ticks.X[len(ticks.X)-1],
			Y1: yOrigin,
			Y2: yOrigin,
		},
		Y: Line{
			X1: xOrigin,
			X2: xOrigin,
			Y1: ticks.Y[0],
			Y2: ticks.Y[len(ticks.Y)-1],
		},
	}
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
		scatter.Ticks.XLabel[i] = float.RemoveTrailingZeros(scatter.Ticks.X[i])
		scatter.Ticks.X[i] = scaleXValue(scatter.Ticks.X[i])
	}
	scatter.Ticks.YLabel = make([]string, len(scatter.Ticks.Y))
	for i := range scatter.Ticks.Y {
		scatter.Ticks.YLabel[i] = float.RemoveTrailingZeros(scatter.Ticks.Y[i])
		scatter.Ticks.Y[i] = scaleYValue(scatter.Ticks.Y[i])
	}

	scatter.Axes.X.X1 = scaleXValue(scatter.Axes.X.X1)
	scatter.Axes.X.X2 = scaleXValue(scatter.Axes.X.X2)
	scatter.Axes.X.Y1 = customMath.Round(axisLength-scaleYValue(scatter.Axes.X.Y1), 0.01)
	scatter.Axes.X.Y2 = customMath.Round(axisLength-scaleYValue(scatter.Axes.X.Y2), 0.01)
	scatter.Axes.Y.X1 = scaleXValue(scatter.Axes.Y.X1)
	scatter.Axes.Y.X2 = scaleXValue(scatter.Axes.Y.X2)
	scatter.Axes.Y.Y1 = customMath.Round(axisLength-scaleYValue(scatter.Axes.Y.Y1), 0.01)
	scatter.Axes.Y.Y2 = customMath.Round(axisLength-scaleYValue(scatter.Axes.Y.Y2), 0.01)
}

func getScaler(logBase string, axisLength float64, ticks []float64) func(float64) float64 {
	first := ticks[0]
	last := ticks[len(ticks)-1]
	if logBase != "none" {
		logFunc := math.Log10
		if logBase == "2" {
			logFunc = math.Log2
		}

		segments := len(ticks) - 1
		numNegativeTicks := 0
		for _, tick := range ticks {
			if tick < 0 {
				numNegativeTicks += 1
			}
		}
		numPositiveTicks := len(ticks) - numNegativeTicks
		negAxisLength := float64(0)
		if numNegativeTicks > 0 {
			negAxisLength = axisLength * ((float64(numNegativeTicks) - 1) / float64(segments))
		}
		posAxisLength := float64(0)
		if numPositiveTicks > 0 {
			posAxisLength = axisLength * ((float64(numPositiveTicks) - 1) / float64(segments))
		}

		negativeExtremes := map[string]float64{
			"max": 0,
			"min": math.Inf(1),
		}
		positiveExtremes := map[string]float64{
			"max": 0,
			"min": math.Inf(1),
		}
		for _, tick := range ticks {
			absoluteTick := math.Abs(tick)
			if tick < 0 && absoluteTick > negativeExtremes["max"] {
				negativeExtremes["max"] = absoluteTick
			}
			if tick < 0 && absoluteTick < negativeExtremes["min"] {
				negativeExtremes["min"] = absoluteTick
			}
			if tick > positiveExtremes["max"] {
				positiveExtremes["max"] = tick
			}
			if tick > 0 && tick < positiveExtremes["min"] {
				positiveExtremes["min"] = tick
			}
		}

		voidSpace := float64(0)
		if numPositiveTicks > 0 && numNegativeTicks > 0 {
			voidSpace = axisLength / (float64(len(ticks)) - 1)
		}

		kNeg := float64(0)
		if negAxisLength > 0 {
			kNeg = negAxisLength / (logFunc(negativeExtremes["max"]) - logFunc(negativeExtremes["min"]))

		}
		kPos := float64(0)
		if posAxisLength > 0 {
			kPos = posAxisLength / (logFunc(positiveExtremes["max"]) - logFunc(positiveExtremes["min"]))

		}
		cNeg := -1 * kNeg * logFunc(negativeExtremes["min"])
		cPos := -1 * kPos * logFunc(positiveExtremes["min"])
		scaleLinear := func(point float64) float64 {
			return customMath.Round(
				negAxisLength+(((point+negativeExtremes["min"])/(positiveExtremes["min"]+negativeExtremes["min"]))*voidSpace),
				0.01,
			)
		}
		return func(point float64) float64 {
			if point >= 0 {
				if point < positiveExtremes["min"] {
					return scaleLinear(point)
				}
				return customMath.Round(kPos*logFunc(point)+cPos+negAxisLength+voidSpace, 0.01)
			}
			if point > -negativeExtremes["min"] {
				return scaleLinear(point)
			}
			return customMath.Round(negAxisLength-(kNeg*logFunc(math.Abs(point))+cNeg), 0.01)
		}
	}
	return func(point float64) float64 {
		return customMath.Round(axisLength*(point-first)/(last-first), 0.01)
	}
}
