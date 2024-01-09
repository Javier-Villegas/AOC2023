package days

import (
  "os"
  "fmt"
  "sort"
  "strings"
  "strconv"
)


type handt struct {
  cards []int
  hand_type int
  main_card int
  secondary_card int
  bid int
}


type kv struct {
  key int
  val int
}

var conversion = map[rune]int{'A':14,'K':13,'Q':12,'J':11,'T':10,'9':9,'8':8,'7':7,'6':6,'5':5,'4':4,'3':3,'2':2}

func ProcessJoker(h handt) handt {
  num_jokers := 0
  for i := range h.cards {
    if h.cards[i] == 11 {
      h.cards[i] = -1
      num_jokers += 1
    }
  }
  if num_jokers == 0 {return h}

  counter := map[int]int{14:0,13:0,12:0,-1:0,10:0,9:0,8:0,7:0,6:0,5:0,4:0,3:0,2:0}

  for i := range h.cards {
    counter[h.cards[i]] += 1
  }

  var ss []kv
  for k,v := range counter {ss = append(ss, kv{key:k,val:v})}
  sort.Slice(ss, func(i, j int) bool {return ss[i].val > ss[j].val})
  
  if ss[0].key == -1 {
    ss = ss[1:]
  }else if ss[1].key == -1 {
    ss[1] = ss[0]
    ss = ss[1:]
  }

  ss[0].val += counter[-1]

  h.main_card = ss[0].key
  if ss[1].val > 0 {h.secondary_card = ss[1].key}

  if ss[0].val == 5 {h.hand_type = 6} else
  if ss[0].val == 4 {h.hand_type = 5} else
  if ss[0].val == 3 && ss[1].val == 2 {h.hand_type = 4} else
  if ss[0].val == 3 {h.hand_type = 3} else
  if ss[0].val == 2 && ss[1].val == 2 {h.hand_type = 2} else
  if ss[0].val == 2 {h.hand_type = 1} else
  {h.hand_type = 0}

  return h
}

func ProcessHand(h string) handt {
  aux := strings.Split(h, " ")
  bid,_ := strconv.Atoi(aux[1])
  card := make([]int,5)
  for i := range aux[0] {card[i],_ = conversion[rune(aux[0][i])]}
  hand := handt{cards: card, hand_type: -1, main_card: -1, secondary_card: -1, bid: bid}
  counter := map[int]int{14:0,13:0,12:0,11:0,10:0,9:0,8:0,7:0,6:0,5:0,4:0,3:0,2:0}

  for i := range hand.cards {
    counter[hand.cards[i]] += 1
  }

  var ss []kv
  for k,v := range counter {ss = append(ss, kv{key:k,val:v})}
  sort.Slice(ss, func(i, j int) bool {return ss[i].val > ss[j].val})
  
  hand.main_card = ss[0].key
  if ss[1].val > 0 {hand.secondary_card = ss[1].key}

  if ss[0].val == 5 {hand.hand_type = 6} else
  if ss[0].val == 4 {hand.hand_type = 5} else
  if ss[0].val == 3 && ss[1].val == 2 {hand.hand_type = 4} else
  if ss[0].val == 3 {hand.hand_type = 3} else
  if ss[0].val == 2 && ss[1].val == 2 {hand.hand_type = 2} else
  if ss[0].val == 2 {hand.hand_type = 1} else
  {hand.hand_type = 0}
  return hand
}

func CompHands(h1 handt, h2 handt) bool { 
  if h1.hand_type != h2.hand_type {
    return h1.hand_type > h2.hand_type
  } else {
    for i := range h1.cards {
      if h1.cards[i] != h2.cards[i] {
        return h1.cards[i] > h2.cards[i]
      }
    }
  }
  return true
}

func Day07() bool {
  b,_ := os.ReadFile("input/day07.txt")
  input := string(b)
  lines := strings.Split(input[:len(input)-1],"\n")

  hands := make([]handt,len(lines))

  for i := range hands {hands[i] = ProcessHand(lines[i])}

  sort.Slice(hands, func (i, j int) bool {return CompHands(hands[j],hands[i])})


  new_hands := make([]handt,len(hands))
  for i := range new_hands {new_hands[i] = ProcessJoker(hands[i])}
  sort.Slice(new_hands, func (i, j int) bool {return CompHands(new_hands[j],new_hands[i])})

  p1 := 0
  p2 := 0
  for i := range hands {
    p1 += (i+1)*hands[i].bid
    p2 += (i+1)*new_hands[i].bid
  }

  fmt.Printf("Part 1: %d\nPart 2: %d\n", p1, p2)
  return true
}
