package main

import (
	"fmt"
	"os"

	ruleEngine "github.com/rgehrsitz/go-rules-engine/rule_engine"
)

func main() {

	// jsonRule1, err := os.ReadFile("rule1.json")
	// if err != nil {
	// 	panic(err)
	// }
	jsonRule2, err := os.ReadFile("rule2.json")
	if err != nil {
		panic(err)
	}

	// Define the data to be used in the conditions
	data1 := map[string]interface{}{
		"Sensor1Temp": 9,
		"Sensor2Temp": 999,
		"Sensor3Temp": 100,
	}

	data2 := map[string]interface{}{
		"Sensor2Temp":     177,
		"Sensor2Pressure": 60,
	}

	engine := ruleEngine.New(&ruleEngine.EvaluatorOptions{
		AllowUndefinedVars: true,
	})

	// engine.AddRule(string(jsonRule1))
	engine.AddRule(string(jsonRule2))
	// engine.AddRules(string(jsonRule1), string(jsonRule2))

	// Evaluate the conditions using the EvaluateCondition function
	result1 := engine.EvaluateRules(data1)
	result2 := engine.EvaluateRules(data2)

	// Print the result
	fmt.Println(result1)
	fmt.Println(result2)

}
