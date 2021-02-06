# MEMORANDUM

#### DATE: February 5, 2021
#### TO: Ferdinand Vesely, Project Manager
#### FROM: William Victoria and Luke Novak, Software Developers
#### SUBJECT: Report on Traveller Specification Implementation

---

Overall, we were impressed by the implementation of our `Traveller` specification that we received. Especially considering we were the only group who chose to use Go for our project.

## Specification Implementation Details

The other team that implemented our specification did implement all the data types and functions that we specified. They did make some minor changes in their implementation that they documented clearly in their attached memo. The changes they made involved using pointers in both the data structures and function parameters that allowed for mutation without changing the function signatures.

## Ease of Integration

To integrate the supplied server code with our client code would have required some minor modifications. The most impactful in terms of work required would be changing the function parameters to be pointers to correspond to the changes the other group made. However, this would still be relatively simple to change in our client code.

It would have been useful if the supplied server code was organized into a Go module other than `main`. This would allow it to be imported nicely from our client code. This could be done by simply adding a `go.mod` file to the folder that specified the package name.

## Improvements to Specification

- use pointers as they contractors did
- included some redundant functions and unnecessary functions to complete task
- fix purpuse statement of dunction that changed its functionality
