In Go you define classes or objects as structures. These are like C structures and can only contain data. For the members that access those data types, you write functions passing the object type as parameter.
Let's say I'm defining a CPU object. Its architectures can be X86 32-bit, ARM 64-bit etc. You define it as follows -

```golang
type CPU struct {
    architecture string
}

func (cpu *CPU) set_architecture(arch string) {
    cpu.architecture = arch
}

func (cpu *CPU) get_architecture() string {
    return cpu.architecture
}
```

One quirky thing to note about Go. With a CPU object declared like this, you'd think that to instantiate an object of type CPU, you'd write -
```golang
my_cpu := CPU()
```
Based on your C++ or Java experience.

Actually in Go, that line would be
```golang
my_cpu := CPU{}
```
Note the curly braces in place of parenthesis.

