package main

import "fmt"

func main() {
	var N int
	var sim string
	_, _ = fmt.Scan(&N)
	_, _ = fmt.Scan(&sim)
	sandglass(N, sim)
}

func sandglass(N int, sim string) {
	seredina := N/2 - 1
	for j := 1; j <= N; j++ { // osnovanie verh
		fmt.Printf("%s", sim)
		if j == N {
			fmt.Println()
		}
	}
	nacalo, konec := 1, N-2
	probel := " "
	for i := 1; i <= konec; i++ {
		if i == seredina { // centre glass
			for j := 0; j <= nacalo-1; j++ {
				fmt.Printf("%s", probel)
			}
			for j := nacalo - 1; j <= konec-1; j++ {
				fmt.Printf("%s", sim)
			}
			fmt.Println()
		} else if i < seredina { // verh glass
			for j := 0; j <= konec; j++ {
				if (j == nacalo || j == konec) && nacalo != konec {
					fmt.Printf("%s", sim)
				} else {
					fmt.Printf("%s", probel)
				}
			}
			nacalo++
			konec--
			fmt.Println()
		} else {
			for i := 1; i <= seredina-1; i++ {
				for j := 0; j <= konec+1; j++ {
					if (j == nacalo-1 || j == konec+1) && nacalo != konec {
						fmt.Printf("%s", sim)
					} else {
						fmt.Printf("%s", probel)
					}
				}
				nacalo--
				konec++
				fmt.Println()
			}
			break
		}
	}
	for j := 1; j <= N; j++ { // osnovanie niz
		fmt.Printf("%s", sim)
		if j == N {
			fmt.Println()
		}
	}
}
