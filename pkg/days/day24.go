package days

import (
	"fmt"
	"os"
	"strings"
  "regexp"
  "gonum.org/v1/gonum/mat"

	"github.com/davidkleiven/gononlin/nonlin"
)


type R3 struct {
  x,y,z float64
}
type hailstone struct {
  p,v R3
}

type matrix struct {
  val [][]float64
}


func stringArrToFloatArr(ss []string) []float64 {
  nums := make([]float64, len(ss))
  for i,s := range ss {
    nums[i] = float64(stoi(s))
  }
  return nums
}

func getHailstones(s string) []hailstone {
  lines := strings.Split(s, "\n")
  hs := make([]hailstone, len(lines))
  r := regexp.MustCompile("(-\\d+)|(\\d+)")

  var line []string
  var p,v []float64
  for l := range lines {
    line = strings.Split(lines[l], " @ ")
    p = stringArrToFloatArr(r.FindAllString(line[0],-1))
    v = stringArrToFloatArr(r.FindAllString(line[1],-1))
    hs[l] = hailstone{p: R3{p[0],p[1],p[2]}, v: R3{v[0],v[1],v[2]}}
  }

  return hs
}

func getSubmatrix(m matrix, i, j int) matrix {
  nm := make([][]float64, len(m.val)-1)
  for r := range nm {
    nm[r] = make([]float64, len(m.val[0])-1)
  }
  in,jn := 0,0
  for ii,row := range m.val {
    if ii == i { continue }
    for jj,v := range row {
      if jj == j { continue }
      nm[in][jn] = v
      jn++
    }
    jn = 0
    in++
  }

  return matrix{nm}

}

func matrixDet(m matrix) float64 {
  if len(m.val) == 1 {
    return float64(m.val[0][0])
  }
  var det float64 = 0
  var sgn float64 =  1.0
  for j := 0; j < len(m.val[0]); j++ {
    det += sgn*float64(m.val[0][j])*matrixDet(getSubmatrix(m, 0, j))
    sgn *= -1
  }
  return det
}

func intersectLine2D(h1, h2 hailstone, x, y interval) bool {
  //h1.v.x*t - h2.v.x*s = h2.p.x - h1.p.x
  //h1.v.y*t - h2.v.y*s = h2.p.y - h1.p.y

  detc := matrixDet(matrix{[][]float64{{h1.v.x, -h2.v.x},{h1.v.y, -h2.v.y}}})

  if detc == 0 {
    return false
  }

  dett := matrixDet(matrix{[][]float64{{h2.p.x-h1.p.x, -h2.v.x}, {h2.p.y-h1.p.y, -h2.v.y}}})
  t := dett/detc
  if t < 0 {
    return false
  }
  
  dets := matrixDet(matrix{[][]float64{{h1.v.x, h2.p.x-h1.p.x}, {h1.v.y, h2.p.y-h1.p.y}}})
  s := dets/detc
  if s < 0 {
    return false
  }
  xt := h1.p.x+h1.v.x*t
  yt := h1.p.y+h1.v.y*t

  return float64(x.min) <= xt && float64(x.max) >= xt && float64(y.min) <= yt && float64(y.max) >= yt
}

func replaceColumn(m, c matrix, col int) matrix {
  nm := make([][]float64, len(m.val))
  for i := range m.val {
    nm[i] = make([]float64, len(m.val[i]))
    copy(nm[i], m.val[i])
  }
  fmt.Printf("nm: %v\n", nm)
  fmt.Printf("m.val: %v\n", m.val)
  for i := range c.val[0] {
    nm[i][col] = c.val[0][i]
  }
  fmt.Printf("nm.val: %v\n", nm)


  return matrix{nm}
}

