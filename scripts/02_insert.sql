DO $$
DECLARE
i INT := 1;
BEGIN
    WHILE i <= 10000 LOOP
        INSERT INTO items (field_index, field_no_index)
        VALUES ('value_' || i, 'value_' || i);
        i := i + 1;
END LOOP;
END $$;
