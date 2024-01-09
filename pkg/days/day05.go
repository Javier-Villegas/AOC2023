package days

import (
	"os"
	"strings"
	"fmt"
	"strconv"
	"slices"
)

type data struct {
	src int
	dst int
	rng int
}
func simplify(table []data) []data {

  size := 0
  for size != len(table) {
    var new_table []data
    size = len(table)
    for i,s1 := range table {
      for _,s2 := range table[i+1:] {
        // [    s1  
        //    [ s2 
        if s1.src <= s2.src {
          // [    s1    ]
          //    [ s2 ] 
          if s1.src+s1.rng >= s2.src+s2.rng {
            new_table = append(new_table, data{src: s1.src, dst: 0, rng: s1.rng})
        // [    s1 ]
        //    [ s2    ]
          } else if s1.src+s1.rng < s2.src+s2.rng && s1.src+s1.rng >= s2.src {
            new_table = append(new_table, data{src: s1.src, dst: 0, rng: s2.src+s2.rng-s1.src})
          }
        //    [ s1 
        // [    s2 
        } else {
          //    [ s1   ]
          // [    s2 ]
          if s1.src+s1.rng >= s2.src+s2.rng && s2.src+s2.rng >= s1.src {
            new_table = append(new_table, data{src: s2.src, dst: 0, rng: s1.src+s1.rng-s2.src})
          //    [ s1 ]
          // [    s2    ]
          } else if s1.src+s1.rng < s2.src+s2.rng {
            new_table = append(new_table, data{src: s2.src, dst: 0, rng: s2.rng})
          }
        }
      }
    }
    if len(new_table) > 0 {
      table = new_table  
    }
  }

  return table
}

func Day05() bool {
	b,_ := os.ReadFile("input/sample05")
	input := string(b)
	
	var maps [][]data

	input_parts := strings.Split(input[:len(input)-1], "\n\n")
	var seeds []int
	for _,s := range strings.Split(strings.Split(input_parts[0], ": ")[1], " ") {
		seed,_ := strconv.Atoi(s)
		seeds = append(seeds,seed)
	}

	fmt.Printf("Seeds: %v\n", seeds)

	for _,part:= range input_parts[1:] {
		var new_map []data
		for _,inst := range strings.Split(part, "\n")[1:] {
			line := strings.Split(inst, " ")
			dst,_ := strconv.Atoi(line[0])
			src,_ := strconv.Atoi(line[1])
			rng,_ := strconv.Atoi(line[2])
			new_map = append(new_map, data{src: src, dst: dst, rng: rng})
		}
		maps = append(maps, new_map)
	}
	table := make([][]int,len(maps)+1)
	for i := range table {table[i] = make([]int, len(seeds))}
	for i := range table[0] {table[0][i] = seeds[i]}

	for i,m := range maps {
		for j,s := range table[i] {
			for _,d := range m {
				if s >= d.src && s <= d.src+d.rng {
					table[i+1][j] = s-d.src+d.dst
					break
				}
			}
			if table[i+1][j] == 0 {table[i+1][j] = table[i][j]}
		}
	}


	table2 := make([][]data, len(maps)+1)
	table2[0] = make([]data, len(seeds)/2)

	for i := 0; i < len(seeds)/2; i++ {
		table2[0][i] = data{src: seeds[i*2], dst: 0, rng: seeds[i*2+1]}
	}

  for i,m := range maps {
    fmt.Printf("%v\n", m)
		for _,s := range table2[i] {
      found := false
		  for _,d := range m {
        if s.src >= d.src {
          //    [  S    ]
          // [     M      ]
          if s.src+s.rng < d.src+d.rng {
            table2[i+1] = append(table2[i+1],
                                 data{src: d.dst, dst: 0, rng: s.rng})
            found = true
          //    [   S    ]
          // [   M    ]
          }else if s.src+s.rng >= d.src+d.rng  && s.src <= d.src+d.rng{
            table2[i+1] = append(table2[i+1],
                                 data{src: d.dst, dst: 0, rng: d.src+d.rng-s.src})
            found = true
          } 

        } else if s.src < d.src {
          // [   S    ]
          //    [ M ]
          if s.src+s.rng >= d.src+d.rng {
            table2[i+1] = append(table2[i+1],
                                 data{src: d.src-s.src+d.dst, dst: 0, rng: d.rng})
            found = true
          //   [   S   ]
          //    [    M      ]
          } else if s.src+s.rng >= d.src && s.src+s.rng <= d.src+d.rng {
            table2[i+1] = append(table2[i+1],
                                 data{src: d.dst+d.src-s.src, dst: 0, rng: s.src+s.rng-d.src})
            found = true
          } 
        }
			}
		println(found)
		}
    fmt.Printf("%v\n", table2[i+1])
    table2[i+1] = simplify(table2[i+1])
    fmt.Printf("%v\n", table2[i+1])
	}


	fmt.Printf("Part 1: %d\n Part 2: %d\n", slices.Min(table[len(table)-1]), 0)
	return true
}
