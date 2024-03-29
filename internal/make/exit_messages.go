package make

var makeExitMessages = map[int]string{
	0:  "Normal exit with no errors.",
	1:  "General purpose error if no other explicit error is known.",
	2:  "There was an error in the makefile.",
	3:  "A shell line had a non-zero status.",
	4:  "Make ran out of memory.",
	5:  "The program specified on the shell line was not executable.",
	6:  "The shell line was longer than the command processor allowed.",
	7:  "The program specified on the shell line could not be found.",
	8:  "There was not enough memory to execute the shell line.",
	9:  "The shell line produced a device error.",
	10: "The program specified on the shell line became resident.",
	11: "The shell line produced an unknown error.",
	15: "There was a problem with the memory miser.",
	16: "The user hit CTRL+C or CTRL+BREAK..",
}
