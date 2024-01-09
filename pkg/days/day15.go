package days

import (
  "os"
  "strings"
)

type lense struct {
  label string
  focal int
}

func hash(inst string) int {
  h := 0
  for _,i := range inst {
    h += int(i)
    h *= 17
    h %= 256
  }
  return h % 256
}

func atoi(r rune) int {
  return int(r)-int('0')
}

func add_modify_lense(lenses []lense, l lense) []lense {
  for ol := range lenses {
    if lenses[ol].label == l.label {
      lenses[ol].focal = l.focal
      return lenses
    }
  }
  return append(lenses, l)
}


func remove_lense(lenses []lense, l string) []lense {
  var new_lenses []lense
  for ol := range lenses {
    if lenses[ol].label != l {
      new_lenses = append(new_lenses, lenses[ol])
    }
  }
  return new_lenses
}

func Day15() bool {
  b,_ := os.ReadFile("input/day15.txt")
  inst := strings.Split(string(b[:len(b)-1]), ",")
  boxes := make(map[int][]lense)
  p1 := 0
  for i := range inst {
    p1 += hash(inst[i])
    for j,r := range inst[i] {
      if r == '-' || r == '=' {
        label := hash(inst[i][:j])
        val, ok := boxes[label]
        if r == '=' {
          l := lense{label: inst[i][:j], focal: atoi(rune(inst[i][j+1]))}
          if ok {
            boxes[label] = add_modify_lense(val, l)
          } else {
            boxes[label] = []lense{l}
          }

        } else {
          val, ok := boxes[label]
          if ok {
            new_val := remove_lense(val, inst[i][:j])
            if len(new_val) == 0 {
              delete(boxes, label)
            } else {
              boxes[label] = new_val
            }
          }
        }
      }
    }
  }
  p2 := 0
  for k,v := range boxes {
    for l := range v {
      p2 += (k + 1) * (l + 1) * v[l].focal
    }
  }

  print("Part 1: ")
  println(p1)
  print("Part 2: ")
  println(p2)

  return true
}
