# watchfor

### Description

Cli tool for running arbitrary bash commands when notified of a change in the file system. 

### Usage

`watchfor DIRECTORY 'COMMAND'`

`watchfor ProjectDir 'docker build . -t myimage'`

`watchfor ProjectDir 'terraform plan'`

`watchfor ProjectDir 'dotnet run'`

