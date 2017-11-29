package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_areSentencesSimilar(t *testing.T) {
	assert := assert.New(t)

	assert.False(areSentencesSimilar([]string{"great"}, []string{}, [][]string{}))

	assert.True(areSentencesSimilar([]string{"great"}, []string{"great"}, [][]string{}))

	assert.True(areSentencesSimilar(
		[]string{"great", "acting", "skills"},
		[]string{"fine", "drama", "talent"},
		[][]string{{"great", "fine"}, {"acting", "drama"}, {"skills", "talent"}}))

	assert.False(areSentencesSimilar(
		[]string{"great", "acting", "skills"},
		[]string{"fine", "painting", "talent"},
		[][]string{{"great", "fine"}, {"acting", "drama"}, {"skills", "talent"}}))
}
