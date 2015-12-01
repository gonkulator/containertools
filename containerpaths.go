package containertools

import (
  "io/ioutil"
  "strings"
)

const croot = "/proc/1/cgroup" // No idea what we do for Windows.

// Containerpaths returns a pointer to a list of cgroups, one per entry in
//  /proc/1/cgroup.  If their GroupIDs are all "/", you're not running
//  in a container.  If you get an error, you are not running a Linux that
//  supports containerization, or the format of /proc/1/cgroup isn't what
//  we were expecting.  If the file is empty, you will get no error, but
//  a nil pointer.  That is the case, for instance, on Debian Wheezy. 
func Containerpaths() (*[]Cgroup, error) {
  b, err := ioutil.ReadFile(croot)
  if err != nil {
    return nil, err
  }
  inp := strings.Split(string(b),"\n")
  var r []Cgroup
  for _,v := range inp {
    s := strings.TrimSpace(v)
    if s == "" {
      continue
    }
    pcg, err := Parse(s)
    if err != nil {
      return nil, err
    }
    r = append(r,*pcg)
  }
  if len(r) == 0 {
    return nil,nil
  }
  return &r, nil
}
