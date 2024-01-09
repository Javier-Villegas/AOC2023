package days

import (
	"fmt"
	"os"
	"strings"
)

type Module struct {
  input []string
  output []string
  mem []int
  in func(i int, im string, m []int, inl []string)
  out func(i int, om string, m []int, onl []string) int
}

type Pulse struct {
  src string
  dst string
  val int
}

func inFlipFlop(i int,im string,  mem []int, inl []string) {
  if i == 0 {
    if mem[0] == 0 {
      mem[0] = 1
    } else {
      mem[0] = 0
    }
  }
}

func outFlipFlop(i int, om string, mem []int, onl []string) int {
  if i == 0 {
    return mem[0]
  }
  return -1
}


func inConjunction(i int, im string, mem []int, inl []string) {
  var iter int
  for iter = 0; iter < len(inl); iter++ {
    if inl[iter] == im {
      mem[iter] = i
      return
    }
  }
}

func outConjunction(i int, om string, mem []int, onl []string) int {
  for iter := range mem {
    if mem[iter] == 0 {
      return 1
    }
  }
  return 0
}

func inForward(i int, om string, mem []int, onl []string) {
  return
}

func outForward(i int, im string, mem []int, inl []string) int {
  return i
}

func Day20() bool {
  b,_ := os.ReadFile("input/day20.txt")
  input := strings.Split(string(b[:len(b)-1]),"\n")


  moduleList := make(map[string]*Module)
  //connection := make(map[string]*Module)

  var fi func(i int, im string, m []int, inl []string)
  var fo func(i int, om string, m []int, onl []string) int
  var name string

  for _,line := range input {
    in_out := strings.Split(line, " -> ")
    if rune(in_out[0][0]) == '%'{
      fi = inFlipFlop
      fo = outFlipFlop
      name = in_out[0][1:]
    } else if rune(in_out[0][0]) == '&' {
      fi = inConjunction
      fo = outConjunction
      name = in_out[0][1:]
    } else {
      fi = inForward
      fo = outForward
      name = in_out[0]
    }

    moduleList[name] = &Module{input: []string{}, output: strings.Split(in_out[1], ", "), mem: []int{}, in: fi, out: fo}
  }

  //moduleList["output"] = &Module{input: []string{}, output: []string{}, mem: []int{}, in: nil, out: nil}
  moduleList["button"] = &Module{input: []string{}, output: []string{"broadcaster"}, mem: []int{}, in: inForward, out: outForward}
  (*moduleList["broadcaster"]).input = append((*moduleList["broadcaster"]).input, "button")
  for k,valptr := range moduleList {
    for i := range (*valptr).output {
      if moduleList[(*valptr).output[i]] == nil {
        println((*valptr).output[i])
        moduleList[(*valptr).output[i]] = &Module{input: []string{}, output: []string{}, mem: []int{}, in: inForward, out: outForward}
      }
      (*moduleList[(*valptr).output[i]]).input = append((*moduleList[(*valptr).output[i]]).input, k)
      (*moduleList[(*valptr).output[i]]).mem = append((*moduleList[(*valptr).output[i]]).mem, 0)
    }
  }


  var queue []Pulse
  var current,newPulse Pulse
  var mod Module
  var mem []int
  var il, ol []string
  res := [2]int{0,0}
  p2 := 0
  isP2Done := false

  for iter := 0; iter < 1000; iter++ {
    queue = append(queue, Pulse{"button", "broadcaster", 0})
    if !isP2Done {
      p2++
    }
    for len(queue) > 0 {
      //fmt.Printf("%v\n", queue)
      current = queue[0]
      queue = queue[1:]
      res[current.val]++
      //fmt.Printf("%v\n", current)
      if current.dst == "rx" && current.val == 0 {
        print("Part 2: ")
        println(p2)
        
        isP2Done = true
      }
      
      mod = *(moduleList[current.dst])
      mem = mod.mem
      il = mod.input
      ol = mod.output
      //fmt.Printf("%v\n", mem)
      //println(current.val)
      mod.in(current.val, current.src, mem, il)
      //fmt.Printf("%v\n", mem)

      for o := range ol {
        newPulse.dst = ol[o]
        newPulse.src = current.dst
        newPulse.val = mod.out(current.val, ol[o], mem, ol)
        if newPulse.val == -1 {
          continue
        }
        //fmt.Printf("%s %d -> %s\n", newPulse.src, newPulse.val, ol[o])
        queue = append(queue,newPulse)
      }
    }
  }
  for !isP2Done {
    queue = append(queue, Pulse{"button", "broadcaster", 0})
    if !isP2Done {
      p2++
    }
    for len(queue) > 0 {
      //fmt.Printf("%v\n", queue)
      current = queue[0]
      queue = queue[1:]
      //fmt.Printf("%v\n", current)
      if current.dst == "rx" && current.val == 0 {
        print("Part 2: ")
        println(p2)
        
        isP2Done = true
      }
      
      mod = *(moduleList[current.dst])
      mem = mod.mem
      il = mod.input
      ol = mod.output
      //fmt.Printf("%v\n", mem)
      //println(current.val)
      mod.in(current.val, current.src, mem, il)
      //fmt.Printf("%v\n", mem)

      for o := range ol {
        newPulse.dst = ol[o]
        newPulse.src = current.dst
        newPulse.val = mod.out(current.val, ol[o], mem, ol)
        if newPulse.val == -1 {
          continue
        }
        //fmt.Printf("%s %d -> %s\n", newPulse.src, newPulse.val, ol[o])
        queue = append(queue,newPulse)
      }
    }
  }
  fmt.Printf("%v\n", res)
  fmt.Printf("Part 1: %d\n", res[0]*res[1])
  return true
}
