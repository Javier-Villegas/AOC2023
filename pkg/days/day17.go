package days

import (
	"fmt"
	"os"
	"strings"
)

type Crucible struct {
  pos R2
  dir R2
  heatloss int
  inertia int
  history []string
}

func lesseq(c1 Crucible, c2 Crucible) bool {
  return c1.heatloss < c2.heatloss || (c1.heatloss == c2.heatloss && c1.inertia < c2.inertia)
}

func insertSort[T any](s []T, e T, less func(T,T) bool) []T {
  ns := make([]T, len(s)+1)
  for i := range s {
    if less(s[i], e) {
      ns[i] = s[i]
    } else {
      ns[i] = e
      for j := i+1; j < len(ns); j++ {
        ns[j] = s[i]
        i++
      }
      return ns
    }
  }
  ns[len(s)] = e

  return ns
}


func getIslandMap(s string) [][]int {
  lines := strings.Split(s, "\n")
  islandMap := make([][]int, len(lines))

  for i := range lines {
    islandMap[i] = make([]int, len(lines[i]))
    for j,r := range lines[i] {
      islandMap[i][j] = int(r)-int('0')
    }
  }


  return islandMap
}

func getNextCrucible(c Crucible, m [][]int, v[][]int) (Crucible,bool) {
  var nc Crucible

  return nc,true
}

func dfs(m [][]int) int{
  //limits := R2{y: len(m), x: len(m[0])}
  //inertial_limit := 3

  v:= make([][]int, len(m))
  for i := range v{
    v[i] = make([]int, len(m[0]))
  }
  
  remaining := []Crucible{
    Crucible{pos: R2{x:0,y:0}, dir: R2{x:0,y:0}, heatloss: 0, inertia: 0, history: nil},
  }

  var current, next Crucible
  var ok bool
  for len(remaining) > 0 {
    current = remaining[0]
    next,ok = getNextCrucible(current, m, v)
    if ok {
      remaining = insertSort(remaining, next, lesseq)
    } else {
      remaining = remaining[1:]
    }

  }

  



  return 0

}

func bfs(pos R2, target R2,  m [][]int) int {
  heatloss := 0
  limits := R2{y: len(m), x: len(m[0])}
  inertial_limit := 3

  visited := make([][]int, len(m))
  for i := range visited {
    visited[i] = make([]int, len(m[0]))
  }

  remaining := []Crucible{
    Crucible{pos: R2{x: pos.x, y: pos.y}, dir: R2{x: 0, y: 0}, heatloss: 0, inertia: 0, history: nil},
  }

  var current Crucible
  var new_inertia int
  for len(remaining) > 0 {
    current = remaining[0]
    remaining = remaining[1:]
    if CheckBounds(current.pos, limits) && (visited[current.pos.y][current.pos.x] == 0 ||visited[current.pos.y][current.pos.x] >= current.heatloss) {
      visited[current.pos.y][current.pos.x] = current.heatloss
      if current.pos.x == target.x && current.pos.y == target.y {
        fmt.Printf("%v\n", current.history)
        return current.heatloss
      }

      
      // Rightward
      if current.dir.x == 0 || (current.dir.x == 1 && current.inertia < inertial_limit) {
        if current.dir.x == 0 {
          new_inertia = 1
        } else {
          new_inertia = current.inertia + 1
        }
        new_pos := R2{x: current.pos.x+1, y: current.pos.y}
        if CheckBounds(new_pos, limits) && (visited[new_pos.y][new_pos.x] == 0 || visited[new_pos.y][new_pos.x] > current.heatloss+m[new_pos.y][new_pos.x]) {
          remaining = insertSort(remaining,
                              Crucible{ pos: new_pos,
                              dir: R2{x: 1, y: 0}, heatloss: current.heatloss+m[new_pos.y][new_pos.x], inertia: new_inertia,
                              history: append(current.history, "R")},
                              lesseq)
        }
      
      // Leftward 
      }
      if current.dir.x == 0 || (current.dir.x == -1 && current.inertia < inertial_limit) {
        if current.dir.x == 0 {
          new_inertia = 1
        } else {
          new_inertia = current.inertia + 1
        }
        new_pos := R2{x: current.pos.x-1, y: current.pos.y}
        if CheckBounds(new_pos, limits) && (visited[new_pos.y][new_pos.x] == 0 || visited[new_pos.y][new_pos.x] > current.heatloss+m[new_pos.y][new_pos.x]) {
          remaining = insertSort(remaining,
                              Crucible{ pos: new_pos,
                                        dir: R2{x: -1, y: 0}, heatloss: current.heatloss+m[new_pos.y][new_pos.x], inertia: new_inertia, 
                              history: append(current.history, "L")},
                              lesseq)
          }
      
      // Downward
      }
      if current.dir.y == 0 || (current.dir.y == 1 && current.inertia < inertial_limit) {
        if current.dir.y == 0 {
          new_inertia = 1
        } else {
          new_inertia = current.inertia + 1
        }
        new_pos := R2{x: current.pos.x, y: current.pos.y+1}
        if CheckBounds(new_pos, limits) && (visited[new_pos.y][new_pos.x] == 0 || visited[new_pos.y][new_pos.x] > current.heatloss+m[new_pos.y][new_pos.x]) {
          remaining = insertSort(remaining,
                              Crucible{ pos: new_pos,
                                        dir: R2{x: 0, y: 1}, heatloss: current.heatloss+m[new_pos.y][new_pos.x], inertia: new_inertia,
                                        history: append(current.history, "D")},
                              lesseq)
        }
      
      // Upward
      }
      if current.dir.y == 0 || (current.dir.y == -1 && current.inertia < inertial_limit) { 
        if current.dir.y == 0 {
          new_inertia = 1
        } else {
          new_inertia = current.inertia + 1
        }
        new_pos := R2{x: current.pos.x, y: current.pos.y-1}
        if CheckBounds(new_pos, limits) && (visited[new_pos.y][new_pos.x] == 0 || visited[new_pos.y][new_pos.x] > current.heatloss+m[new_pos.y][new_pos.x]) {
          remaining = insertSort(remaining,
                              Crucible{ pos: new_pos,
                                        dir: R2{x: 0, y: -1}, heatloss: current.heatloss+m[new_pos.y][new_pos.x], inertia: new_inertia,
                                        history: append(current.history, "U")},
                              lesseq)
        }
      } 
    }
  }
  


  return heatloss
}

func Day17() bool {
  b,_ := os.ReadFile("input/sample17")
  island := getIslandMap(string(b[:len(b)-1]))

  for i := range island {
    for j := range island[i] {
      print(island[i][j])
    }
    println()
  }

  p1 := bfs(R2{x: 0, y: 0}, R2{x: len(island[0])-1, y: len(island)-1}, island)
  println(p1)
  fmt.Printf("")

  return true
}
