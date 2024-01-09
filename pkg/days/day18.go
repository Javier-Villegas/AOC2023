package days

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type digPlan struct {
  dir rune
  dist int
  rgb int
}

type holeEdge struct {
  R2
  r int
}

var hex_to_dec = map[rune]int{'0':0,'1':1,'2':2,'3':3,'4':4,'5':5,'6':6,'7':7,'8':8,'9':9,'a':10,'b':11,'c':12,'d':13,'e':14,'f':15}

func getRGB(color string) int {
  res := 0
  for i,r := range color {
    res += int(math.Pow(2,float64(len(color)-1-i)))*hex_to_dec[r]
  }


  return res
}

func stoi(s string) int {
  i,_ := strconv.Atoi(s)
  return i
}

func less(r1 R2, r2 R2) bool {
  return r1.y < r2.y || (r1.y == r2.y && r1.x < r2.x)
}

func sortTrenchMap(m map[R2]int) []holeEdge {
  var sm []holeEdge
  for k,v := range m {
    sm = append(sm, holeEdge{k,v})
  }

  sort.Slice(sm, func(i, j int) bool {return less(sm[i].R2, sm[j].R2)})

  return sm
}

func colorTrench(t [][]int) {
  var l []R2
  limits := R2{x:len(t[0]),y:len(t)}
  for j := range t[0] {
    if t[0][j] == 1 && t[1][j] == 0 {
      t[1][j] = 2
      l = append(l, R2{x:j,y:1})
      break
    }
  }
  var pos R2
  for len(l) > 0 {
    pos = l[0]
    l = l[1:]
    if CheckBounds(R2{x:pos.x+1,y:pos.y},limits) && t[pos.y][pos.x+1] == 0 {
      t[pos.y][pos.x+1] = 2
      l = append(l, R2{x:pos.x+1,y:pos.y})
    }
    if CheckBounds(R2{x:pos.x-1,y:pos.y},limits) && t[pos.y][pos.x-1] == 0 {
      t[pos.y][pos.x-1] = 2
      l = append(l, R2{x:pos.x-1,y:pos.y})
    }
    if CheckBounds(R2{x:pos.x,y:pos.y+1},limits) && t[pos.y+1][pos.x] == 0{
      t[pos.y+1][pos.x] = 2
      l = append(l, R2{x:pos.x,y:pos.y+1})
    }
    if CheckBounds(R2{x:pos.x,y:pos.y-1},limits) && t[pos.y-1][pos.x] == 0{
      t[pos.y-1][pos.x] = 2
      l = append(l, R2{x:pos.x,y:pos.y-1})
    }
  }

}

func Day18() bool {
  b,_ := os.ReadFile("input/day18.txt")
  lines := strings.Split(string(b[:len(b)-1]), "\n")

  digplan := make([]digPlan, len(lines))
  trench := make(map[R2]int)

  trench2 := make(map[R2]int)


  r := regexp.MustCompile("[a-f0-9]{6}")

  var minx,miny,maxx,maxy int

  current := R2{x:0,y:0}
  dir := R2{x:0,y:0}
  var inst []string
  for l := range lines {
    inst = strings.Split(lines[l], " ")
    digplan[l] = digPlan{dir: rune(inst[0][0]), dist: stoi(inst[1]), rgb: getRGB(r.FindAllString(inst[2],-1)[0])}

    if digplan[l].dir == 'R' {
      dir.x = 1
      dir.y = 0
    } else if digplan[l].dir == 'L' {
      dir.x = -1
      dir.y = 0
    } else if digplan[l].dir == 'D' {
      dir.x = 0
      dir.y = 1
    } else {
      dir.x = 0
      dir.y = -1
    }
    
    //trench2[current] = 1
    if dir.y == 0 {
      trench2[current] = 1
      trench2[R2{x: current.x+dir.x*digplan[l].dist, y: current.y+dir.y*digplan[l].dist}] = 1
    } 
    for i := 0; i < digplan[l].dist; i++ {
      current.x += dir.x
      current.y += dir.y
      trench[current] = digplan[l].rgb
      if dir.x == 0 {
        trench2[current] = 0

      }
    }
    if current.x < minx {
      minx = current.x
    } else if current.x > maxx {
      maxx = current.x
    }
    if current.y < miny {
      miny = current.y
    } else if current.y > maxy {
      maxy = current.y
    }
  }

  trench_arr := make([][]int, maxy-miny+1)
  for i := range trench_arr {
    trench_arr[i] = make([]int, maxx-minx+1)
  }
  for k,_ := range trench {
    trench_arr[k.y-miny][k.x-minx] = 1
  }

  colorTrench(trench_arr)
  p1 := 0
  for i := range trench_arr {
    for j := range trench_arr[i] {
      print(trench_arr[i][j])
      if trench_arr[i][j] > 0 {
        p1++
      }
    }
    println()
  }
  println(len(trench))
  println(p1)

  slicedTrench := sortTrenchMap(trench2)
  p2 := 0
  prev := 0
  for i := 0; i < len(slicedTrench); i++ {
    if i < len(slicedTrench)-1 && slicedTrench[i].y == slicedTrench[i+1].y {
      p2 += (slicedTrench[i+1].x - slicedTrench[i].x + 1)
      if slicedTrench[i+1].r == 0 {
        if prev == 1 {
          p2--
          prev = 0
        }
        i++
      } else {
        if prev == 1 {
          p2 -= 2
          prev = 0
        } else {
          prev = 1
        }
      } 
    } else if i < len(slicedTrench)-1 && slicedTrench[i].y != slicedTrench[i+1].y {
      prev = 0
    } 
  }
  println(prev)
  println(p2)
  fmt.Printf("")

  return true
}
