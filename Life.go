// Just another one implementation of John Conway's "Life" game
// Made with golang
// 2022

package main
 
import (
    "fmt"
    "math/rand"
//    "time"
///	"os"
////	"os/exec"
)
 
const (
    width  			= 220 	// размеры расчитаны под мой дисплей, когда консоль (окно cmd) раздвинуто на максисмум. 
    height 			= 61  	// больше 63 строк в мой экран не вмещается 
	watchPeriod 	= 1500 	// продолжительность наблюдения
)
 
// ThePlanet является двухмерным полем клеток.
type ThePlanet [][]bool
 
// метод ЭтойПланеты -> CreateNewPlanet создает пустую планету.
func CreateNewPlanet() ThePlanet {
    planet := make(ThePlanet, height)
    for i := range planet {
        planet[i] = make([]bool, width)
    }
    return planet
}
 
// метод ЭтойПланеты -> Seed заполняет планету случайными живыми клетками.
func (planet ThePlanet) Seed() {
    for i := 0; i < (width * height / 4); i++ {
        planet.Set(rand.Intn(width), rand.Intn(height), true)
    }
}
 
// метод ЭтойПланеты -> Set устанавливает состояние конкретной клетки.
func (planet ThePlanet) Set(x, y int, b bool) {
    planet[y][x] = b
}
 
// метод ЭтойПланеты -> Alive сообщает, является ли клетка живой.
// Если координаты за пределами видимой поляны планеты, возвращаемся к началу (планета имеет форму эллипсоида и гауссову кривизну >0).
func (planet ThePlanet) Alive(x, y int) bool {
    x = (x + width) % width
    y = (y + height) % height
    return  planet[y][x]
}
 
// метод ЭтойПланеты -> Neighbors считает живых соседей.
func (planet ThePlanet) Neighbors(x, y int) int {
    n := 0
    for v := -1; v <= 1; v++ {
        for h := -1; h <= 1; h++ {
            if !(v == 0 && h == 0) && planet.Alive(x+h, y+v) {
                n++
            }
        }
    }
    return  n
}
 
// Next возвращает состояние конкретной клетки на следующем шаге.
func (planet ThePlanet) Next(x, y int) bool {
    n := planet.Neighbors(x, y)
    return n == 3 || (n == 2 && planet.Alive(x, y)) // либо 3 соседа у любой клетки: {из мертвой делают живую, а живую не убивают}, либо 2 соседа у живой клетки -> остается живой
	
}
 
// метод ЭтойПланеты -> String возвращает планету в одну строку с ретурнами "\n". Чтоб в методе печати плямс! - и все на экране.
func (planet ThePlanet) String() string {
    var b byte
    buffer := make([]byte, 0, (width+1)*height)
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            b = ' '
            if planet[y][x] {
                b = '*'
            }
            buffer = append(buffer, b)
        }
        buffer = append(buffer, '\n') // пИрИвод к0ретки
    }
    return string(buffer)
}
 
// метод ЭтойПланеты -> Show очищает экран и рисует планету.
func (planet ThePlanet) Show() {
//    fmt.Print("\x0c", planet.String())
//    fmt.Print("\033c", planet.String())
//    fmt.Print("\f", planet.String())
	fmt.Print("\033[H\033[J", planet.String()) // SIC! и еще раз SIC! Только эта последовательность возвращает курсовр в левый верхний угол \033[H и чистит эран от этой позиции \033[J
	// честно, почти час провел в безуспешных поисках истины

}
 
// Step обновляет состояние следующей планеты (b) из
// текущей планеты (a).
func Step(a, b ThePlanet) {
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            b.Set(x, y, a.Next(x, y))
        }
    }
}
 
func main() {
    a, b := CreateNewPlanet(), CreateNewPlanet() // В начале сотворил Бог небо и землю. Земля же была безвидна и пуста, и тьма над бездною, и Дух Божий носился ...
lbla: // да. дикий кайф воткнуть метку для безусловного перехода. так нельзя, но если хочется, то <<ВотЪ>>
//	cmd := exec.Command("cmd", "/c", "cls") // чистит экран только для Windows 
//    cmd.Stdout = os.Stdout
//    cmd.Run()								// в принципе не обязательно. Потому што удалось найти правильную ESC последовательность :*Р
	
	a.Seed() // И сказал Бог: да произрастит земля зелень, траву, сеющую семя, дерево плодовитое, приносящее по роду своему плод, в котором семя его на земле. И стало так.
			 // И сотворил Бог рыб больших и всякую душу животных пресмыкающихся, которых произвела вода, по роду их, и всякую птицу пернатую по роду её. И увидел Бог, что это хорошо.


    for i := 0; i < watchPeriod; i++ { 	// И благословил их Бог, говоря: плодитесь и размножайтесь, и наполняйте воды в морях, и птицы да размножаются на земле. И так watchPeriod раз.
        Step(a, b)
        a.Show()
 //       time.Sleep(time.Second /4 )
        a, b = b, a // меняем местами текущий и будущий имидж ThePlanet
    }
	goto lbla // гото и вечный цЫкл. просто. брутально. без затей.
}