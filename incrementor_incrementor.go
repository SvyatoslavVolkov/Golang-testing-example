package incrementor

import "math"

type incrementor struct {
	number       int32
	maximumValue int32
}

// Структура предоставляет интерфейс счётчика с ограничением максимального
// значения
type Incrementor struct {
	inc incrementorInt
}

type incrementorInt interface {
	getNumber() int32
	incrementNumber()
	setMaximumValue(maximumValue int32)
}

// Создаёт новый экземпляр структуры
func NewIncrementor() *Incrementor {
	i := &incrementor{number: 0, maximumValue: math.MaxInt32}
	inc := &Incrementor{inc: i}
	return inc
}

// Возвращает текущее число
func (i *Incrementor) GetNumber() int32 {
	return i.inc.getNumber()
}

// Увеличивает текущее число на один. После каждого вызова этого метода
// GetNumber() будет возвращать число на один больше.
func (i *Incrementor) IncrementNumber() {
	i.inc.incrementNumber()
}

// Устанавливает максимальное значение текущего числа. Когда при вызове
// IncrementNumber() текущее число достигает этого значения, оно обнуляется.
func (i *Incrementor) SetMaximumValue(maximumValue int32) {
	i.inc.setMaximumValue(maximumValue)
}

func (i *incrementor) getNumber() int32 {
	return i.number
}

func (i *incrementor) incrementNumber() {
	if i.number < i.maximumValue {
		i.number++
	} else {
		i.number = 0
	}
}

func (i *incrementor) setMaximumValue(maximumValue int32) {
	if maximumValue >= 0 {
		i.maximumValue = maximumValue
		if maximumValue < i.number {
			i.number = 0
		}
	}
}