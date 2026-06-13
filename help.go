package main

func split(sum int) (int , int){
	x :=sum *4/9
	y := sum - x
	return x,y
}