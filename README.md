# Test Case Sagara

> Version 1.0

## Path Table

| Method | Path | Description |
| --- | --- | --- |
| POST | [/auth/login](#postauthlogin) | Login |
| POST | [/auth/register](#postauthregister) | User Registration |
| GET | [/api/users](#getapiusers) | List All Users |
| GET | [/api/users/{userId}](#getapiusersuserid) | Get User By Id |
| GET | [/api/products](#getapiproducts) | List All Products |
| POST | [/api/products](#postapiproducts) | Create Product |
| GET | [/api/products/{productId}](#getapiproductsproductid) | Get Product By Id |
| PUT | [/api/products/{productId}](#putapiproductsproductid) | Update Product By Id |
| DELETE | [/api/products/{productId}](#deleteapiproductsproductid) | Delete Product |

## Reference Table

| Name | Path | Description |
| --- | --- | --- |
| model.LoginBody | [#/components/schemas/model.LoginBody](#componentsschemasmodelloginbody) |  |
| model.LoginResponse | [#/components/schemas/model.LoginResponse](#componentsschemasmodelloginresponse) |  |
| model.Tokens | [#/components/schemas/model.Tokens](#componentsschemasmodeltokens) |  |
| model.User | [#/components/schemas/model.User](#componentsschemasmodeluser) |  |
| model.UserBody | [#/components/schemas/model.UserBody](#componentsschemasmodeluserbody) |  |
| model.Product | [#/components/schemas/model.Product](#componentsschemasmodelproduct) |  |
| model.ProductBody | [#/components/schemas/model.ProductBody](#componentsschemasmodelproductbody) |  |

## Path Details

***

### [POST]/auth/login

- Summary  
Login

- Description  
Login

#### RequestBody

- application/json

```ts
{
  username?: string
  password?: string
}
```

#### Responses

- 200 Login Success

`application/json`

```ts
{
  code?: number
  status?: string
  data: {
    tokens: {
      access_token?: string
      refresh_token?: string
    }
    user: {
      id?: number
      fullname?: string
      username?: string
      created_at?: string
      updated_at?: string
    }
  }
}
```

***

### [POST]/auth/register

- Summary  
User Registration

- Description  
User Registration

#### RequestBody

- application/json

```ts
{
  fullname?: string
  username?: string
  password?: string
}
```

#### Responses

- 200 Registration Success

`application/json`

```ts
{
  code?: number
  status?: string
  data: {
    id?: number
    fullname?: string
    username?: string
    created_at?: string
    updated_at?: string
  }
}
```

***

### [GET]/api/users

- Summary  
List All Users

- Description  
List All Users

#### Headers

```ts
Authorization: string
```

#### Responses

- 200 Success get all users

`application/json`

```ts
{
  code?: number
  status?: string
  data: {
    id?: number
    fullname?: string
    username?: string
    created_at?: string
    updated_at?: string
  }[]
}
```

***

### [GET]/api/users/{userId}

- Summary  
Get User By Id

- Description  
Get User By Id

#### Headers

```ts
Authorization: string
```

#### Responses

- 200 Success get user by id

`application/json`

```ts
{
  code?: number
  status?: string
  data: {
    id?: number
    fullname?: string
    username?: string
    created_at?: string
    updated_at?: string
  }
}
```

***

### [GET]/api/products

- Summary  
List All Products

- Description  
List All Products

#### Headers

```ts
Authorization: string
```

#### Responses

- 200 Success get all products

`application/json`

```ts
{
  code?: number
  status?: string
  data: {
    id?: number
    name?: string
    description?: string
    price?: number
    image?: string
    created_at?: string
    updated_at?: string
  }[]
}
```

***

### [POST]/api/products

- Summary  
Create Product

- Description  
Create Product

#### Headers

```ts
Authorization: string
```

#### RequestBody

- application/json

```ts
{
  name?: string
  description?: string
  price?: number
  image?: string
}
```

#### Responses

- 200 Success create product

`application/json`

```ts
{
  code?: number
  status?: string
  data: {
    id?: number
    name?: string
    description?: string
    price?: number
    image?: string
    created_at?: string
    updated_at?: string
  }
}
```

***

### [GET]/api/products/{productId}

- Summary  
Get Product By Id

- Description  
Get Product By Id

#### Headers

```ts
Authorization: string
```

#### Responses

- 200 Success get product by id

`application/json`

```ts
{
  code?: number
  status?: string
  data: {
    id?: number
    name?: string
    description?: string
    price?: number
    image?: string
    created_at?: string
    updated_at?: string
  }
}
```

***

### [PUT]/api/products/{productId}

- Summary  
Update Product By Id

- Description  
Update Product By Id

#### Headers

```ts
Authorization: string
```

#### RequestBody

- application/json

```ts
{
  name?: string
  description?: string
  price?: number
  image?: string
}
```

#### Responses

- 200 Success update product by id

`application/json`

```ts
{
  code?: number
  status?: string
  data: {
    id?: number
    name?: string
    description?: string
    price?: number
    image?: string
    created_at?: string
    updated_at?: string
  }
}
```

***

### [DELETE]/api/products/{productId}

- Summary  
Delete Product

- Description  
Delete Product

#### Headers

```ts
Authorization: string
```

#### Responses

- 200 Success delete product

`application/json`

```ts
{
  code?: number
  status?: string
}
```

## References

### #/components/schemas/model.LoginBody

```ts
{
  username?: string
  password?: string
}
```

### #/components/schemas/model.LoginResponse

```ts
{
  tokens: {
    access_token?: string
    refresh_token?: string
  }
  user: {
    id?: number
    fullname?: string
    username?: string
    created_at?: string
    updated_at?: string
  }
}
```

### #/components/schemas/model.Tokens

```ts
{
  access_token?: string
  refresh_token?: string
}
```

### #/components/schemas/model.User

```ts
{
  id?: number
  fullname?: string
  username?: string
  created_at?: string
  updated_at?: string
}
```

### #/components/schemas/model.UserBody

```ts
{
  fullname?: string
  username?: string
  password?: string
}
```

### #/components/schemas/model.Product

```ts
{
  id?: number
  name?: string
  description?: string
  price?: number
  image?: string
  created_at?: string
  updated_at?: string
}
```

### #/components/schemas/model.ProductBody

```ts
{
  name?: string
  description?: string
  price?: number
  image?: string
}
```
