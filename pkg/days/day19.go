package days

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
)


type gear struct {
  x,m,a,s int
}

type rule struct {
  operators []rune
  fields []string
  values []int
  next []string
  exit string
}

type gearRange struct {
  x,m,a,s interval
}

type interval struct {
  min, max int
}

func getField (g *gear, field string) int {
  r := reflect.ValueOf(g)
  f := reflect.Indirect(r).FieldByName(field)
  return int(f.Int())
}


func getGears(s string, r *regexp.Regexp) []gear {
  lines := strings.Split(s,"\n")
  gears := make([]gear, len(lines))

  var xmas []string
  for l := range lines {
    xmas = r.FindAllString(lines[l],-1)
    gears[l] = gear{x: stoi(xmas[0]), m: stoi(xmas[1]), a: stoi(xmas[2]), s: stoi(xmas[3])}
  }
  return gears
}

func getRules(s string) map[string]rule {
  lines := strings.Split(s, "\n")
  rules := make(map[string]rule)

  var name string
  var  operators []rune
  var next,fields,s1,s2,s3 []string
  var exit string
  var values []int
  var i int
  for l := range lines {
    next,fields,operators,values = nil,nil,nil,nil
    s1 = strings.Split(lines[l], "{")
    name = s1[0]
    s2 = strings.Split(s1[1], ",")
    
    for i = 0; i < len(s2)-1; i++ {
      s3 = strings.Split(s2[i],":")
      fields = append(fields, string(s3[0][0]))
      operators = append(operators, rune(s3[0][1]))
      values = append(values, stoi(s3[0][2:]))
      next = append(next, s3[1])
    }
    exit = s2[i][:len(s2[i])-1]
    rules[name] = rule{operators: operators, fields: fields, values: values, next: next, exit: exit}


  }
  return rules
}

func applyRule(g gear, r rule) string {
  for ir := range r.operators {
    if (r.operators[ir] == '<' && getField(&g, r.fields[ir]) < r.values[ir]) ||
        (r.operators[ir] == '>' && getField(&g, r.fields[ir]) > r.values[ir]) {
      return r.next[ir]
    }
  }
  return r.exit
}

func sumField(g gear) int {
  return g.x + g.m + g.a + g.s
}

func getValidIntervals(gr gearRange, r rule, rs map[string]rule) []gearRange {
  var res []gearRange
  var ngr gearRange
  for ir := range r.operators {
    ngr = gr
    operator := r.operators[ir]
    if r.next[ir] == "R" {
      continue
    }
    if operator == '<' {
      if r.fields[ir] == "x" {
        if gr.x.min >= r.values[ir]{ 
          continue
        } else if gr.x.max >= r.values[ir]{
          ngr.x.max = r.values[ir]-1 
        }
      } else if r.fields[ir] == "m" {
        if gr.m.min >= r.values[ir]{ 
          continue
        } else if gr.m.max >= r.values[ir]{
          ngr.m.max = r.values[ir]-1 
        }
      } else if r.fields[ir] == "a" {
        if gr.a.min >= r.values[ir]{ 
          continue
        } else if gr.a.max >= r.values[ir]{
          ngr.a.max = r.values[ir]-1 
        }
      } else {
        if gr.s.min >= r.values[ir]{ 
          continue
        } else if gr.s.max >= r.values[ir]{
          ngr.s.max = r.values[ir]-1 
        }
      }
    } else {
      if r.fields[ir] == "x" {
        if gr.x.max <= r.values[ir]{ 
          continue
        } else if gr.x.min <= r.values[ir]{
          ngr.x.min = r.values[ir]-1 
        }
      } else if r.fields[ir] == "m" {
        if gr.m.max <= r.values[ir]{ 
          continue
        } else if gr.m.min <= r.values[ir]{
          ngr.m.min = r.values[ir]-1 
        }
      } else if r.fields[ir] == "a" {
        if gr.a.max <= r.values[ir]{ 
          continue
        } else if gr.a.min <= r.values[ir]{
          ngr.a.min = r.values[ir]-1 
        }
      } else {
        if gr.s.min >= r.values[ir]{ 
          continue
        } else if gr.s.max >= r.values[ir]{
          ngr.s.max = r.values[ir]-1 
        }
      }


    }
    if r.next[ir] == "A" {
      res = append(res, gr)
    } else {
      res = append(res, getValidIntervals(ngr, rs[r.next[ir]], rs)...)
    } 
  }  
  return nil
}

func Day19() bool {
  b,_ := os.ReadFile("input/day19.txt")
  input := strings.Split(string(b[:len(b)-1]), "\n\n")

  r := regexp.MustCompile("\\d+")

  gears := getGears(input[1], r)
  rules := getRules(input[0])


  p1 := 0
  var ri string 
  for g := range gears {
    ri = "in"
    for ri != "A" && ri != "R" {
      ri = applyRule(gears[g], rules[ri])
    }
    if ri == "A" {
      p1 += sumField(gears[g])
    }
  }
  print("Part 1: ")
  println(p1)

  testi := interval{1,10}
  test := gearRange{x: testi, m: testi , a: testi , s: testi}
  fmt.Printf("%v\n", test)
  //setField(&test, "x", "min", 5)
  fmt.Printf("%v\n", test)

  return true
}
