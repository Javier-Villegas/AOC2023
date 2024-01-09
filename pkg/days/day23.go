package days

import (
  "os"
)

type hiker struct {
  p R2
  v map[R2]struct{}
  steps int
}

func getMapCopy[T any](s map[R2]T) map[R2]T {
  c := make(map[R2]T)
  for i := range s {
    c[i] = s[i]
  }
  return c
}

func getNextTrail(c hiker, si [][]rune) []hiker {
  dirs := []R2{{1,0},{-1,0},{0,1},{0,-1}}
  limits := R2{y:len(si), x:len(si[0])}
  var nh []hiker
  var nv map[R2]struct{}

  var np R2
  for _,d := range dirs {
    np.x = c.p.x+d.x
    np.y = c.p.y+d.y
    _,ok := c.v[np]

    if !ok && CheckBounds(np,limits) && si[np.y][np.x] != '#' {
      nv = getMapCopy(c.v)
      if si[np.y][np.x] == '.' {
        nv[np] = struct{}{}
        nh = append(nh, hiker{p: np, v: nv, steps: c.steps+1})
      } else if si[np.y][np.x] == '>' && d.x != -1 {
        nv[np] = struct{}{}
        np.x++
        nv[np] = struct{}{}
        nh = append(nh, hiker{p: np, v: nv, steps: c.steps+2})
      } else if si[np.y][np.x] == '<' && d.x != 1 {
        nv[np] = struct{}{}
        np.x--
        nv[np] = struct{}{}
        nh = append(nh, hiker{p: np, v: nv, steps: c.steps+2})
      } else if si[np.y][np.x] == '^' && d.y != 1 {
        nv[np] = struct{}{}
        np.y--
        nv[np] = struct{}{}
        nh = append(nh, hiker{p: np, v: nv, steps: c.steps+2})
      } else if si[np.y][np.x] == 'v' && d.y != -1 {
        nv[np] = struct{}{}
        np.y++
        nv[np] = struct{}{}
        nh = append(nh, hiker{p: np, v: nv, steps: c.steps+2})
      }
    }
  }
  return nh
}


func getNextTrail2(c hiker, si [][]rune) []hiker {
  dirs := []R2{{1,0},{-1,0},{0,1},{0,-1}}
  limits := R2{y:len(si), x:len(si[0])}
  var nh []hiker
  var nv map[R2]struct{}

  var np R2
  for _,d := range dirs {
    np.x = c.p.x+d.x
    np.y = c.p.y+d.y
    _,ok := c.v[np]

    if !ok && CheckBounds(np,limits) && si[np.y][np.x] != '#' {
      nv = getMapCopy(c.v)
      nv[np] = struct{}{}
      nh = append(nh, hiker{p: np, v: nv, steps: c.steps+1})
    }
  }
  return nh
}


func insertHiker(h hiker, hl []hiker) []hiker {
  nhl := make([]hiker, len(hl)+1)
  inserted := false
  var i int
  for  i < len(hl) {
    if hl[i].steps < h.steps {
      nhl[i] = h
      inserted = true
      break
    }
    nhl[i] = hl[i]
    i++
  }
  if inserted {
    for i < len(hl) {
      nhl[i+1] = hl[i]
      i++
    }
  } else {
    nhl[i] = h
  }
  return nhl
}


func mergeHikers(h1, h2[]hiker) []hiker {
  for _,h := range h2 {
    h1 = insertHiker(h,h1)
  }
  return h1
}


func getInitialTrail(si [][]rune)  hiker{
  var v  = make(map[R2]struct{})
  var p R2
  for j := range si[0] {
    if si[0][j] == '.' {
      p = R2{x: j, y: 0}
      v[p] = struct{}{}
      break
    }
  }
  return hiker{p:p,v:v,steps: 0}
}


func getLastTrail(si [][]rune)  hiker{
  var v  = make(map[R2]struct{})
  var p R2
  for j := range si[len(si)-1] {
    if si[len(si)-1][j] == '.' {
      p = R2{x: j, y: len(si)-1}
      v[p] = struct{}{}
      break
    }
  }
  return hiker{p:p,v:v,steps: 0}
}


func searchHikingPaths(p hiker, e hiker, si [][]rune, f func(h hiker, s [][]rune) []hiker) hiker {
  var nhl,ends []hiker
  hl := []hiker{p}
  var current hiker
  for len(hl) > 0{
    current = hl[0]
    hl = hl[1:]
    nhl = f(current, si)
    if len(nhl) == 0 && current.p == e.p{
      ends = insertHiker(current, ends)
    } else {
      hl = mergeHikers(hl, nhl)
    }
  }

  return ends[0]
}

func Day23() bool {
  b,_ := os.ReadFile("input/day23.txt")
  snowIsland := getGardenPlots(string(b[:len(b)-1]))

  p1 := searchHikingPaths(getInitialTrail(snowIsland), getLastTrail(snowIsland),snowIsland, getNextTrail)
  println(p1.steps)
  p2 := searchHikingPaths(getInitialTrail(snowIsland),getLastTrail(snowIsland),snowIsland, getNextTrail2)
  println(p2.steps)

  //for i,l := range snowIsland {
  //  for j,r := range l {
  //    _,ok := p2.v[R2{x:j,y:i}]
  //    if ok {
  //      print("O")
  //    } else {
  //      print(string(r))
  //    }
  //  }
  //  println()
  //}


  return true
}
