package countserversthatcommunicate

func countServers(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	connected := make(map[server]bool, rows*cols)

	for y := 0; y <= rows-1; y++ {
		temp := make([]server, 0, cols)

		for x := 0; x <= cols-1; x++ {
			pos := grid[y][x]

			if pos == 0 {
				continue
			}

			s := server{
				x: x,
				y: y,
			}

			temp = append(temp, s)
		}

		if len(temp) > 1 {
			for _, s := range temp {
				connected[s] = true
			}
		}
	}

	for x := 0; x <= cols-1; x++ {
		temp := make([]server, 0, rows)

		for y := 0; y <= rows-1; y++ {
			pos := grid[y][x]

			if pos == 0 {
				continue
			}

			s := server{
				x: x,
				y: y,
			}

			temp = append(temp, s)
		}

		if len(temp) > 1 {
			for _, s := range temp {
				connected[s] = true
			}
		}
	}

	return len(connected)
}

type server struct {
	x, y int
}
