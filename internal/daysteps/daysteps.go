package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {

	activity := strings.Split(datastring, ",")

	if len(activity) != 2 { //проверка на правильность преобразования
		return errors.New("error spliting data")
	}

	steps, err := strconv.Atoi(activity[0]) //преобзование шагов в int, и проверка на ошибки
	if err != nil {
		return err
	}

	if steps <= 0 { //проверка, что шагов - положительное число
		return errors.New("steps cannot be zero or less")
	}

	//преобразование времени
	preTime := activity[1]
	duration, err := time.ParseDuration(preTime)
	// проверка времени
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("duration can't be zero or less")
	}
	//сохранение значений
	ds.Steps = steps
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// нахождение переменных
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	ccal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	//формирование вывода
	result := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		ds.Steps, distance, ccal)
	return result, nil
}
