package classfile_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/jvmakine/goarm/classfile"
	"github.com/stretchr/testify/require"
)

func TestRoundTrip(t *testing.T) {
	t.Run("round trips empty class correctly", func(t *testing.T) {
		data, err := ioutil.ReadFile("../testdata/Empty.class")
		require.NoError(t, err)

		classFile, err := classfile.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		buf := bytes.Buffer{}
		err = classfile.Write(classFile, &buf)
		require.NoError(t, err)

		require.Equal(t, data, buf.Bytes())
	})
	t.Run("round trips hello class correctly", func(t *testing.T) {
		data, err := ioutil.ReadFile("../testdata/Hello.class")
		require.NoError(t, err)

		classFile, err := classfile.Parse(bytes.NewReader(data))
		require.NoError(t, err)

		buf := bytes.Buffer{}
		err = classfile.Write(classFile, &buf)
		require.NoError(t, err)

		require.Equal(t, data, buf.Bytes())
	})
}
