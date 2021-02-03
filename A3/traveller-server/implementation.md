# MEMORANDUM

#### DATE: February 2, 2021
#### TO: Ferdinand Vesely, Project Manager
#### FROM: William Victoria and Luke Novak, Software Developers
#### SUBJECT: Issues Implementing Traveller Specification

---

We encountered several issues trying to implement this package specification, many of which would have made it impossible to return working code had we followed the spec to the letter. We attempted to write a package that delivered the functionality we believe the specification was attempting to dictate.

## Blocking Concerns
This section outlines the multiple concerns that blocked us from implementing the spec
as written. 

#### Interfaces with fields
Our first and foremost concern is with request for interfaces with
fields, as mentioned for all three requested `interfaces` (`Character`, `Town`,
and `Traveler`). The request also asks for "setters" and "getters" for these
fields. As Java interfaces do not have instance fields, we cannot possibly
fulfill this request. While Java interfaces *may* generally have fields, any field in
an interface is `public static final`, which makes it a constant. Given the
context in which these fields are used, they *need* to be mutable, which is
not possible if they are `static` and `final`. Thus, any potential interface-based
field would be functionally useless. One example of this request can
be found in the spec for the `Town` interface, where the spec mentioned that
> The interface should have a Boolean type field `is_occupied`

#### Town interface has no public methods
Our second major issue pertains specifically to the requested `Town` interface,
which is outlined in its entirety as follows:

> * The interface should have a Boolean type field `is_occupied`, which tells us if a `Character` occupies the `Town` or not. This can be initialized as *false*.
> * The interface should have a String type field `name`, which represents the name of the `Town`. This can be initialized as null.

Given our first concern, we can disregard both of these requests as being
impossible. There are now no remaining properties associated with the `Town`
interface, meaning that the interface as requested would have no content:
```Java
// represents a location in the game and a node in the game's world.
public interface Town {

}
```
The issues with this empty stub interface are twofold. 

First, the requested functionality of the `Town` is not possible to implement without any public methods, rendering the `Traveller` also impossible to implement beyond the first concern.

Second, even if we ignored the first concern and added additional public methods not outlined by the spec, we would have no class to delegate the implementation of those methods to. We could delegate the required functionality of the `Town` to an implementing class, were it not for our third concern.

#### No Implementations Requested
Given that this spec frequently relies on implementations of the three
requested interfaces, as return values, parts of object fields, and method
parameters, we need a specification for implementations of these interfaces
so that they can be instantiated on the client side. No such specification
has been given. However, the document currently contains plenty of the information that *should* be in this spec.
The main issue with this information is that it is fully encompassed in the
description of the three `interfaces`. It would not be possible to implement
the requests of the spec outside of `class`es, however the specification does
not call for any. Just that 
> interfaces for `Character`, `Town`, and `Traveller` must be created

In order for us to be able to deliver these implementations, we would need
classes, and either requests for a `Builder` or `ClassFactory`, or a constructor
with the correct input parameters. As it stands, the current spec requests that
*some* fields be initialized to `null`, but does not give any information as to the
constructor that will be initializing them.

#### world_graph Unspecified Return Type
Compared to the previous three, this request is more minor as a reasonable solution could be implied. However strictly speaking, the `world_graph` field in `Traveller`, which is specified as a `Map<Town, List>` should have a type specified for the `List` that acts as the value in the `Map`, i.e. `Map<Town, List<Town>>`, or `Map<Town, List<String>>`. While we can imply that one of these
two options are going to be the desired return type, we cannot ascertain which one
of these two types the client actually needs.

The Quote from the spec is as follows:
  - > The interface should have a Map<Town, List> type field `world_graph`


## Other Minor Concerns

Not all of our concerns were blocking, which is to say they did not prevent our
implementation of the spec. However, we felt given the opportunity to improve
the design, we should outline some other smaller issues with the design so that
the authors could make improvements for better Java.


1. The current spec does not follow Java naming conventions:
    - uses snake_case instead of camelCase, as shown here:
      > The interface should have a `add_to_network` method
2. `place_character` should not be in the Traveller interface, as it does not relate to the `Traveller`, it should be in either the `Town` or `Character` interface and called from that interface
- `IllegalStateException` is not the idiomatically correct exception to use in `place_character`.
It should be `IllegalArgumentException`, or some other custom exception


## Attempting a deliverable
  This spec specifically asked for three interfaces, we tried to reconcile this 
  with the poor implementation details by adding getters/setters to the interfaces and implementing classes for each interface. We hope that satisfies the group's
  requirements, however it deviates from the Spec rather significantly


**\#\#**