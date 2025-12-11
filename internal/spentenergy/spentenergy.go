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
	if steps <= 0 {
		return 0, errors.New("steps calculation error")
	}
	if weight <= 0 {
		return 0, errors.New("weight calculation error")
	}
	if height <= 0 {
		return 0, errors.New("height calculation error")
	}
	if duration <= 0 {
		return 0, errors.New("duration calculation error")
	}
	speed := MeanSpeed(steps, height, duration)
	walkingCalories := (weight * speed * duration.Minutes()) / minInH
	walkingCaloriesWithCoef := walkingCalories * walkingCaloriesCoefficient
	return walkingCaloriesWithCoef, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps calculation error")
	}
	if weight <= 0 {
		return 0, errors.New("weight calculation error")
	}
	if height <= 0 {
		return 0, errors.New("height calculation error")
	}
	if duration <= 0 {
		return 0, errors.New("duration calculation error")
	}
	speed := MeanSpeed(steps, height, duration)
	runningCalories := (weight * speed * duration.Minutes()) / minInH
	return runningCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	if steps <= 0 {
		return 0
	}
	distanceKilometers := Distance(steps, height)
	return distanceKilometers / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	length := stepLength * float64(steps)
	lengthInKm := length / mInKm
	return lengthInKm
}
