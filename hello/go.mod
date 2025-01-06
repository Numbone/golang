module example.com/hello

go 1.21.1

replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000

require golang.org/x/example/hello v0.0.0-20241216154601-40afcb705d05 // indirect
