package main

import "fmt"

func main() {
	sales :=[]int{12_000, 8_000, 15_000, 8_000}
	sum:=0
/*#######################################
//#########								#
//#  V.1. #								#
//#########								#
	sales[0]+=sales[0]*5/100//			#
	sales[2]+=sales[2]*5/100//			#
	sum = sales[0]+sales[2]//			#
	fmt.Println(sum)//					#
#######################################*/



//#######################################
//#########								#
//#  V.2. #								#
//#########								#
	for i:=0; i<len(sales); i++ {//		#
		if sales[i]>=10_000 {//			#
			sales[i]+=sales[i]*5/100//	#
			sum+=sales[i]//				#
		}//								#
	}//									#
	fmt.Println(sum)//					#
//#######################################
}