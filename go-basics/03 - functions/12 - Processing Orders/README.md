# PROCESSING ORDERS

Management thinks our branding is so creative that our SaaS customers will pay for Textio merch.

## ASSIGNMENT

Complete the `placeOrder` function.

It returns a `bool` indicating whether the order was successful (`true` is a success) and a `float64` representing the user's balance after the order.

The `amountInStock` and `calcPrice` functions can be used to look up the current stock and price of an item.

* If the quantity is greater than the amount in stock, the order should be rejected.
* If the user doesn't have enough money in their account, the order should be rejected.
* Otherwise, the order should be accepted and you should return the new balance.