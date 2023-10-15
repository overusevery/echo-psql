INSERT INTO todos 
    (
        ID, 
        Content,
        Status,
        UpdatedAt,
        CreatedAt
    )
    VALUES 
    (
        $1,
        $2,
        true,
        NOW(),
        NOW()
    )