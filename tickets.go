package main

import( "fmt"
        "time"
		"sync"
)
var (
	totalTickets = 10 //available tickets
	mutex sync.Mutex //prevent race condition
	n
)

//bookTicket function to book tickets
func BookTicket(user string, tickets int, wg * sync.WaitGroup, ch chan string) {
	defer wg.Done() //decrement the counter when done

	mutex.Lock() //lock the mutex to avoid concurrnt modification
	defer mutex.Unlock() //unlock the mutex after modification
	if totalTickets >= tickets{
		fmt.Printf("%s booked %d tickets successfully!\n", user, tickets)
		totalTickets -= tickets
		ch <- fmt.Sprintf("%s booked %d tickets succesfuly!\n", user, tickets)
	}else{
		fmt.Printf("sorry %s, only %d tickets left!\n", user, totalTickets)
		ch <- fmt.Sprintf("%s failed to book %d tickets", user, tickets)
	}
	//mutex.Unlock() //unlock the mutex after modification

}

func main(){
	var wg sync.WaitGroup
	bookingChannel := make(chan string, 5) //buffered channel to store booking results
	//simulating multiple users trying to book tickets concurrently
	users := []struct{
		Name string
		Tickets int
	}{
		{"Alice", 3},
		{"Bob", 2},
		{"manoti", 4},
		{"james", 5},
		{"jane", 1},
	}
	for _, user := range users{
		wg.Add(1)
		go BookTicket(user.Name, user.Tickets, &wg, bookingChannel)
		time.Sleep(200 * time.Millisecond) //sleep for 200ms to simulate real world scenario
	}
	wg.Wait() //wait for all goroutines to finish
	close(bookingChannel) //close the channel

	fmt.Println("\nBooking results:")
	for result :=range bookingChannel{
		fmt.Println(result)
	}
	fmt.Println("\nTickets left:", totalTickets)

}