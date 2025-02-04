package models
import (
    "errors"
)
// 
type AttackSimulationOperationType int

const (
    CREATESIMUALATION_ATTACKSIMULATIONOPERATIONTYPE AttackSimulationOperationType = iota
    UPDATESIMULATION_ATTACKSIMULATIONOPERATIONTYPE
    UNKNOWNFUTUREVALUE_ATTACKSIMULATIONOPERATIONTYPE
)

func (i AttackSimulationOperationType) String() string {
    return []string{"createSimualation", "updateSimulation", "unknownFutureValue"}[i]
}
func ParseAttackSimulationOperationType(v string) (any, error) {
    result := CREATESIMUALATION_ATTACKSIMULATIONOPERATIONTYPE
    switch v {
        case "createSimualation":
            result = CREATESIMUALATION_ATTACKSIMULATIONOPERATIONTYPE
        case "updateSimulation":
            result = UPDATESIMULATION_ATTACKSIMULATIONOPERATIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ATTACKSIMULATIONOPERATIONTYPE
        default:
            return 0, errors.New("Unknown AttackSimulationOperationType value: " + v)
    }
    return &result, nil
}
func SerializeAttackSimulationOperationType(values []AttackSimulationOperationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
