package Strategy

type TextSquare struct {
	PrintOutput
}

/*func (t *TextSquare) Print() error {
	r := bytes.NewReader([]byte("Circle"))
	io.Copy(t.Writer, r)
	return nil
}*/

func (t *TextSquare) Print() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}
