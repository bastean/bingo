package internet

import (
	"fmt"

	"github.com/bastean/bingo/config"
	"github.com/bastean/bingo/service"
	"github.com/spf13/cobra"
)

var InternetFlags = []*config.Flag{
	{Name: "url", Run: func() {
		fmt.Printf("URL: %s\n", service.Create.URL())
	}},
	{Name: "domainName", Run: func() {
		fmt.Printf("Domain Name: %s\n", service.Create.DomainName())
	}},
	{Name: "domainSuffix", Run: func() {
		fmt.Printf("Domain Suffix: %s\n", service.Create.DomainSuffix())
	}},
	{Name: "macAddress", Run: func() {
		fmt.Printf("Mac Address: %s\n", service.Create.MacAddress())
	}},
	{Name: "userAgent", Run: func() {
		fmt.Printf("User Agent: %s\n", service.Create.UserAgent())
	}},
}

var InternetCmd = &cobra.Command{
	Use: "internet",
	Run: func(cmd *cobra.Command, args []string) {
		service.RunCommandFlags(cmd, InternetFlags)
	},
}
