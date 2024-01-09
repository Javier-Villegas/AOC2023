package days

import (
  "os"
  "strings"
)

func getPattern(s string) [][]rune {
  lines := strings.Split(s, "\n")
  pattern := make([][]rune, len(lines))
  for l := range lines {
    pattern[l] = []rune(lines[l])
  }
  return pattern
}

func getHorizontalLine(p [][]rune) int {
  lor := -1
  for l := 0; l < len(p)-1; l++ {
    found := true
    i := l
    ir := l+1
    
    out:
    for  i >= 0 && ir < len(p) {
      for j := 0; j < len(p[0]); j++ {
        if p[i][j] != p[ir][j] {
          found = false
          break out
        }
      }
      
      i--
      ir++
    }

    if found {
      lor = l
      break
    }
  }

  return lor+1
}

func getVerticalLine(p [][]rune) int {
  lor := -1
  for l := 0; l < len(p[0])-1; l++ {
    found := true
    j := l
    jr := l+1
    
    out:
    for  j >= 0 && jr < len(p[0]) {
      for i := 0; i < len(p); i++ {
        if p[i][j] != p[i][jr] {
          found = false
          break out
        }
      }
      
      j--
      jr++
    }

    if found {
      lor = l
      break
    }
  }

  return lor+1
}

func fixHorizontalLine(p [][]rune) int {
  lor := -1
  for l := 0; l < len(p)-1; l++ {
    count := 0
    found := true
    i := l
    ir := l+1
    
    out:
    for  i >= 0 && ir < len(p) {
      for j := 0; j < len(p[0]); j++ {
        if p[i][j] != p[ir][j] {
          if count == 0 {
            count++
          } else {
            found = false
            break out
          }
        }
      }
      
      i--
      ir++
    }

    if found && count == 1 {
      lor = l
      break
    }
  }

  return lor+1
}

func fixVerticalLine(p [][]rune) int {
  lor := -1
  for l := 0; l < len(p[0])-1; l++ {
    count := 0
    found := true
    j := l
    jr := l+1
    
    out:
    for  j >= 0 && jr < len(p[0]) {
      for i := 0; i < len(p); i++ {
        if p[i][j] != p[i][jr] {
          if count == 0 {
            count += 1
          } else {
            found = false
            break out
          }
        }
      }
      
      j--
      jr++
    }

    if found && count == 1{
      lor = l
      break
    }
  }

  return lor+1
}

func Day13() bool {
  b,_ := os.ReadFile("input/day13.txt")
  input := string(b[:len(b)-1])
  patterns_str := strings.Split(input, "\n\n")

  p1 := 0
  p2 := 0
  for p := range patterns_str {
    pattern := getPattern(patterns_str[p])
    hlor := getHorizontalLine(pattern)
    vlor := getVerticalLine(pattern)

    p1 += vlor + 100 * hlor

    vlor = fixVerticalLine(pattern)
    hlor = fixHorizontalLine(pattern)
    
    p2 += vlor + 100 * hlor
  }

  print("Part 1: ")
  println(p1)
  print("Part 2: ")
  println(p2)

  return true
}
