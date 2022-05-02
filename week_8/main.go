package main

import (
	"context"
	"fmt"
	"strconv"
)

func main() {
	ctx := context.Background()
	rTest := NewRedisTest()

	used, usedHuman := rTest.InfoMemory(ctx)
	fmt.Println("before: ", used, usedHuman)
	rTest.InfoKeyspace(ctx)

	batch := 20 * 10000
	rTest.WriteValue(ctx, 20, batch)

	usedAfter, usedHumanAfter := rTest.InfoMemory(ctx)
	usedEnd, _ := strconv.ParseInt(usedAfter, 10, 64)
	usedStart, _ := strconv.ParseInt(used, 10, 64)
	totalByte := usedEnd - usedStart
	fmt.Println("after: ", usedAfter, usedHumanAfter)
	rTest.InfoKeyspace(ctx)
	fmt.Printf("total: %d total human: %0.2fM avg: %0.2f", totalByte, float64(totalByte)/1024/1024, float64(totalByte)/float64(batch))
}
