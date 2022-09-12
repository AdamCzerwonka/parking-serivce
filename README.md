# System for managing parking
System designed to mangage a network of city parkings.

# How to run

# Docs
### User
#### Create new user
Path: /api/v1/user
Method: POST
Input (as json): 
```json
{
	"firstName": "John",
	"lastNmae": "Doe",
	"email": "johnd@example.com",
	"password": "SuperSecretPassword",
	"password2": "SUperSecretPassword"
}
```
Returns: 201, 400, 500
