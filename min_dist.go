//Минимальное расстояние
package main
import (
	"fmt" // стандартный пакет, там ввод/вывод
	"github.com/skorobogatov/input" // пакет, написанный Скорбом, его нужно отдельно ставить
)						//в лекциях написано как..но там для Linux...как в винде поставить не знаю...
func main (){
	var ( // переменные можно объявлять блоком
		s string 
		nx, ny, md int
		x, y, symb rune // в GO, как я поняла, нет типа char. Есть вот этот rune. Он синоним int. В общем, хранит значение юникода буквы
	)
	input.Scanf ("%s%c%c%c%c", &s, &symb, &x, &symb, &y) // вот такой ввод из стндрт потока. Написал Скорб, ибо стандартный 
										//скан не считывает отдельно буквы..вроде так
	arrS := []rune(s) // чудо-операция, которая делает из строки срез, чтобы можно было поэлементно к ней обращаться
	nx = -1
	ny = -1
	md = len(arrS)
	for i, l := range arrS { // особый вид for-а просматривает все элементы "массива"/среза
		switch l {
			case x :
				nx = i
				if  (ny != -1) && (nx - ny < md) {
					md = nx - ny
				}
			case y :
				ny = i
				if (nx != -1) && (ny - nx < md) {
					md = ny - nx
				}
		}
	}
	fmt.Printf ("%d\n", md - 1)
}
