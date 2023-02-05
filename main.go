package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Taille max de notre tableau de bytes
	const SIZE int = 2048
	// Taille max de notre chaîne en argument
	const MAXOP int = 4096

	if len(os.Args) != 2 || len(os.Args[1]) > MAXOP {
		return
	}

	// Création et allocation de notre tableau de bytes
	brainfuck := make([]byte, SIZE)

	// Initialisation de tous les (2048) éléments du tableau à 0
	for i := 0; i < SIZE; i++ {
		brainfuck[i] = 0
	}

	// Variable pour stocker la position du pointeur
	pointerIndex := 0
	// Au début, le pointeur stocke l'adresse du premier élément du tableau
	pointer := &brainfuck[pointerIndex]

	// Compteur pour parcourir la chaîne fournie en argument
	counter := 0
	// Récupération et stockage de ladite chaîne dans un tableau de bytes
	op := []byte(os.Args[1])
	// Taille de la chaîne (Nombre d'opérations) == Taille du tableau
	opNumb := len(op)
	// Tant que notre compteur est inférieur à la taille du tableau on entre dans la boucle
	for counter < opNumb {
		switch op[counter] {
		case '>':
			// Incrémentation du pointeur; i.e stocke l'adresse de l'élément qui suit dans le tableau brainfuck
			pointerIndex++
			if pointerIndex >= 0 && pointerIndex < SIZE {
				pointer = &brainfuck[pointerIndex]
			}
		case '<':
			// Décrémentation du pointeur; i.e stocke l'adresse de l'élément qui précède dans le tableau brainfuck
			pointerIndex--
			if pointerIndex >= 0 && pointerIndex < SIZE {
				pointer = &brainfuck[pointerIndex]
			}
		case '+':
			// Incrémentation de la valeur contenue à l'adresse stockée par le pointeur
			if *pointer <= 255 {
				*pointer += 1
			}
		case '-':
			// Décrémentation de la valeur contenue à l'adresse stockée par le pointeur
			if *pointer <= 255 {
				*pointer -= 1
			}
		case '.':
			// Affichage de la valeur contenue à l'adresse stockée par le pointeur
			z01.PrintRune(rune(*pointer))
		case '[':
			/*
			* Si la valeur contenue à l'adresse stockée par le pointeur **est nulle**,
			* nous cherchons le crochet fermant (]) correspondant et la boucle continue
			* à partir de l'opérateur (du caractère) suivant
			 */
			if *pointer == 0 {
				/*
				* Ces variables nous permettent de déterminer à partir d'un crochet ouvrant ([)
				* quel est le crochet fermant correspondant (])
				 */
				openBrackets := 0
				closeBrackets := 0
				for counter < opNumb {
					if op[counter] == '[' {
						openBrackets++
					} else if op[counter] == ']' {
						closeBrackets++
					}
					/*
					* Si le nombre de crochet.s ouvrant.s est égal au nombre de crochet.s fermant.s
					* alors on a atteint le crochet fermant correspondant au crochet ouvrant sur
					* lequel nous sommes positionnés. On sort donc de notre boucle (break)
					 */
					if openBrackets == closeBrackets {
						break
					}
					counter++
				}
				/*
				* Si nous sortons de la boucle sans trouver de crochet fermant correspondant,
				* nous arrêtons le programme
				 */
				if counter >= opNumb {
					return
				}
			}
		case ']':
			/*
			* Si la valeur contenue à l'adresse stockée par le pointeur **n'est pas nulle**,
			* nous cherchons le crochet ouvrant ([) correspondant et la boucle continue
			* à partir de l'opérateur (du caractère) suivant
			 */
			if *pointer != 0 {
				/*
				* Ces variables nous permettent de déterminer à partir d'un crochet fermant (])
				* quel est le crochet ouvrant correspondant ([)
				 */
				openBrackets := 0
				closeBrackets := 0
				for counter < opNumb {
					if op[counter] == '[' {
						openBrackets++
					} else if op[counter] == ']' {
						closeBrackets++
					}
					/*
					* Si le nombre de crochet.s ouvrant.s est égal au nombre de crochet.s fermant.s
					* alors on a atteint le crochet ouvrant correspondant au crochet fermant sur
					* lequel nous sommes positionnés. On sort donc de notre boucle (break)
					 */
					if openBrackets == closeBrackets {
						break
					}
					counter--
				}
				/*
				* Si nous sortons de la boucle sans trouver de crochet ouvrant correspondant,
				* nous arrêtons le programme
				 */
				if counter < 0 {
					return
				}
			}
		}
		// On passe à l'opérateur (au caractère) suivant
		counter++
	}
}
