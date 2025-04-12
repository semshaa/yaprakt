 package spentcalories

import (
    "fmt"
    "strconv"
    "strings"
    "time"
)

const (
    lenStep                          = 0.7
    mInKm                            = 1000
    runningCaloriesMeanSpeedMultiplier = 18.0
    runningCaloriesMeanSpeedShift      = 1.2
    walkingCaloriesWeightMultipler     = 0.035
    walkingSpeedHeightMultiplier       = 0.029
    minInH                          = 60
)

// parseTraining парсит строку данных вида "3456,Ходьба,3h00m"
func parseTraining(data string) (int, string, time.Duration, error) {
    parts := strings.Split(data, ",")
    if len(parts) != 3 {
        return 0, "", 0, fmt.Errorf("invalid training data")
    }

    steps, err := strconv.Atoi(parts[0])
    if err != nil {
        return 0, "", 0, fmt.Errorf("steps parse error: %w", err)
    }

    activity := strings.TrimSpace(parts[1])
    duration, err := time.ParseDuration(parts[2])
    if err != nil {
        return 0, "", 0, fmt.Errorf("duration parse error: %w", err)
    }

    return steps, activity, duration, nil
}

// distance вычисляет дистанцию в км
func distance(steps int) float64 {
    return float64(steps) * lenStep / mInKm
}

// meanSpeed вычисляет среднюю скорость (км/ч)
func meanSpeed(steps int, duration time.Duration) float64 {
    if duration <= 0 {
        return 0
    }
    dist := distance(steps)
    hours := duration.Hours()
    return dist / hours
}

// RunningSpentCalories расчёт калорий для бега
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
    speed := meanSpeed(steps, duration)
    return (runningCaloriesMeanSpeedMultiplier*speed - runningCaloriesMeanSpeedShift) * weight
}

// WalkingSpentCalories расчёт калорий для ходьбы
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
    speed := meanSpeed(steps, duration)
    hours := duration.Hours()
    return (walkingCaloriesWeightMultipler*weight + (speed*speed/height)*walkingSpeedHeightMultiplier) * hours
}

// TrainingInfo формирует отчёт о тренировке
func TrainingInfo(data string, weight, height float64) string {
    steps, activity, duration, err := parseTraining(data)
    if err != nil {
        return fmt.Sprintf("Ошибка: %v", err)
    }

    var calories float64
    dist := distance(steps)
    speed := meanSpeed(steps, duration)

    switch activity {
    case "Бег":
        calories = RunningSpentCalories(steps, weight, duration)
    case "Ходьба":
        calories = WalkingSpentCalories(steps, weight, height, duration)
    default:
        return "неизвестный тип тренировки"
    }

    return fmt.Sprintf(
        "Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
        activity, duration.Hours(), dist, speed, calories,
    )
}

// KalkingSpentCalories - функция для пакета daysteps
func KalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
    speed := meanSpeed(steps, duration)
    return (walkingCaloriesWeightMultipler*weight + (speed*speed/height)*walkingSpeedHeightMultiplier) * duration.Hours()
}
