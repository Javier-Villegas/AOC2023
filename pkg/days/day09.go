package days

import (
  "os"
  "strings"
  "regexp"
  "strconv"
  "fmt"
)

func DifferentialExtrapolation(seq []int) (int,int) {
 diff_seq := make([]int,len(seq)-1)
  acc := 0
  for i := 0; i < len(seq)-1; i++ {
    diff_seq[i] = seq[i+1]-seq[i]
    acc += diff_seq[i]
  }
  if acc == 0 {
    next := seq[len(seq)-1]
    prev := seq[0]
    return next,prev
  } else {
    n,p := DifferentialExtrapolation(diff_seq)
    next := seq[len(seq)-1]+n
    prev := seq[0]-p
    return next,prev
  }
}

func Day09() bool {
  b,_ := os.ReadFile("input/day09.txt")
  input := string(b)
  lines := strings.Split(input[:len(input)-1],"\n")
  r := regexp.MustCompile("[-]*[0-9]+")

  p1 := 0
  p2 := 0
  for _,l := range lines {
    seq_str := r.FindAllString(l,-1)
    seq := make([]int,len(seq_str))
    for i := range seq {seq[i],_ = strconv.Atoi(seq_str[i])}

    n,p := DifferentialExtrapolation(seq)
    p1 += n
    p2 += p
  }
  
  fmt.Printf("Part 1: %d\nPart 2: %d\n",p1,p2)

  return true
}
