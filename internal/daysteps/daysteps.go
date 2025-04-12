 package daysteps

import (
    "fmt"
    "strconv"
    "strings"
    "time"
    "github.com/Yandex-Practicum/go1fl-4-sprint-final/spentcalories"
)

const StepLength = 0.7 // Константа длины шага в метрах

// parsePackage парсит строку данных вида "678,0h50m"
func parsePackage(data string) (int, time.Duration, error) {
    parts := strings.Split(data, ",")
    if len(parts) != 2 {
        return 0, 0, fmt.Errorf("invalid data format")
    }

    steps, err := strconv.Atoi(parts[0])
    if err != nil {
        return 0, 0, fmt.Errorf("steps parse error: %w", err)
    }

    if steps <= 0 {
        return 0, 0, fmt.Errorf("steps must be positive")
    }

    duration, err := time.ParseDuration(parts[1])
    if err != nil {
        return 0, 0, fmt.Errorf("duration parse error: %w", err)
    }

    return steps, duration, nil
}

// DayActionInfo формирует строку с информацией о дневной активности
func DayActionInfo(data string, weight, height float64) string {
    steps, duration, err := parsePackage(data)
    if err != nil {
        return fmt.Sprintf("Error: %v", err)
    }

    distanceKm := float64(steps) * StepLength / 1000
    calories := spentcalories.KalkingSpentCalories(steps, weight, height, duration)

    return fmt.Sprintf(
        "Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
        steps, distanceKm, calories,
    )
}
