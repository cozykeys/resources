# To Do List

- [x] Keyboard
- [x] Constants
- [x] Constant (LEAF)
- [x] Layers
- [x] Layer
- [x] Groups
- [x] Group
- [~] Children
- [x] Path
- [x] Components
- [x] AbsoluteLineTo
- [x] AbsoluteMoveTo
- [x] Point (EndPoint, ControlPoint) (LEAF)
- [x] Legend (LEAF)
- [x] Text (LEAF)
- [x] Spacer (LEAF)
- [x] Circle (LEAF)
- [x] Key

# Test Template

```
func Test_unmarshalFoo(t *testing.T) {
	testCases := map[string]struct {
		xml      []byte
		expected *models.Foo
	}{
		"happy_path": {
			xml: []byte(`<Foo />`),
			expected: &models.Foo{
                // TODO
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			doc := etree.NewDocument()

			err := doc.ReadFromBytes(testCase.xml)
			require.Nil(t, err)

			foo, err := unmarshalFoo(doc.Root())
			require.Nil(t, err)
			require.Equal(t, testCase.expected, foo)
		})
	}
}
```
