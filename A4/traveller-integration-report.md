# MEMORANDUM

#### DATE: February 5, 2021
#### TO: Ferdinand Vesely, Project Manager
#### FROM: William Victoria and Luke Novak, Software Developers
#### SUBJECT: Report on Traveller Specification Implementation

---

Overall, we were impressed by the implementation of our `Traveller` specification that we received. Especially considering we were the only group who chose to use Go for our project.

## Specification Implementation Details

The other team that implemented our specification did implement all the data types and functions that we specified. They did make some minor changes in their implementation that they documented clearly in their attached memo. The changes they made involved using pointers in both the data structures and function parameters that allowed for mutation without changing the function signatures. These changes improve the quality of the code and make the implementation of the specified functionality more simple. We believe they were well justified.

## Ease of Integration

To integrate the supplied server code with our client code would have required some minor modifications. The most impactful in terms of work required would be changing the data types and function parameters to use pointers to correspond to the changes the other group made. However, this would still be relatively simple to change in our client code.

It would have been useful if the supplied server code was organized into a Go module other than `main`. This would allow it to be imported nicely from our client code. This could be done by simply adding a `go.mod` file to the folder that specified the package name.

## Improvements to Specification

Our specification could be improved by making the changes that the other team specified in their memo. Namely, using pointers in several places to allow for elegant mutation. There were some redundant functions that were not required to solve to task. With a better understanding of the problem statement we could have eliminated these from the specification and saved the other team work. We also noticed a type-o in one of our purpose statements that obfuscated the meaning of one of our functions. Fortunately, the other team understood the meaning behind our purpose statement and implemented the function as we intended. However, this could have caused confusion and code that did not function the way we would have expected.
