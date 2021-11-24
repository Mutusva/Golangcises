package algorithms


/*
Finds the first unique character in a string
*/

func UniqueCharacter(s string) int {

	for i,v := range s {
	   isUnique := true
      for j,u := range s {
		  if v == u && i != j {
			isUnique = false
			continue
		  }
	  }
	  if isUnique {
		  return i + 1
	  }
	}
	return -1
}

func UniqueCharacterImproved(s string) int {
    m := map[rune]int{}
	for _ ,char := range s {
	 v,ok := m[char]
	 if ok {
		 m[char] = v + 1
	 } else {
		 m[char] = 1
	 }
	}

	for i, r := range s {
		if m[r] == 1 {
			return i + 1
		}
	}
	return -1
}