"Polymorphism is the ability to write code that can take on different behavior through the
 implementation of types. Once a type implements an interface, an entire world of
 functionality can be opened up to values of that type."
 - Bill Kennedy

"Interfaces are types that just declare behavior. This behavior is never implemented by the
 interface type directly, but instead by user-defined types via methods. When a
 user-defined type implements the set of methods declared by an interface type, values of
 the user-defined type can be assigned to values of the interface type. This assignment
 stores the value of the user-defined type into the interface value.

 If a method call is made against an interface value, the equivalent method for the
 stored user-defined value is executed. Since any user-defined type can implement any
 interface, method calls against an interface value are polymorphic in nature. The
 user-defined type in this relationship is often called a concrete type, since interface values
 have no concrete behavior without the implementation of the stored user-defined value."
  - Bill Kennedy

Receivers       Values
-----------------------------------------------
(t T)           T and *T
(t *T)          *T

Values          Receivers
-----------------------------------------------
T               (t T)
*T              (t T) and (t *T)


SOURCE:
Go In Action
William Kennedy
/////////////////////////////////////////////////////////////////////////

Interface types express generalizations or abstractions about the behaviors of other types.
By generalizing, interfaces let us write functions that are more flexible and adaptable
because they are not tied to the details of one particular implementation.

Many object-oriented lagnuages have some notion of interfaces, but what makes Go's interfaces
so distinctive is that they are SATISIFIED IMPLICITLY. In other words, there's no need to declare
all the interfaces that a given CONCRETE TYPE satisifies; simply possessing the necessary methods
is enough. This design lets you create new interfaces that are satisifed by existing CONCRETE TYPES
without changing the existing types, which is particularly useful for types defined in packages that
you don't control.

All the types we've looked at so far have been CONCRETE TYPES. A CONCRETE TYPE specifies the exact
representation of its values and exposes the intrinsic operations of that representation, such as
arithmetic for numbers, or indexing, append, and range for slices. A CONCRETE TYPE may also provide
additional behaviors through its methods. When you have a value of a CONCRETE TYPE, you know exactly
what is IS and what you can DO with it.

There is another kind of type in Go called an INTERFACE TYPE. An interface is an ABSTRACT TYPE. It doesn't
expose the representation or internal structure of its values, or the set of basic operations they support;
it reveals only some of their methods. When you have a value of an interface type, you know nothing about
what it IS; you know only what it can DO, or more precisely, what BEHAVIORS ARE PROVIDED BY ITS METHODS.

-------------------

type ReadWriter interface {
    Reader
    Writer
}

This is called EMBEDDING an interface.


-------------------

A type SATISFIES an interface if it possesses all the methods the interface requires.

-------------------

Conceptually, a value of an interface type, or INTERFACE VALUE, has two components,
    a CONCRETE TYPE and a
    VALUE OF THAT TYPE.
These are called the interface's
    DYNAMIC TYPE and
    DYNAMIC VALUE.

For a statically typed language like Go, types are a compile-time concept, so a type is not a value.
In our conceptual model, a set of values called TYPE DESCRIPTORS provide information about each type,
such as its name and methods. In an interface value, the type component is represented by the appropriate
type descriptor.


var w io.Writer
w = os.Stdout
w = new(bytes.Buffer)
w = nil


var w io.Writer
w
type: nil
value: nil

w = os.Stdout
w
type: *os.File
value: the address where a value of type os.File is stored

w = new(bytes.Buffer)
w
type: *bytes.Buffer
value: the address where a value of type bytes.Buffer is stored

w = nil
w
type: nil
value: nil

-------------------
The Go Programming Language
Donovan and Kernighan

Caplitalization and identation mine.