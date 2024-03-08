package main

import (
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	lakers := Team{
		Name:    "Los Angeles Lakers",
		Players: []string{"Lebron James", "Anthony Davis"},
	}
	clippers := Team{
		Name:    "Los Angeles Clippers",
		Players: []string{"Russell Westbrook", "Kawhi Leonard", "Paul George"},
	}
	heat := Team{
		Name:    "Miami Heat",
		Players: []string{"Jimmy Butler", "Tyler Herro"},
	}

	nba := League{
		Teams: []Team{lakers, clippers, heat},
		Wins:  map[string]int{},
	}

	nba.MatchResult(lakers.Name, 100, clippers.Name, 120)
	nba.MatchResult(lakers.Name, 100, heat.Name, 110)
	nba.MatchResult(heat.Name, 100, clippers.Name, 120)

	RankPrinter(nba, os.Stdout)
}

// 7.1 You have been asked to manage a basketball league and are going to write a program to help you. Define two types. The first one, called Team, has a field for the name of the team and a field for the player names. The second type is called League and has a field called Teams for the teams in the league and a field called Wins that maps a team’s name to its number of wins.

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

// 7.2 Add two methods to League. The first method is called MatchResult. It takes four parameters: the name of the first team, its score in the game, the name of the second team, and its score in the game. This method should update the Wins field in League. Add a second method to League called Ranking that returns a slice of the team names in order of wins. Build your data structures and call these methods from the main function in your program using some sample data.

func (l League) MatchResult(name1 string, score1 int, name2 string, score2 int) {
	_, ok1 := l.Wins[name1]
	_, ok2 := l.Wins[name2]

	if !ok1 {
		l.Wins[name1] = 0
	}
	if !ok2 {
		l.Wins[name2] = 0
	}

	l.Wins[name1] = l.Wins[name1] + score1
	l.Wins[name2] = l.Wins[name2] + score2
}

func (l League) Ranking() []string {
	temp := make([]Team, len(l.Teams))
	copy(temp, l.Teams)

	sort.Slice(temp, func(i int, j int) bool {
		return l.Wins[temp[i].Name] > l.Wins[temp[j].Name]
	})

	res := make([]string, len(l.Teams))

	for i, v := range temp {
		res[i] = v.Name
	}

	return res
}

// 7.3 Define an interface called Ranker that has a single method called Ranking that returns a slice of strings. Write a function called RankPrinter with two parameters, the first of type Ranker and the second of type io.Writer. Use the io.Write String function to write the values returned by Ranker to the io.Writer, with a newline separating each result. Call this function from main.
type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	ranking := r.Ranking()
	res := strings.Join(ranking, "\n")
	w.Write([]byte(res))
}
