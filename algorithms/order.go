package algorithms

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		s := target - v
		if j, ok := m[s]; ok {
			return []int{i, j}
		}
		m[v] = i
	}
	return nil
}

func mergeAlternately(word1 string, word2 string) string {
	m, n := len(word1), len(word2)
	ans := make([]byte, 0, m+n)
	for i := 0; i < m || i < n; i++ {
		if i < m {
			ans = append(ans, word1[i])
		}
		if i < n {
			ans = append(ans, word2[i])
		}
	}
	return string(ans)
}

func gcdOfStrings(str1 string, str2 string) string {
	l1, l2 := len(str1), len(str2)
	if l1 < l2 {
		return gcdOfStrings(str2, str1)
	}
	if l1 == l2 {
		if str1 == str2 {
			return str1
		}
		return ""
	}
	for i := 0; i < l2; i++ {
		if str1[i] != str2[i] {
			return ""
		}
	}
	return gcdOfStrings(str1[l2:], str2)
}

func gcdOfStrings2(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[:gcd(len(str1), len(str2))]
}

func gcd(a, b int) int {
	temp := 0
	for b != 0 {
		temp = a % b
		a, b = b, temp
	}
	return a
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	ans := make([]bool, len(candies))
	max := 0
	for _, v := range candies {
		if max < v {
			max = v
		}
	}
	for i, v := range candies {
		ans[i] = v+extraCandies >= max
	}
	return ans
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	i := 0
	l := len(flowerbed)
	for i < l && n > 0 {
		if flowerbed[i] == 1 {
			i += 2
		} else {
			if i+1 == l || flowerbed[i+1] == 0 {
				n--
				i += 2
			} else {
				i++
			}
		}

	}

	return n == 0
}
