# Todo List API

## Description
An API for a Todo List Application created as the 2nd assignment of admin oprec from SE Lab (RPL). The code is divided into three parts which are as follows:
- `structs` which functions like entity of Clean Architecture that is filled with structs used in the API
- `handleFunc` which functions like the combination of controller and service of Clean Architecture that handle requests by the specified routes
- `dbFunc` which functions like repository of Clean Architecture that do direct contacts with the database

## Tech Stack
- Golang using `net/http` (Back-end)
- PostgreSQL (Database)

## Features
- Supports CRUD operations for Todos
    - Create (C)
        - Adding New Todo with validation included
    - Read (R)
        - Get All Todos Available
        - Search for Todos by query
        - Filter Todos by start and finish time
    - Update (U)
        - Edit Todo
    - Delete
        - Delete Todo

## Hardships I Felt
- Me, who at first never used Golang and therefore am not experienced in the Golang environment have to first learn quite a lot about how to work in Golang, especially using `net/http`
- In my first assignment, I only used one main.go file for the full code which is very bad. So, I decided to change things up and improve by trying to separate each codes according to their functionality. That said, I have to first learn about packages and dependencies in Golang which are very hazy for me at the beginning. But after learning for some hours, I finally got it and able to figure out the usage.

## Things I Learned
Doing this assignment pushes me to be able to learn a lot of things in creating API using Golang without Framework, which for me feels like quite a hassle to create just a single feature.
<br>
But, overall, completing this assignment give me a lot of ideas regarding how to work in the Golang environment by creating functional programming using a still quite bad copy of the Clean Architecture.
<br>
I've also learned a lot about packages in Golang and how they can be used to link everything together as one fully functional API, which improve the readability and clear the functionality of each code even further.