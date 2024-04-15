# Billing System

Textio is making some changes to how they bill users for bulk messages. The system should now calculate the bill based on:

1. The number of messages sent
2. The cost per message
3. Thresholds for discounts

## ASSIGNMENT

Complete the `calculateFinalBill` and the `calculateDiscount` functions.

### CALCULATEFINALBILL

Use the `calculateBaseBill` function to get the cost for the messages sent. Then, use the `calculateDiscount` function to get the discount rate. Finally, calculate the final bill by applying the discount to the base bill and return the result.

The discount is a percentage represented by a float. `10%` = `.1`

### CALCULATEDISCOUNT

This function should take the number of messages sent, and return the relevant discount based on the following discount rates:

* `10%` for more than `1000` messages.
* `20%` for more than `5000` messages.
* `0%` for anything less.


