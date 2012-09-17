// Package dice ...
// TODO add documentation
package dice

import (
	"fmt"
	"regexp"
	"strconv"
)

type DiceRoll struct {
	NumberOfDice int
	DieFaces int
	Adder    int
	Half bool
	Rolls    []int
	RawTotal int
	Total    int
}

func (dr *DiceRoll) Description() string {
	var adder, half string
	if dr.Adder < 0 {
		adder = strconv.Itoa(dr.Adder)
	} else if dr.Adder == 0 {
		adder = ""
	} else {
		adder = "+" + strconv.Itoa(dr.Adder)
	}
	if dr.Half {
		half = ".5"
	} else {
		half = ""
	}
	return fmt.Sprintf("%v%vd%v%v", dr.NumberOfDice, half, dr.DieFaces, adder)
}

func newDiceRollP(number, faces, adder int, half bool, rng intRng) *DiceRoll {
	d := &DiceRoll{NumberOfDice: number, DieFaces: faces, Adder: adder, Half: half}
	d.Rolls = make([]int, number)
	for i := 0; i < number; i++ {
		d.Rolls[i] = rng.Intn(faces) + 1
		d.RawTotal += d.Rolls[i]
	}
	if half {
		halfDie := rng.Intn(faces/2) + 1
		d.Rolls = append(d.Rolls, halfDie)
		d.RawTotal += halfDie
	}
	d.RawTotal += d.Adder
	if d.RawTotal < 1 {
		d.Total = 1
	} else {
		d.Total = d.RawTotal
	}
	return d
}

func RollP(number, faces, adder int, half bool) *DiceRoll {
	return newDiceRollP(number, faces, adder, half, localRng)
}

var diceExp = regexp.MustCompile(`^([1-9][0-9]*)?(\.5)?[dD]([1-9][0-9]*)([+-][1-9][0-9]*)?$`)

func Roll(description string) (*DiceRoll, error) {
	return newDiceRoll(description, localRng)
}

func newDiceRoll(description string, rng intRng) (*DiceRoll, error) {
	parts := diceExp.FindStringSubmatch(description)
	if parts == nil {
		return nil, fmt.Errorf("Bad description: %v", description)
	}

	numS, halfS, facesS, adderS := parts[1], parts[2], parts[3], parts[4]
	var number, faces, adder int
	var half bool
	var err error
	if numS == "" {
		number = 1
	} else if number, err = strconv.Atoi(numS); err != nil {
		return nil, fmt.Errorf("Bad number of dice: %v", numS)
	}
	half = halfS != ""
	if faces, err = strconv.Atoi(facesS); err != nil {
		return nil, fmt.Errorf("Bad die faces: %v", facesS)
	}
	if adderS == "" {
		adder = 0
	} else if adder, err = strconv.Atoi(adderS); err != nil {
		return nil, fmt.Errorf("Bad adder: %v", adderS)
	}
	return newDiceRollP(number, faces, adder, half, rng), nil
}
