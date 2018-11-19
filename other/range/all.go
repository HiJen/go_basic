package main

func ForSlice(s []string) {
	len := len(s)
	for i := 0; i < len; i++ {
		_, _ = i, s[i]
	}
}

//func RangeForSlice(s []string) {
//	for i, v := range s {
//		_, _ = i, v
//	}
//}

//Performance booster, for get elment by s[i]
func RangeForSlice(s []string) {
	for i, _ := range s {
		_, _ = i, s[i]
	}
}
