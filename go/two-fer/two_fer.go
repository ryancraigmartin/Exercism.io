package twofer

// ShareWith accepts a name and returns a string
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
