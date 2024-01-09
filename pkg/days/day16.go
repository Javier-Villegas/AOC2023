package days

import (
	"os"
	"strings"
)


type Ray struct {
  position R2
  direction R2
  history map[[4]int]int
}

type R2 struct {
  x int
  y int
}

func CheckBounds(pos R2, lims R2) bool {
  return pos.x >= 0 && pos.x < lims.x && pos.y >= 0 && pos.y < lims.y
}

func inclinedReflexion(r *Ray, reverse bool) {
  if r.direction.x != 0 {
    if r.direction.x == 1 {
      r.direction.y = -1
    } else {
      r.direction.y = 1
    }
    r.direction.x = 0
  } else {
    if r.direction.y == 1 {
      r.direction.x = -1
    } else {
      r.direction.x = 1
    }
    r.direction.y = 0
  }
  if reverse {
    r.direction.x *= -1
    r.direction.y *= -1
  }

}

func propagateRay(mirrors [][]rune, energy [][]int, r Ray, duplicates []Ray, val int) []Ray{
  limits := R2{x: len(mirrors[0]), y: len(mirrors)}
  r.position.x += r.direction.x
  r.position.y += r.direction.y
  for CheckBounds(r.position, limits) {
    char := mirrors[r.position.y][r.position.x]
    energy[r.position.y][r.position.x] = val
    _,ok := r.history[[4]int{r.position.x, r.position.y, r.direction.x, r.direction.y}]
    if char != '.' && ok {
      return duplicates
    } else {
      if char != '.' {
        r.history[[4]int{r.position.x, r.position.y, r.direction.x, r.direction.y}] = 1
      }
    } 
    if char == '|' && r.direction.x != 0 {
      duplicates = append(duplicates, Ray{position: R2{x: r.position.x, y: r.position.y},
                                        direction: R2{x: 0, y: -1},
                                        history: r.history})
      r.direction.x = 0
      r.direction.y = 1
    } else if char == '-' && r.direction.y != 0 {
      duplicates = append(duplicates, Ray{position: R2{x: r.position.x, y: r.position.y},
                                        direction: R2{x: -1, y: 0},
                                        history: r.history})
      r.direction.x = 1
      r.direction.y = 0
    } else if char == '/' {
      inclinedReflexion(&r, false)
    } else if char == '\\' {
      inclinedReflexion(&r, true)
    }
    r.position.x += r.direction.x
    r.position.y += r.direction.y

  }
  return duplicates
}


func getMirrors(s string) ([][]rune, [][]int) {
  lines := strings.Split(s,"\n")

  mirrors := make([][]rune, len(lines))
  visited := make([][]int, len(lines))

  for i,l := range lines {
    mirrors[i] = []rune(l)
    visited[i] = make([]int, len(l))
  }

  return mirrors, visited
}

func countTiles(e [][]int, val int) int {
  c := 0
  for i := range e {
    for j := range e[i] {
      if e[i][j] == val {
        c++
      }
    }
  }
  return c
}

func Day16() bool {
  b,_ := os.ReadFile("input/sample16")
  input := string(b[:len(b)-1])
  
  mirrors, energy := getMirrors(input)
  energy[0][0] = 1

  // Part 1
  ray := Ray{ position: R2{x: 0, y: 0},
              direction: R2{x: 1, y: 0},
              history: make(map[[4]int]int)}
  var duplicates []Ray

  val := 1
  duplicates = propagateRay(mirrors, energy, ray, duplicates, val)
  for len(duplicates) > 0 {
    ray = duplicates[0]
    duplicates = duplicates[1:]
    duplicates = propagateRay(mirrors, energy, ray, duplicates, val)
  }

  p1 := countTiles(energy, 1)
  
  // Part 2
  val = 2
  max := p1
  var aux_count int
  //Rows 
  for i := range mirrors {
    // Rightward
    ray = Ray{position: R2{x: -1, y: i}, direction: R2{x: 1, y: 0}, history: make(map[[4]int]int)}
    duplicates = propagateRay(mirrors, energy, ray, duplicates, val)
    for len(duplicates) > 0 {
      ray = duplicates[0]
      duplicates = duplicates[1:]
      duplicates = propagateRay(mirrors, energy, ray, duplicates, val)
    }
    
    aux_count = countTiles(energy, val)
    if aux_count > max {
      max = aux_count
    }

    val++

    //Leftward
    ray = Ray{position: R2{x: len(mirrors[i]), y: i}, direction: R2{x: -1, y: 0}, history: make(map[[4]int]int)}
    duplicates = propagateRay(mirrors, energy, ray, duplicates, val)
    for len(duplicates) > 0 {
      ray = duplicates[0]
      duplicates = duplicates[1:]
      duplicates = propagateRay(mirrors, energy, ray, duplicates, val)
    }

    aux_count = countTiles(energy, val)
    if aux_count > max {
      max = aux_count
    }
    val++
  }
  for j := range mirrors {
    // Downward
    ray = Ray{position: R2{x: j, y: -1}, direction: R2{x: 0, y: 1}, history: make(map[[4]int]int)}
    duplicates = propagateRay(mirrors, energy, ray, duplicates, val)
    for len(duplicates) > 0 {
      ray = duplicates[0]
      duplicates = duplicates[1:]
      duplicates = propagateRay(mirrors, energy, ray, duplicates, val)
    }
    
    aux_count = countTiles(energy, val)
    if aux_count > max {
      max = aux_count
    }

    val++

    //Upward
    ray = Ray{position: R2{x: j, y: len(mirrors)}, direction: R2{x: 0, y: -1}, history: make(map[[4]int]int)}
    duplicates = propagateRay(mirrors, energy, ray, duplicates, val)
    for len(duplicates) > 0 {
      ray = duplicates[0]
      duplicates = duplicates[1:]
      duplicates = propagateRay(mirrors, energy, ray, duplicates, val)
    }

    aux_count = countTiles(energy, val)
    if aux_count > max {
      max = aux_count
    }
    val++

  }

  print("Part 1: ")
  println(p1)
  print("Part 2: ")
  println(max)
  

  return true
}
