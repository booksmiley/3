package engine

import (
	"github.com/mumax/3/cuda"
	"github.com/mumax/3/data"
	"github.com/mumax/3/script"
	"github.com/mumax/3/util"
	"math"
	"reflect"
)

// An excitation, typically field or current,
// can be defined region-wise plus extra mask*multiplier terms.
// TODO: per-component mulmasks
type excitation struct {
	name       string
	perRegion  VectorParam // Region-based excitation
	extraTerms []mulmask   // add extra mask*multiplier terms
}

type mulmask struct {
	mul  func() float64
	mask *data.Slice
}

func (e *excitation) init(name, unit, desc string) {
	e.name = name
	e.perRegion.init("_"+name+"_perRegion", unit, "(internal)") // name starts with underscore: unexported
	DeclLValue(name, e, cat(desc, unit))
}

func (e *excitation) AddTo(dst *data.Slice) {
	if !e.perRegion.isZero() {
		cuda.RegionAddV(dst, e.perRegion.gpuLUT(), regions.Gpu())
	}
	for _, t := range e.extraTerms {
		var mul float32 = 1
		if t.mul != nil {
			mul = float32(t.mul())
		}
		cuda.Madd2(dst, dst, t.mask, 1, mul)
	}
}

func (e *excitation) isZero() bool {
	return e.perRegion.isZero() && len(e.extraTerms) == 0
}

func (e *excitation) Slice() (*data.Slice, bool) {
	buf := cuda.Buffer(e.NComp(), e.Mesh().Size())
	cuda.Zero(buf)
	e.AddTo(buf)
	return buf, true
}

// Add an extra mask*multiplier term to the excitation.
func (e *excitation) Add(mask *data.Slice, f script.ScalarFunction) {
	var mul func() float64
	if f != nil {
		if Const(f) {
			val := f.Float()
			mul = func() float64 {
				return val
			}
		} else {
			mul = func() float64 {
				return f.Float()
			}
		}
	}
	if mask != nil {
		checkNaN(mask, e.Name()+".add()") // TODO: in more places
		mask = data.Resample(mask, e.Mesh().Size())
		mask = assureGPU(mask)
	}
	e.extraTerms = append(e.extraTerms, mulmask{mul, mask})
}

func assureGPU(s *data.Slice) *data.Slice {
	if s.GPUAccess() {
		return s
	} else {
		return cuda.GPUCopy(s)
	}
}

// user script: has to be 3-vector
func (e *excitation) SetRegion(region int, f script.VectorFunction) {
	e.perRegion.SetRegion(region, f)
}

// for gui (nComp agnostic)
func (e *excitation) setRegion(region int, value []float64) {
	e.perRegion.setRegion(region, value)
}

// does not use extramask!
func (e *excitation) getRegion(region int) []float64 {
	return e.perRegion.getRegion(region)
}

func (e *excitation) average() []float64 {
	return qAverageUniverse(e)
}

func (e *excitation) Average() data.Vector {
	return unslice(qAverageUniverse(e))
}

func (e *excitation) IsUniform() bool {
	return e.perRegion.IsUniform()
}

func (e *excitation) SetValue(v interface{}) {
	e.perRegion.SetValue(v) // allows function of time
}

func (e *excitation) Set(v data.Vector) {
	e.perRegion.setRegions(0, NREGION, slice(v))
}

func (e *excitation) Name() string            { return e.name }
func (e *excitation) Unit() string            { return e.perRegion.Unit() }
func (e *excitation) NComp() int              { return e.perRegion.NComp() }
func (e *excitation) Mesh() *data.Mesh        { return Mesh() }
func (e *excitation) Region(r int) *vOneReg   { return vOneRegion(e, r) }
func (e *excitation) Comp(c int) *comp        { return Comp(e, c) }
func (e *excitation) Eval() interface{}       { return e }
func (e *excitation) Type() reflect.Type      { return reflect.TypeOf(new(excitation)) }
func (e *excitation) InputType() reflect.Type { return script.VectorFunction_t }

//func (e *excitation) Comp(c int) *comp        { return Comp(e, c) }

func checkNaN(s *data.Slice, name string) {
	h := s.Host()
	for _, h := range h {
		for _, v := range h {
			if math.IsNaN(float64(v)) || math.IsInf(float64(v), 0) {
				util.Fatal("NaN or Inf in", name)
			}
		}
	}
}
