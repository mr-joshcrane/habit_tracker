package habit

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Store interface {
	GetHabit(string, string) (*Habit, bool)
	UpdateHabit(*Habit) error
}

type Tracker struct {
	store Store
}

type Habit struct {
	HabitName     string
	Streak        int
	LastPerformed time.Time
	Username	  string
}

type TimeOption func() time.Time

func NewTracker(store Store) *Tracker {
	return &Tracker{
		store: store,
	}
}

var Now = time.Now

func (t *Tracker) GetHabit(habitname, username string) (*Habit, bool) {
	return t.store.GetHabit(habitname, username)
}

func (h *Habit) performedPreviousDay(d time.Time) bool {
	previousDay := d.AddDate(0, 0, -1)
	return h.LastPerformed.Day() == previousDay.Day()
}

func (h *Habit) Perform() {
	t := Now()
	if h.performedPreviousDay(t) {
		h.Streak++
	} else if !h.LastPerformed.Equal(t) {
		h.Streak = 1
	}
	h.LastPerformed = t
}

func RunCLI(s Store) {
	args := os.Args[1:]
	username, ok := os.LookupEnv("USER")
	if !ok {
		username = "unknown"
	}
	if len(args) == 0 {
		fmt.Fprintf(os.Stdout, "Pass the name of the habit you performed today\nExample: %s played violin\n", os.Args[0])
		os.Exit(0)
	}
	t := NewTracker(s)
	habit := strings.Join(args, " ")
	h, ok := t.GetHabit(habit, username)
	h.Perform()
	if !ok {
		fmt.Fprintf(os.Stdout, "Well done, you started the new habit: %s!\n", habit)
	} else {
		fmt.Fprintf(os.Stdout, "Well done, you continued working on habit: %s!\n", habit)
		fmt.Fprintf(os.Stdout, "You've been performing this for a streak of %d day(s)!\n", h.Streak)
	}
	err := t.store.UpdateHabit(h)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
