package cerrors

type CErr struct {
	Msg string
}
func (self *CErr) Error() string {
	return self.Msg
}
