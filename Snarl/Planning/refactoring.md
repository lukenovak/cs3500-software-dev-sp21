# Milestone 6 - Refactoring Report

**Team members:** William "Rigby" Victoria, Luke "Mordecai" Novak

**Github team/repo:** https://github.ccs.neu.edu/CS4500-S21/Ormegland


## Plan

Prior to this milestone, we already identified a number of github issues outstanding with our codebase, and we
will be tackling those. Our first priority will be lining up our level representation with the JSON input and
output that represents a level. This input and output is currently represented in row/column form, but our
internal representation works in the opposite way (cartesian coordinates). Beyond that, we would like to introduce
more local functions to cut down on duplicate code and decrease the overall size to our codebase.


## Changes

We accomplished our main refactoring goal, the conversion of levels to row/column format.
In addition to that goal, we also made some small improvements to rendering, as well as adding
some abstraction to some of the level generation functions by using local functions. There
is still more room for improvement with how our rendering works, and out game state might also
need some cleaning, but we accomplished our main goal.

#### Full list of changes:
- Levels work on row column format
  - Tests adjusted
  - Test harness adapters eliminated
  - Level rendering adjusted for r/c format
- Text rendering code now abstracted with local function
- Bug fix: all players at exit to end a level, not just one

## Future Work

If we have extra time, we would also like to improve our rendering. Doing this would require some playing around
with the Fyne library that we are currently using to render our UI, but will make testing game development easier
as we get further into the development process.


## Conclusion

We are taking a *bit* of a down week this week in terms of work load however we have identified a number of
areas that need work, and we will be improving these areas.

Our level code is much easier to work with as of this week thanks both to the adjustment of the
format (which makes our layout room generation function much more logical), and thanks to a number
of abstractions introduced in the level generation and rendering portions of our code.