# goChess

-This is a project I am undertaking in order to learn the Go programming language.

-My current goal is to have a fully functional chess engine, most likely contained in the command line.

-The engine will hopefully be an improvement on my previous chess engines in terms of speed, as I am implementing 
the move generation logic using bitboards instead of multidimensional arrays.

-The code written here is mine, although I am following the general ideas and architecture from Chess Programming 
on YouTube, which was written in C. I hope to expand on their design and add more features, but my concern for now is
getting some experience writing in Go

- 6/11/22
- Found a bug in popBit func which caused the rook attacks to be generated incorrectly.

- 6/17/22
- Now it is generating more rook attacks, but attacking it's own pieces

- 6/18/22
- Fixed rook attack bug by adding a check for ally pieces at the target quare