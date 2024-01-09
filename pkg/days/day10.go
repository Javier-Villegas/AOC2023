package days

import (
	"os"
	"strings"
  "fmt"
  "slices"
)

type pos struct {
  y int
  x int
  w int
}

func insertSorted(arr []pos, el pos) []pos {
  i := 0
  for  i < len(arr) && arr[i].w < el.w {i++}
  return slices.Insert(arr, i, el)
}

func connectsNorth(el rune) bool {
  return (el == '|' || el == 'F' || el == '7')
}

func connectsSouth(el rune) bool {
  return (el == '|' || el == 'J' || el == 'L')
}

func connectsWest(el rune) bool {
  return (el == '-' || el == 'F' || el == 'L')
}

func connectsEast(el rune) bool {
  return (el == '-' || el == 'J' || el == '7')
}

func getNext(pipes [][]rune,visited [][]int, p pos) (pos,bool) {
  visited[p.y][p.x] = p.w
  var n pos
  ctrl := false
  if pipes[p.y][p.x] == '|' {
    if p.y > 0 && visited[p.y-1][p.x] < 0 && connectsNorth(pipes[p.y-1][p.x]){
      n = pos{y: p.y-1, x: p.x, w: p.w+1}
      ctrl = true
    } else if p.y < len(pipes)-1 && visited[p.y+1][p.x] < 0 && connectsSouth(pipes[p.y+1][p.x]){
      n = pos{y: p.y+1 ,x: p.x ,w: p.w+1}
      ctrl = true
    }
  } else if pipes[p.y][p.x] == '-' {
    if p.x > 0 && visited[p.y][p.x-1] < 0 && connectsWest(pipes[p.y][p.x-1]){
      n = pos{y: p.y, x: p.x-1, w: p.w+1}
      ctrl = true
    } else if p.x < len(pipes[p.y])-1 && visited[p.y][p.x+1] < 0 && connectsEast(pipes[p.y][p.x+1]){
      n = pos{y: p.y, x: p.x+1, w: p.w+1}
      ctrl = true
    }
  } else if pipes[p.y][p.x] == 'L' {
    if p.y > 0 && visited[p.y-1][p.x] < 0 && connectsNorth(pipes[p.y-1][p.x]){
      n = pos{y: p.y-1, x: p.x, w: p.w+1}
      ctrl = true
    } else if p.x < len(pipes[p.y])-1 && visited[p.y][p.x+1] < 0 && connectsEast(pipes[p.y][p.x+1]){
      n = pos{y: p.y, x:p.x+1, w: p.w+1}
      ctrl = true
    }
  } else if pipes[p.y][p.x] == 'J' {
    if p.y > 0 && visited[p.y-1][p.x] < 0 && connectsNorth(pipes[p.y-1][p.x]){
      n = pos{y: p.y-1, x: p.x, w: p.w+1}
      ctrl = true
    } else if p.x > 0 && visited[p.y][p.x-1] < 0 && connectsWest(pipes[p.y][p.x-1]){
      n = pos{y: p.y, x: p.x-1, w: p.w+1}
      ctrl = true
    }
  } else if pipes[p.y][p.x] == '7' {
    if p.y < len(pipes)-1 && visited[p.y+1][p.x] < 0 && connectsSouth(pipes[p.y+1][p.x]){
      n = pos{y: p.y+1,x: p.x, w: p.w+1}
      ctrl = true
    } else if p.x > 0  && visited[p.y][p.x-1] < 0 && connectsWest(pipes[p.y][p.x-1]){
      n = pos{y: p.y, x: p.x-1, w: p.w+1}
      ctrl = true
    }
  } else if pipes[p.y][p.x] == 'F' {
    if p.y < len(pipes)-1  && visited[p.y+1][p.x] < 0 && connectsSouth(pipes[p.y+1][p.x]){
      n = pos{y: p.y+1, x: p.x, w: p.w+1}
      ctrl = true
    } else if p.x < len(pipes[p.y])-1 && visited[p.y][p.x+1] < 0 && connectsEast(pipes[p.y][p.x+1]){
      n = pos{y: p.y, x: p.x+1, w: p.w+1}
      ctrl = true
    }
  }
  return n,ctrl
}

