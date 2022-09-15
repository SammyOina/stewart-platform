package kinematics

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func rotZ(psi float64) *mat.Dense {

	//psi = psi * math.Pi / 180

	rotRes := mat.NewDense(3, 3, nil)
	rotRes.Set(0, 0, math.Cos(psi))
	rotRes.Set(0, 1, math.Sin(psi))
	rotRes.Set(0, 2, 0)
	rotRes.Set(1, 0, -1*math.Sin(psi))
	rotRes.Set(1, 1, math.Cos(psi))
	rotRes.Set(1, 2, 0)
	rotRes.Set(2, 0, 0)
	rotRes.Set(2, 1, 0)
	rotRes.Set(2, 2, 1)

	return rotRes
}

func rotY(theta float64) *mat.Dense {
	//theta = theta * math.Pi / 180
	rotRes := mat.NewDense(3, 3, nil)
	rotRes.SetRow(0, []float64{math.Cos(theta), 0, -1 * math.Sin(theta)})
	rotRes.SetRow(1, []float64{0, 1, 0})
	rotRes.SetRow(2, []float64{math.Sin(theta), 0, math.Cos(theta)})
	return rotRes
}

func rotX(phi float64) *mat.Dense {
	//phi = phi * math.Pi / 180

	rotRes := mat.NewDense(3, 3, nil)
	rotRes.SetRow(0, []float64{1, 0, 0})
	rotRes.SetRow(1, []float64{0, math.Cos(phi), math.Sin(phi)})
	rotRes.SetRow(2, []float64{0, -1 * math.Sin(phi), math.Cos(phi)})

	return rotRes
}

func d2r(angle float64) float64 {
	return angle * math.Pi / 180
}

func r2d(angle float64) float64 {
	return angle * 180 / math.Pi
}
