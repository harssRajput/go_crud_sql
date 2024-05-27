for tab icons, modern browser made auto-request to favicon.ico ignore such request or provide the icon.


Integration testcases:
- GetAccount TCs
invalid accountId: abc
empty accountId: ""
valid accountId but account not found: 123
valid accountId and account found: 1

- createAccount TCs
invalid documentNumber: abc
empty documentNumber: ""
11 chars invalid documentNumber: "x2345678901"
valid documentNumber but already exist: 12345678901
valid documentNumber with unknown fields: {"documentNumber":"12345678902","notexistfield":"value"} -> (should get ignored)

for README
- what is the project.
- what covered in it so far. (business assumptions if taken any)
- what optional features added so far. (like env based configuration, graceful shutdown, logging, validation, sub-routing, middleware, etc.)
- explain directory structure.
- explain how to run the project.
- sample request/response for each API.
- what is pending/what can be added further.
