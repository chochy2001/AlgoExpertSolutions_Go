package main

import "fmt"

const Team_won = 1

func TournamentWinner(competitions [][]string, results []int) string {
	//0 gana el visitante
	//1 gana el local
	//["local","visitante"]
	bestTeam := ""
	scores := map[string]int{bestTeam: 0}
	for i, v := range competitions {
		result := results[i]
		local, visit := v[0], v[1]

		winTeam := visit
		if result == Team_won {
			winTeam = local
		}
		updateScores(winTeam, 3, scores)
		if scores[winTeam] > scores[bestTeam] {
			bestTeam = winTeam
		}
	}
	return bestTeam
}

func updateScores(team string, points int, scores map[string]int) {
	scores[team] += points
}

func main() {
	entrada := [][]string{
		{"HTML", "Java"},
		{"Java", "Python"},
		{"Python", "HTML"},
		{"C#", "Python"},
		{"Java", "C#"},
		{"C#", "HTML"},
		{"SQL", "C#"},
		{"HTML", "SQL"},
		{"SQL", "Python"},
		{"SQL", "Java"},
	}
	results := []int{0, 0, 0, 0, 0, 0, 1, 0, 1, 1}
	result := TournamentWinner(entrada, results)
	fmt.Println(result)
}
