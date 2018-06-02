package multiline_test

import (
	"testing"

	"github.com/fernandoocampo/katas/multiline"
)

type data struct {
	param  string
	result []string
}

var cases = []data{
	{"Hitchhiker", []string{"Hitchhiker's Guide to the Galaxy"}},
	{"Douglas", []string{"Hitchhiker's Guide to the Galaxy"}},
	{"Dune", []string{"Dune"}},
	{"the", []string{"Hitchhiker's Guide to the Galaxy", "Dune", "A Song Of Ice And Fire Series"}},
	{"asdfasfff", []string{}},
	{"Arrakis", []string{"Dune"}},
	{"winter", []string{"A Song Of Ice And Fire Series"}},
	{"demolished", []string{"Hitchhiker's Guide to the Galaxy"}},
}

var content string = `TITLE: Hitchhiker's Guide to the Galaxy
AUTHOR: Douglas Adams
DESCRIPTION: Seconds before the Earth is demolished to make way for a galactic freeway,
Arthur Dent is plucked off the planet by his friend Ford Prefect, a researcher for the
revised edition of The Hitchhiker's Guide to the Galaxy who, for the last fifteen years,
has been posing as an out-of-work actor.

TITLE: Dune
AUTHOR: Frank Herbert
DESCRIPTION: The troubles begin when stewardship of Arrakis is transferred by the
Emperor from the Harkonnen Noble House to House Atreides. The Harkonnens don't want to
give up their privilege, though, and through sabotage and treachery they cast young
Duke Paul Atreides out into the planet's harsh environment to die. There he falls in
with the Fremen, a tribe of desert dwellers who become the basis of the army with which
he will reclaim what's rightfully his. Paul Atreides, though, is far more than just a
usurped duke. He might be the end product of a very long-term genetic experiment
designed to breed a super human; he might be a messiah. His struggle is at the center
of a nexus of powerful people and events, and the repercussions will be felt throughout
the Imperium.

TITLE: A Song Of Ice And Fire Series
AUTHOR: George R.R. Martin
DESCRIPTION: As the Seven Kingdoms face a generation-long winter, the noble Stark family
confronts the poisonous plots of the rival Lannisters, the emergence of the
White Walkers, the arrival of barbarian hordes, and other threats.

`

func TestSearch(t *testing.T) {
	multiline.Index(&content)
	for _, item := range cases {
		caseresult := multiline.Search(item.param)

		if len(caseresult) != len(item.result) {
			for i := 0; i < len(item.result); i++ {
				if item.result[i] != caseresult[i] {
					t.Errorf("Given parameter %s expects %v, but it got %v", item.param, item.result, caseresult)
				}
			}
		}

	}
}
