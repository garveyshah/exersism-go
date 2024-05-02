package main

import "testing"

func TestParseCard(t *testing.T) {
	tt := struct {
		name string
		card string
		want int
	}{
		name: "parse ace",
		card: "ace",
		want: 11,
	}

	if got := ParseCard(tt.card); got != tt.want {
		t.Errorf("ParseCard(%s) = %d, want %d", tt.card, got, tt.want)
	}
}

func TestParseCard2(t *testing.T) {
	tt := struct {
		name string
		card string
		want int
	}{
		name: "parse ace",
		card: "ace",
		want: 11,
	}

	if got := ParseCard(tt.card); got != tt.want {
		t.Errorf("ParseCard(%s) = %d, want %d", tt.card, got, tt.want)
	}
}

func TestParseCard3(t *testing.T) {
	tt := struct {
		name string
		card string
		want int
	}{
		name: "parse three",
		card: "three",
		want: 3,
	}

	if got := ParseCard(tt.card); got != tt.want {
		t.Errorf("ParseCard(%s) = %d, want %d", tt.card, got, tt.want)
	}
}

func TestParseCard4(t *testing.T) {
	tt := struct {
		name string
		card string
		want int
	}{
		name: "parse four",
		card: "four",
		want: 4,
	}

	if got := ParseCard(tt.card); got != tt.want {
		t.Errorf("ParseCard(%s) = %d, want %d", tt.card, got, tt.want)
	}
}

func TestParseCard5(t *testing.T) {
	tt := struct {
		name string
		card string
		want int
	}{
		name: "parse five",
		card: "five",
		want: 5,
	}

	if got := ParseCard(tt.card); got != tt.want {
		t.Errorf("ParseCard(%s) = %d, want %d", tt.card, got, tt.want)
	}
}

func TestParseCard6(t *testing.T) {
	tt := struct {
		name string
		card string
		want int
	}{
		name: "parse six",
		card: "six",
		want: 6,
	}

	if got := ParseCard(tt.card); got != tt.want {
		t.Errorf("ParseCard(%s) = %d, want %d", tt.card, got, tt.want)
	}
}

func TestParseCard7(t *testing.T) {
	tt := struct {
		name string
		card string
		want int
	}{
		name: "parse seven",
		card: "seven",
		want: 7,
	}

	if got := ParseCard(tt.card); got != tt.want {
		t.Errorf("ParseCard(%s) = %d, want %d", tt.card, got, tt.want)
	}
}

func TestParseCard8(t *testing.T) {
	tt := struct {
		name string
		card string
		want int
	}{
		name: "parse seven",
		card: "seven",
		want: 7,
	}

	if got := ParseCard(tt.card); got != tt.want {
		t.Errorf("ParseCard(%s) = %d, want %d", tt.card, got, tt.want)
	}
}

func TestParseCard9(t *testing.T) {
	tt := struct {
		name string
		card string
		want int
	}{
		name: "parse nine",
		card: "nine",
		want: 9,
	}

	if got := ParseCard(tt.card); got != tt.want {
		t.Errorf("ParseCard(%s) = %d, want %d", tt.card, got, tt.want)
	}
}

func TestParseCard10(t *testing.T) {
	tt := struct {
		name string
		card string
		want int
	}{
		name: "parse ten",
		card: "ten",
		want: 10,
	}

	if got := ParseCard(tt.card); got != tt.want {
		t.Errorf("ParseCard(%s) = %d, want %d", tt.card, got, tt.want)
	}
}

func TestParseCard11(t *testing.T) {
	tt := struct {
		name string
		card string
		want int
	}{
		name: "parse jack",
		card: "jack",
		want: 10,
	}

	if got := ParseCard(tt.card); got != tt.want {
		t.Errorf("ParseCard(%s) = %d, want %d", tt.card, got, tt.want)
	}
}

func TestParseCard12(t *testing.T) {
	tt := struct {
		name string
		card string
		want int
	}{
		name: "parse queen",
		card: "queen",
		want: 10,
	}

	if got := ParseCard(tt.card); got != tt.want {
		t.Errorf("ParseCard(%s) %d, wand %d", tt.card, got, tt.want)
	}
}

func TestParseCard23(t *testing.T) {
	tt := struct {
		name string
		card string
		want int
	}{
		name: "parse king",
		card: "king",
		want: 10,
	}
	if got := ParseCard(tt.card); got != tt.want {
		t.Errorf("ParseCard(%s) = %d, want %d", tt.card, got, tt.want)
	}
}

func TestParseCard14(t *testing.T) {
	tt := struct {
		name string
		card string
		want int
	}{
		name: "parse other",
		card: "joker",
		want: 0,
	}

	if got := ParseCard(tt.card); got != tt.want {
		t.Errorf("ParseCard(%s) = %d, want %d", tt.card, got, tt.want)
	}
}

