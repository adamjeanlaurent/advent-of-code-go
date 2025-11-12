DAY=$1

mkdir -p "./2025/day$DAY"
touch "./2025/day$DAY/main.go"
touch "./2025/day$DAY/debug.txt"
touch "./2025/day$DAY/input.txt"

cat <<EOF > "./2025/day$DAY/main.go"
package main

import "github.com/adamjeanlaurent/advent-of-code-go/util"

func main() {
    util.RunDailyProblemSet(part1, part2)
}

func part1(input string) {

}

func part2(input string) {

}

EOF
