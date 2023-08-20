# Oh My Gotch

This is a gotch "hello world" to verify `gotch` is installed correctly and as expected.

To test, do the following:

```bash
git clone https://github.com/sugarme/oh-my-gotch

cd oh-my-gotch

go mod tidy

go run main.go

```

Output should be the following if CUDA is installed and libtorch for CUDA installed correctly.

```bash
TENSOR INFO:
        Shape:          [3 4 5]
        DType:          Float
        Device:         {CUDA 0}
        Defined:        true
```

