ALTER TABLE posts
    ADD CONSTRAINT FK_Users
        FOREIGN KEY (user_id) REFERENCES users (id);