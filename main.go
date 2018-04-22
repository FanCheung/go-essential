package main

import ("fmt")
var println=fmt.Println
// accept pointer
func foo(arr *[4]*string){
println("pass by reference:",arr)
}
func arrays(){
/** Array *********************************/
// empty array
var arrA [9]int
fmt.Println("With an initial length:",arrA)
// initiated with value 
arrB:=[6]int{3,4,3,2,6}
fmt.Println("Initated:",arrB)
arrC:=[...]int{1,2,13,4,3,2,6}
// pass by reference

fmt.Println("Not fixed sized array:",arrC)

// pointer array with *
arrD:=[4]*string{new(string),new(string)}
*arrD[0]="hello"

// pass a arry by reference
foo(&arrD)

fmt.Println("Pointer type array - ","memory: ",arrD[0],"value: ",*arrD[0])
// multi dimensional array
arrE:= [4][2]int{{1,1},{1,1},{1,1},{1,1}}
println("Multi dimensional array: ",arrE)
var arrF [4][2]int
// copy array 

arrF=arrE
println(arrF)

}

func main(){
/* Slice ************************************/
// declare a slice very similar to array but not specifying index
var sliceC []int

var arrB=[3]int{1,2,3}
fmt.Println("Not to fix length:",sliceC)

// append item
sliceC=append(sliceC,3,1)
fmt.Println("Append item",sliceC)
// this won't work. arrB is not slice
// fmt.Println("Append item",append(arrB,1))
// however this works
fmt.Println("Append item to an array",append(arrB[0:],1))

// make slice length of 8 and capacity of 10
sliceD:=make([]int, 8,10)
// copy slice 
copy(sliceD,sliceC)
fmt.Println("Copy slice:",sliceD)

//slice array, last index exclusive
fmt.Println("Slice array:",arrB[1:3])
fmt.Println("Slice array to index 2:",arrB[:3])
fmt.Println("Slice array to the end from 0:",arrB[0:])

/* Map  *************************************/
// only declare not initialized
// var mapA map[string]int
// that will result in error since it's not initialized
mapB:= make(map[string]string)
mapB["hi"]="there"
mapB["hello"]="there"
fmt.Println("A string map:",mapB)
mapA:=make(map[int]string)
mapA[1]="hi"
fmt.Println("An int map:",mapA)
// delete a non existing item doesnt have effect
delete(mapA,0)
delete(mapA,1)
fmt.Println("deleted an item from int map:",mapA)

// declare and init a map
mapC:=map[string]int{
"MON":1,
"TUE":2,
"WED":3,
}
fmt.Println("Declared and initialied:",mapC)
//Map can return two values 

// declare and init a nested map
mapD:= map[string]map[string]int{
	"a":{"b":1},
}

fmt.Println("Nested map:",mapD)

// check existence
value,exist:=mapB["notExist"]
if !exist {
	fmt.Println("Two value return from an non existing Map item:",value,exist)}
 // declare and assign value
var words="hello"
//  declare a constant
const pi float32=3.14
fmt.Println(pi)
// declare multiple var
var (
	a=1 
	b=2 
)
// constant 
 c:=3.3333
fmt.Println(int(c))
// casting

fmt.Print(a,b)
 // same as words:="hello"
 // length of string
	fmt.Print("hello length",len(words))
	// accessing the first char
	fmt.Print("hello"[0])
	// string concatenation
	fmt.Print("hello"+" world!")
	// fmt.Scanf("%s",k)
	arrays()
}

func controls(){
	// for loop
	for i:=0; i<5; i++{
		// if 
		if i%2 ==0 {
		fmt.Println(i)
		}
		// switch
		switch i{
			case 1:fmt.Println("prime one")
			case 3:fmt.Println("prime two")
			default:fmt.Println("not prime")
		}
	}

}
