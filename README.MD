Postel’s Law
A good maxim for writing good software is Postel’s Law. It’s often said as

> “Be conservative with what you do, be liberal with you accept”

Translated to Go, the saying is as goes:

> “Accept interfaces, return structs”
  


* cycle import error
package a depends on package b, package b depends on package a