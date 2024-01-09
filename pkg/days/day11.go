package days

import (
  "os"
  "strings"
  "fmt"
  "slices"
  "math"
)


type pair struct {
  a pos
  b pos
}

func fill_void(universe [][]rune) [][]rune {
  for i := 0; i < len(universe); i++ {
    add_row := true
    for j := 0; j < len(universe[i]); j++ {
      if universe[i][j] == '#' {
        add_row = false
      }
    }
    if add_row {
      i++
      universe = slices.Insert(universe, i, make([]rune, len(universe[i])))
      for j := range universe[i] {
        universe[i][j] = '.'
      }
    }
  }
  for j := 0; j < len(universe[0]); j++ {
    add_col := true
    for i := 0; i < len(universe); i++ {
      if universe[i][j] == '#' {
        add_col = false
      }
    }
    if add_col {
      j++
      for i := range universe {
        universe[i] = slices.Insert(universe[i],j,'.')
      }
    }
  }

  return universe
}

func search_pairs(universe [][]rune) []pair {
  var pairs []pair
  var galaxies []pos

  for i := 0; i < len(universe); i++ {
    for j := 0; j < len(universe[i]); j++ {
      if universe[i][j] == '#' {
        galaxies = append(galaxies, pos{y: i, x: j, w: 0})
      }
    }
  }

  for i := 0; i < len(galaxies); i++ {
    for j := i+1; j < len(galaxies); j++ {
      pairs = append(pairs, pair{a: galaxies[i], b: galaxies[j]})
    }
  }

  return pairs
}

func count_void(universe [][]rune) ([]int,[]int){
  var ver []int
  var hor []int

  for i := 0; i < len(universe); i++ {
    add_row := true
    for j := 0; j < len(universe[i]); j++ {
      if universe[i][j] == '#' {
        add_row = false
      }
    }
    if add_row {
      hor = append(hor, i)
    }
  }
  for j := 0; j < len(universe[0]); j++ {
    add_col := true
    for i := 0; i < len(universe); i++ {
      if universe[i][j] == '#' {
        add_col = false
      }
    }
    if add_col {
      ver = append(ver, j)
    }
  }

  return ver,hor
}


func get_distance(p pair, ver_exp []int, hor_exp []int, fac_exp int) int {
  var min_x, max_x, min_y, max_y int
  if p.a.x > p.b.x {
    min_x = p.b.x
    max_x = p.a.x
  } else {
    min_x = p.a.x
    max_x = p.b.x
  }

  if p.a.y > p.b.y {
    min_y = p.b.y
    max_y = p.a.y
  } else {
    min_y = p.a.y
    max_y = p.b.y
  }

  extra_ver := 0
  for _,v := range ver_exp {
    if min_x < v && max_x > v {
      extra_ver += 1
    }
  }
  
  extra_hor := 0
  for _,h := range hor_exp {
    if min_y < h && max_y > h {
      extra_hor += 1
    }
  }

  base_distance := int(math.Abs(float64(p.a.y-p.b.y))+math.Abs(float64(p.a.x-p.b.x)))
  return base_distance  + (fac_exp - 1) * (extra_hor + extra_ver )
}

func Day11() bool {
  b,_ := os.ReadFile("input/day11.txt")
  input := string(b[:len(b)-1])
  rows := strings.Split(input, "\n")

  universe := make([][]rune, len(rows))
  for i,row := range rows {
    universe[i] = make([]rune, len(row))
    for j,r := range row {
      universe[i][j] = r
    }
  }


  //universe = fill_void(universe)
  ver,hor := count_void(universe)

  pairs := search_pairs(universe)
  p1 := 0
  p2 := 0
  for _,p := range pairs {
    dist1 := get_distance(p, ver, hor, 2)
    dist2 := get_distance(p, ver, hor, 1000000)
    p1 += dist1
    p2 += dist2
  }

  fmt.Printf("Part 1: %d\nPart 2: %d\n", p1, p2)
  return true
}