func getDirection(p pos, prev pos) int {
  // West-East
  if p.y == prev.y {
    return (p.x - prev.x)*2 // 2 == East; -2 == West
  } else {
    return (p.y - prev.y)   // 1 == South; -1 == North
  }

}
// -2 == Left
// -3 == Right
func colorSides(pipes [][]rune, visited[][]int, p pos, dir int) {
  dir_offset := 0
  if pipes[p.y][p.x] == '|' {
    if dir == 1 {
      dir_offset = 1
    }
    if p.x > 0  && visited[p.y][p.x-1] == -1 { visited[p.y][p.x-1] = -2 -dir_offset}
    if p.x < len(pipes[p.y])-1  && visited[p.y][p.x+1] == -1 { visited[p.y][p.x+1] = -3 +dir_offset}
  } else if pipes[p.y][p.x] == '-' {
    if dir == -2 {
      dir_offset = 1
    }
    if p.y > 0  && visited[p.y-1][p.x] == -1 { visited[p.y-1][p.x] = -2 -dir_offset}
    if p.y > len(pipes)-1 && visited[p.y+1][p.x] == -1 { visited[p.y+1][p.x] = -3 +dir_offset}
  } else if pipes[p.y][p.x] == 'L' {
    if dir == 1 {
      dir_offset = 1
    }
    if p.y < len(pipes)-1  && visited[p.y+1][p.x] == -1 { visited[p.y+1][p.x] = -2 -dir_offset}
    if p.x > 0  && visited[p.y][p.x-1] == -1 { visited[p.y][p.x-1] = -2 -dir_offset}
    if p.y < len(pipes)-1 && p.x > 0 && visited[p.y+1][p.x-1] == -1 { visited[p.y+1][p.x-1] = -2 -dir_offset}
    if p.y > 0 && p.x < len(pipes[p.y])-1 && visited[p.y-1][p.x+1] == -1 { visited[p.y-1][p.x+1] = -3 +dir_offset}
  } else if pipes[p.y][p.x] == 'J' {
    if dir == 1 {
      dir_offset = 1
    }
    if p.y < len(pipes)-1 && visited[p.y+1][p.x] == -1 { visited[p.y+1][p.x] = -3 +dir_offset}
    if p.x < len(pipes[p.y])-1 && visited[p.y][p.x+1] == -1 { visited[p.y][p.x+1] = -3 +dir_offset}
    if p.y < len(pipes)-1 && p.x < len(pipes[p.y])-1 && visited[p.y+1][p.x+1] == -1 { visited[p.y+1][p.x+1] = -3 +dir_offset}
    if p.y > 0 && p.x > 0 && visited[p.y-1][p.x-1] == -1 { visited[p.y-1][p.x-1] = -2 -dir_offset}
  } else if pipes[p.y][p.x] == '7' {
    if dir == 2 {
      dir_offset = 1
    }
    if p.x < len(pipes[p.y]) && visited[p.y][p.x+1] == -1 { visited[p.y][p.x+1] = -3 +dir_offset}
    if p.y > 0 && visited[p.y-1][p.x] == -1 { visited[p.y-1][p.x] = -3 +dir_offset}
    if p.y > 0 && p.x < len(pipes[p.y])-1 && visited[p.y-1][p.x+1] == -1 { visited[p.y-1][p.x+1] = -3 +dir_offset}
    if p.y < len(pipes)-1 && p.x > 0 && visited[p.y+1][p.x-1] == -1 { visited[p.y+1][p.x-1] = -2 -dir_offset}
  } else if pipes[p.y][p.x] == 'F' {
    if dir == -2 {
      dir_offset = 1
    }
    if p.y > 0 && visited[p.y-1][p.x] == -1 { visited[p.y-1][p.x] = -2 -dir_offset}
    if p.x > 0 && visited[p.y][p.x-1] == -1 { visited[p.y][p.x-1] = -2 -dir_offset}
    if p.y > 0 && p.x > 0 && visited[p.y-1][p.x-1] == -1 { visited[p.y-1][p.x-1] = -2 -dir_offset}
    if p.y < len(pipes)-1 && p.x < len(pipes[p.y])-1 && visited[p.y+1][p.x+1] == -1 { visited[p.y+1][p.x+1] = -3 +dir_offset}
  }
}

