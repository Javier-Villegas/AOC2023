package days

import (
	"fmt"
	"os"
	"strings"
)


type block3D struct {
  x,y,z interval
  supports []int
  supported []int
  id int
}

func blockFactory(s string, i int) block3D {
  split := strings.Split(s,"~")
  b1 := strings.Split(split[0],",")
  b2 := strings.Split(split[1],",")

  return block3D{ x: interval{min:stoi(b1[0]), max:stoi(b2[0])},
                  y: interval{min:stoi(b1[1]), max:stoi(b2[1])},
                  z: interval{min:stoi(b1[2]), max:stoi(b2[2])},
                  supports: []int{},
                  supported: []int{},
                  id: i,
                }
}

func lessBlock3D(b1, b2 block3D) bool {
  return b1.z.max < b2.z.max
}

func insertBlock(ss []block3D, b block3D) []block3D {
  nss := make([]block3D, len(ss)+1)
  i := 0
  inserted := false
  for i < len(ss) {
    if !lessBlock3D(ss[i],b) {
      nss[i] = b
      inserted = true
      break
    }
    nss[i] = ss[i]
    i++
  }
  if !inserted {
    nss[i] = b
  } else {
    for i < len(ss) {
      nss[i+1] = ss[i]
      i++
    }
  }
  

  return nss
}

func placeBlock(ss []block3D, b block3D) []block3D {
  placed := false
  var nss []block3D
  z := b.z.max
  for i := len(ss)-1; i > -1; i-- {
    if lessBlock3D(ss[i], b) && intersectBlock(ss[i], b) && (!placed || z <= ss[i].z.max){
      if !placed {
        placed = true
        z = ss[i].z.max
        //nss = insertBlock(ss, b)
      }
      ss[i].supports = append(ss[i].supports,b.id)
      b.supported = append(b.supported, ss[i].id)
    }

  }
  if !placed {
    b.z.max = (b.z.max-b.z.min)+1
    b.z.min = 1

  } else {
    b.z.max = (b.z.max-b.z.min)+z+1
    b.z.min = z+1
  }
  nss = insertBlock(ss, b)
  return nss
}

func intersect(i interval, j interval) (interval, bool) {
  // [     i 
  //    [  j
  if i.min <= j.min {
    if i.max < j.min {
      return interval{-1,-1}, false
    } else {
      if j.max <= i.max {
        return interval{j.min, j.max}, true
      } else {
        return interval{j.min, i.max}, true
      }
    }

  //     [   i
  // [       j
  } else {
    if j.max < i.min {
      return interval{-1,-1}, false
    } else {
      if i.max <= j.max {
        return interval{i.min, i.max}, true
      } else {
        return interval{i.min, j.max}, true
      }
    }
  }
}
func intersectBlock(b1, b2 block3D) bool {
  _,okx := intersect(b1.x, b2.x)
  _,oky := intersect(b1.y, b2.y)
  return okx && oky
}

func searchBlockByID(id int, bl []block3D) block3D {
  for i := 0; i < len(bl); i++ {
    if bl[i].id == id {
      return bl[i]
    }
  }
  return block3D{}
}

func destroyBlocks(initial block3D, bl []block3D) int {
  counter := 0
  var tl = make([]int, len(initial.supports))
  dl := []int{initial.id}
  copy(tl, initial.supports)

  var target int
  var next block3D
  for len(tl) > 0 {
    target = tl[0]
    tl = tl[1:]
    next = searchBlockByID(target, bl)
    if isFalling(next, dl) {
      counter++
      dl = append(dl, target)
      tl = mergeListWithoutRepeats(tl, next.supports)
    }
  }
  return counter
}

func isFalling(b block3D,  destroyed[]int) bool {
  counter := 0
  for _,s := range b.supported {
    for _,id := range destroyed{
      if id == s {
        counter += 1
        break
      }
    }
  }
  return  counter == len(b.supported)
}

func insertWithoutRepetition(e int, l []int) []int {
  for i := 0; i < len(l); i++ {
    if l[i] == e {
      return l
    }
  }
  return append(l, e)
}

func mergeListWithoutRepeats(src []int, dst []int) []int {
  for _,e := range dst {
    src = insertWithoutRepetition(e, src)
  }
  return src
}

func Day22() bool {
  //b,_ := os.ReadFile("input/sample22")
  b,_ := os.ReadFile("input/day22.txt")
  input := strings.Split(string(b[:len(b)-1]),"\n")

  var blocks []block3D

  for i,line := range input {
    blocks = insertBlock(blocks, blockFactory(line, i))
  }


  var placedBlocks []block3D
  for _,b := range blocks {
    placedBlocks = placeBlock(placedBlocks, b)
  }

  p1 := 0
  for _,b := range placedBlocks {
    if len(b.supports) == 0 {
      p1 += 1
    } else {
      inc := true
      for _,ib := range b.supports {
        for i := range placedBlocks {
          if placedBlocks[i].id == ib {
            if len(placedBlocks[i].supported) == 1 {
              inc = false
            }
            break
          } 
        }
      }
      if inc {
        p1 += 1
      }
    }
  }

  p2 := 0
  for _,b := range placedBlocks {
    p2 += destroyBlocks(b, placedBlocks)
  }

  fmt.Printf("Part 1: %d\nPart 2: %d\n", p1, p2)
  return true
}
