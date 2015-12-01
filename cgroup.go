package containertools

import (
	"fmt"
	"strconv"
	"strings"
)

// Cgroup represents an entry in (Linux) /proc/<pid>/cgroups
type Cgroup struct {
	HierarchyID int
	Subsystems  string
	GroupID     string
}

// String gives back a representation of a Cgroup in the form in which you
//  would find it in /proc/<pid>/cgroups
func (c Cgroup) String() {
	fmt.Printf("%d:%s:%s", c.HierarchyID, c.Subsystems, c.GroupID)
}

// Parse takes a string (as if it came from /proc/<pid>/cgroups) and
//  returns a pointer to a Cgroup (or nil if it won't parse, but that will
//  also return an error)
func Parse(inp string) (*Cgroup, error) {
	s := strings.TrimSpace(inp)
	fl := strings.Split(s, ":")
	lfl := len(fl)
	if lfl != 3 {
		return nil, fmt.Errorf("String %s not of form ID:Subsystems:GroupID", s)
	}
	hid, err := strconv.Atoi(fl[0])
	if err != nil {
		return nil, err
	}
	r := &Cgroup{
		HierarchyID: hid,
		Subsystems:  fl[1],
		GroupID:     fl[2],
	}
	return r, err
}
