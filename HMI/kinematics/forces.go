package kinematics

import (
	"gonum.org/v1/gonum/mat"
)

func (p *StewartPlatform) GetForceTransformationMatrix() {
	if p.LegUnitVector != nil && p.Bvector != nil {
		bi := mat.NewDense(3, 6, nil)
		bi.MulElem(p.LegVector, p.Bvector)
		ftm := mat.NewDense(6, 6, nil)
		for i := 0; i < 3; i++ {
			row := p.LegUnitVector.RawRowView(i)
			ftm.SetRow(i, row)
		}
		for i := 3; i < 6; i++ {
			row := bi.RawRowView(i - 3)
			ftm.SetRow(i, row)
		}
		p.ForceTransformationMatrix = ftm
	}
}

func (p *StewartPlatform) GetForceAndMoments(f1 float64, f2 float64, f3 float64, f4 float64, f5 float64, f6 float64) []float64 {
	forces := mat.NewDense(6, 1, []float64{f1, f2, f3, f4, f5, f6})
	res := mat.NewDense(6, 1, nil)
	res.Mul(p.ForceTransformationMatrix, forces)
	return res.RawMatrix().Data
}
