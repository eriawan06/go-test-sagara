# Test Case Sagara

> Version 1.0

## Path Table

| Method | Path | Description |
| --- | --- | --- |
| GET | [/products](#getproducts) | List All Products |
| POST | [/products](#postproducts) | Create Product |
| GET | [/products/{productId}](#getproductsproductid) | Get Product By Id |
| PUT | [/products/{productId}](#putproductsproductid) | Update Product By Id |
| DELETE | [/products/{productId}](#deleteproductsproductid) | Delete Product |

## Reference Table

| Name | Path | Description |
| --- | --- | --- |
| model.TProduct | [#/components/schemas/model.TProduct](#componentsschemasmodeltproduct) |  |
| model.TProductBody | [#/components/schemas/model.TProductBody](#componentsschemasmodeltproductbody) |  |

## Path Details

***

### [GET]/products

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

### [POST]/products

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

### [GET]/products/{productId}

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

### [PUT]/products/{productId}

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

### [DELETE]/products/{productId}

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

### #/components/schemas/model.TProduct

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

### #/components/schemas/model.TProductBody

```ts
{
  name?: string
  description?: string
  price?: number
  image?: string
}
```
