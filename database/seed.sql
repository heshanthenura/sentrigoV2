INSERT INTO
    users (
        username,
        email,
        password_hash,
        roles
    )
VALUES (
        'admin',
        'admin@admin.com',
        '$2y$10$P2mXi55n7UHLEVg8v3Igj.3Cpj8FIXfyFMu8rq1Rvqtz1BcsmDsKm',
        ARRAY['admin']
    );