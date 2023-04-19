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
		data, err := ioutil.ReadFile("../testdata/Hello.class")
		require.NoError(t, err)
		classFile, err := classfile.Parse(bytes.NewReader(data))
		require.NoError(t, err)
		clazz := class.NewClass(classFile)
		require.Equal(t, map[class.AccessFlag]bool{
			class.ACC_PUBLIC: true,
			class.ACC_FINAL:  true,
			class.ACC_SUPER:  true,
		}, clazz.AccessFlags())
	})
}
