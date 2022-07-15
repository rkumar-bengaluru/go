package main

import (
	"encoding/json"
	"fmt"
)

type Manager struct {
	FullName       string `json:"full_name"`
	Position       string `json:"position"`
	Age            int32  `json:"age"`
	YearsInCompany int32  `json:"years_in_company"`
}

type NManager struct {
	FullName       string `json:"full_name"`
	Position       string `json:"position"`
	Age            int32  `json:"age"`
	YearsInCompany int32  `json:"years_in_company,omitempty"`
}

func EncodeManager(manager *Manager) {
	var nm NManager
	if manager.YearsInCompany == 0 {
		nm = NManager{FullName: manager.FullName, Position: manager.Position, Age: manager.Age}
	} else {
		nm = NManager{
			FullName:       manager.FullName,
			Position:       manager.Position,
			Age:            manager.Age,
			YearsInCompany: manager.YearsInCompany,
		}
	}
	b, _ := json.Marshal(nm)
	fmt.Println(string(b))
}

func main() {
	m := Manager{FullName: "Jack", Position: "CEO", Age: 43, YearsInCompany: 23}
	EncodeManager(&m)
}
