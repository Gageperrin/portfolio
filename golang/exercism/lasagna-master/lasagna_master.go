package lasagna

func CountOccurences(ingreds []string) map[string]int {
    dict := make(map[string]int)
    for _, v := range ingreds {
        dict[v]++
    }
    return dict
}

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avg int) (t int) {
	if avg == 0 {
		avg = 2
	}
	return len(layers) * avg
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	d := CountOccurences(layers)
	return (d["noodles"] * 50), (float64(d["sauce"]) * 0.2)
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(listOne, listTwo []string) {
	listTwo[len(listTwo)-1] = listOne[len(listOne)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(q []float64, portions int) (amt []float64) {
	for i, _ := range q {
        amt = append(amt, (q[i] * float64(portions) / 2))
    }
	return
}