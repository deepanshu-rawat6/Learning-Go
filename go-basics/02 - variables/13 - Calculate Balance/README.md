# CALCULATE BALANCE

We need to be able to calculate costs of entire batches of messages at once.

## ASSIGNMENT

Using the variables that have been declared write conditional statements to calculate and update the variables.

First, set `finalCost` to the `bulkMessageCost`. If the user is a premium user, then the discountRate should be applied to the finalCost.

Next, if the user has enough money in their `accountBalance`:

1. Process the payment by deducting the finalCost from their accountBalance
2. Print the purchaseSuccessMessage
3. Otherwise, just print the insufficientFundMessage.