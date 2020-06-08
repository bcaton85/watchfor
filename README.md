# watchfor

### Description

Cli tool for running arbitrary bash commands when notified of a change in the file system. 

### Usage

`watchfor DIRECTORY -- COMMAND COMMAND_ARG_1 COMMAND_ARG_2 COMMAND_ARG_3...`

`watchfor ProjectDir -- docker build . -t myimage`

`watchfor ProjectDir -- dotnet run`

