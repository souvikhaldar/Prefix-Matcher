# Prefix Matcher

## Problem statement
Given a list of string prefixes of variable length, the assignment is to implement a method that takes a string as a parameter
and returns the longest prefix that matches the input string. A prefix matches if the input string starts with that prefix.


## Instructions
1. Install golang.   
2. Run the command `go run main.go -file <path to file containing prefixes>` at the root of the repository.  
3. Pre-processing of the prefix file would be done (reverse indexing each substring of the all prefixes). This step might take some time as it is dependent on size of the file. This step can be optimized further using goroutines that is in the development roadmap.  
4. Now you can provide any string as input and it would return the longest prefix that matched the input string from the list of prefixes that we have in near constant time.
