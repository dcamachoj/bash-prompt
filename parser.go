package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"strconv"
	"strings"
)

// Parser type
type Parser struct {
	data []string
	defs map[string]int
}

// NewParser constructor
func NewParser() (*Parser, error) {
	var err error
	var data map[string]int
	var list []string

	var p = &Parser{
		defs: map[string]int{},
	}

	data, err = p.loadMap("defs")
	if err != nil {
		return nil, err
	}
	for n, v := range data {
		p.defs[n] = v
	}

	data, err = p.loadMap("col16")
	if err != nil {
		return nil, err
	}
	for n, v := range data {
		p.defs[n] = v
	}

	list, err = p.loadLines("col256")
	if err != nil {
		return nil, err
	}
	for k, n := range list {
		p.defs[n] = k
	}

	return p, nil
}

func (p *Parser) loadLines(name string) ([]string, error) {
	var err error
	var data []byte
	var filename = fmt.Sprintf("assets/%s.txt", name)

	data, err = Asset(filename)
	if err != nil {
		return nil, err
	}

	var lines = strings.Split(string(data), "\n")
	var result []string
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		result = append(result, line)
	}
	return result, nil
}

func (p *Parser) loadMap(name string) (map[string]int, error) {
	var lines, err = p.loadLines(name)
	var value int
	var parts []string
	var res = map[string]int{}
	for k, line := range lines {
		parts = strings.Split(line, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("Definition %s (%d) %s has wrong format", path.Base(name), k, line)
		}
		value, err = strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return nil, fmt.Errorf("Definition %s (%d) %s has wrong format. %s", path.Base(name), k, line, err)
		}
		res[parts[0]] = value
	}

	return res, nil
}

// ParseFile filename
func (p *Parser) ParseFile(filename string) error {
	var err error
	var data []byte
	var cmd, val string
	p.data = nil
	data, err = ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var lines = strings.Split(string(data), "\n")
	var parts []string
	for k, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts = strings.SplitN(line, ":", 2)
		cmd = strings.ToLower(parts[0])
		val = ""
		if len(parts) > 1 {
			val = parts[1]
		}
		switch cmd {
		case "nop":
			//NOP
		case "txt":
			err = p.parseTxt(parts[1])
		case "frm":
			err = p.parseFrm(parts[1])
		case "brd":
			err = p.parseTxt("")
			switch val {
			case "end":
			default:
				err = p.parseTxt(" ")
			}
		default:
			err = fmt.Errorf("Unsupported command at line %d: %s", k+1, parts[0])
		}
		if err != nil {
			return fmt.Errorf("Error at line %d. %s", k+1, err)
		}
	}
	return nil
}

// String implementation
func (p *Parser) String() string {
	return strings.Join(p.data, "")
}

func (p *Parser) isInt(value string) bool {
	var _, err = strconv.Atoi(value)
	return err == nil
}

func (p *Parser) parseTxt(value string) error {
	p.data = append(p.data, value)
	return nil
}
func (p *Parser) parseFrm(value string) error {
	var err error
	var cmds []string
	var values []int

	values, err = p.parseCmd(value)
	if err != nil {
		return err
	}

	if len(values) > 0 {
		cmds = make([]string, len(values))
		for k, v := range values {
			cmds[k] = strconv.Itoa(v)
		}
		p.data = append(p.data, fmt.Sprintf("\\[\\e[%sm\\]", strings.Join(cmds, ";")))
	}
	return nil
}
func (p *Parser) parseCmd(cmd string) ([]int, error) {
	var err error
	var cmds []string
	var values, subValues []int

	cmds = strings.Split(strings.TrimSpace(cmd), ",")
	for k, cmd := range cmds {
		subValues, err = p.parseSubCmd(cmd)
		if err != nil {
			return nil, fmt.Errorf("Error at sub-command %d. %s", k+1, err)
		}
		values = append(values, subValues...)
	}
	return values, nil
}
func (p *Parser) parseSubCmd(cmd string) ([]int, error) {
	var err error
	var parts []string
	var name, str string
	var ok bool
	var val int

	parts = strings.SplitN(cmd, ":", 2)
	switch len(parts) {
	case 0:
		return nil, nil
	case 1:
		str = strings.TrimSpace(parts[0])
		val, ok = p.defs[str]
		if !ok {
			val, err = strconv.Atoi(str)
			if err != nil {
				return nil, fmt.Errorf("Wrong Format %s", cmd)
			}
		}
		return []int{val}, nil
	case 2:
		name = strings.ToLower(strings.TrimSpace(parts[0]))
		str = strings.TrimSpace(parts[1])
		val, ok = p.defs[str]
		if !ok {
			val, err = strconv.Atoi(str)
			if err != nil {
				return nil, fmt.Errorf("Wrong Format for '%s'. %s not found", cmd, str)
			}
		}
		switch name {
		case "set":
			return formatSet(val), nil
		case "rst":
			return formatRst(val), nil
		case "fg":
			return formatFg(val), nil
		case "bg":
			return formatBg(val), nil
		case "fg256":
			return formatFg(val), nil
		case "bg256":
			return formatBg(val), nil
		default:
			return nil, fmt.Errorf("Wrong format for subcommand %s. %s not supported", cmd, name)
		}
	}
	return nil, nil
}

// Names getter
func (p *Parser) Names() []string {
	var names = make([]string, len(p.defs))
	var k = 0
	for n := range p.defs {
		names[k] = n
		k++
	}
	return names
}
