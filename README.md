# auth-service

- Inspired by https://gist.github.com/azagniotov/a4b16faf0febd12efbc6c3d7370383a6
- RESTful APIs running on `localhost:7999`
- gRPC APIs running on `localhost:50052`

---

### Creating new JWT token

<details>
<summary><code>POST</code> <code><b>/generate</b></code> <code>(Generate a new JWT token given the info)</code></summary>

##### Body (application/json or application/x-www-form-urlencoded)

> | key  | required | data type | description |
> | ---- | -------- | --------- | ----------- |
> | id   | true     | string    | N/A         |
> | mail | true     | string    | N/A         |
> | name | true     | string    | N/A         |

##### Responses

> | http code    | content-type       | response                                               |
> | ------------ | ------------------ | ------------------------------------------------------ |
> | `200`        | `application/json` | `{"message": "Success", "token": "Bearer your_token"}` |
> | `400`, `500` | `application/json` | `{"message": "Failed", "error":"Error messages"}`      |

</details>

---

### Verifying existing JWT token

<details>
<summary><code>GET</code> <code><b>/verify</b></code> <code>(Verify a existing JWT token)</code></summary>

##### Headers

> | key           | value         | description                 |
> | ------------- | ------------- | --------------------------- |
> | Authorization | The JWT token | Starts with `Bearer<space>` |

##### Responses

```typescript
type jwtContent = {
  id: string
  mail: string
  name: string
}
```

> | http code | content-type       | response                                           |
> | --------- | ------------------ | -------------------------------------------------- |
> | `200`     | `application/json` | `{"message": "Success", "jwtContent": jwtContent`} |
> | `400`     | `application/json` | `{"message": "Failed", "error":"Error messages"}`  |

</details>
