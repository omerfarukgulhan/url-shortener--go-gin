package id

import "fmt"

func GetUniqueId() (int64, error) {
	snowflake, err := NewSnowflake(1, 1)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}

	return snowflake.GenerateID(), nil
}
