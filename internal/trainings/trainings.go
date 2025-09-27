package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration

	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {

	activity := strings.Split(datastring, ",")

	if len(activity) != 3 { //проверка на правильность преобразования
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
	preTime := activity[2]
	duration, err := time.ParseDuration(preTime)
	// проверка времени
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("duration can't be zero or less")
	}
	//сохранение значений
	t.Steps = steps
	t.TrainingType = activity[1]
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	var err error
	var ccal float64
	// нахождение переменных
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	//парсинг по типу тренеровки
	switch t.TrainingType {
	case "Бег":
		ccal, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

	case "Ходьба":
		ccal, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

	default:
		return "", errors.New("undefined train type")
	}
	//проверка на ошибки
	if err != nil {
		return "", err
	}
	//формирование вывода
	result := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, ccal)
	return result, nil
}
