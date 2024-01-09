package days

import (
  "os"
  "strings"
  "regexp"
  "fmt"
)

type node struct {
  left string
  right string
}

func GCD(a int, b int) int {
  for b != 0 {
    t := b
    b = a%b
    a = t
  }
  return a
}

func LCM(a int, b int, integers []int) int {
  result := a * b / GCD(a,b)

  for i := 0; i < len(integers); i++ {
    result = LCM(result, integers[i], nil)
  }
  return result
}

func Day08() bool {
  b,_ := os.ReadFile("input/day08.txt")
  input := string(b)
  parts := strings.Split(input, "\n\n")
  inst := parts[0]
  
  r := regexp.MustCompile("[0-9,A-Z]{3}")
  paths := make(map[string]node)
  for _,m := range strings.Split(parts[1][:len(parts[1])-1], "\n") {
    res := r.FindAllString(m,-1)
    paths[res[0]] = node{res[1],res[2]}
  }

  current := "AAA"
  inst_num := len(inst)
  p1 := 0
  i := 0
  for current != "ZZZ" {
    p1 += 1
    if inst[i] == 'L' {
      current = paths[current].left
    } else {
      current = paths[current].right
    }
    i += 1
    if i >= inst_num {i = 0}
  }

  p2 := 0
  i = 0
  var starting []string
  for k,_ := range paths {
    if k[2] == 'A' {starting = append(starting, k)}
  }

  var lcm_slice []int
  for len(starting) > 0 {
    p2 += 1
    if inst[i] == 'L' {
      for i := 0; i < len(starting); i++ {
        starting[i] = paths[starting[i]].left
        if starting[i][2] == 'Z' {
          starting[i] = starting[len(starting)-1]
          starting = starting[:len(starting)-1]
          lcm_slice = append(lcm_slice, p2)
          i -= 1
        }
      }
    } else {
      for i := 0; i < len(starting); i++ {
        starting[i] = paths[starting[i]].right
        if starting[i][2] == 'Z' {
          starting[i] = starting[len(starting)-1]
          starting = starting[:len(starting)-1]
          lcm_slice = append(lcm_slice, p2)
          i -= 1
        }
      }
    }

    i += 1
    if i >= inst_num {i = 0}
  }
  fmt.Printf("Part 1: %d\nPart 2: %d\n",p1,LCM(1,1,lcm_slice))
  return true
}
