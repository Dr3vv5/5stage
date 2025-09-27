package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// проверка вводных данных
	if steps <= 0 || height <= 0 || duration <= 0 || weight <= 0 {
		return 0, errors.New("incorrect	data")
	}
	durationInMinutes := duration.Minutes()
	return ((weight * MeanSpeed(steps, height, duration) * durationInMinutes) / minInH) * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// проверка вводных данных
	if steps <= 0 || height <= 0 || duration <= 0 || weight <= 0 {
		return 0, errors.New("incorrect	data")
	}
	durationInMinutes := duration.Minutes()
	return (weight * MeanSpeed(steps, height, duration) * durationInMinutes) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// проверка вводных данных
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}
	durationInHours := duration.Hours()
	return Distance(steps, height) / durationInHours
}

func Distance(steps int, height float64) float64 {
	// проверка вводных данных
	if steps <= 0 || height <= 0 {
		return 0
	}
	return ((height * stepLengthCoefficient) * float64(steps) / mInKm)
}
