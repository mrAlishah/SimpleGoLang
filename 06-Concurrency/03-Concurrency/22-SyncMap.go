//28. Sync Map Exercise One Solution

package main

/*
	type Map
		func (m *Map) Delete(key any)
		func (m *Map) Load(key any) (value any, ok bool)
		func (m *Map) LoadAndDelete(key any) (value any, loaded bool)
		func (m *Map) LoadOrStore(key, value any) (actual any, loaded bool)
		func (m *Map) Range(f func(key, value any) bool)
		func (m *Map) Store(key, value any)
*/

import (
	"fmt"
	"sync"
	"time"
)

func printMap(m *sync.Map, title string){
	<-time.After(5 * time.Microsecond)
	fmt.Println("\n-------------------------------",title)
	//m.Range(func(key, value interface{}) bool{
	m.Range(func(key, value any) bool{
		fmt.Printf("Key= %v -> Value= %v\n",key,value)
		return true
	})
}

func main(){

	m := sync.Map{}

	m.Store("key 1","value 1")
	m.Store("key 2","value 2")
	m.Store("key 3","value 3")
	m.Store("key 4","value 4")
	m.Store("key 5","value 5")
	printMap(&m,"m.Store")

	los1, ok := m.LoadOrStore("key 6","value 6")
	printMap(&m,"m.LoadOrStore")
	if ok { fmt.Println("Find and loaded")} else { fmt.Println("Created item") }
	fmt.Printf("Key 6: %v  -> ok: %v\n",los1, ok)

	m.Delete("key 5")
	printMap(&m,"m.Delete")

	los2, ok := m.Load("key 3")
	printMap(&m,"m.Load")
	if ok { fmt.Printf("Key 3: %v \n",los2)} else { fmt.Println("Key 3 not find!")}


}