func completeColoring(visited [][]int) (int,int) {
  l := 0
  r := 0

  var i int
  var j int
  var check []pos
  for i = 0; i < len(visited); i++ {
    for j = 0; j < len(visited[i]); j++ {
      if visited[i][j] < -1 {
        check = append(check, pos{y: i, x: j, w: visited[i][j]})
      }
    }
  }
  var next pos
  for len(check) > 0 {
    next = check[0]
    check = check[1:]
    if next.w == -2 {
      l += 1
    } else {
      r += 1
    }
    if next.y > 0 && visited[next.y-1][next.x] == -1 {
      visited[next.y-1][next.x] = next.w
      check = append(check, pos{y: next.y-1, x: next.x, w: next.w})
    }
    if next.y < len(visited)-1 && visited[next.y+1][next.x] == -1 {
      visited[next.y+1][next.x] = next.w
      check = append(check, pos{y: next.y+1, x: next.x, w: next.w})
    }
    if next.x > 0 && visited[next.y][next.x-1] == -1 {
      visited[next.y][next.x-1] = next.w
      check = append(check, pos{y: next.y, x: next.x-1, w: next.w})
    }
    if next.x < len(visited[next.y])-1 && visited[next.y][next.x+1] == -1 {
      visited[next.y][next.x+1] = next.w
      check = append(check, pos{y: next.y, x: next.x+1, w: next.w})
    }
  }
  return l,r
}


func Day10() bool {
  b,_ := os.ReadFile("input/day10.txt")
  input := string(b)
  lines := strings.Split(input[:len(input)-1],"\n") 

  pipes := make([][]rune,len(lines))
  for i,p := range lines {
    pipes[i] = make([]rune,len(p))
    for j,r := range p {
      pipes[i][j] = r
    }
  }

  visited := make([][]int,len(pipes))
  for i := range visited {
    visited[i] = make([]int,len(pipes[i]))
    for j := range pipes[i] {
      visited[i][j] = -1
    }
  }

  var next []pos
  var prev pos
  for i := 0; i < len(pipes); i++ {
    for j := 0; j < len(pipes[i]); j++ {
      if pipes[i][j] == 'S' {
        visited[i][j] = 0
        prev = pos{y: i, x: i, w: 0}
        if i > 0 && connectsNorth(pipes[i-1][j]){
          next = append(next, pos{y: i-1, x: j, w: 1})
        } else if i < len(pipes)-1 && connectsSouth(pipes[i+1][j]){
          next = append(next, pos{y: i+1, x: j, w: 1})
        } else if j > 0 && connectsWest(pipes[i][j-1]) {
          next = append(next, pos{y: i, x: j-1, w: 1})
        } else if j < len(pipes[i])-1 && connectsEast(pipes[i][j+1]) {
          next = append(next, pos{y: i, x: j+1, w: 1})
        }

        j = len(pipes[i])
        i = len(pipes)
        break
      }
    }
  }
  
  max := 0
  var current pos
  for len(next) > 0 {
    current = next[0]
    next = next[1:]
    next_pos, ctrl := getNext(pipes, visited, current)
    dir := getDirection(current, prev)
    colorSides(pipes, visited, current, dir)
    if ctrl {
      if next_pos.w > max {max = next_pos.w}
      next = insertSorted(next, next_pos)
    }
    prev = current
  }

  l,r := completeColoring(visited)


  fmt.Printf("Left: %d Right: %d\n", l, r)
  fmt.Printf("Part 1: %d\nPart 2: %d\n", max/2+max%2, l)

  return true
}
