package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// Вычисление калорий при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 { // проверка входных данных на корректность
		return 0, fmt.Errorf("input parameters can't be less or equal zero")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return ((weight * meanSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient, nil // расчет и возврат количества калорий
}

// Вычисление калорий при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 { // проверка входных данных на корректность
		return 0, fmt.Errorf("input parameters can't be less or equal zero")
	}
	meanSpeed := MeanSpeed(steps, height, duration) // расчет средней скорости
	durationInMinutes := duration.Minutes()
	return (weight * meanSpeed * durationInMinutes) / minInH, nil // расчет и возврат количества калорий
}

// Вычисление средней скорости
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 { // проверка продолжительности
		return 0
	}
	if steps <= 0 { // проверка количества шагов
		return 0
	}
	return Distance(steps, height) / duration.Hours() // вычисление и возврат средней скорости
}

// Вычисление дистанции
func Distance(steps int, height float64) float64 {
	leghtStep := height * stepLengthCoefficient
	totalDistanceKm := (leghtStep * float64(steps)) / mInKm
	return totalDistanceKm
}
