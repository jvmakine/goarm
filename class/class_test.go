package class_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/jvmakine/goarm/class"
	"github.com/jvmakine/goarm/classfile"
	"github.com/stretchr/testify/require"
)

func TestAtrributes(t *testing.T) {
	t.Run("returns attributes for a public concrete class", func(t *testing.T) {
		clazz := classFrom(t, "../testdata/Hello.class")
		access := clazz.AccessFlags()

		require.True(t, access.IsPublic())
		require.True(t, access.IsFinal())
		require.True(t, access.IsSuper())
		require.False(t, access.IsAbstract())
		require.False(t, access.IsAnnotation())
		require.False(t, access.IsEnum())
		require.False(t, access.IsInterface())
		require.False(t, access.IsSynthetic())
	})
	t.Run("changes attributes", func(t *testing.T) {
		clazz := classFrom(t, "../testdata/Hello.class")
		clazz.AccessFlags().SetSynthetic(true)

		require.True(t, clazz.AccessFlags().IsSynthetic())
	})
}

func TestClassInfo(t *testing.T) {
	t.Run("returns name of the class", func(t *testing.T) {
		clazz := classFrom(t, "../testdata/Hello.class")
		require.Equal(t, "com/github/jvmakine/test/Hello", clazz.ThisClass().Name().Text())
	})
	t.Run("returns name of the super class", func(t *testing.T) {
		clazz := classFrom(t, "../testdata/Hello.class")
		require.Equal(t, "java/lang/Object", clazz.SuperClass().Name().Text())
	})
	t.Run("returns interfaces correctly", func(t *testing.T) {
		clazz := classFrom(t, "../testdata/Hello.class")
		var names []string
		for _, ci := range clazz.Interfaces() {
			names = append(names, ci.Name().Text())
		}
		require.Equal(t, []string{"java/io/Serializable"}, names)
	})
}

func classFrom(t *testing.T, path string) *class.Class {
	t.Helper()

	data, err := ioutil.ReadFile(path)
	require.NoError(t, err)
	classFile, err := classfile.Parse(bytes.NewReader(data))
	require.NoError(t, err)
	return class.NewClass(classFile)
}
