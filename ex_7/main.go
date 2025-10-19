package main
/*
Реализовать безопасную для конкуренции запись данных в структуру map.

Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).

Проверьте работу кода на гонки (util go run -race).
*/


import (
	"fmt"
	"sync"
)

func main() {
	var syncMap sync.Map
	syncMap.Store("key", "value from syncmap")
	syncMap.Store(1, 3)

	val, _ := syncMap.Load("key")
	val2, _ := syncMap.Load(1)

	fmt.Println(val)
	fmt.Println(val2)

}