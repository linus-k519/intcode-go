

# Usage

```bash
intcode run program.ic
```



# Intcode

## Opcodes

| Opcode | Params | Description                                                  |
| ------ | ------ | ------------------------------------------------------------ |
| 01     | 3      | Adds the first two arguments and stores the result in the third argument. |
| 02     | 3      | Multiplies the first two arguments and stores the result in the third argument. |
| 03     | 1      | Inputs an integer and stores it in the first argument.       |
| 04     | 1      | Outputs the first argument.                                  |
| 05     | 2      | If the first argument is non-zero, sets the instruction pointer to the second argument. |
| 06     | 2      | If the first argument is zero, sets the instruction pointer to the second argument. |
| 07     | 3      | If the first argument is less than the second argument, sets the third argument to 1. If not less, sets it to 0. |
| 08     | 3      | If the first argument is equal to the second argument, sets the third argument to 1. If not equal, sets it to 0. |
| 09     | 1      | Adds the first argument to the relative base register.       |
| 99     | 0      | Ends the program                                             |
> From https://esolangs.org/wiki/Intcode

## Parameter Modes

| Mode | Name           | Description                                                  |
| ---- | -------------- | ------------------------------------------------------------ |
| 0    | Position Mode  | The parameter is the address of the value.                   |
| 1    | Immediate Mode | The parameter is the value itself (Not used for writing).    |
| 2    | Relative Mode  | The parameter is added to the relative base register, which results in the memory address of the value. |

> From https://esolangs.org/wiki/Intcode
>

```bash
ABCDE
 1002

DE - two-digit opcode,      02 == opcode 2
 C - mode of 1st parameter,  0 == position mode
 B - mode of 2nd parameter,  1 == immediate mode
 A - mode of 3rd parameter,  0 == position mode,
                                  omitted due to being a leading zero
```

> From https://adventofcode.com/2019/day/5
>

## Examples

```bash
1, 3, 4, 2, 99
```

