package containertools

// Containerized attempts to determine whether you're running in a container.
//  It does this by firing off Containerpaths(), and if that comes back with
//  an error, or if it comes back with a nil result, you're not in a container.
//  If it has a non-nil result, you are only not in a container if the GroupID
//  of every entry is "/".
func Containerized() bool {
  m,err := Containerpaths()
  if err != nil || m == nil {
    return false
  }  
  for _, v := range *m {
    if v.GroupID != "/" {
      return true
    }
  }
  return false
}

