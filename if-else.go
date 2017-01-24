package main

func equal(a, b int) bool {
	return a == b
}

func test_if(a int) byte {
	if a == 100 {
		return 'A'
	} else if a > 85 {
		return 'B'
	} else if a > 60 {
		return 'C'
	} else {
		return 'D'
	}
}
