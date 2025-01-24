// https://www.hackerearth.com/practice/basic-programming/input-output/basics-of-input-output/practice-problems/algorithm/favourite-singer-a18e086a/
package easy

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FavoriteSinger() {
	var n int
	buf := bufio.NewReader(os.Stdin)
	num_of_songs, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	n, err = strconv.Atoi(strings.Trim(strings.TrimRight(num_of_songs, "\n"), " "))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	line, err := buf.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		fmt.Println(err)
		panic(err)
	}

	song_singers := strings.TrimRight(line, "\n")
	lst_song_singers := strings.Split(strings.Trim(song_singers, " "), " ")
	m := make(map[string]int, n)

	max := 0
	for _, s := range lst_song_singers {
		m[s] = m[s] + 1
		if m[s] > max {
			max = m[s]
		}
	}

	numberOfFavSinger := 0
	for _, i := range lst_song_singers {
		if count, ok := m[i]; ok && count == max {
			numberOfFavSinger += 1
			delete(m, i)
		}
	}
	fmt.Println(numberOfFavSinger)
}
