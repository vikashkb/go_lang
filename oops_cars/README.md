In Go you define classes or objects as structures. These are like C structures and can only contain data. For the members that access those data types, you write functions passing the object type as parameter.
Let's say I'm defining a CPU object. Its architectures can be X86 32-bit, ARM 64-bit etc. You define it as follows -

```
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
```
my_cpu := CPU()
```
Based on your C++ or Java experience.

Actually in Go, that line would be
```
my_cpu := CPU{}
```
Note the curly braces in place of parenthesis.

Let's start with a simple exercise. We will define 4 objects in Go, a car object,an engine object, a maker object, a car_type object and a seat object. Engine objects are like v4/v6/turbo_v4 etc. Maker objects are like Mercedes, BMW, Lexus etc. Car_type objects are like sedan, minivan, SUV etc. Seat objects are like leather, faux leather, cloth etc. And a car object has all of these underneath - maker, engine, car_type, and seat.
