package days 

import (
  "os"
  "strings"
  "strconv"
)

func get_code(cs []string) []int {
  c := make([]int, len(cs))
  for i := range c {
    c[i],_ = strconv.Atoi(cs[i])
  }
  return c
}

func get_input(line string) ([]rune, []int) {
  s := strings.Split(line, " ")
  return []rune(s[0]),get_code(strings.Split(s[1], ","))
}

func mapsClear[M ~map[K]V, K comparable, V any](m M) {
	for k := range m {
		delete(m, k)
	}
}

func repeatSlice[T any](s []T, num int) []T{
  new_slice := make([]T,num*len(s))
  for i := 0; i < num; i++ {
    for j := 0; j < len(s); j++ {
      new_slice[len(s)*i+j] = s[j]
    }
  }
  return new_slice
}

func count_possible(s []rune, c []int) int {
  arrangements := 0
  cstates := map[[3]int]int{{0,0,0}:1}
  new_states := map[[3]int]int{}



  for s_index := 0; s_index < len(s); s_index++ {
    s_rune := s[s_index]
    for state, num := range cstates {
      c_index, c_int, expdot := state[0], state[1], state[2]
      switch {
      case (s_rune == '#' || s_rune == '?') && c_index < len(c) && expdot == 0:
        if s_rune == '?' && c_int == 0 {
          new_states[[3]int{c_index, c_int, expdot}] += num
        }
        c_int++
        if c_int == c[c_index] {
          c_index, c_int, expdot = c_index+1, 0, 1
        }
        new_states[[3]int{c_index, c_int, expdot}] += num
      case (s_rune == '.' || s_rune == '?') && c_int == 0:
        expdot = 0
        new_states[[3]int{c_index, c_int, expdot}] += num
      }
    }
    cstates, new_states = new_states, cstates
    mapsClear(new_states)
  }

  for state, value := range cstates {
    if state[0] == len(c) {
      arrangements += value
    }
  }
  return arrangements
}

func Day12_2() bool {
  b,_ := os.ReadFile("input/day12.txt")
  lines := strings.Split(string(b[:len(b)-1]), "\n")

  p1 := 0
  p2 := 0
  for l := range lines {
    s,c := get_input(lines[l])
    str := string(s)
    p1 += count_possible(s, c)
    p2 += count_possible([]rune(str+"?"+str+"?"+str+"?"+str+"?"+str), repeatSlice[int](c, 5))
  }
  println(p1)
  println(p2)

  return true
}
