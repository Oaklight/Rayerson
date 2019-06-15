# Rayerson: a CPU-based ray tracing engine
Peng Ding [dingpeng@uchicago.edu]

![img](file://C:/Users/dingp/Projects/Rayerson/assets/outComplex.png?lastModify=1560567591)

## Problem Description
When looking at an image, people usually can tell whether it is taken by a camera, or it is synthetic by some program. One of the reasons that the synthetic images look unreal is due to lighting. Lighting is the reason that people can see the world. Everything either performs as a light source, e.g. the bulbs, the sun, or serves as a "mirror", which does not necessarily need to be a perfect reflector. In most cases, light is scattered in random directions by the diffuse object. Some substances can let light pass through them, like glass and water, which are considered to be transparent. Light traveling in these substances will be refracted and the intensity may be attenuated.

During my undergraduate study, I took a class in Computer Graphics. One of the project was to implement a ray tracer, using OpenGL in C++. That project was somehow beyond my capability at that time, so the outcome was just okay rather than good. After two years, I think maybe it is time to tie the loose end, so I decided to work on a ray tracer as the final project.

The project aims to implement a ray tracer from scratch using Golang. All dependencies are just standard libraries. The following features are included:
- sphere is provided as a primitive object
- 3 major materials are provided:
  - Diffuse / Lambertian
  - Metallic
  - Dielectric: e.g. glass
- Tunable Camera model with defocus


## System Implementation

`Rayerson` package contains several supporting libraries, constituting the cornerstones of the rendering task.

### `vector` module

Implements the `Vec3`  class, with basic vector arithmetic provided in two fashions: one is called by class methods, the other is general function call, with the first elements as the pivot elements.

For example, `Add`:

```go
func (v1 *Vec3) Add(vs ...*Vec3) *Vec3 {
	e0, e1, e2 := v1.X, v1.Y, v1.Z
	for _, v := range vs {
		e0 += v.X
		e1 += v.Y
		e2 += v.Z
	}
	return &Vec3{e0, e1, e2}
}
...
func Abs(v1 *Vec3) *Vec3 {
	return v1.Abs()
}
```



### `ray` module

### `primitives` module

### `render` module

### Parallelism
- Data Decomposition
- Granularity
- 

### Advanced Features
- Load Balance
- Channel Patterns
- 

## System Specs

## Experiments and Performance

### Hotspots

### Bottlenecks

### Speedup limitation
