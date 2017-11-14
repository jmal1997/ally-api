# ally-api

Ally Invest API bindings for Go.

## Status

This project is currently under development.

TODO:
* Refactor `lib` to be the root of the project.
* Define all response types to coincide with the API documentation.

The bindings take advantage of the XML output of the Ally Invest API. I
do not plan on supporting anything else. They specifically state in
their API documentation that XML is output by default. Furthermore, they
only provide sample XML output (not JSON). Therefore, I'm assuming XML
is what Ally Invest prefers to support.


## Usage

At this stage of development, I recommend you look at the CLI client
tool found in the root of the project. It is the primary interface for
testing the library.