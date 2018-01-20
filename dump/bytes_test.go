package dump

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytes(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		var suites = []struct {
			Actual   interface{}
			Expected []byte
		}{
			{
				uint8(2),
				[]byte{2},
			},
			{
				uint16(3),
				[]byte{0, 3},
			},
			{
				uint32(4),
				[]byte{0, 0, 0, 4},
			},
			{
				uint64(5),
				[]byte{0, 0, 0, 0, 0, 0, 0, 5},
			},
			{
				int8(1),
				[]byte{1},
			},
			{
				int16(1),
				[]byte{0, 1},
			},
			{
				int32(1),
				[]byte{0, 0, 0, 1},
			},
			{
				int64(1),
				[]byte{0, 0, 0, 0, 0, 0, 0, 1},
			},
			{
				float32(1),
				[]byte{63, 128, 0, 0},
			},
			{
				float64(1),
				[]byte{63, 240, 0, 0, 0, 0, 0, 0},
			},
			{
				struct{ A uint8 }{5},
				[]byte{5},
			},
			{
				struct{ A uint16 }{5},
				[]byte{0, 5},
			},
			{
				struct{ A, B uint16 }{5, 5},
				[]byte{0, 5, 0, 5},
			},
		}

		for _, suite := range suites {
			data, err := Bytes(binary.BigEndian, suite.Actual)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, suite.Expected, data)
		}
	})

	t.Run("Fail", func(t *testing.T) {
		var suites = []interface{}{
			uint(2),
			int(2),
			uintptr(2),
			reflect.Value{},
		}

		for _, suite := range suites {
			_, err := Bytes(binary.BigEndian, suite)
			assert.Error(t, err)
		}
	})
}

func TestBytesTo(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		var suites = []struct {
			Actual   interface{}
			Expected []byte
		}{
			{
				uint8(2),
				[]byte{2},
			},
			{
				uint16(3),
				[]byte{0, 3},
			},
			{
				uint32(4),
				[]byte{0, 0, 0, 4},
			},
			{
				uint64(5),
				[]byte{0, 0, 0, 0, 0, 0, 0, 5},
			},
			{
				int8(1),
				[]byte{1},
			},
			{
				int16(1),
				[]byte{0, 1},
			},
			{
				int32(1),
				[]byte{0, 0, 0, 1},
			},
			{
				int64(1),
				[]byte{0, 0, 0, 0, 0, 0, 0, 1},
			},
			{
				float32(1),
				[]byte{63, 128, 0, 0},
			},
			{
				float64(1),
				[]byte{63, 240, 0, 0, 0, 0, 0, 0},
			},
			{
				struct{ A uint8 }{5},
				[]byte{5},
			},
			{
				struct{ A uint16 }{5},
				[]byte{0, 5},
			},
			{
				struct{ A, B uint16 }{5, 5},
				[]byte{0, 5, 0, 5},
			},
		}

		for _, suite := range suites {
			buf := bytes.NewBuffer(nil)
			err := BytesTo(buf, binary.BigEndian, suite.Actual)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, suite.Expected, buf.Bytes())
		}
	})

	t.Run("Fail", func(t *testing.T) {
		var suites = []interface{}{
			uint(2),
			int(2),
			uintptr(2),
			reflect.Value{},
		}

		buf := bytes.NewBuffer(nil)

		for _, suite := range suites {
			err := BytesTo(buf, binary.BigEndian, suite)
			assert.Error(t, err)
		}
	})
}
