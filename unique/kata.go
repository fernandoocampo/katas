package unique

func FindUniq(arr []float32) float32 {
	result := float32(0)
	filter := make(map[float32]int, 2)

	for i := 0; i < len(arr); i++ {
		filter[arr[i]]++
	}

	for key, amount := range filter {
		if amount > 1 {
			continue
		}
		result = key
	}

	return result
}

func FindUniqTwo(arr []float32) float32 {
	if arr[0] != arr[1] && arr[1] == arr[2] {
		return arr[0]
	}
	for i, v := range arr[1:] {
		if v != arr[i] {
			return v
		}
	}
	return 0
}

func FindUniqThree(arr []float32) float32 {
	alreadyChecked := make(map[float32]bool)
	ch := make(chan float32)
	length := len(arr)

	for i, number := range arr {
		if !alreadyChecked[number] {
			alreadyChecked[number] = true
			go func(curr float32, currI int, ch chan float32) {
				for currI < length {
					if arr[currI] == curr {
						return
					}
					currI++
				}
				ch <- curr
				return
			}(number, i+1, ch)
		}
	}

	num := <-ch
	return num
}
