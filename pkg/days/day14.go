package days

import (
	"os"
	"strings"
)

func getRocks(s string) [][]rune {
  lines := strings.Split(s, "\n")
  rocks := make([][]rune, len(lines))
  for l := range lines {
    rocks[l] = []rune(lines[l])
  }

  return rocks
}

func tiltNorth(r [][]rune) {
  for j := 0; j < len(r[0]); j++ {
    anchor := 0
    for r[anchor][j] != '.' {
      anchor++
    }
    for i := anchor+1; i < len(r); i++ {
      if r[i][j] == 'O' {
        r[anchor][j] = 'O'
        r[i][j] = '.'
        anchor++
      } else if r[i][j] == '#' {
        anchor = i+1
        for anchor < len(r) && r[anchor][j] != '.' {
          anchor++
        }
        i = anchor
      }
    }
  }
}

func tiltSouth(r [][]rune) {
  for j := 0; j < len(r[0]); j++ {
    anchor := len(r)-1
    for r[anchor][j] != '.' {
      anchor--
    }
    for i := anchor-1; i >= 0; i-- {
      if r[i][j] == 'O' {
        r[anchor][j] = 'O'
        r[i][j] = '.'
        anchor--
      } else if r[i][j] == '#' {
        anchor = i-1
        for anchor > 0 && r[anchor][j] != '.' {
          anchor--
        }
        i = anchor
      }
    }
  }
}

func tiltWest(r [][]rune) {
  for i := 0; i < len(r); i++ {
    anchor := 0
    for r[i][anchor] != '.' {
      anchor++
    }
    for j := anchor+1; j < len(r[0]); j++ {
      if r[i][j] == 'O' {
        r[i][anchor] = 'O'
        r[i][j] = '.'
        anchor++
      } else if r[i][j] == '#' {
        anchor = j+1
        for anchor < len(r[0]) && r[i][anchor] != '.' {
          anchor++
        }
        j = anchor
      }
    }
  }
}


func tiltEast(r [][]rune) {
  for i := 0; i < len(r); i++ {
    anchor := len(r[0])-1
    for r[i][anchor] != '.' {
      anchor--
    }
    for j := anchor-1; j >= 0; j-- {
      if r[i][j] == 'O' {
        r[i][anchor] = 'O'
        r[i][j] = '.'
        anchor--
      } else if r[i][j] == '#' {
        anchor = j-1
        for anchor > 0 && r[i][anchor] != '.' {
          anchor--
        }
        j = anchor
      }
    }
  }
}

func cycle(r [][]rune) {
  tiltNorth(r)
  tiltWest(r)
  tiltSouth(r)
  tiltEast(r)
}

func computeLoad(r [][]rune) int {
  load := 0
  for i := 0; i < len(r); i++ {
    for j := 0; j < len(r[0]); j++ {
      if r[i][j] == 'O' {
        load += len(r)-i
      }
    }
  }
  return load
}
func hashRocks(r [][]rune) string {
  str := ""
  for row := range r {
    str += string(r[row])
  }
  return str
}

func Day14() bool {
  b,_ := os.ReadFile("input/day14.txt")
  rocks := getRocks(string(b[:len(b)-1]))

  tiltNorth(rocks)
  print("Part 1: ")
  println(computeLoad(rocks))

  tiltWest(rocks)
  tiltSouth(rocks)
  tiltEast(rocks)

  history := make(map[string][2]int)
  history[hashRocks(rocks)] = [2]int{computeLoad(rocks), 1}
  var period, base int
  i := 2
  for {
    cycle(rocks)
    
    hash := hashRocks(rocks)
    val, ok := history[hash]
    if ok {
      base = val[1]
      period = i-val[1]
      break
    }
    history[hash] = [2]int{computeLoad(rocks), i}
    i++
  }
  
  entry := (1000000000 - base) % period + base
  var p2 int
  for _,v := range history {
    if v[1] == entry {
      p2 = v[0]
      break
    }
  }


  print("Part 2: ")
  println(p2)

  return true
}
