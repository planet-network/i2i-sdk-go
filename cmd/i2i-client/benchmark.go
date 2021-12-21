package main

import (
	"fmt"
	"time"

	"github.com/planet-network/i2i-sdk-go/client"
	"github.com/spf13/cobra"
)

type result struct {
	best       time.Duration
	worst      time.Duration
	average    time.Duration
	total      time.Duration
	name       string
	iterations int
	errorCount int
}

func run(f func() error, name string) {
	result := &result{
		best:       time.Hour * 24,
		iterations: 1000,
	}

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Millisecond * 1)
		now := time.Now()
		err := f()
		dur := time.Since(now)
		if err != nil {
			result.errorCount++
		}
		result.total = result.total + dur

		if dur < result.best {
			result.best = dur
		}

		if dur > result.worst {
			result.worst = dur
		}
	}

	result.average = result.total / time.Duration(result.iterations)
	fmt.Println("===========================")
	fmt.Println("Name         :", name)
	fmt.Println("Iterations   :", result.iterations)
	fmt.Println("Total time   :", result.total)
	fmt.Println("Best time    :", result.best)
	fmt.Println("Worst time   :", result.worst)
	fmt.Println("Average time :", result.average)
	fmt.Println("Error count  :", result.errorCount)
	fmt.Println("===========================")

	fmt.Println()

}

func benchmark(cmd *cobra.Command, args []string) {
	node, err := activeNode()
	if err != nil {
		fail(err)
	}

	i2iClient := client.New(client.Opt{
		Token:    node.Meta.Hosting.UnlockToken,
		Address:  node.Meta.NodeAddress,
		Acl:      node.Meta.APIToken,
		Keychain: node.Keychain,
	})

	fConversations := func() error {
		_, err := i2iClient.Conversations("")
		return err
	}

	run(fConversations, "conversations")

	fPing := func() error {
		return i2iClient.Ping()
	}

	run(fPing, "ping")
}
