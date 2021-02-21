

//Slice = list
//map = dict

//You will see some simalarities to C here, we declare a variable along with its type and contents (1)
// we then point to the location of our var in memory (2)
//next we must derefence the address from whichever registry is being used so that we can call a function
//now we reassign ptr to "100"
var count = int(42) 
ptr := &count
fmt.Println(*ptr) x*ptr = 100
fmt.Println(count)


//here we create a new data type for further use down the track, we say that the name is Person, with two fields(name and age)
//next we define a method attached to the type called SayHello which will print a defined string aswell as whatever is in the persons Name field

type Person struct { Name string
	Age int }
	func (p *Person) SayHello() { fmt.Println("Hello,", p.Name)
	   }
	   func main() {
	var guy = new(Person) {guy.Name = "Dave"
	| guy.SayHello()
	}


//an interface is a way of a type fulfilling a contract to be considered another type EG:
type Friend interface { SayHello()
}
//as above if say, a newly created "person" or "dog" type was able to say hello (does call the SayHello method), the would then be considered a friend and can be referenced as such in future code, let's look at this further.
//first we create a function that says if the type calling greet meets being a being a friend then go ahead with saying hello
func Greet(f Friend) { 
	f.SayHello()
}
//remember that new data type we create with struct and the method we implemented in it earlier? well now let's create a variable called guy which will be a new TYPE of PERSON with his name field as dave before we GREET him.

func main() {
	var guy = new(Person) guy.Name = "Dave" Greet(guy)
	}

//now notice that we are able to greet, as the datatype person is able to SayHello because we coded this as a method when we declared the type

//but we don't JUST want to say hello to people, we like dogs too!, let's create another datatype and implement a SayHello method:

type Dog struct {}
func (d *Dog) SayHello() {
	fmt.Println("Woof woof") 
}
func main() {
	var guy = new(Person) guy.Name = "Dave"
	Greet(guy)
	var dog = new(Dog)
	Greet(dog)
}

//closing  off interfaces and structs for now I will write a program utilising types and lists aswell as user input to demonstrate knowledge.package golang


//loop example, using boolean true, looks cleaner than bash ans JS, similar to python
if x == 1 {
	fmt.Println("X is equal to 1")
	} else {
	fmt.Println("X is not equal to 1")
	}

//here is what i would call an elif conditional equivalent, again, looks clean, the check is initiated by "switch" and "elif" appears to be replaced with "case" and finally "else" with default (side note, go infers x to be a string) - it was worth noting that this is handled differently to elif, as switch is actually taking these cases and COMPARING them against x.

switch x{
	case "foo":
		fmt.Println("Found foo") 
	case "bar":
		fmt.Println("Found bar") 
	default:
		fmt.Println("Default case")
}

//a variant on our friend mr switch here is a "type switch", in whch we check for the type of case, not the contents. The below example is attempting to retrieve the type of interface. The switch syntax is v := i.(type), meaning we declare v as being equal to i and it's type. the cases then get compared to switch and handle accordingly on a match or no match

func foo(i interface{}) { 
	switch v := i.(type) {
	case int:
		fmt.Println("I'm an integer!") 
	case string:
		fmt.Println("I'm a string!") 
	default:
		fmt.Println("Unknown type!") 
	}
}

//for loop, same logica as seen in bash and js, explained [HERE](<link to explaination)

for i := 0; i < 10; i++ { 
	fmt.Println(i)
}

//a more complex forloop example looping over a list and dict, or as go calls them, a slice and a map.
//a list called nums is declared, with the contents being 4 integers
//we then use range to iterate over our list, for each index and it's corresponding value, they will be printed.
//so, the range of our list is 4, so the output would be as follows: 1:2, 3:4, 6:7, 8:9

nums := []int{2,4,6,8}
for idx, val:= range nums {
fmt.Println(idx, val) 
}

//one of the things that makes go powerful - concurrency through "goroutines".
//"goroutines" can be functions or methods able to run simultaneously (sometimes described as lightweight threads as the cost of creating them is minimal compared to actual threads)
//example below, when we jump into main notice we prface our dummy function with go, meaning that main() will CONTINUE WITHOUT WAITING FOR F TO COMPLETE

func f() {
	fmt.Println("f function")
}

func main() { 
	go f()
	time.Sleep(1 * time.Second)
	fmt.Println("main function") 
}

//go contains a data type called channels that provide a mechanism through which goroutines can synchronise their execution and communicate with one another.
//below we define a var c of type chan int. We choose int as the type for channel because we will be passing through the lengths of various strings between goroutines.
//we declare whether data is flowing to or from a channel with <-
//our function accepts a string as an argument aswell as a channel, it then places the length of the string into the channel.
//the main function pieces everything together

func strlen(s string, c chan int) {
	c <- len(s)
}

func main() {
	c := make(chan int)
	go strlen("Salutations", c)
	go strlen("World", c) 
	x, y := <-c, <-c
	fmt.Println(x, y, x+y) 
}

