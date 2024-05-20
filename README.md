## Aldifaa Magazine Pictures Archive

go get: 

	`go get github.com/gorilla/mux v1.8.1`
    `go get github.com/jmoiron/sqlx v1.3.5`
	`go get github.com/lib/pq v1.10.9`
	`go get golang.org/x/crypto/bcrypt`


<table>
<thead>

<tr>
<th>Route</th>
<th>Method</th>
<th>Description</th>
<th>Parameters</th>
<th>Expected Response</th>
</tr>
</thead>

<tbody>
<tr>
<td>'/login'</td>
<td>POST</td>
<td>use it to login user create serect token</td>
<td>{"username":"...", "password":"..."}</td>
<td>200 ok if user found</td>
</tr>
<tr>
<td>'/signup'</td>
<td>POST</td>
</tr>
<tr>
<td>'/signup'</td>
<td>POST</td>
</tr>
<tr>
<td>'/users'</td>
<td>GET</td>
</tr>
<tr>
<td>'/users/{id}'</td>
<td>GET</td>
</tr>
<tr>
<td>'/photos'</td>
<td>GET</td>
</tr>
<tr>
<td>'/photos/{id}'</td>
<td>GET</td>
</tr>
</tbody>
</table>
