package main

import (
	"fmt"
	"strconv"
)

// logs2 = [
//     ["300", "user_1", "resource_3"],
//     ["599", "user_1", "resource_3"],
//     ["900", "user_1", "resource_3"],
//     ["1199", "user_1", "resource_3"],
//     ["1200", "user_1", "resource_3"],
//     ["1201", "user_1", "resource_3"],
//     ["1202", "user_1", "resource_3"]
// ]

// desired output:
// {'user_1': [300, 1202]}

type SortedLogs map[string][]int

func userSessions(logs [][]string) SortedLogs {
	sortedLogs := make(SortedLogs)
	for i := range logs {
		key := logs[i][1]
		err, _ := sortedLogs[key]
		if err != nil {
			val := sortedLogs[key]
			parseInt, _ := strconv.Atoi(logs[i][0])
			if parseInt < val[0] {
				sortedLogs[key][0] = parseInt
			} else {
				sortedLogs[key][1] = parseInt
			}
		} else {
			parseInt, _ := strconv.Atoi(logs[i][0])
			sortedLogs[key] = append(sortedLogs[key], parseInt)
			sortedLogs[key] = append(sortedLogs[key], parseInt)
		}
	}

	return sortedLogs
}

func main() {
	logs := [][]string{
		{"300", "user_1", "resource_3"},
		{"599", "user_1", "resource_3"},
		{"900", "user_1", "resource_3"},
		{"1199", "user_1", "resource_3"},
		{"1200", "user_1", "resource_3"},
		{"1200", "user_1", "resource_3"},
		{"1201", "user_1", "resource_3"},
		{"1202", "user_1", "resource_3"},
	}

	// logs = [][]string{
	// 	{"58523", "user_1", "resource_1"},
	// 	{"62314", "user_2", "resource_2"},
	// 	{"54001", "user_1", "resource_3"},
	// 	{"200", "user_6", "resource_5"},
	// 	{"215", "user_6", "resource_4"},
	// 	{"54060", "user_2", "resource_3"},
	// 	{"53760", "user_3", "resource_3"},
	// 	{"58522", "user_22", "resource_1"},
	// 	{"53651", "user_5", "resource_3"},
	// 	{"2", "user_6", "resource_1"},
	// 	{"100", "user_6", "resource_6"},
	// 	{"400", "user_7", "resource_2"},
	// 	{"100", "user_8", "resource_6"},
	// 	{"54359", "user_1", "resource_3"},
	// }

	sortedLogs := userSessions(logs)
	fmt.Printf("Sorted Logs: %v", sortedLogs)
}
