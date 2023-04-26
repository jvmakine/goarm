package class_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConstants(t *testing.T) {
	t.Run("returns all strings from a class", func(t *testing.T) {
		clazz := classFrom(t, "../testdata/com/github/jvmakine/test/Hello.class")
		var result []string
		for _, s := range clazz.Constants().Strings() {
			result = append(result, s.Text())
		}
		require.Len(t, result, 24)
		require.Equal(t, []string{
			"java/lang/Object",
			"<init>",
			"()V",
		}, result[:3])
	})
	t.Run("returns all class infos", func(t *testing.T) {
		clazz := classFrom(t, "../testdata/com/github/jvmakine/test/Hello.class")
		var result []string
		for _, s := range clazz.Constants().ClassInfos() {
			result = append(result, s.Name().Text())
		}
		require.Equal(t, []string{
			"java/lang/Object",
			"com/github/jvmakine/test/Hello",
			"java/lang/System",
			"java/io/PrintStream",
			"java/io/Serializable",
		}, result)
	})
}