func TestFirstTurn1(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "pair of aces",
		hand:   hand{card1: "ace", card2: "ace"},
		dealer: "ace",
		want:   "P",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn2(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "pair of jacks",
		hand:   hand{card1: "jack", card2: "jack"},
		dealer: "ace",
		want:   "S",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn3(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "pair of kings",
		hand:   hand{card1: "king", card2: "king"},
		dealer: "ace",
		want:   "S",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn4(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "pair of twos",
		hand:   hand{card1: "two", card2: "two"},
		dealer: "ace",
		want:   "H",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn5(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "pair of fives",
		hand:   hand{card1: "five", card2: "five"},
		dealer: "ace",
		want:   "H",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn6(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "blackjack with ace for dealer",
		hand:   hand{card1: "ace", card2: "jack"},
		dealer: "ace",
		want:   "S",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn7(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "blackjack with queen for dealer",
		hand:   hand{card1: "king", card2: "ace"},
		dealer: "queen",
		want:   "S",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn8(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "blackjack with five for dealer",
		hand:   hand{card1: "ace", card2: "ten"},
		dealer: "five",
		want:   "W",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn9(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "blackjack with nine for dealer",
		hand:   hand{card1: "ace", card2: "king"},
		dealer: "nine",
		want:   "W",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn11(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 20",
		hand:   hand{card1: "ten", card2: "king"},
		dealer: "ace",
		want:   "S",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn12(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 19",
		hand:   hand{card1: "ten", card2: "nine"},
		dealer: "ace",
		want:   "S",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn13(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 18",
		hand:   hand{card1: "ten", card2: "eight"},
		dealer: "ace",
		want:   "S",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn14(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 17",
		hand:   hand{card1: "seven", card2: "king"},
		dealer: "ace",
		want:   "S",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn15(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 16 with six for dealer",
		hand:   hand{card1: "ten", card2: "six"},
		dealer: "six",
		want:   "S",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn16(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 16 with seven for dealer",
		hand:   hand{card1: "ten", card2: "six"},
		dealer: "seven",
		want:   "H",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn17(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 16 with ace for dealer",
		hand:   hand{card1: "ten", card2: "six"},
		dealer: "ace",
		want:   "H",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn18(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 15 with six for dealer",
		hand:   hand{card1: "ten", card2: "five"},
		dealer: "six",
		want:   "S",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn19(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 15 with seven for dealer",
		hand:   hand{card1: "ten", card2: "five"},
		dealer: "seven",
		want:   "H",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn20(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 15 with king for dealer",
		hand:   hand{card1: "ten", card2: "five"},
		dealer: "king",
		want:   "H",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn21(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 14 with six for dealer",
		hand:   hand{card1: "ten", card2: "four"},
		dealer: "six",
		want:   "S",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn22(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 14 with seven for dealer",
		hand:   hand{card1: "ten", card2: "four"},
		dealer: "seven",
		want:   "H",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn23(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 14 with queen for dealer",
		hand:   hand{card1: "ten", card2: "four"},
		dealer: "queen",
		want:   "H",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn24(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 13 with six for dealer",
		hand:   hand{card1: "ten", card2: "three"},
		dealer: "six",
		want:   "S",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn25(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 13 with seven for dealer",
		hand:   hand{card1: "ten", card2: "three"},
		dealer: "seven",
		want:   "H",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn26(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 13 with queen for dealer",
		hand:   hand{card1: "ten", card2: "three"},
		dealer: "queen",
		want:   "H",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn27(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 12 with six for dealer",
		hand:   hand{card1: "ten", card2: "two"},
		dealer: "six",
		want:   "S",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn28(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 12 with seven for dealer",
		hand:   hand{card1: "ten", card2: "two"},
		dealer: "seven",
		want:   "H",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn29(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 12 with queen for dealer",
		hand:   hand{card1: "ten", card2: "two"},
		dealer: "queen",
		want:   "H",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn30(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 11 with queen for dealer",
		hand:   hand{card1: "nine", card2: "two"},
		dealer: "queen",
		want:   "H",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn31(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 10 with two for dealer",
		hand:   hand{card1: "eight", card2: "two"},
		dealer: "two",
		want:   "H",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn32(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 5 with queen for dealer",
		hand:   hand{card1: "three", card2: "two"},
		dealer: "queen",
		want:   "H",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}

func TestFirstTurn33(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tt := struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		name:   "score of 4 with five for dealer",
		hand:   hand{card1: "two", card2: "two"},
		dealer: "five",
		want:   "H",
	}

	if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
		t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
	}
}
