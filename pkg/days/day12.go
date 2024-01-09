package days 

import (
  "os"
  "strings"
  "fmt"
  "strconv"
  "regexp"
)
func match_left(s string, code []int) (string, []int) {
  for len(code) > 0 {
    r := regexp.MustCompile(fmt.Sprintf("^[.]*[#]{%d}([.]+|$)", code[0]))
    split := r.Split(s, 2)  
    if len(split) < 2 {
      return s,code
    } else {
      s = split[1]
      code = code[1:]
    }
  }

  return s,code
}

func match_right(s string, code []int) (string, []int) {
  for len(code) > 0 {
    r := regexp.MustCompile(fmt.Sprintf("(^|[.]+)[#]{%d}[.]*$", code[len(code)-1]))
    split := r.Split(s, 2)  
    if len(split) < 2 {
      return s,code
    } else {
      s = split[0]
      code = code[:len(code)-1]
    }
  }

  return s,code
}


func check(s string, code []int) bool {
  sres := strings.Clone(s)
  cres := make([]int, len(code))
  copy(cres, code)


  sres,cres = match_left(sres, cres)
  return sres == "" && len(cres) == 0

}
func replace_at_index(in string, r rune, i int) string {
  out := []rune(in)
  out[i] = r
  return string(out)
}

func is_possible(s string, code []int) bool {
  r := "^[.\\?]*?"
  rsep := "[.\\?]+?"
  rend := "[.\\?]*?$"
  for i,c := range code {
    r += fmt.Sprintf("[#\\?]{%d}",c)
    if i < len(code)-1 {
      r += rsep
    } else {
      r+= rend
    }
  }
  reg := regexp.MustCompile(r)
  return reg.MatchString(s)
}

func fill_data(s string, c []int) int {
  res := 0
  if check(s, c) {
    return 1
  }
  if !is_possible(s, c) {
    return 0
  }

  for i,r := range s {
    if r == '?' {
      s1 := replace_at_index(s, '.', i)
      s2 := replace_at_index(s, '#', i)
      res += fill_data(s1, c)
      res += fill_data(s2, c)
      break
    }
  }
  return res
}


func Day12() bool {
  b,_ := os.ReadFile("input/day12.txt")
  input := string(b[:len(b)-1])

  lines := strings.Split(input, "\n")
  p1 := 0
  p2 := 0
  for _,s := range lines[1:2] {
    parts := strings.Split(s, " ")

    var code []int
    for _,num := range strings.Split(parts[1], ",") {
      c,_ := strconv.Atoi(num)
      code = append(code, c)
    }
    sum := len(code)-1
    for i := range code {
      sum += code[i]
    }

    springs := parts[0]
    springs2 := springs + "?" + springs
    code2 := append(code, code...)
    
    println(fill_data(springs2+"?"+springs,
    append(code2,code...)))
    println(fill_data(springs+"?",
    append(code, []int{}...)))
    if sum == len(parts[0]){
      p2 += 1
      p1 += 1
    } else {

      f1 := fill_data(springs, code)
      f2 := fill_data(springs2, code2)/f1
      p1 += f1
      p2 += f2*f2*f1

    }



  }


  fmt.Printf("Part 1: %d\nPart 2: %d\n", p1, p2)
  return true
}
