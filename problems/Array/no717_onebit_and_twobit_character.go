package arr

//这道题是看最后一个字符 是否能组成单独的0,那么就需要看，他的前一位是否为1
//并且该1 不能和n-2位组成组合
func IsOneBitCharacter(bits []int) bool {
	if bits[len(bits)-1] != 0 {
		return false
	}
	for i := 0; i < len(bits)-1; i++ {
		if bits[i] == 0 {
			continue
		}
		if bits[i] == 1 {
			if i == len(bits)-2 {
				return false
			} else {
				i++
			}
		}
	}
	return true
}
