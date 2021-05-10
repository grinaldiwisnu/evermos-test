#EVERMOS CART CASE

##PROBLEM
We are members of the engineering team of an online store. When we look at ratings for our online store application, we received the following
facts:
- Customers were able to put items in their cart, check out, and then pay. After several days, many of our customers received calls from
our Customer Service department stating that their orders have been canceled due to stock unavailability.
- These bad reviews generally come within a week after our 12.12 event, in which we held a large flash sale and set up other major
discounts to promote our store.

After checking in with our Customer Service and Order Processing departments, we received the following additional facts:
- Our inventory quantities are often misreported, and some items even go as far as having a negative inventory quantity.
- The misreported items are those that performed very well on our 12.12 event.
- Because of these misreported inventory quantities, the Order Processing department was unable to fulfill a lot of orders, and thus
requested help from our Customer Service department to call our customers and notify them that we have had to cancel their orders.
  

##SOLUTION
We need to sync the actual stock with item that added to cart, so we can do reduce every customer add to cart, so if the cart is 0, customer wont abble to add item to cart.
We can restore the stock when some customer remove item from cart, so the stock will increase.
So if we do 12.12 sale, when customer add item to cart, we can prevent it from out of stock, because there's item stock checking in API.

##HO TO RUN
1. clone this repository
2. get all required library
3. config .env with your database (using mysql)
4. run the main.go
5. its will auto migrate the database from models and there are seeders to fill the database


##DOCUMENTATION
You can check the API documentation in https://www.postman.com/collections/35e9aa21b57cf9fd1339

##MISSING
Because of time limits and i have some problem on weekends, so the flow is just to decrease item stock when add to cart, and we can get all items on cart using user ID.
Also the test doesnt filled because the time limits too.
