package kinematics

import (
	"fmt"
	"math"

	"github.com/sammyoina/stewart-platform-ui/models"
	"gonum.org/v1/gonum/mat"
)

type StewartPlatform struct {
	Beta                      *mat.Dense
	PsiBase                   *mat.Dense
	PsiPlatform               *mat.Dense
	BaseAttachment            mat.Matrix
	PlatformAttachment        mat.Matrix
	PlatformRadius            float64
	BaseRadius                float64
	GammaBase                 float64
	GammaPlatform             float64
	ServoHornLength           float64
	RodLength                 float64
	HomePosition              *mat.Dense
	LegUnitVector             *mat.Dense
	LegVector                 mat.Matrix
	Bvector                   *mat.Dense
	ForceTransformationMatrix *mat.Dense
}

func NewStewartPlatform(baseRadius float64, PlatformRadius float64, GammaBase float64, GammaPlatform float64, ServoHornLength float64, RodLength float64, offset float64) StewartPlatform {
	var Platform StewartPlatform
	Platform.Beta = mat.NewDense(1, 6, nil)
	Platform.Beta.SetRow(0, []float64{
		math.Pi/2 + math.Pi,
		math.Pi / 2,
		2*math.Pi/3 + math.Pi/2 + math.Pi,
		2*math.Pi/3 + math.Pi/2,
		4*math.Pi/3 + math.Pi/2 + math.Pi,
		4*math.Pi/3 + math.Pi/2,
	})
	Platform.PsiBase = mat.NewDense(1, 6, nil)
	Platform.PsiBase.SetRow(0, []float64{
		-1 * GammaBase,
		GammaBase,
		2*math.Pi/3 - GammaBase,
		2*math.Pi/3 + GammaBase,
		2*math.Pi/3 + 2*math.Pi/3 - GammaBase,
		2*math.Pi/3 + 2*math.Pi/3 + GammaBase,
	})
	Platform.PsiPlatform = mat.NewDense(1, 6, nil)
	Platform.PsiPlatform.SetRow(0, []float64{
		math.Pi/3 + 2*math.Pi/3 + 2*math.Pi/3 + GammaPlatform,
		math.Pi/3 + -1*GammaPlatform,
		math.Pi/3 + GammaPlatform,
		math.Pi/3 + 2*math.Pi/3 - GammaPlatform,
		math.Pi/3 + 2*math.Pi/3 + GammaPlatform,
		math.Pi/3 + 2*math.Pi/3 + 2*math.Pi/3 - GammaPlatform,
	})
	baseAttachment := mat.NewDense(6, 3, nil)
	baseAttachment.SetRow(0, []float64{
		math.Cos(Platform.PsiBase.At(0, 0)),
		math.Sin(Platform.PsiBase.At(0, 0)), 0,
	})
	baseAttachment.SetRow(1, []float64{
		math.Cos(Platform.PsiBase.At(0, 1)),
		math.Sin(Platform.PsiBase.At(0, 1)), 0,
	})
	baseAttachment.SetRow(2, []float64{
		math.Cos(Platform.PsiBase.At(0, 2)),
		math.Sin(Platform.PsiBase.At(0, 2)), 0,
	})
	baseAttachment.SetRow(3, []float64{
		math.Cos(Platform.PsiBase.At(0, 3)),
		math.Sin(Platform.PsiBase.At(0, 3)), 0,
	})
	baseAttachment.SetRow(4, []float64{
		math.Cos(Platform.PsiBase.At(0, 4)),
		math.Sin(Platform.PsiBase.At(0, 4)), 0,
	})
	baseAttachment.SetRow(5, []float64{
		math.Cos(Platform.PsiBase.At(0, 5)),
		math.Sin(Platform.PsiBase.At(0, 5)), 0,
	})
	baseAttachment.Apply(func(i, j int, v float64) float64 {
		return v * baseRadius
	}, baseAttachment)
	Platform.BaseAttachment = baseAttachment.T()

	platformAttachment := mat.NewDense(6, 3, nil)
	platformAttachment.SetRow(0, []float64{
		math.Cos(Platform.PsiPlatform.At(0, 0)),
		math.Sin(Platform.PsiPlatform.At(0, 0)), 0,
	})
	platformAttachment.SetRow(1, []float64{
		math.Cos(Platform.PsiPlatform.At(0, 1)),
		math.Sin(Platform.PsiPlatform.At(0, 1)), 0,
	})
	platformAttachment.SetRow(2, []float64{
		math.Cos(Platform.PsiPlatform.At(0, 2)),
		math.Sin(Platform.PsiPlatform.At(0, 2)), 0,
	})
	platformAttachment.SetRow(3, []float64{
		math.Cos(Platform.PsiPlatform.At(0, 3)),
		math.Sin(Platform.PsiPlatform.At(0, 3)), 0,
	})
	platformAttachment.SetRow(4, []float64{
		math.Cos(Platform.PsiPlatform.At(0, 4)),
		math.Sin(Platform.PsiPlatform.At(0, 4)), 0,
	})
	platformAttachment.SetRow(5, []float64{
		math.Cos(Platform.PsiPlatform.At(0, 5)),
		math.Sin(Platform.PsiPlatform.At(0, 5)), 0,
	})
	platformAttachment.Scale(PlatformRadius, platformAttachment)
	Platform.PlatformAttachment = platformAttachment.T()

	Platform.BaseRadius = baseRadius
	Platform.PlatformRadius = PlatformRadius
	Platform.GammaBase = GammaBase
	Platform.GammaPlatform = GammaPlatform
	Platform.ServoHornLength = ServoHornLength
	Platform.RodLength = RodLength

	sub1 := math.Pow(RodLength, 2) + math.Pow(ServoHornLength, 2)
	attachmentDiff1 := mat.NewDense(1, 6, nil)
	attachmentDiff2 := mat.NewDense(1, 6, nil)
	attachmentDiff1.Sub(platformAttachment.ColView(0).T(), baseAttachment.ColView(0).T())
	attachmentDiff2.Sub(platformAttachment.ColView(1).T(), baseAttachment.ColView(1).T())
	attachmentDiff1.Apply(elementWiseSquare, attachmentDiff1)
	attachmentDiff2.Apply(elementWiseSquare, attachmentDiff2)
	summedAttachDiff := mat.NewDense(1, 6, nil)
	summedAttachDiff.Add(attachmentDiff1, attachmentDiff2)
	z := mat.NewDense(1, 6, nil)
	z.Apply(func(i, j int, v float64) float64 {
		return math.Sqrt(sub1 - v)
	}, summedAttachDiff)
	Platform.HomePosition = mat.NewDense(1, 3, nil)
	fmt.Println("Home: ", z.At(0, 0))
	Platform.HomePosition.SetRow(0, []float64{0, 0, z.At(0, 0)})
	Platform.Bvector = mat.NewDense(3, 6, nil)
	Platform.LegUnitVector = mat.NewDense(3, 6, nil)
	return Platform
}

