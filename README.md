# restapi

This is an example rest api written in Go.

## How to run this api locally

- Clone this repository in your machine.

- Open your terminal and navigate to the project directory/folder.

- Create a `.env` file in the root directory and add the following configuration fields:

  ```
    SERVER_PORT = 3000

    DB_DSN = ../../storage/sqlite/api.db
  ```

- Run `cd cmd/app` to navigate to `cmd/app` directory. Then build with project by running `go build`. This will generate your executable in the current directory.

- Run the executable in your terminal by running `./app`

## The following endpoints are available

- ### Create new user

  ```JSON
  URL: http://localhost:3000/api/users
  Method: POST
  Body:
    {
        "first_name": "Derek",
        "last_name": "Shoffer",
        "email": "dshofer@email.com",
        "password": "ThisIsRawDime"
    }

  Example response:
    {
        "description": "User created successfully",
        "user": {
            "id": 9,
            "first_name": "Derek",
            "last_name": "Shoffer"
        }
    }
  ```

- ### Get user by id

  ```JSON
  URL: http://localhost:3000/api/users/9
  Method: GET
  Example response:
    {
        "description": "found 1 user",
        "user": {
            "id": 9,
            "first_name": "Derek",
            "last_name": "Shoffer"
        }
    }

  ```

- ### Update a user

  ```JSON
  URL: http://localhost:3000/api/users/9
  Method: PUT
  Body:
    {
        "last_name": "shoferrrrr",
        "email": "newEmail@email.com"
    }

  Example response:
    {
        "description": "user updated successfully",
        "user": {
            "id": 9,
            "first_name": "Derek",
            "last_name": "Shofferrrrrr"
        }
    }
  ```

  - ### Delete a user

  ```JSON
  URL: http://localhost:3000/api/users/9
  Method: DELETE
  Example response:
    {
        "description": "user with id 9 was deleted successfully"
    }
  ```

  - ### Get all users

  ```JSON
  URL: http://localhost:3000/api/users
  Method: GET
  Example response:
    {
        "description": "found 4 users",
        "users": [
            {
                "id": 1,
                "first_name": "Bruce",
                "last_name": "Banner"
            },
            {
                "id": 6,
                "first_name": "Garret",
                "last_name": "Mike"
            },
            {
                "id": 7,
                "first_name": "Monsignor",
                "last_name": "Makiaveli"
            },
            {
                "id": 8,
                "first_name": "Peter",
                "last_name": "Petrelli"
            }
        ]
    }
  ```
