# Table of Contents

- [Table of Contents](#table-of-contents)
- [Getting Started with Postgres with Go](#getting-started-with-postgres-with-go)
  - [Setup](#setup)
    - [Local Environment](#local-environment)
  - [Methods](#methods)
    - [Method 1: Using `chi`, `goose` and `sqlc`](#method-1-using-chi-goose-and-sqlc)
- [References](#references)

# Getting Started with Postgres with Go

PostgreSQL or Postgres is a ACID Compliant Relational Database that uses a superset of SQL as a Query Language and is the most popular SQL-based database.

To make Postgres work with Go we need the following:

1. A Server that can route requests and respond to requests
2. A Postgres Database setup either locally or on the Cloud (either on a remote server or using a managed provider like AWS Aurora or RDS)
3. Using `pgadmin` to manage Postgres using an Admin interface
4. Optional Dependencies that make it simpler to work with: E.g. `Goose` and `Sqlc`

---

## Setup

### Local Environment

If you want to setup Postgres in a Local Environment, here's a setup guide:

**Mac**:

1. [Download `pgadmin4` from Brew](https://formulae.brew.sh/cask/pgadmin4)
2. [Pull the Postgres Docker Image from DockerHub](https://hub.docker.com/_/postgres)
3. Run the docker command:

   ```bash
   docker run --name [PostgresContainerName] --restart always -d -p 5432:5432 -e POSTGRES_PASSWORD=[PASSWORD] postgres
   ```

   > **Note**: This is only for local development. There are better ways to mask passwords and persist to disk etc.

4. In `pgAdmin4`, **Register a new Server** using the following info under the `Connection` tab:

   ```yaml
   Host name/address: localhost
   Port: 5432
   Username: postgres
   Password: [POSTGRES_PASSWORD]
   ```

   This should connect to the `postgres` container we ran using `docker run`, considering the port mapping is done correctly using `-p` flag as given above.

5. Test this: Right click on the appropriate database and click the **Query Tool** option from the Drop Down. This launches a SQL client, we can use to run SQL Queries on the database. Run a basic:

   ```sql
   SELECT version()
   ```

   If you get a response, this means the database is connected and we can create tables and run operations on it.

---

## Methods

1. Using `chi`, `goose` and `sqlc`
2.

### Method 1: Using `chi`, `goose` and `sqlc`

---

# References

- [DockerHub](https://hub.docker.com/_/postgres)
-