func (p *StewartPlatform) Calculate(yaw float64, roll float64, pitch float64, transx float64, transy float64, transz float64) (*models.ServoPositionEvent, error) {
	rots := mat.NewDense(3, 1, nil)
	rots.SetCol(0, []float64{yaw, pitch, roll})
	transl := mat.NewDense(3, 1, nil)
	transl.SetCol(0, []float64{transx, transy, transz})
	R := mat.NewDense(3, 3, nil)
	R.Mul(rotZ(rots.At(2, 0)), rotY(rots.At(1, 0)))
	R.Mul(R, rotX(rots.At(0, 0)))
	translRep := mat.NewDense(3, 6, nil)
	homePosRep := mat.NewDense(3, 6, nil)
	translRep.SetCol(0, []float64{transx, transy, transz})
	translRep.SetCol(1, []float64{transx, transy, transz})
	translRep.SetCol(2, []float64{transx, transy, transz})
	translRep.SetCol(3, []float64{transx, transy, transz})
	translRep.SetCol(4, []float64{transx, transy, transz})
	translRep.SetCol(5, []float64{transx, transy, transz})
	homePosRep.SetCol(0, []float64{p.HomePosition.At(0, 0), p.HomePosition.At(0, 1), p.HomePosition.At(0, 2)})
	homePosRep.SetCol(1, []float64{p.HomePosition.At(0, 0), p.HomePosition.At(0, 1), p.HomePosition.At(0, 2)})
	homePosRep.SetCol(2, []float64{p.HomePosition.At(0, 0), p.HomePosition.At(0, 1), p.HomePosition.At(0, 2)})
	homePosRep.SetCol(3, []float64{p.HomePosition.At(0, 0), p.HomePosition.At(0, 1), p.HomePosition.At(0, 2)})
	homePosRep.SetCol(4, []float64{p.HomePosition.At(0, 0), p.HomePosition.At(0, 1), p.HomePosition.At(0, 2)})
	homePosRep.SetCol(5, []float64{p.HomePosition.At(0, 0), p.HomePosition.At(0, 1), p.HomePosition.At(0, 2)})
	l := mat.NewDense(3, 6, nil)
	l.Mul(R, p.PlatformAttachment)
	l.Sub(l, p.BaseAttachment)
	l.Add(l, translRep)
	l.Add(l, homePosRep)
	lll := mat.NewDense(1, 6, nil)
	for i := 0; i < 6; i++ {
		lll.Set(0, i, mat.Norm(l.ColView(i), 2))
	}
	L := mat.NewDense(3, 6, nil)
	L.Add(l, p.BaseAttachment)
	p.LegVector = L

	for column := 0; column < 6; column++ {
		li := GetUnitVector([]float64{p.LegVector.At(0, column), p.LegVector.At(1, column), p.LegVector.At(2, column)})
		p.LegUnitVector.SetCol(column, li)
	}
	lx := l.RowView(0).T()
	ly := l.RowView(1).T()
	lz := l.RowView(2).T()
	g := mat.NewDense(1, 6, nil)
	g.Apply(elementWiseSquare, lll)
	g.Apply(func(i, j int, v float64) float64 {
		return v + +math.Pow(p.ServoHornLength, 2) - math.Pow(p.RodLength, 2)
	}, g)
	e := mat.NewDense(1, 6, nil)
	e.Apply(func(i, j int, v float64) float64 {
		return 2 * p.ServoHornLength * v
	}, lz)
	angles := make([]float64, 6)

	for i := 0; i < 6; i++ {
		fk := 2 * p.ServoHornLength * (math.Cos(p.Beta.At(0, i)*lx.At(0, i) + math.Sin(p.Beta.At(0, i)*ly.At(0, i))))
		angles[i] = math.Asin(g.At(0, i)/math.Sqrt(math.Pow(e.At(0, i), 2)+math.Pow(fk, 2))) - math.Atan2(fk, e.At(0, i))
		p.Bvector.SetCol(i, []float64{
			p.ServoHornLength*math.Cos(angles[i])*math.Cos(p.Beta.At(0, i)) + p.BaseAttachment.At(0, i),
			p.ServoHornLength*math.Cos(angles[i])*math.Sin(p.Beta.At(0, i)) + p.BaseAttachment.At(1, i),
			p.ServoHornLength * math.Sin(angles[i]),
		})
		angles[i] = R2d(angles[i])
		if math.IsInf(angles[i], 0) || math.IsNaN(angles[i]) {
			return nil, fmt.Errorf("nan or inf angle on servo %d", i)
		}
		if angles[i]+90 < 0 || angles[i] > 180 {
			return nil, fmt.Errorf("out of range angle on servo %d", i)
		}
	}

	var angs models.ServoPositionEvent
	angs.Servo1 = float32(angles[0]) + 90
	angs.Servo2 = float32(angles[1]) + 90
	angs.Servo3 = float32(angles[2]) + 90
	angs.Servo4 = float32(angles[3]) + 90
	angs.Servo5 = float32(angles[4]) + 90
	angs.Servo6 = float32(angles[5]) + 90

	return &angs, nil
}

func elementWiseSquare(i, j int, v float64) float64 {
	return math.Pow(v, 2)
}
func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func GetUnitVector(vector []float64) []float64 {
	magnitude := math.Sqrt(math.Pow(vector[0], 2) + math.Pow(vector[1], 2) + math.Pow(vector[2], 2))
	return []float64{vector[0] / magnitude, vector[1] / magnitude, vector[2] / magnitude}
}
