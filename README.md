## Pictures Archive

This is a  very simple archive manager for photos (only CRUD) and it just for refreshing my information about golang.

if you want to take this repository and complete this api you are welcome ^-^

# Go get: 

	`go get github.com/golang-jwt/jwt/v5 v5.2.1`
	`go get github.com/gorilla/mux v1.8.1`
    `go get github.com/jmoiron/sqlx v1.3.5`
	`go get github.com/lib/pq v1.10.9`
	`go get golang.org/x/crypto/bcrypt`


# Routes:

<table>
<thead>

<tr>
<th>Route</th>
<th>Method</th>
<th>Description</th>
<th>Json body required</th>
<th>Expected Response</th>
</tr>
</thead>

<tbody>
<!-- /login-->
<tr>
<td>'/login'</td>
<td>POST</td>
<td>use it to login user and create serect token</td>
<td>`{"username":"...",
 "password":"..."
 }`</td>
<td>200 ok if user found</td>
</tr>
<!-- /adduser-->
<tr>
<td>'/adduser'</td>
<td>POST</td>
<td>use it to register a new user</td>
<td>`{"name":"...",
 "password":"...",
 "email":"...",
 "mobile":"...",
 "password":"...",
 "permissions":"..."
 }`</td>
<td>200 ok</td>
</tr>
<!-- /addphoto-->
<tr>
<td>'/addphoto'</td>
<td>POST</td>
<td>use it to register a new photo</td>
<td>`{"name":"...",
 "photographer":"...",
 "tags":"..."
 }`</td>
<td>200 ok</td>
</tr>
<!-- /photographer-->
<tr>
<td>'/addphotographer'</td>
<td>POST</td>
<td>use it to register a new photographer</td>
<td>`{"name":"...",
 "username":"...",
 "password":"..."
 }`</td>
<td>200 ok</td>
</tr>
<!-- /users-->
<tr>
<td>'/users'</td>
<td>GET</td>
<td>use it to get a list of all users</td>
<td></td>
<td>200 ok</td>
</tr>
<!-- /users/{id}-->
<tr>
<td>'/users/{id}'</td>
<td>GET</td>
<td>use it to get a specific user by its id</td>
<td></td>
<td>200 ok</td>
</tr>
<!-- /photos-->
<tr>
<td>'/photos'</td>
<td>GET</td>
<td>use it to get a list of all photos</td>
<td></td>
<td>200 ok if user found</td>
</tr>
<!-- /photos/{id}-->
<tr>
<td>'/photos/{id}'</td>
<td>GET</td>
<td>use it to get a specific photo by its id</td>
<td></td>
<td>200 ok</td>
</tr>
<!-- /photographers-->
<tr>
<td>'/photoographers'</td>
<td>GET</td>
<td>use it to get a list of all photographers data</td>
<td></td>
<td>200 ok</td>
</tr>
<!-- /photographers/{id}-->
<tr>
<td>'/photographers/{id}'</td>
<td>GET</td>
<td>use it to get a specific photographer by its id</td>
<td></td>
<td>200 ok</td>
</tr>
</tbody>
</table>

# Note:
This is not public software, so only admins can add new users to the system.