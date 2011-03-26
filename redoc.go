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
    group       = flag.Bool("g", true, "display group")
    since       = flag.Bool("si", false, "display since")
    summary     = flag.Bool("s", true, "display summary")

    // general
    colors  = flag.Bool("c", true, "use colors")
    listCommands  = flag.Bool("lc", false, "list all available commands")
    listGroups  = flag.Bool("lg", false, "list all available groups")
)

func usage() {
    fmt.Fprintf(os.Stderr, "usage: redoc [flags] [command|group]\n")
    flag.PrintDefaults()
    os.Exit(2)
}

func (c *Command) printCommand() {
    var formats map[string]string

    if *colors {
        formats = map[string]string {
            "name": "  \x1b[1m%s\x1b[0m \x1b[90m%s\x1b[0m\n",
            "summary": "  \x1b[38;5;31msummary:\x1b[0m %s\n",
            "since": "  \x1b[38;5;31msince:\x1b[0m %s\n",
            "group": "  \x1b[38;5;31mgroup:\x1b[0m %s\n",
            "description": "  \x1b[38;5;31mdescription:\x1b[0m\n\n\x1b[90m%s\x1b[0m\n",
        }
    } else {
        formats = map[string]string {
            "name": "  %s %s\n",
            "summary": "  summary: %s\n",
            "since": "  since: %s\n",
            "group": "  group: %s\n",
            "description": "  description:\n\n%s\n",
        }
    }

    fmt.Fprintf(os.Stdout, formats["name"],strings.ToUpper(c.Name), c.Arguments)

    if *summary {
        fmt.Fprintf(os.Stdout, formats["summary"], c.Summary)
    }

    if *since {
        fmt.Fprintf(os.Stdout, formats["since"], c.Since)
    }

    if *group {
        fmt.Fprintf(os.Stdout, formats["group"], c.Group)
    }

    if *description {
        fmt.Fprintf(os.Stdout, formats["description"], c.Description)
    }

    fmt.Fprintf(os.Stdout, "\n")
}

func printAll() {
    for _, c := range Commands {
        c.printCommand()
    }
}

func printSingle(name string) bool {
    if c, ok := Commands[name]; ok {
        c.printCommand()
        return ok
    } 
    return false
}

func printGroup(name string) bool {
    g := orderByGroup()

    for k, v := range g {
        if k == name {
            printGrouped(map[string][]string{k: v,})
        }
    }

    return false
}

func printCommands() {
    var commands []string

    for k, _ := range Commands {
        commands = append(commands, k)
    }

    sort.SortStrings(commands)

    for _, name := range commands {
        fmt.Fprintf(os.Stdout, "%s ", name)
    }

    fmt.Fprintf(os.Stdout, "\n")
}

func printGroups() {
    found := map[string] bool{}
    var groups []string

    for _, v := range Commands {
        if _, ok := found[v.Group]; !ok {
            found[v.Group] = true
            groups = append(groups, v.Group)
        }
    }

    sort.SortStrings(groups)

    for _, name := range groups {
        fmt.Fprintf(os.Stdout, "%s ", name)
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
        sort.SortStrings(v)
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

    if len(args) == 1 {
        typ := strings.ToLower(args[0])
        if !printSingle(typ) {
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
