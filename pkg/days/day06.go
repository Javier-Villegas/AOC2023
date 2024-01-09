package days

import (
  "os"
  "regexp"
  "strings"
  //"fmt"
  "strconv"
)

func GetDist(t_load int, t_max int, distance int) int{
  if t_load*(t_max-t_load) > distance {
    return 1
  } else {
    return 0
  }
}

func Day06() bool {
  b,_ := os.ReadFile("input/day06.txt")
  input := string(b)
  input = input[:len(input)-1]
  lines := strings.Split(input, "\n")
  

  r := regexp.MustCompile("[0-9]+")

  time := r.FindAllString(lines[0],-1)
  dist := r.FindAllString(lines[1],-1)

  p1 := 1
  for i := range time {
    valid := 0
    t_max,_ := strconv.Atoi(time[i])
    distance,_ := strconv.Atoi(dist[i])
    for j := 0; j < t_max; j++ {
      valid += GetDist(j, t_max, distance)
    }
    p1 *= valid
  }

  t_max,_ := strconv.Atoi(strings.Join(time,""))
  distance,_ := strconv.Atoi(strings.Join(dist,""))
  t0 := 1
  t1 := t_max
  for GetDist(t0, t_max, distance) == 0 { t0++}
  for GetDist(t1, t_max, distance) == 0 { t1--}


  print("Part 1: ")
  println(p1)
  print("Part 2: ")
  println(t1-t0+1)
  return true
}
