package habit

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/glamour"
)

type Store interface {
	GetHabit(Username, HabitID) (*Habit, error)
	UpdateHabit(*Habit) error
	ListHabits(Username) ([]*Habit, error)
	GetBattle(BattleCode) (*Battle, error)
	UpdateBattle(*Battle) error
	ListBattlesByUser(Username) ([]*Battle, error)
}

type Tracker interface {
	PerformHabit(Username, HabitID) (int, error)
	DisplayHabits(Username) []string
	RegisterBattle(Username, HabitID) (BattleCode, error)
	JoinBattle(BattleCode, Username, HabitID) error
	GetBattleAssociations(Username, HabitID) []BattleCode
}

type Username string
type HabitID string
type BattleCode string
type Habit struct {
	HabitName     string
	Streak        int
	LastPerformed time.Time
	Username      string
}

type Battle struct {
	HabitOne *Habit
	HabitTwo *Habit
	Code     BattleCode
	Winner   string
}

type TimeOption func() time.Time

var Now = time.Now
var BattleCodeGenerator = generateBattleCode

func (h *Habit) performedPreviousDay(d time.Time) bool {
	previousDay := d.AddDate(0, 0, -1)
	return h.LastPerformed.Day() == previousDay.Day()
}

func (h *Habit) Perform() {
	t := Now()
	if h.performedPreviousDay(t) {
		h.Streak++
	} else if h.LastPerformed.Before(t.AddDate(0, 0, -1)) {
		h.Streak = 1
	}
	h.LastPerformed = t
}

func CreateChallenge(h *Habit) *Battle {
	return &Battle{
		HabitOne: h,
		HabitTwo: &Habit{},
		Code:     BattleCodeGenerator(),
	}
}

func JoinBattle(h *Habit, b *Battle) (*Battle, error) {
	if b.HabitOne == h || b.HabitTwo == h {
		return nil, fmt.Errorf("already enrolled in this battle")
	}
	if b.HabitOne.HabitName != "" && b.HabitTwo.HabitName != "" {
		return nil, fmt.Errorf("battle already has two participants")
	}
	if b.HabitOne.Username == h.Username {
		return nil, fmt.Errorf("participant is already registered in this battle")
	}
	b.HabitTwo = h
	return b, nil
}

func (b *Battle) DetermineWinner() string {
	t := Now()
	t1 := b.HabitOne.LastPerformed
	t2 := b.HabitTwo.LastPerformed
	if t1.After(t.AddDate(0, 0, -1)) && t2.After(t.AddDate(0, 0, -1)) {
		return ""
	}
	if t1.After(t2.AddDate(0, 0, -1)) {
		return b.HabitOne.Username
	}
	return b.HabitTwo.Username
}

func (b *Battle) IsPending() bool {
	if b.HabitOne == nil || b.HabitTwo == nil {
		return true
	}
	return false
}

func generateBattleCode() BattleCode {
	length := 5
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(length)]
	}
	return BattleCode(b)
}

func RunCLI(p Tracker) {
	output := ""
	createChallenge := flag.Bool("c", false, "Create a new challenge")
	joinChallenge := flag.String("j", "", "Join an existing challenge")
	flag.Parse()
	args := flag.Args()
	username, ok := os.LookupEnv("USER")
	if !ok {
		username = "unknown"
	}
	if len(args) == 0 {
		output += fmt.Sprintf("Pass the name of the habit you performed today\n\nExample: %s played violin\n\n", os.Args[0])
		Render(output, os.Stdout)
		os.Exit(0)
	}
	habitID := strings.Join(args, " ")
	streak, err := p.PerformHabit(Username(username), HabitID(habitID))
	if err != nil {
		output += fmt.Sprintf("issues performing habit: %v", err)
		Render(output, os.Stdout)
		os.Exit(1)
	}
	if streak == 1 {
		output += fmt.Sprintf("New streak started for a new habit: %s!\n\n", habitID)
	} else {
		output += fmt.Sprintf("Well done, you continued working on habit: %s!\n\nYou've been performing this for a streak of %d day(s)!\n\n", habitID, streak)
	}
	hList := p.DisplayHabits(Username(username))
	output += fmt.Sprintf("All your current habits: %s!\n\n", hList)
	if *createChallenge {
		code, err := p.RegisterBattle(Username(username), HabitID(habitID))
		if err != nil {
			output += err.Error()
			Render(output, os.Stderr)
			os.Exit(1)
		}
		output += fmt.Sprintf("Created challenge: %s\n\n", code)
	}
	if *joinChallenge != "" {
		err := p.JoinBattle(BattleCode(*joinChallenge), Username(username), HabitID(habitID))
		if err != nil {
			output += err.Error()
			Render(output, os.Stderr)
			os.Exit(1)
		}
		output += fmt.Sprintf("Joined challenge: %s\n\n", *joinChallenge)
	}
	b := p.GetBattleAssociations(Username(username), HabitID(habitID))
	output += fmt.Sprintf("Your current battles: %s!\n\n", b)
	Render(output, os.Stdout)
}

func Render(input string, std *os.File) {
	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),

	)
	out, _ := r.Render(input)
	fmt.Println(out)
}
