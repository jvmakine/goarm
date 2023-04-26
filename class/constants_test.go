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
		require.Len(t, result, 45)
		require.Equal(t, []string{
			"com/github/jvmakine/test/Hello",
			"java/lang/Object",
			"java/io/Serializable",
		}, result[:3])
	})
	t.Run("returns all class infos", func(t *testing.T) {
		clazz := classFrom(t, "../testdata/com/github/jvmakine/test/Hello.class")
		var result []string
		for _, s := range clazz.Constants().ClassInfos() {
			result = append(result, s.Name().Text())
		}
		require.Equal(t, []string{
			"com/github/jvmakine/test/Hello",
			"java/lang/Object",
			"java/io/Serializable",
			"java/lang/System",
			"java/io/PrintStream",
			"kotlin/jvm/internal/Intrinsics",
		}, result)
	})
}