func intersect3D(h1, h2, h3 hailstone) {
  m := matrix{[][]float64{
    {h2.v.y-h1.v.y, h1.v.x-h2.v.x, 0, h1.p.y-h2.p.y, h2.p.x-h1.p.x, 0},
    {h3.v.y-h1.v.y, h1.v.x-h3.v.x, 0, h1.p.y-h3.p.y, h3.p.x-h1.p.x, 0},

    {h2.v.z-h1.v.z, 0, h1.v.x-h2.v.x, h1.p.z-h2.p.z, 0, h2.p.x-h1.p.x},
    {h3.v.z-h1.v.z, 0, h1.v.x-h3.v.x, h1.p.z-h3.p.z, 0, h3.p.x-h1.p.x},

    {0, h1.v.z-h2.v.z, h2.v.y-h1.v.y, 0, h2.p.z-h1.p.z, h1.p.y-h2.p.y},
    {0, h1.v.z-h3.v.z, h3.v.y-h1.v.y, 0, h3.p.z-h1.p.z, h1.p.y-h3.p.y},
  }}

  s := matrix{[][]float64{{
      h2.p.x*h2.v.y -h2.p.y*h2.v.x -h1.p.x*h1.v.y +h1.p.y*h1.v.x,
      h3.p.x*h3.v.y -h3.p.y*h3.v.x -h1.p.x*h1.v.y +h1.p.y*h1.v.x,

      h2.p.x*h2.v.z -h2.p.z*h2.v.x -h1.p.x*h1.v.z +h1.p.z*h1.v.z,
      h3.p.x*h3.v.z -h3.p.z*h3.v.x -h1.p.x*h1.v.z +h1.p.z*h1.v.z,

      -h2.v.z*h2.p.y +h2.v.y*h2.p.z -h1.v.z*h1.p.x +h1.v.x*h1.p.z,
      -h3.v.z*h3.p.y +h3.v.y*h3.p.z -h1.v.z*h1.p.x +h1.v.x*h1.p.z,
    }}}
  mi := []float64{
    h2.v.y-h1.v.y, h1.v.x-h2.v.x, 0, h1.p.y-h2.p.y, h2.p.x-h1.p.x, 0,
    h3.v.y-h1.v.y, h1.v.x-h3.v.x, 0, h1.p.y-h3.p.y, h3.p.x-h1.p.x, 0,

    h2.v.z-h1.v.z, 0, h1.v.x-h2.v.x, h1.p.z-h2.p.z, 0, h2.p.x-h1.p.x,
    h3.v.z-h1.v.z, 0, h1.v.x-h3.v.x, h1.p.z-h3.p.z, 0, h3.p.x-h1.p.x,

    0, h1.v.z-h2.v.z, h2.v.y-h1.v.y, 0, h2.p.z-h1.p.z, h1.p.y-h2.p.y,
    0, h1.v.z-h3.v.z, h3.v.y-h1.v.y, 0, h3.p.z-h1.p.z, h1.p.y-h3.p.y,
  }

  si :=[]float64{
      h2.p.x*h2.v.y -h2.p.y*h2.v.x -h1.p.x*h1.v.y +h1.p.y*h1.v.x,
      h3.p.x*h3.v.y -h3.p.y*h3.v.x -h1.p.x*h1.v.y +h1.p.y*h1.v.x,

      h2.p.x*h2.v.z -h2.p.z*h2.v.x -h1.p.x*h1.v.z +h1.p.z*h1.v.z,
      h3.p.x*h3.v.z -h3.p.z*h3.v.x -h1.p.x*h1.v.z +h1.p.z*h1.v.z,

      -h2.v.z*h2.p.y +h2.v.y*h2.p.z -h1.v.z*h1.p.x +h1.v.x*h1.p.z,
      -h3.v.z*h3.p.y +h3.v.y*h3.p.z -h1.v.z*h1.p.x +h1.v.x*h1.p.z,
    }
  fmt.Printf("s: %v\n", s)
  mvx := replaceColumn(m,s,0)
  mvy := replaceColumn(m,s,1)
  mvz := replaceColumn(m,s,2)
  mpx := replaceColumn(m,s,3)
  mpy := replaceColumn(m,s,4)
  mpz := replaceColumn(m,s,5)
  detm := matrixDet(m)
  detmvx := matrixDet(mvx)
  detmvy := matrixDet(mvy)
  detmvz := matrixDet(mvz)
  detmpx := matrixDet(mpx)
  detmpy := matrixDet(mpy)
  detmpz := matrixDet(mpz)
  println(detmvx/detm)
  println(detmvy/detm)
  println(detmvz/detm)
  println(detmpx/detm)
  println(detmpy/detm)
  println(detmpz/detm)
  fmt.Printf("mvx: %v\n", mvx)

  mm := mat.NewDense(6,6, mi)
  ss := mat.NewDense(6,1, si)
  fmt.Printf("mm: %v\n", mm)
  fmt.Printf("ss: %v\n", ss)

  ss.Solve(mm,ss)
  fmt.Printf("ss: %v\n", ss)





  problem := nonlin.Problem{
    F : func(out, x []float64) {
      out[0] = x[0]-h1.p.x+(x[3]-h1.v.x)*x[6]
      out[1] = x[1]-h1.p.y+(x[4]-h1.v.y)*x[6]
      out[2] = x[2]-h1.p.z+(x[5]-h1.v.z)*x[6]

      out[3] = x[0]-h2.p.x+(x[3]-h2.v.x)*x[7]
      out[4] = x[1]-h2.p.y+(x[4]-h2.v.y)*x[7]
      out[5] = x[2]-h2.p.z+(x[5]-h2.v.z)*x[7]

      out[6] = x[0]-h3.p.x+(x[3]-h3.v.x)*x[8]
      out[7] = x[1]-h3.p.y+(x[4]-h3.v.y)*x[8]
      out[8] = x[2]-h3.p.z+(x[5]-h3.v.z)*x[8]
    },
  }
   x0 := []float64{0.0,0.0,0.0,0.0,0.0,0.0,1.0,1.0,1.0 }

   solver := nonlin.NewtonKrylov{Maxiter: 100000000000, StepSize: 1e-6, Tol: 1e-2}
   println(solver.Maxiter)
   res := solver.Solve(problem, x0)
   fmt.Printf("res: %v\n", res)
   
}

func throwRock(rock hailstone, hl []hailstone, v []int) {

}

func Day24() bool {
  b,_ := os.ReadFile("input/sample24")
  hl := getHailstones(string(b[:len(b)-1]))


  x := interval{min: 200000000000000, max: 400000000000000}
  p1 := 0
  for i := 0; i < len(hl); i++ {
    for j := i+1; j < len(hl); j++ {
      res := intersectLine2D(hl[i], hl[j], x, x)
      if res {
        p1 += 1
      }
    }
  }

  p2 := 0
  fmt.Printf("Part 1: %d\nPart 2: %d\n", p1, p2)
  //rock := hailstone{p: R3{0,0,0}, v: R3{0,0,0}}
  intersect3D(hl[4], hl[1], hl[2])

  return true
}
