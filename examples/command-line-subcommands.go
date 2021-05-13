package examples

import (
	"flag"
	"fmt"
	"os"
)

func CommandLineSubcommands() {
	fooSubCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooCmd1 := fooSubCmd.String("foo1", "foo1", "Foo value")

	barSubCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barCmd1 := barSubCmd.Int("bar1", 1, "Bar value")

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Expected foo or bar subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fmt.Println("foo subcommand")
		fmt.Println("foo1", *fooCmd1)
		fmt.Println("tails", fooSubCmd.Args())
	case "bar":
		fmt.Println("bar subcommand")
		fmt.Println("bar1", *barCmd1)
		fmt.Println("tails", barSubCmd.Args())
	default:
		fmt.Fprintln(os.Stderr, "Expected foo or bar subcommands")
		os.Exit(1)
	}
}
