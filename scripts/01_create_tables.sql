CREATE TABLE items
(
    id             BIGSERIAL PRIMARY KEY,
    field_index    varchar(255) not null
        constraint field_uniq_constraint unique,
    field_no_index varchar(255) not null
);

