package days

import (
	"fmt"
	"os"
	"strings"
)

func getGardenPlots(s string) [][]rune {
  lines := strings.Split(s, "\n")
  gp := make([][]rune, len(lines))
  for l := range lines {
    gp[l] = []rune(lines[l])
  }
  return gp
}

func startingPoint(r rune, gp [][]rune) R2 {
  for i := range gp {
    for j := range gp[i] {
      if gp[i][j] == r {
        return R2{x:j,y:i}
      }
    }
  }
  return R2{0,0}
}

func addToSet(s map[R2]struct{}, r R2) {
  s[r] = struct{}{}
}

func getNextGardenPlots(p R2, gp [][]rune, n map[R2]struct{}) {
  limits := R2{x:len(gp[0]), y:len(gp)}
  inc := [4]R2{R2{x:0,y:1},R2{x:0,y:-1},R2{x:1,y:0},R2{x:-1,y:0}}

  var np R2
  for _,i := range inc {
    np.x = p.x+i.x
    np.y = p.y+i.y
    if CheckBounds(np, limits) && gp[np.y][np.x] == '.' {
      addToSet(n, np)
    }
  }
}

func getNextInfiniteGardenPlots(p R2, gp [][]rune, n map[R2]struct{}) {
  limits := R2{x:len(gp[0]), y:len(gp)}
  inc := [4]R2{R2{x:0,y:1},R2{x:0,y:-1},R2{x:1,y:0},R2{x:-1,y:0}}

  var np, rnp R2
  for _,i := range inc {
    rnp.x = p.x+i.x
    np.x = rnp.x % limits.x
    if np.x < 0 {
      np.x += limits.x
    }
    rnp.y = p.y+i.y
    np.y = rnp.y % limits.y
    if np.y < 0 {
      np.y += limits.y
    }
    if gp[np.y][np.x] == '.' {
      addToSet(n, rnp)
    }
  }
}

func takeStep(gp [][]rune, pl map[R2]struct{}) map[R2]struct{} {
  npl := make(map[R2]struct{})
  for p := range pl {
    getNextGardenPlots(p, gp, npl)
  }
  return npl
}

func takeStepInfinite(gp [][]rune, pl map[R2]struct{}) map[R2]struct{} {
  npl := make(map[R2]struct{})
  for p := range pl {
    getNextInfiniteGardenPlots(p, gp, npl)
  }
  return npl
}

func Day21() bool {
  b,_ := os.ReadFile("input/sample21")
  gp := getGardenPlots(string(b[:len(b)-1]))

  next := make(map[R2]struct{})
  next2 := make(map[R2]struct{})
  pos := startingPoint('S', gp)
  gp[pos.y][pos.x] = '.'
  addToSet(next, pos)
  addToSet(next2, pos)

  for i := 0; i < 64; i++ {
    next = takeStep(gp, next)
  }
  p1 := len(next)
  fmt.Printf("Part 1: %d\nPart 2: %d\n",p1,0)


  prev := 0
  for i := 0; i < 100; i++ {
    next2 = takeStepInfinite(gp, next2)
    println(len(next2)-prev)
    prev = len(next2)
  }
  println(len(next2))
  println(-100%10)


  return true
}
