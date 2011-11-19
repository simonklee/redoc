package main

import (
	"fmt"
	"flag"
	"os"
	"strings"
	"sort"
)

type Command struct {
	Name        string
	Summary     string
	Arguments   string
	Group       string
	Since       string
	Description string
}

var (
	// what to display for each command
	description = flag.Bool("d", false, "display long description")
	since       = flag.Bool("s", false, "display since")

	// general
	colors       = flag.Bool("c", true, "use colors")
	listCommands = flag.Bool("lc", false, "list available Redis commands")
	listGroups   = flag.Bool("lg", false, "list available Redis groups")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: redoc [flags] [command|@group]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func (c *Command) printCommand() {
	var formats map[string]string

	if *colors {
		formats = map[string]string{
			"name":        "  \x1b[1m%s\x1b[0m \x1b[90m%s\x1b[0m\n",
			"summary":     "  \x1b[38;5;31msummary:\x1b[0m %s\n",
			"since":       "  \x1b[38;5;31msince:\x1b[0m %s\n",
			"group":       "  \x1b[38;5;31mgroup:\x1b[0m %s\n",
			"description": "  \x1b[38;5;31mdescription:\x1b[0m\n\n\x1b[90m%s\x1b[0m\n",
		}
	} else {
		formats = map[string]string{
			"name":        "  %s %s\n",
			"summary":     "  summary: %s\n",
			"since":       "  since: %s\n",
			"group":       "  group: %s\n",
			"description": "  description:\n\n%s\n",
		}
	}

	fmt.Fprintf(os.Stdout, formats["name"], strings.ToUpper(c.Name), c.Arguments)
	fmt.Fprintf(os.Stdout, formats["summary"], c.Summary)

	if *since {
		fmt.Fprintf(os.Stdout, formats["since"], c.Since)
	}

	fmt.Fprintf(os.Stdout, formats["group"], c.Group)

	if *description {
		fmt.Fprintf(os.Stdout, formats["description"], c.Description)
	}

	fmt.Fprintf(os.Stdout, "\n")
}

func printName(name string) bool {
	if c, ok := Commands[name]; ok {
		c.printCommand()
		return ok
	}
	return false
}

func printCommands() {
	var commands []string

	for k := range Commands {
		commands = append(commands, k)
	}

	sort.Strings(commands)

	var format string

	if *colors {
		format = "%s\x1b[38;5;196m|\x1b[0m"
	} else {
		format = "%s "
	}

	for _, name := range commands {
		fmt.Fprintf(os.Stdout, format, name)
	}

	fmt.Fprintf(os.Stdout, "\n")
}

func printGroup(name string) bool {
	g := orderByGroup()

	for k, v := range g {
		if k == name {
			printGrouped(map[string][]string{k: v})
		}
	}

	return false
}

func printGroups() {
	found := map[string]bool{}
	var groups []string

	for _, v := range Commands {
		if _, ok := found[v.Group]; !ok {
			found[v.Group] = true
			groups = append(groups, v.Group)
		}
	}

	sort.Strings(groups)

	var format string

	if *colors {
		format = "%s\x1b[38;5;196m|\x1b[0m"
	} else {
		format = "%s "
	}

	for _, name := range groups {
		fmt.Fprintf(os.Stdout, format, name)
	}

	fmt.Fprintf(os.Stdout, "\n")
}

func printGrouped(grouped map[string][]string) {
	for _, v := range grouped {
		for _, name := range v {
			c := Commands[name]
			c.printCommand()
		}
	}
}

func orderByGroup() map[string][]string {
	grouped := make(map[string][]string)

	for k, v := range Commands {
		grouped[v.Group] = append(grouped[v.Group], k)
	}

	for _, v := range grouped {
		sort.Strings(v)
	}

	return grouped
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	if *listGroups {
		printGroups()
		os.Exit(0)
	}

	if *listCommands {
		printCommands()
		os.Exit(0)
	}

	if len(args) > 0 {
		typ := strings.ToLower(strings.Join(args, " "))

		if strings.HasPrefix(typ, "@") && printGroup(typ[1:]) {
			os.Exit(0)
		}

		if !printName(typ) {
			if !printGroup(typ) {
				os.Exit(1)
			}
		}
		os.Exit(0)
	}

	g := orderByGroup()
	printGrouped(g)

	os.Exit(0)
}
