package main

/*
#include <stdio.h>

int chengfa99_C() {
    int i, j;
    for (i=1; i<=9; ++i) {
        for (j=1; j<i; ++j) {
            printf("%d*%d=%-3d", j, i, j*i);
        }
        printf("%d*%d=%d\n", j, i, j*i);
    }
	return 33;
}
*/
//import "C"
import (
	"fmt"
)

func chengfa99_GO() {
	var i, j int
	for i = 1; i <= 9; i++ {
		for j = 1; j < i; j++ {
			fmt.Printf("%dx%d=%-2d ", j, i, j*i)
		}
		fmt.Printf("%dx%d=%d\n", j, i, j*i)
	}

}
