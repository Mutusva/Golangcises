package algorithms

func ShareWith(name string) string {
    if len(name) == 0 {
        name = "you"
    }
	return "One for " + name + ", one for me."
}

func ShareWithTwo(name string) string {
    if name == "" {
        name = "you"
    }
	return "One for " + name + ", one for me."
}