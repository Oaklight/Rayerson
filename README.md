# Rayerson
A Golang implementation of ray tracing rendering engine

## How to run the code

```
# compile render
go build render.go

# run with a scene file and designate the output image path
# render [-p[=nThread]] <path to csv file> <output path>
render -p test/sceneSimple.csv outSimple.png
```

## Dataset and result

All three datasets are in `./test/` folder, corresponding results are in the same folder.

You may modify the code in function `RandomScene` in `./render/scene.go` to generate randomly placed sphere with varying materials.

## Dependencies

All dependencies are standard libraries:

- math
- math/rand
- fmt
- image
- image/color
- image/png
- sync
- os
- time
- bufio
- strconv
- runtime
- strings