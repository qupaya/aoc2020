package day7

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRule(t *testing.T) {
	bag, content := ParseRule("faded blue bags contain no other bags.")
	assert.Equal(t, "faded blue", bag)
	assert.Equal(t, map[string]int{}, content)

	bag, content = ParseRule("light red bags contain 1 bright white bag, 2 muted yellow bags.")
	assert.Equal(t, "light red", bag)
	assert.Equal(t, map[string]int{"bright white": 1, "muted yellow": 2}, content)

	bag, content = ParseRule("light red bags contain 2 bright white bags, 2 muted yellow bags.")
	assert.Equal(t, "light red", bag)
	assert.Equal(t, map[string]int{"bright white": 2, "muted yellow": 2}, content)

	bag, content = ParseRule("bright white bags contain 1 shiny gold bag.")
	assert.Equal(t, "bright white", bag)
	assert.Equal(t, map[string]int{"shiny gold": 1}, content)
}

func TestParseRules(t *testing.T) {
	input := strings.Split(`faded blue bags contain no other bags.
light red bags contain 1 bright white bag, 2 muted yellow bags.
bright white bags contain 1 shiny gold bag.`, "\n")

	assert.Equal(t, Bags{"faded blue": {}, "light red": {"bright white": 1, "muted yellow": 2}, "bright white": {"shiny gold": 1}}, ParseRules(&input))
}

func TestCountPossibilities(t *testing.T) {
	input := strings.Split(`muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.`, "\n")
	assert.Equal(t, 1, CountPossibilities("shiny gold", ParseRules(&input)))

	input = strings.Split(`light red bags contain 1 bright white bag, 2 muted yellow bags
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.`, "\n")
	assert.Equal(t, 2, CountPossibilities("shiny gold", ParseRules(&input)))

	input = strings.Split(`dark orange bags contain 3 bright white bags, 4 light red bags.
light red bags contain 1 bright white bag, 2 muted yellow bags
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.`, "\n")
	assert.Equal(t, 3, CountPossibilities("shiny gold", ParseRules(&input)))

	input = strings.Split(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`, "\n")
	assert.Equal(t, 4, CountPossibilities("shiny gold", ParseRules(&input)))
}

func TestCountBags(t *testing.T) {
	input := strings.Split(`muted yellow bags contain no other bags.`, "\n")
	assert.Equal(t, 0, CountBags("muted yellow", ParseRules(&input)))

	input = strings.Split(`muted yellow bags contain 1 bright white bag.
bright white bags contain no other bags.`, "\n")
	assert.Equal(t, 1, CountBags("muted yellow", ParseRules(&input)))

	input = strings.Split(`muted yellow bags contain 2 bright white bag.
bright white bags contain no other bags.`, "\n")
	assert.Equal(t, 2, CountBags("muted yellow", ParseRules(&input)))

	input = strings.Split(`muted yellow bags contain 1 bright red bag.
bright red bags contain 1 bright white bag.
bright white bags contain no other bags.`, "\n")
	assert.Equal(t, 2, CountBags("muted yellow", ParseRules(&input)))

	input = strings.Split(`shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`, "\n")
	assert.Equal(t, 126, CountBags("shiny gold", ParseRules(&input)))

}
