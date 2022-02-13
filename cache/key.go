package cache

import (
	"fmt"
	"strconv"
)

// @Author KHighness
// @Update 2022-02-13

const (
	RankTaskKey = "rank:task"
)

// key: Task点击数
func TaskViewKey(id uint) string {
	return fmt.Sprintf("view:task:%s", strconv.Itoa(int(id)))
}

