package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Tally parses tournament input file and writes a summary
func Tally(r io.Reader, w io.Writer) error {

	var mapTeam = map[string]int{"": 0}
	//https://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) < 1 || line[0] == '#' {
			continue
		}
		lst := strings.Split(line, ";")
		if len(lst) != 3 {
			return errors.New("invalid game input")
		}

		team1, team2, result := lst[0], lst[1], lst[2]

		//parse input
		switch {
		case result == "win":
			// team1 won
			if _, ok := mapTeam[team1+"__MP"]; ok {
				mapTeam[team1+"__MP"]++
				mapTeam[team1+"__W"]++
				mapTeam[team1+"__P"] += 3
			} else {
				mapTeam[team1+"__MP"] = 1
				mapTeam[team1+"__W"] = 1
				mapTeam[team1+"__P"] = 3
				mapTeam[team1+"__L"] = 0
				mapTeam[team1+"__D"] = 0
			}

			// team2 lost
			if _, ok := mapTeam[team2+"__MP"]; ok {
				mapTeam[team2+"__MP"]++
				mapTeam[team2+"__L"]++
			} else {
				mapTeam[team2+"__MP"] = 1
				mapTeam[team2+"__L"] = 1
				mapTeam[team2+"__P"] = 0
				mapTeam[team2+"__W"] = 0
				mapTeam[team2+"__D"] = 0
			}

		case lst[2] == "loss":
			// team2 won
			if _, ok := mapTeam[team2+"__MP"]; ok {
				mapTeam[team2+"__MP"]++
				mapTeam[team2+"__W"]++
				mapTeam[team2+"__P"] += 3
			} else {
				mapTeam[team2+"__MP"] = 1
				mapTeam[team2+"__W"] = 1
				mapTeam[team2+"__P"] = 3
				mapTeam[team2+"__L"] = 0
				mapTeam[team2+"__D"] = 0
			}

			// team1 lost
			if _, ok := mapTeam[team1+"__MP"]; ok {
				mapTeam[team1+"__MP"]++
				mapTeam[team1+"__L"]++
			} else {
				mapTeam[team1+"__MP"] = 1
				mapTeam[team1+"__L"] = 1
				mapTeam[team1+"__P"] = 0
				mapTeam[team1+"__W"] = 0
				mapTeam[team1+"__D"] = 0
			}

		case lst[2] == "draw":
			// team2
			if _, ok := mapTeam[team2+"__MP"]; ok {
				mapTeam[team2+"__MP"]++
				mapTeam[team2+"__D"]++
				mapTeam[team2+"__P"]++
			} else {
				mapTeam[team2+"__MP"] = 1
				mapTeam[team2+"__D"] = 1
				mapTeam[team2+"__P"] = 1
				mapTeam[team2+"__W"] = 0
				mapTeam[team2+"__L"] = 0
			}

			// team1
			if _, ok := mapTeam[team1+"__MP"]; ok {
				mapTeam[team1+"__MP"]++
				mapTeam[team1+"__D"]++
				mapTeam[team1+"__P"]++
			} else {
				mapTeam[team1+"__MP"] = 1
				mapTeam[team1+"__D"] = 1
				mapTeam[team1+"__P"] = 1
				mapTeam[team1+"__L"] = 0
				mapTeam[team1+"__W"] = 0
			}
		default:
			return errors.New("invalid game result")
		}

	}
	if err := scanner.Err(); err != nil {
		//log.Fatal(err)
		return err
	}

	// prepare sortList
	sortList := []string{}

	for k, v := range mapTeam {
		if strings.HasSuffix(k, "__P") {
			sortList = append(sortList, fmt.Sprintf("%05d__%s", (10000-v), k))
		}
	}
	//var strSlice sort.StringSlice = sortList
	//sort.Sort(sort.Reverse(strSlice[:]))

	sort.Strings(sortList)

	tallies := []string{"Team                           | MP |  W |  D |  L |  P\n"}
	for _, item := range sortList {
		lst := strings.Split(item, "__")
		team := lst[1]
		tallies = append(tallies, fmt.Sprintf("%-31v| %2d | %2d | %2d | %2d | %2d\n", team, mapTeam[team+"__MP"], mapTeam[team+"__W"], mapTeam[team+"__D"], mapTeam[team+"__L"], mapTeam[team+"__P"]))
	}

	_, err := w.Write([]byte(strings.Join(tallies[:], "")))

	return err

}
