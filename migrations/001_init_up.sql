-- ++ goose Up
CREATE TABLE donations (
    id SERIAL PRIMARY KEY,
    food_type VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    expiration DATE NOT NULL,
    location VARCHAR(255) NOT NULL
);


-- ++ goose Down
DROP TABLE donations;