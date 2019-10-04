## Development

Postgres setup:

```
CREATE ROLE blocks WITH login password 'blocks';
CREATE DATABASE blocks_test OWNER blocks;
CREATE DATABASE blocks_dev OWNER blocks;
```
