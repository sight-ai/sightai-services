## Getting start with DB schemas
Please refer the below for the styling guide of DB schema.

### Configuration
Always use the following configuration for all tables:
```cassandraql
ENGINE = InnoDB
CHARSET = utf8mb4
COLLATE utf8mb4_unicode_ci;
```

### Naming
- Table name is `singular`, `lowercase`, `underscored` without prefix. 
- Always use `id` as the primary key which is `BIGINT UNSIGNED` and `AUTO_INCREMENT`.
- In most cases, tables should have `created_at`, `updated_at`, and `deleted_at` with types of `DATETIME NOT NULL` for both `created_at` and `updated_at`, and `DATETIME DEFAULT NULL`, respectively. 
- Index name is `lowercase`, `underscored` with prefix `idx_`, e.g., `idx_email`. For combined index, the name is the concatenated by all related fields, e.g., `idx_name_email`. 
- Do not use foreign keys at all times. 


### General guidelines
- Always use `id` as the primary key which is `BIGINT UNSIGNED` and `AUTO_INCREMENT`.
- Add a column `display_id` as another unique identifier for the same row, for entities that will be exposed to outside world.
  Generate it using UUID4. For users to see in user portals, in reports, dashboards, emails, etc.
- In most cases, tables should have `created_at`, `updated_at`, and `deleted_at` with types of `DATETIME NOT NULL` for both `created_at` and `updated_at`, and `DATETIME DEFAULT NULL`, respectively. 
- Do not use foreign keys constraints.
- For this project, we've decided to not use `meta` JSON fields, so create separate fields in the table.
  Look at `activation`/`refund` table for example.  
- If you use `BIGINT UNSIGNED` for columns in schema that can be `NULL`, in Golang you'll have to use `Null.Int64` (signed type) because the sql/driver itself only interprets ints as int64:
  * https://github.com/golang/go/issues/9373
  * https://github.com/go-sql-driver/mysql/issues/715

### Misc

* Quick test for migrations locally:  
- To install "migrate":
    
    ```
    go get -u -d github.com/mattes/migrate/cli 
    go get -u -d github.com/go-sql-driver/mysql
    go build -tags 'mysql' -o /usr/local/bin/migrate github.com/mattes/migrate/cli
    ```

- To migrate, type these cmd under migration folder

  `migrate -source file:// -database mysql://root:BM17siPDm6XZf@tcp\(34.84.94.161\)/sightai-services up`  
  
   
* please turn off NO_ZERO_DATE for mysql > 5.7
  `https://"github.com/jinzhu/gorm"/issues/595`
 