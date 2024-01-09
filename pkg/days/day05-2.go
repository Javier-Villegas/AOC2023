package days

import (
  "os"
  "strings"
  "fmt"
)


type almanac struct {
  src, dst interval
}

func intersectWithRest(a interval, b almanac) (interval,[]interval,bool) {
  //      [  a
  //  [      b
  if a.min >= b.src.min {
    if b.src.max < a.min {
      //return interval{},[]interval{{min: a.min, max: a.max}},false
      return interval{},[]interval{},false
    }
    //    [  a   ]
    // [     b     ]
    if a.max <= b.src.max {
      return interval{min: a.min-b.src.min+b.dst.min, max: a.max-b.src.min+b.dst.min},[]interval{},true
    //    [   a   ]
    // [      b ]
    } else {
      return interval{min: a.min-b.src.min+b.dst.min, max: b.dst.max},[]interval{{min: b.src.max+1,max: a.max}}, true
    }
  //  [      a
  //     [   b
  } else {
    if a.max < b.src.min {
      //return interval{},[]interval{{min: a.min, max: a.max}}, false
      return interval{},[]interval{}, false
    }
    // [     a   ]
    //    [   b    ]
    if b.src.max >= a.max {
      return interval{min: b.dst.min, max: a.max-b.src.min+b.dst.min},[]interval{{min: a.min, max: b.src.min-1}}, true
    // [     a     ]
    //    [  b  ]
    } else {
      return interval{min: b.dst.min, max: b.dst.max},[]interval{{min: a.min, max: b.src.min-1},{min: b.src.max+1, max: a.max}}, true
    }
  }
}

func Day05_2() bool{
	b,_ := os.ReadFile("input/sample05")
	input := string(b)
	

	input_parts := strings.Split(input[:len(input)-1], "\n\n")
	var seeds []interval
  var seed interval
  ss := strings.Split(strings.Split(input_parts[0], ": ")[1], " ")
  for i := 0; i < len(ss); i += 2 {
    seed = interval{min: stoi(ss[i]), max: stoi(ss[i])+stoi(ss[i+1])-1}
		seeds = append(seeds,seed)
	}

	fmt.Printf("Seeds: %v\n", seeds)
  almanacs := make([][]almanac, 7)
  var almanac_def, line []string

  var src, dst, rng int
	for i,part:= range input_parts[1:] {
    almanac_def = strings.Split(part, "\n")[1:]
    almanacs[i] = make([]almanac, len(almanac_def))
		for j,inst := range almanac_def {
      line = strings.Split(inst, " ")
      src = stoi(line[1])
      dst = stoi(line[0])
      rng = stoi(line[2])
      almanacs[i][j] = almanac{src: interval{min: src, max: src+rng-1}, dst: interval{min: dst, max: dst+rng-1}}
    }
	}

  for i := range almanacs {
    fmt.Printf("almanacs[i]: %v\n", almanacs[i])
  }

  var rest,next []interval
  var new_seed interval
  var ok bool
  fmt.Printf("seeds: %v\n", seeds)
  for i := range almanacs {
    for s := 0; s < len(seeds); s++{
      for j := range almanacs[i] {
        seed = seeds[s]
        new_seed,rest,ok = intersectWithRest(seed, almanacs[i][j])
        if len(rest) > 0 {
          seeds = append(seeds, rest...)
        }
        if ok {
          next = append(next, new_seed)
          seeds = append(seeds[:s], seeds[s+1:]...)
          s--
          break
        }

        
      }
    }
    fmt.Printf("next: %v\n", next)
    seeds = append(next, seeds...)
  }
  fmt.Printf("seeds: %v\n", seeds)
  
  min := seeds[0].min
  for s := range seeds {
    if min > seeds[s].min {
      min = seeds[s].min
    }
  }
  println(min)

  return true
}

