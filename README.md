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

- create transaction
cannot create transaction with amount 0
cannot create trx with invalid accId
cannot create trx with non-exist accId
cannot create trx with invalid opTypeId
cannot create trx with non-exist opTypeId
cannot create trx with invalid amount
cannot create trx with -ve amount for credit voucher(4)
cannot create trx with +ve amount for other than credit voucher(4)

for README
- what is the project.
- what covered in it so far. (business assumptions if taken any)
- what optional features added so far. (like env based configuration, graceful shutdown, logging, validation, sub-routing, middleware, etc.)
- explain directory structure.
- explain how to run the project.
- explain how to run the unit testcases.
- explain how to run the integration testcases.
- explain how to run the coverage. (optional)
- sample request/response for each API.
- what is pending/what can be added further.
- - no boundary objects created for each API. using entity objects only. once entity/request objects gets complicated we can introduce boundary objects.
- what is the approach taken for the project.
