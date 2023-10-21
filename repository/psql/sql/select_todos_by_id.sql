SELECT 
        ID, 
        Content,
        Status,
        UpdatedAt,
        CreatedAt
FROM todos
WHERE
    1 = 1
    AND ID = $1
;