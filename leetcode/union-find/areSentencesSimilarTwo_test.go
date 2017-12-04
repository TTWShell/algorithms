package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_areSentencesSimilarTwo(t *testing.T) {
	assert := assert.New(t)

	assert.False(areSentencesSimilarTwo([]string{"great"}, []string{}, [][]string{}))

	assert.True(areSentencesSimilarTwo([]string{"great"}, []string{"great"}, [][]string{}))
	assert.False(areSentencesSimilarTwo([]string{"great"}, []string{"great1"}, [][]string{}))

	assert.True(areSentencesSimilarTwo(
		[]string{"great", "acting", "skills"},
		[]string{"fine", "drama", "talent"},
		[][]string{{"great", "fine"}, {"acting", "drama"}, {"skills", "talent"}}))

	assert.False(areSentencesSimilarTwo(
		[]string{"great", "acting", "skills"},
		[]string{"fine", "painting", "talent"},
		[][]string{{"great", "fine"}, {"acting", "drama"}, {"skills", "talent"}}))

	assert.False(areSentencesSimilarTwo(
		[]string{"fine", "painting", "talent"},
		[]string{"great", "acting", "skills"},
		[][]string{{"great", "fine"}, {"acting", "drama"}, {"skills", "talent"}}))

	assert.True(areSentencesSimilarTwo(
		[]string{"great", "acting", "skills"},
		[]string{"fine", "drama", "talent"},
		[][]string{{"great", "good"}, {"fine", "good"}, {"acting", "drama"}, {"skills", "talent"}}))

	words1 := []string{"this", "summer", "thomas", "get", "really", "very", "rich", "and", "have", "any", "actually", "wonderful", "and", "good", "truck", "every", "morning", "he", "drives", "an", "extraordinary", "truck", "around", "the", "nice", "city", "to", "eat", "some", "extremely", "extraordinary", "food", "as", "his", "meal", "but", "he", "only", "entertain", "an", "few", "well", "fruits", "as", "single", "lunch", "he", "wants", "to", "eat", "single", "single", "and", "really", "healthy", "life"}
	words2 := []string{"this", "summer", "thomas", "get", "very", "extremely", "rich", "and", "possess", "the", "actually", "great", "and", "wonderful", "vehicle", "every", "morning", "he", "drives", "unique", "extraordinary", "automobile", "around", "unique", "fine", "city", "to", "drink", "single", "extremely", "nice", "meal", "as", "his", "super", "but", "he", "only", "entertain", "a", "few", "extraordinary", "food", "as", "some", "brunch", "he", "wants", "to", "take", "any", "some", "and", "really", "healthy", "life"}
	pairs := [][]string{{"good", "wonderful"}, {"nice", "well"}, {"fine", "extraordinary"}, {"excellent", "good"}, {"wonderful", "nice"}, {"well", "fine"}, {"extraordinary", "excellent"}, {"great", "wonderful"}, {"one", "the"}, {"a", "unique"}, {"single", "some"}, {"an", "one"}, {"the", "a"}, {"unique", "single"}, {"some", "an"}, {"any", "the"}, {"car", "wagon"}, {"vehicle", "car"}, {"auto", "vehicle"}, {"automobile", "auto"}, {"wagon", "automobile"}, {"truck", "wagon"}, {"have", "have"}, {"take", "take"}, {"eat", "eat"}, {"drink", "drink"}, {"entertain", "entertain"}, {"meal", "food"}, {"lunch", "breakfast"}, {"super", "brunch"}, {"dinner", "meal"}, {"food", "lunch"}, {"breakfast", "super"}, {"brunch", "dinner"}, {"fruits", "food"}, {"own", "own"}, {"have", "have"}, {"keep", "keep"}, {"possess", "own"}, {"very", "very"}, {"super", "super"}, {"really", "really"}, {"actually", "actually"}, {"extremely", "extremely"}}
	assert.False(areSentencesSimilarTwo(words1, words2, pairs))
}
