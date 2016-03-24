#doitlive shell: /bin/bash
#doitlive prompt: minimal_color

### setup

. .envrc
cd src

### show the src layout

la
# show the main in vim
go run main.go 1 2 3
cd demo
la
# show demo/demo_test in vim

### attaching

# attach to main
cd ..
dlv debug # without optimization
exit
dlv trace SumStrings

# with optimizations
go build -o main
dlv exec ./main
exit
dlv attach

# attach to test
cd demo
dlv test

### debug tests

help # at any time you can type help to get a list of commands
c
# show runtime breakpoint
p sum
r
b start demo.go:5
bp
c

args
locals

p nums
s
n
n
p number
n
p number
b return 15
bp
c
p total
n
n
p total
set total = 6
c
exit

### debug main

cd ..
dlv debug
b main.go:12
c

stack
threads
goroutines

ls 25

b 25
b 21
c
goroutines

regs
disassemble

sources demo
vars demo
funcs demo
