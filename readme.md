<h1>EVERMOS CART CASE</h1>

<h3>PROBLEM</h3><br>
We are members of the engineering team of an online store. When we look at ratings for our online store application, we received the following
facts:<br>
- Customers were able to put items in their cart, check out, and then pay. After several days, many of our customers received calls from<br>
our Customer Service department stating that their orders have been canceled due to stock unavailability.<br>
- These bad reviews generally come within a week after our 12.12 event, in which we held a large flash sale and set up other major
discounts to promote our store.<br>
<br>
After checking in with our Customer Service and Order Processing departments, we received the following additional facts:<br>
- Our inventory quantities are often misreported, and some items even go as far as having a negative inventory quantity.<br>
- The misreported items are those that performed very well on our 12.12 event.<br>
- Because of these misreported inventory quantities, the Order Processing department was unable to fulfill a lot of orders, and thus
requested help from our Customer Service department to call our customers and notify them that we have had to cancel their orders.<br>

<h3>SOLUTION</h3><br>
We need to sync the actual stock with item that added to cart, so we can do reduce every customer add to cart, so if the cart is 0, customer wont abble to add item to cart.
We can restore the stock when some customer remove item from cart, so the stock will increase.
So if we do 12.12 sale, when customer add item to cart, we can prevent it from out of stock, because there's item stock checking in API.

<h3>HOW TO RUN API</h3><br>
1. clone this repository<br>
2. get all required library<br>
3. config .env with your database (using mysql)<br>
4. run the main.go<br>
5. its will auto migrate the database from models and there are seeders to fill the database<br>


<h3>DOCUMENTATION</h3><br>
You can check the API documentation in https://documenter.getpostman.com/view/3935201/TzRSgnkx

<h3>MISSING</h3><br>
Because of time limits and i have some problem on weekends, so the flow is just to decrease item stock when add to cart, and we can get all items on cart using user ID.
Also the test doesnt filled because the time limits too.
