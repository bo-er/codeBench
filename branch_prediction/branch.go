package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Duration(time.Now().UnixNano())
	defer fmt.Println(time.Duration(time.Now().UnixNano()) - start)
	a := []int{1, 3, 11, 19, 29, 31, 40, 46, 47, 54, 55, 60, 62, 65, 75, 84, 89, 91, 97, 105, 112, 118, 125, 134, 143, 151, 159, 167, 176, 177, 183, 185, 194, 202, 204, 214, 221, 229, 231, 237, 244, 248, 249, 254, 258, 262, 270, 279, 284, 294, 298, 306, 308, 318, 328, 329, 335, 344, 353, 357, 363, 365, 366, 372, 379, 386, 395, 397, 400, 404, 411, 415, 422, 425, 434, 442, 447, 455, 459, 466, 467, 471, 475, 483, 487, 489, 499, 503, 507, 509, 512, 521, 528, 535, 543, 544, 548, 551, 555, 561, 570, 576, 578, 584, 592, 600, 601, 602, 608, 609, 612, 621, 625, 627, 630, 635, 643, 651, 659, 661, 666, 673, 676, 678, 688, 695, 696, 700, 707, 717, 719, 722, 728, 734, 735, 743, 753, 762, 771, 776, 780, 785, 793, 796, 799, 806, 816, 817, 824, 826, 833, 834, 836, 841, 846, 852, 856, 858, 859, 862, 871, 879, 881, 890, 895, 898, 902, 905, 914, 918, 927, 934, 935, 941, 942, 947, 950, 958, 964, 966, 976, 977, 981, 982, 992, 993, 997, 998}
	target100 := []int{}
	target200 := []int{}
	target300 := []int{}
	target400 := []int{}
	target500 := []int{}
	target600 := []int{}
	target700 := []int{}
	target800 := []int{}
	target900 := []int{}
	target1000 := []int{}
	for _, i := range a {
		if i <= 100 {
			target100 = append(target100, i)
		} else if i <= 200 {
			target200 = append(target200, i)
		} else if i <= 300 {
			target300 = append(target300, i)
		} else if i <= 400 {
			target400 = append(target400, i)
		} else if i <= 500 {
			target500 = append(target500, i)
		} else if i <= 600 {
			target600 = append(target600, i)
		} else if i <= 700 {
			target700 = append(target700, i)
		} else if i <= 800 {
			target800 = append(target800, i)
		} else if i <= 900 {
			target900 = append(target900, i)
		} else if i <= 1000 {
			target1000 = append(target1000, i)
		}
	}
}
