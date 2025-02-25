# Demo project

##
1. postgres indexes
2. golang channels

## Setting up
```
docker-compose up -d
```

### Prerequisites
- docker
- docker-compose

## Postgres indexes

#### Indexed field
```
explain analyze
select *
from items
where field_index = 'abc';
```
Result:
```
Index Scan using field_uniq_constraint on items  (cost=0.29..8.30 rows=1 width=28) (actual time=0.088..0.089 rows=0 loops=1)
  Index Cond: ((field_index)::text = 'abc'::text)
Planning Time: 0.194 ms
Execution Time: 0.124 ms

```

#### Unindexed field
```
explain analyze
select *
from items
where field_no_index = 'abc';
```
Result:
```
Seq Scan on items  (cost=0.00..199.00 rows=1 width=28) (actual time=2.044..2.045 rows=0 loops=1)
  Filter: ((field_no_index)::text = 'abc'::text)
  Rows Removed by Filter: 10000
Planning Time: 0.146 ms
Execution Time: 2.077 ms

```