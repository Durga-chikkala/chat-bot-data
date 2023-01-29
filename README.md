﻿# chat-bot-data
 
 #### To Run this service we first need to run the Docker use the below command to run the docker 
 docker run --name chatbot -e POSTGRES_PASSWORD=aditya -d postgres

This service uses postgres Database to store the data 

and it contains the following Endpoints


#### user endpoints:

/user/login:

- GET Method 
- this takes 2 query param which are required params
  - Email
  - Password 
- It returns the data if it is present otherwise return empty data which means the user is not registered
<hr/>

/user/signup:
- POST Method 
- takes request body
   - firstName
   - lastName
   - dob
   - gender
   - email
   - phone 
   - password
- we need to provide them in the right format all the validations are taking place 
<hr/>

/user/{id}:
- GET method 
- here id refers to uuid which is generated in create 
- we need to provide the uuid and using that we can fetch the data 
- return 404 if id is not found 
<hr/>

/user/{id}:
- PATCH method
- we can update password 
- password is sent in the request body 
- id here is uuid which is unique to all the records 


#### query endpoints:

/chatbot:

- POST method
- The Request Body is :
  - question
  - solution 
- return the data with ID which uuid generated inside
<hr/>

/chatbot/{question}:
- GET Method 
- takes question as path param 
- since we are making search empty spaces between the words are converted to '-'
<hr/>

/chatbot:
- GET Method
- return all the data 

/chatbot/{question}:
- PATCH method 
- updates the count


 
