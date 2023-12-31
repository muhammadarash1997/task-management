CREATE TABLE tasks (
    id INT NOT NULL AUTO_INCREMENT,
    user_id VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL UNIQUE,
    description VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);