# Battleship

The battle ship is a game played by two people. And the goal is to 
sink all the enemy ships by guessing the position of them. 

## Input

Here is the input is in the format 

```

M   // Grid Size 
S   // total ships 
1p1,1p2 ...  // P1 ship positions 
2p1,2p2 ...  // ship positions  
T   // Total missiles
1m1,1m2 ...  // P1 missiles attack position 
2m2,2m2 ...  // P2 missiles attack position 

```

Eg. 


```

5
5
1:1,2:0,2:3,3:4,4:3
0:1,2:3,3:0,3:4,4:1
5
0,1:4,3:2,3:3,1:4,1
0,1:0,0:1,2:2,3:4,3

```

## Output 

Output will be 

```

Player1
O O _ _ _
_ B O _ _
B _ _ X _
_ _ _ _ B
_ _ _ X _

Player2
_ X _ _ _
_ _ _ _ _
_ _ _ X _
B O _ _ B
_ X _ O _

P1:2
P2:3
Player 2 wins

```

## How to run 

The program can take input from standard input as well as from a file 

```

go run main.go -input input.txt -ouput output.txt 

```

If no `input` or `output` flags are provided then program will take input 
from `stdin` and output to `stdout`.

 


